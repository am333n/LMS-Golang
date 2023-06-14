package managers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type GetManagersResponse struct {
	V []Manager `json:"Response:"`
	Err     string `json:"err,omitempty"`
}
type PostManagerRequest struct {
	manager Manager
}
type PostManagerResponse struct {
	V   string `json:"Result:"`
	Err string `json:"err,omitempty"`
}
//abcd
type GetManagerByIdResponse struct {
	V Manager `json:"Result:"`
	Err     string `json:"err,omitempty"`
}
type DeleteManagerResponse struct{
	V string `json:"Result:"`
	Err string `json:"err,omitempty"`
}
type UpdateManagerRequest struct {
	id int
	manager Manager
}
type UpdateManagerResponse struct{
	V string `json:"Output:"`
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
	request.manager.Status="Pending"
	return request, nil
}
func DecodeGetManagerByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	id ,err:= strconv.Atoi(params["id"])
	if err!= nil {
			return nil, err
		}
	return id, nil

}
func DecodeDeleteManagerRequest(_ context.Context, r *http.Request) (interface{},error) { 
	params:=mux.Vars(r)
	id,err:= strconv.Atoi(params["id"])
	if err!= nil {
				return nil,err
	}
	return id,	nil
}
func DecodeUpdateManagerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UpdateManagerRequest
	params:=mux.Vars(r)
	id,err:=strconv.Atoi(params["id"])
	if err!= nil {
		return nil,err
	}
	request.id=id
	
	if err:=json.NewDecoder(r.Body).Decode(&request.manager);err!=nil{
		return request,err
	}
	return request,nil
}
