package Employee

import (
	"errors"
	dc "lms/Database"

	"gorm.io/gorm"
)

type Employees struct {
	gorm.Model `json:"-"`
	Name       string `json:"employee_name"`
	Address    string `json:"employee_address"`
	Department string `json:"employee_department"`
	Post       string `json:"employee_post"`
	DOB        string `json:"employee_dob"`
	Email      string `json:"employee_email"`
	Phone      int    `json:"employee_phone"`
}

var employee Employees

// servics
type Service interface {
	PostEmployee(employee Employees) (Employees, error)
	GetEmployees() ([]Employees,error)
	GetEmployeeById(id int) (Employees, error)
	DeleteEmployeeById(id int) (string, error)
	UpdateEmployee(id int,employee Employees) (string,Employees, error)
}
type RepoService struct{}


func (RepoService) PostEmployee(employee Employees) (Employees, error) {
	var empty Employees
	if employee.Name == "" {
		return empty, ErrEmpty
	}
	db, err := dc.GetDB()
	if err != nil {
		return empty,err
	}
	err = db.Create(&employee).Error
	if err != nil {
		return empty, err
	}
	return employee, nil
}
func (RepoService) GetEmployees()([]Employees, error){
	var employees []Employees
	db,err:=dc.GetDB()
	if err!=nil{
		return employees,err
	}
	err=db.Find(&employees).Error
	if err!=nil{
		return employees,err
	}
	return employees,nil
}
func (RepoService)GetEmployeeById(id int)(Employees,error){
	var employee Employees
	db,err:=dc.GetDB()
	if err!=nil{
			return employee,err
		}
	err=db.Where("id =?", id).First(&employee).Error
	if err!=nil{
        return employee,err
    }
	return employee,nil
}
func (RepoService)DeleteEmployeeById(id int)(string,error){
	var employee Employees
	db,err:=dc.GetDB()
	if err!=nil{
		return "",err
	}
	err=db.Where("id =?", id).Delete(&employee).Error
	if err!=nil{
		return "",err
	}
	return "The data is deleted", nil
}
func (RepoService)UpdateEmployee(id int,employee Employees) (string, Employees, error) {
	db,err:=dc.GetDB()
	if err!=nil{
        return "",employee,err
    }
	err=db.Where("id=?",id).Updates(&employee).Error
	if err!=nil{
        return "",employee,err
    }
	return "The data is updated", employee, nil
}
// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("empty employee name")
