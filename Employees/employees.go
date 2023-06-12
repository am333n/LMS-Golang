package Employee

// import (
// 	"encoding/json"
// 	dc "lms/Database"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"gorm.io/gorm"
// )

// type Employees struct {
// 	gorm.Model
// 	Name       string `json:"employee_name"`
// 	Address    string `json:"employee_address"`
// 	Department string `json:"employee_department"`
// 	Post       string `json:"employee_post"`
// 	DOB        string `json:"employee_dob"`
// 	Email      string `json:"employee_email"`
// 	Phone      int    `json:"employee_phone"`
// }

// func GetEmployees(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var employees []Employees
// 	dc.DB.Find(&employees)
// 	json.NewEncoder(w).Encode(employees)

// }

// func GetEmployee(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var employees Employees
// 	params := mux.Vars(r)
// 	dc.DB.First(&employees, params["id"])
// 	json.NewEncoder(w).Encode(employees)

// }

// func PostEmployee(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var employee Employees
// 	json.NewDecoder(r.Body).Decode(&employee)
// 	dc.DB.Create(&employee)
// 	json.NewEncoder(w).Encode(employee)
// }

// func PutEmployee(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var employees Employees
// 	params := mux.Vars(r)
// 	dc.DB.First(&employees, params["id"])
// 	json.NewDecoder(r.Body).Decode(&employees)
// 	dc.DB.Save(&employees)
// 	json.NewEncoder(w).Encode(employees)

// }

// func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var employees Employees
// 	params := mux.Vars(r)
// 	dc.DB.Delete(&employees, params["id"])
// 	json.NewEncoder(w).Encode("the user is deleted")
// }
