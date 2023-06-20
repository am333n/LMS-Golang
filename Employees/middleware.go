package Employee

import (
	"time"

	"github.com/go-kit/kit/log"
)

//write a middleware in go kit to authenticate the request

type Middleware func(Service) Service

// LoggingMiddleware for auth verification
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

func (m loggingMiddleware) PostEmployee(employee Employees) (Employees, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "PostEmployee", "ID", employee.ID, "took", time.Since(begin))
	}(time.Now())
	return m.next.PostEmployee(employee)
}
func (m loggingMiddleware) GetEmployees() ([]Employees, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "GetEmployees",  "took", time.Since(begin))
	}(time.Now())
	return m.next.GetEmployees()
}
func (m loggingMiddleware) GetEmployeeById(id int) (Employees, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "GetEmployees",  "took", time.Since(begin))
	}(time.Now())
	return m.next.GetEmployeeById(id) 
}
func (m loggingMiddleware) DeleteEmployeeById(id int) (string, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "GetEmployees",  "took", time.Since(begin))
	}(time.Now())
	return m.next.DeleteEmployeeById(id) 
}
func (m loggingMiddleware) UpdateEmployee(id int, employee Employees) (string, Employees, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "GetEmployees",  "took", time.Since(begin))
	}(time.Now())
	return m.next.UpdateEmployee(id , employee)
}
func (m loggingMiddleware) PostLeaves(id int) (string, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "GetEmployees",  "took", time.Since(begin))
	}(time.Now())
	return m.next.PostLeaves(id)
}
func (m loggingMiddleware) DeleteLeaves(id int) (string, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "GetEmployees",  "took", time.Since(begin))
	}(time.Now())
	return m.next.DeleteLeaves(id)
}
func (m loggingMiddleware) EnterLeaves(id int, leave Leaves) ([]Leaves, error){

	defer func(begin time.Time) {
		m.logger.Log("method", "GetEmployees",  "took", time.Since(begin))
	}(time.Now())
	return m.next.EnterLeaves(id, leave )
}
func (m loggingMiddleware) ApproveEmployee(id int) (string, error){

	defer func(begin time.Time) {
		m.logger.Log("method", "GetEmployees",  "took", time.Since(begin))
	}(time.Now())
	return m.next.ApproveEmployee(id)
}
func (m loggingMiddleware)PostLeaveRequest(request Requests) (string, error){

	defer func(begin time.Time) {
		m.logger.Log("method", "GetEmployees",  "took", time.Since(begin))
	}(time.Now())
	return m.next.PostLeaveRequest(request) 
}
func (m loggingMiddleware)GetLeaveRequest() ([]Requests, error){

	defer func(begin time.Time) {
		m.logger.Log("method", "GetEmployees",  "took", time.Since(begin))
	}(time.Now())
	return m.next.GetLeaveRequest()
}
func (m loggingMiddleware)ApproveLeaveRequest(id int) (string, error){

	defer func(begin time.Time) {
		m.logger.Log("method", "GetEmployees",  "took", time.Since(begin))
	}(time.Now())
	return m.next.ApproveLeaveRequest(id) 
}
func (m loggingMiddleware)DeleteLeaveRequest(id int) (string, error){

	defer func(begin time.Time) {
		m.logger.Log("method", "GetEmployees",  "took", time.Since(begin))
	}(time.Now())
	return m.next.DeleteLeaveRequest(id)
}
func (m loggingMiddleware)GetLeaves()([]Leaves,error){

	defer func(begin time.Time) {
		m.logger.Log("method", "GetEmployees",  "took", time.Since(begin))
	}(time.Now())
	return m.next.GetLeaves()
}
func (m loggingMiddleware)GetLeavesById(id int)(Leaves,error){

	defer func(begin time.Time) {
		m.logger.Log("method", "GetEmployees",  "took", time.Since(begin))
	}(time.Now())
	return m.next.GetLeavesById(id)
}