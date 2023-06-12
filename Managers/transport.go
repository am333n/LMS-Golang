package managers

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetManagersRequest struct {
	manager []Manager
}
type GetManagersResponse struct {
	manager []Manager
	Err     string `json:"err,omitempty"`
}
type PostManagerRequest struct {
	manager Manager
}
type PostManagerResponse struct {
	V   string `json:"Result:"`
	Err string `json:"err,omitempty"`
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func DecodeGetManagersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request struct{}
	return request, nil
}
func DecodePostManagerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request PostManagerRequest
	if err := json.NewDecoder(r.Body).Decode(&request.manager); err != nil {
		return nil, err
	}
	return request, nil
}
