package user

import (
	"context"
	"encoding/json"
	"net/http"

	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
)

type setPreferredLocationRequest struct {
	Username string `json:"username"`
	Location string `json:"location"`
}

type setPreferredLocationResponse struct {
	Status bool   `json:"status"`
	Err    string `json:"err,omitempty"`
}

type getPreferredLocationRequest struct {
	Username string `json:"username"`
}

type getPreferredLocationResponse struct {
	Location string `json:"Location"`
	Err      string `json:"err,omitempty"`
}

//MakeUserEndpoint creates an end point
func MakeUserEndpoint(svc Service) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {
		req := request.(setPreferredLocationRequest)
		claims := context.Value(jwt.JWTClaimsContextKey).(*stdjwt.StandardClaims)
		status, err := svc.SetPreferredLocation(claims.Subject, req.Location)
		if err != nil {
			return setPreferredLocationResponse{status, err.Error()}, nil
		}
		return setPreferredLocationResponse{status, ""}, nil
	}
}

//DecodeUserRequest decodes request
func DecodeUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request setPreferredLocationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//EncodeUserResponse encodes response
func EncodeUserResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

//MakeGetPreferredLocationEndpoint creates an end point
func MakeGetPreferredLocationEndpoint(svc Service) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {
		claims := context.Value(jwt.JWTClaimsContextKey).(*stdjwt.StandardClaims)
		preferredLocation, err := svc.GetPreferredLocation(claims.Subject)
		if err != nil {
			return getPreferredLocationResponse{"", err.Error()}, nil
		}
		return getPreferredLocationResponse{preferredLocation, ""}, nil
	}
}

//DecodePreferredLocationRequest decodes request
func DecodePreferredLocationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getPreferredLocationRequest
	return request, nil
}

//EncodeGetPreferredLocationResponse encodes response
func EncodeGetPreferredLocationResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
