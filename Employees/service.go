package Employee

import (
	"errors"
	dc "lms/Database"

	"gorm.io/gorm"
)

//? --------------------------------- Tables --------------------------------- */

type Employees struct {
	gorm.Model `json:"-"`
	Name       string `json:"employee_name"`
	Address    string `json:"employee_address"`
	Department string `json:"employee_department"`
	Post       string `json:"employee_post"`
	DOB        string `json:"employee_dob"`
	Email      string `json:"employee_email"`
	Phone      int    `json:"employee_phone"`
	//pending or accepted status
	Status string `json:"employee_status"`
}
type Leaves struct {
	LeaveID        int `gorm:"autoIncrement"`
	EmployeeID     int `json:"employee_id"`
	MedicalLeave   int `json:"medical_leave"`
	MedicalTaken   int `json:"medical_taken"`
	AnnualLeave    int `json:"Annual_leave"`
	AnnualTaken    int `json:"Annual_taken"`
	EmergencyLeave int `json:"Emergency_Leave"`
	EmergencyTaken int `json:"Emergency_Taken"`
}
type Requests struct {
	RequestID int `gorm:"autoIncrement"`
	EmployeeID int `json:"employee_id"`
	LeaveID int `json:"leave_id"`
	Type string `json:"type"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	Days int `json:"days"`
	Reason string `json:"reason"`
	Status string `json:"status"`
}


/* -------------------------------------------------------------------------- */
//*=============================== Service =================================*/
/* -------------------------------------------------------------------------- */

type Service interface {
	PostEmployee(employee Employees) (Employees, error)
	GetEmployees() ([]Employees, error)
	GetEmployeeById(id int) (Employees, error)
	DeleteEmployeeById(id int) (string, error)
	UpdateEmployee(id int, employee Employees) (string, Employees, error)
	PostLeaves(id int) (string, error)
	DeleteLeaves(id int) (string, error)
	EnterLeaves(id int,leave Leaves) ([]Leaves, error)
	ApproveEmployee(id int) (string, error)
}
type RepoService struct{}

//* ------------------------- Employee CRUD Functions ------------------------ */

func (RepoService) PostEmployee(employee Employees) (Employees, error) {
	var empty Employees
	if employee.Name == "" {
		return empty, ErrEmpty
	}
	db, err := dc.GetDB()
	if err != nil {
		return empty, err
	}
	err = db.Create(&employee).Error
	if err != nil {
		return empty, err
	}
	return employee, nil
}
func (RepoService) GetEmployees() ([]Employees, error) {
	var employees []Employees
	db, err := dc.GetDB()
	if err != nil {
		return employees, err
	}
	err = db.Find(&employees).Error
	if err != nil {
		return employees, err
	}
	return employees, nil
}
func (RepoService) GetEmployeeById(id int) (Employees, error) {
	var employee Employees
	db, err := dc.GetDB()
	if err != nil {
		return employee, err
	}
	err = db.Where("id =?", id).First(&employee).Error
	if err != nil {
		return employee, err
	}
	return employee, nil
}
func (RepoService) DeleteEmployeeById(id int) (string, error) {
	var employee Employees
	db, err := dc.GetDB()
	if err != nil {
		return "", err
	}
	err = db.Where("id =?", id).Delete(&employee).Error
	if err != nil {
		return "", err
	}
	return "The data is deleted", nil
}
func (RepoService) UpdateEmployee(id int, employee Employees) (string, Employees, error) {
	db, err := dc.GetDB()
	if err != nil {
		return "", employee, err
	}
	err = db.Where("id=?", id).Updates(&employee).Error
	if err != nil {
		return "", employee, err
	}
	return "The data is updated", employee, nil
}
func (RepoService) ApproveEmployee(id int) (string, error){
	db,err:=dc.GetDB()
	if err!=nil{
		return "",err
	}
	var employee Employees
	if err := db.Table("employees").Where("id=?", id).Scan(&employee).Error; err != nil {
		return "No Employee Found", err
	}
	if employee.Status == "pending" || employee.Status == ""{
	
		employee.Status="Approved"
		if err:=db.Where("id=?",id).Updates(&employee).Error; err != nil {
		return "", err
		}
		return "The employee is Approved",nil
	}
	return "The employee is already approved",nil
}
// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("empty employee name")

//* ----------------------------- Leave Functions ---------------------------- */
//new

func (RepoService) PostLeaves(id int) (string, error) {
	db, err := dc.GetDB()
	var leave Leaves
	var employee Employees
	if err != nil {
		return "", err
	}
	//to check for the employee
	if err := db.Session(&gorm.Session{}).Where("id =?", id).First(&employee).Error; err != nil {
		return "Employee Not Found", nil
	}
	//to check for already set entry in leave
	if err := db.Session(&gorm.Session{}).Where("employee_id=?", id).First(&leave).Error; err != nil {
		leave.EmployeeID = id
		leave.AnnualLeave = 0
		leave.AnnualTaken = 0
		leave.MedicalLeave = 0
		leave.MedicalTaken = 0
		leave.EmergencyLeave = 0
		leave.EmergencyTaken = 0
		err = db.Create(&leave).Error
		if err != nil {
			return "", err
		}
		return "Leave Table for the Employee is set", nil

	}
	return "Leave Already Set", err
}
func (RepoService) DeleteLeaves(id int) (string, error) {
	db, err := dc.GetDB()
	var leave Leaves
	if err != nil {
		return "", err
	}
	if err := db.Where("employee_id=?", id).Delete(&leave).Error; err != nil {
		return "", err
	}
	return "Leave Table for the Employee is deleted", nil

}

func (RepoService) EnterLeaves(id int, leave Leaves) ([]Leaves, error) {
	var leaves []Leaves
	db, err := dc.GetDB()
    if err!= nil {
        return nil, err
    }
    err = db.Where("employee_id=?", id).Updates(&leave).Error
	if err!= nil {
        return nil, err
    }
	db.Find(&leaves)
    return leaves, nil
}

/* ---------------------------- Request Functions --------------------------- */

