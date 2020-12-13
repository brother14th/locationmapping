package location

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/brother14th/locationmapping/db"
)

type locationReportRequest struct {
	Location string `json:"location"`
}

type locationReportResponse struct {
	LocationReport db.LocationReport `json:"summary,omitempty"`
	Err            string            `json:"err,omitempty"`
}

//MakeLocationReportEndpoint creates endpoint
func MakeLocationReportEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(locationReportRequest)
		locationReport, err := svc.GetLocationReport(req.Location)
		if err != nil {
			return locationReportResponse{locationReport, err.Error()}, nil
		}
		return locationReportResponse{locationReport, ""}, nil
	}
}

//DecodeLocationReportRequest decodes request
func DecodeLocationReportRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request locationReportRequest
	location := r.URL.Query().Get("location")
	if location != "" {
		request.Location = location
		return request, nil
	}
	return nil, errors.New("location parameter is not set")
}

//EncodeLocationReportResponse encodes response
func EncodeLocationReportResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
