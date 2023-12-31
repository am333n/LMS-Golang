package Employee

import (
	"context"
	"errors"
	"fmt"
	dc "lms/Database"
	"lms/common"

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
	RequestID  int    `gorm:"autoIncrement"`
	EmployeeID int    `json:"employee_id"`
	LeaveID    int    `json:"leave_id"`
	Type       string `json:"type"`
	DateFrom   string `json:"date_from"`
	DateTo     string `json:"date_to"`
	Days       int    `json:"days"`
	Reason     string `json:"reason"`
	Status     string `json:"status"`
}

/* -------------------------------------------------------------------------- */
//*=============================== Service =================================*/
/* -------------------------------------------------------------------------- */

type Service interface {
	PostEmployee(employee Employees) (Employees, error)
	GetEmployees(ctx context.Context) ([]Employees, error)
	GetEmployeeById(id int) (Employees, error)
	DeleteEmployeeById(ctx context.Context,id int) (string, error)
	UpdateEmployee(ctx context.Context,id int, employee Employees) (string, Employees, error)
	PostLeaves(ctx context.Context,id int) (string, error)
	DeleteLeaves(ctx context.Context,id int) (string, error)
	EnterLeaves(ctx context.Context,id int, leave Leaves) ([]Leaves, error)
	ApproveEmployee(ctx context.Context,id int) (string, error)
	PostLeaveRequest(ctx context.Context,request Requests) (string, error)
	GetLeaveRequest(ctx context.Context) ([]Requests, error)
	ApproveLeaveRequest(ctx context.Context,id int) (string, error)
	DeleteLeaveRequest(ctx context.Context,id int) (string, error)
	GetLeaves(ctx context.Context) ([]Leaves, error)
	GetLeavesById(ctx context.Context,id int) (Leaves, error)
}

// type s RepoService struct{}

func NewService(controller Controller) Service {
	return &RepoService{
		controller: controller,
	}
}

//* ------------------------- Employee CRUD Functions ------------------------ */

