package main

import (
	"crypto/subtle"
	"log"
	"net/http"

	"github.com/brother14th/locationmapping/authentication"
	"github.com/brother14th/locationmapping/db"
	"github.com/brother14th/locationmapping/location"
	"github.com/brother14th/locationmapping/user"

	stdjwt "github.com/dgrijalva/jwt-go"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	kithttp "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	basicAuthUser = "prometheus"
	basicAuthPass = "password"
)

func basicAuth(username string, password string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="metrics"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorised\n"))
			return
		}

		h.ServeHTTP(w, r)
	})
}

func methodControl(method string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			h.ServeHTTP(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func main() {

	key := []byte("secret_key")
	keys := func(token *stdjwt.Token) (interface{}, error) {
		return key, nil
	}
	jwtOptions := []kithttp.ServerOption{
		kithttp.ServerBefore(kitjwt.HTTPToContext()),
	}

	fieldKeys := []string{"method", "error"}
	userRepository, _ := db.NewUserRepository()
	locationRepository, _ := db.NewLocationRepository()

	authSvc := authentication.NewService(userRepository, key)
	authSvc = authentication.NewInstrumentingAuthMiddleware(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "auth_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "auth_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		authSvc,
	)

	userSvc := user.NewService(userRepository)
	userSvc = user.NewInstrumentingUserMiddleware(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "user_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "user_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		userSvc,
	)

	locationReportSvc := location.NewService(locationRepository)
	locationReportSvc = location.NewInstrumentingLocationReportMiddleware(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "location_report_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "location_report_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		locationReportSvc,
	)

	authHandler := kithttp.NewServer(
		authentication.MakeAuthenticateEndpoint(authSvc),
		authentication.DecodeAuthRequest,
		authentication.EncodeAuthResponse,
	)
	userHandler := kithttp.NewServer(
		kitjwt.NewParser(keys, stdjwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(user.MakeUserEndpoint(userSvc)),
		user.DecodeUserRequest,
		user.EncodeUserResponse,
		jwtOptions...,
	)
	preferredLocationHandler := kithttp.NewServer(
		kitjwt.NewParser(keys, stdjwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(user.MakeGetPreferredLocationEndpoint(userSvc)),
		user.DecodePreferredLocationRequest,
		user.EncodeGetPreferredLocationResponse,
		jwtOptions...,
	)
	locationReportHandler := kithttp.NewServer(
		kitjwt.NewParser(keys, stdjwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(location.MakeLocationReportEndpoint(locationReportSvc)),
		location.DecodeLocationReportRequest,
		location.EncodeLocationReportResponse,
		jwtOptions...,
	)

	http.Handle("/v1/authenticate", methodControl("POST", authHandler))
	http.Handle("/v1/user", methodControl("PATCH", userHandler))
	http.Handle("/v1/preferredlocation", methodControl("GET", preferredLocationHandler))
	http.Handle("/v1/locationreport", methodControl("GET", locationReportHandler))
	http.Handle("/metrics", basicAuth(basicAuthUser, basicAuthPass, promhttp.Handler()))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
