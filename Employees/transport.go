package Employee

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)
/* ----------------------- employee Request & Response ---------------------- */
type PostEmployeeRequest struct {
	employee Employees
}
type GetEmployeesRequest struct {
	employee []Employees
}
type UpdateEmployeeRequest struct {
	id int
	employee Employees
}
type PostEmployeeResponse struct {
	V   Employees `json:"Result"`
	Err string    `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}
type GetEmployeesResponse struct {
	V   []Employees `json:"Result"`
	Err string      `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}
type GetEmployeeByIdResponse struct{
	V Employees `json:"Resut"`
	Err string    `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}
type DeleteEmployeeByIdResponse struct{
	V string `json:"Output"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}
type UpdateEmployeeResponse struct{
    V string `json:"Result"`
	A Employees `json:"Output"`
    Err string    `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}
/* -------------------- Leave Function Request & Response ------------------- */

type EnterLeaveRequest struct {
	id int
	leave Leaves
}
type EnterLeaveResponse struct {
    V   []Leaves `json:"Result"`
    Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

/* ---------------------- Employee CRUDEmcode & Decode ---------------------- */
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
func DecodePostEmployeeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request PostEmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&request.employee); err != nil {
		return nil, err
	}
	request.employee.Status="Pending"
	return request, nil
}
func DecodeGetEmployeesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request struct{}
	// if err:=json.NewDecoder(r.Body).Decode(&request.employee);err!=nil{
	// 	return nil,err
	// }
	return request, nil
}
func DecodeGetEmployeeByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	id,err:=strconv.Atoi(params["id"])
	if err!=nil{
        return nil,err
    }
	return id, nil
}
func DecodeDeleteEmployeeByIdRequest(_ context.Context, r *http.Request)(interface{},error){
	params:=mux.Vars(r)
	id,err:=strconv.Atoi(params["id"])
	if err!=nil{
        return nil,err
    }
	return id, nil
}
func DecodeUpdateEmployee(_ context.Context, r *http.Request)(interface{},error){
	var request UpdateEmployeeRequest
	params:=mux.Vars(r)
    id,err:=strconv.Atoi(params["id"])
    if err!=nil{
        return nil,err
    }
	request.id=id
	if err:=json.NewDecoder(r.Body).Decode(&request.employee); err!=nil{
		return nil, err
	}
	return request,nil

}
func DecodeEnterLeaveRequest(_ context.Context, r *http.Request)(interface{},error){
	params:=mux.Vars(r)
    id,err:=strconv.Atoi(params["id"])
    if err!=nil{
        return nil,err
    }
	var request EnterLeaveRequest
	request.id=id
	if err:=json.NewDecoder(r.Body).Decode(&request.leave); err!=nil{
        return nil, err
    }
    return request, nil
}