func (s RepoService) PostEmployee(employee Employees) (Employees, error) {
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
func (s RepoService) GetEmployees(ctx context.Context) ([]Employees, error) {
	var employees []Employees
	_,err:=s.controller.CheckIfAdminOrManager(ctx)
	if err!=nil{
		return nil,err
	}
	
	db, err := dc.GetDB()
	if err != nil {
		return employees, err
	}
	err = db.Find(&employees).Error
	if err != nil {
		return employees, common.ErrNoFile
	}
	return employees, nil
}
func (s RepoService) GetEmployeeById(id int) (Employees, error) {
	var employee Employees

	db, err := dc.GetDB()
	if err != nil {
		return employee, err
	}
	err = db.Where("id =?", id).First(&employee).Error
	if err != nil {
		return employee, common.ErrIDNotFound
	}
	return employee, nil
}
func (s RepoService) DeleteEmployeeById(ctx context.Context,id int) (string, error) {
	var employee Employees
	_,err:=s.controller.CheckIfAdmin(ctx)
	if err!=nil{
		return "",common.ErrNoPermission
	}
	db, err := dc.GetDB()
	if err != nil {
		return "", err
	}
	err = db.Where("id =?", id).Delete(&employee).Error
	if err != nil {
		return "", common.ErrIDNotFound
	}
	return "The data is deleted", nil
}
func (s RepoService) UpdateEmployee(ctx context.Context,id int, employee Employees) (string, Employees, error) {
	_,err:=s.controller.CheckIfAdmin(ctx)
	if err!=nil{
		return "",employee,common.ErrNoPermission
	}
	db, err := dc.GetDB()
	if err != nil {
		return "", employee, err
	}
	err = db.Where("id=?", id).Updates(&employee).Error
	if err != nil {
		return "", employee, common.ErrIDNotFound
	}
	return "The data is updated", employee, nil
}
func (s RepoService) ApproveEmployee(ctx context.Context,id int) (string, error) {
	_,err:=s.controller.CheckIfAdmin(ctx)
	if err!=nil{
		return "",common.ErrNoPermission
	}
	db, err := dc.GetDB()
	if err != nil {
		return "", err
	}
	var employee Employees
	if err := db.Table("employees").Where("id=?", id).Scan(&employee).Error; err != nil {
		return "No Employee Found", err
	}
	if employee.Status == "pending" || employee.Status == "" {

		employee.Status = "Approved"
		if err := db.Where("id=?", id).Updates(&employee).Error; err != nil {
			return "", common.ErrIDNotFound
		}
		return "The employee is Approved", nil
	}
	return "The employee is already approved", nil
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("empty employee name")

//* ----------------------------- Leave Functions ---------------------------- */
//new

func (s RepoService) PostLeaves(ctx context.Context,id int) (string, error) {
	_,err:=s.controller.CheckIfAdmin(ctx)
	if err!=nil{
		return "",common.ErrNoPermission
	}
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
	if employee.Status == "Pending" {
		return "", fmt.Errorf("Employee is not Approved %w", err)
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
			return "", common.ErrBadParameter
		}
		return "Leave Table for the Employee is set", nil

	}
	return "Leave Already Set", err
}
func (s RepoService) DeleteLeaves(ctx context.Context,id int) (string, error) {
	_,err:=s.controller.CheckIfAdmin(ctx)
	if err!=nil{
		return "",common.ErrNoPermission
	}
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

func (s RepoService) EnterLeaves(ctx context.Context,id int, leave Leaves) ([]Leaves, error) {
	_,err:=s.controller.CheckIfAdmin(ctx)
	if err!=nil{
		return nil,common.ErrNoPermission
	}
	var leaves []Leaves
	db, err := dc.GetDB()
	if err != nil {
		return nil, err
	}
	err = db.Where("employee_id=?", id).Updates(&leave).Error
	if err != nil {
		return nil, err
	}
	db.Find(&leaves)
	return leaves, nil
}
func (s RepoService) GetLeaves(ctx context.Context) ([]Leaves, error) {
	var leaves []Leaves
	db, err := dc.GetDB()
	if err != nil {
		return nil, err
	}
	if err := db.Find(&leaves).Error; err != nil {
		return nil, fmt.Errorf("Could not get Leaves: %w", err)
	}
	return leaves, nil
}
func (s RepoService) GetLeavesById(ctx context.Context,id int) (Leaves, error) {
	var leaves Leaves
	db, err := dc.GetDB()
	if err != nil {
		return leaves, fmt.Errorf("Could not connect to db")
	}
	if err := db.Where("employee_id=?", id).First(&leaves).Error; err != nil {
		return leaves, fmt.Errorf("Could not find records")
	}
	return leaves, nil
}

/* ---------------------------- Request Functions --------------------------- */

func (s RepoService) PostLeaveRequest(ctx context.Context,request Requests) (string, error) {
	_,err:=s.controller.CheckIfEmployee(ctx)
	if err!=nil{
		return "",common.ErrNoPermission
	}
	var employee Employees
	var leave Leaves
	db, err := dc.GetDB()
	db, err = dc.GetDB()
	if err != nil {
		return "", err
	}
	if err = db.Table("employees").Where("id=?", request.EmployeeID).Scan(&employee).Error; err != nil {
		return "No such employee", nil

	}
	if err = db.Where("leave_id=?", request.LeaveID).Find(&leave).Error; err != nil {
		return "Employee Leave Not Set", nil
	}
	request.Status = "Pending"
	if leave.MedicalLeave > leave.MedicalTaken && request.Type == "Medical" {
		if err = db.Create(&request).Error; err != nil {
			return "", err
		}
		// leave.MedicalTaken = leave.MedicalTaken + request.Days
		// err=db.Updates(&leave).Error
		// if err!= nil {
		// 	return "Could Not Update Leave", err
		// }
		return "medical Request Entered", nil
	}
	if leave.AnnualLeave > leave.AnnualTaken && request.Type == "Annual" {
		if err = db.Create(&request).Error; err != nil {
			return "", err
		}
		// leave.AnnualTaken = leave.AnnualTaken + request.Days
		// err=db.Where("employee_id=?",request.EmployeeID).Updates(&leave).Error
		// if err!= nil {
		// 	return "Could Not Update Leave", err
		// }
		return "Annual Request Entered", nil
	}
	if leave.EmergencyLeave > leave.EmergencyTaken && request.Type == "Emergency" {
		if err = db.Create(&request).Error; err != nil {
			return "", err
		}
		// leave.EmergencyTaken = leave.EmergencyTaken +  request.Days
		// err=db.Updates(&leave).Error
		// if err!= nil {
		// 	return "Could Not Update Leave", err
		// }
		return "Emergency Request Entered", nil
	}
	return "bla", nil
}
func (s RepoService) GetLeaveRequest(ctx context.Context) ([]Requests, error) {
	_,err:=s.controller.CheckIfAdminOrManager(ctx)
	if err!=nil{
		return nil,common.ErrNoPermission
	}
	db, err := dc.GetDB()
	var request []Requests
	if err != nil {
		return nil, err
	}
	if err = db.Find(&request).Error; err != nil {
		return nil, err
	}
	return request, nil
}

// todo check json tags
func (s RepoService) ApproveLeaveRequest(ctx context.Context,id int) (string, error) {
	_,err:=s.controller.CheckIfAdminOrManager(ctx)
	if err!=nil{
		return "",common.ErrNoPermission
	}
	db, err := dc.GetDB()
	if err != nil {
		return "", err
	}

	var request Requests
	if err := db.Where("request_id = ?", id).First(&request).Error; err != nil {
		return "", fmt.Errorf("Could not find the request: %w", err)
	}

	if request.Status == "Approved" {
		return "", errors.New("The request is already approved")
	}

	var leave Leaves
	if err := db.Where("leave_id = ?", request.LeaveID).First(&leave).Error; err != nil {
		return "", fmt.Errorf("Employee leave balance not found: %w", err)
	}

	var leaveTaken int
	switch request.Type {
	case "Emergency":
		leaveTaken = leave.EmergencyTaken
		if leave.EmergencyLeave-leaveTaken < request.Days {
			return "", errors.New("Requested leave is more than available")
		}
		leave.EmergencyTaken += request.Days
	case "Annual":
		leaveTaken = leave.AnnualTaken
		if leave.AnnualLeave-leaveTaken < request.Days {
			return "", errors.New("Requested leave is more than available")
		}
		leave.AnnualTaken += request.Days
	case "Medical":
		leaveTaken = leave.MedicalTaken
		if leave.MedicalLeave-leaveTaken < request.Days {
			return "", errors.New("Requested leave is more than available")
		}
		leave.MedicalTaken += request.Days
	default:
		return "", errors.New("Invalid leave type")
	}

	if err := db.Model(&leave).Where("employee_id = ?", request.EmployeeID).Updates(&leave).Error; err != nil {
		return "", fmt.Errorf("Could not update leave: %w", err)
	}

	request.Status = "Approved"
	if err := db.Model(&request).Where("request_id = ?", id).Updates(&request).Error; err != nil {
		return "", fmt.Errorf("Could not approve request: %w", err)
	}

	return "Successfully Leave Balance Updated", nil
}

// TODO Reject Request
func (s RepoService) DeleteLeaveRequest(ctx context.Context,id int) (string, error) {
	_,err:=s.controller.CheckIfAdminOrManager(ctx)
	if err!=nil{
		return "",common.ErrNoPermission
	}
	db, err := dc.GetDB()
	var request Requests
	if err != nil {
		return "", fmt.Errorf("Could not Connect to Database: %w", err)
	}
	if err := db.Model(&request).Where("request_id=?", id).Delete(&request).Error; err != nil {
		return "", fmt.Errorf("Could not Delete Request: %w", err)
	}
	return "Request Successfully Deleted", nil
}
