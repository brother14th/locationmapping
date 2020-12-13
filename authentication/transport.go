package authentication

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type authenticateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authenticateResponse struct {
	Token string `json:"token"`
	Err   string `json:"err,omitempty"`
}

//MakeAuthenticateEndpoint creates endpoint
func MakeAuthenticateEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(authenticateRequest)
		token, err := svc.Authenticate(req.Username, req.Password)
		if err != nil {
			return authenticateResponse{"", err.Error()}, nil
		}
		return authenticateResponse{token, ""}, nil
	}
}

//DecodeAuthRequest decodes request
func DecodeAuthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request authenticateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//EncodeAuthResponse encodes response
func EncodeAuthResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
