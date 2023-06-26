package Employee

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

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
func (m loggingMiddleware) GetEmployees(ctx context.Context) ([]Employees, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "GetEmployees", "took", time.Since(begin))
	}(time.Now())
	return m.next.GetEmployees(ctx)
}
func (m loggingMiddleware) GetEmployeeById(id int) (Employees, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "GetEmployeesById", "took", time.Since(begin))
	}(time.Now())
	return m.next.GetEmployeeById(id)
}
func (m loggingMiddleware) DeleteEmployeeById(ctx context.Context,id int) (string, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "DeleteEmployeesById", "took", time.Since(begin))
	}(time.Now())
	return m.next.DeleteEmployeeById(ctx ,id )
}
func (m loggingMiddleware) UpdateEmployee(ctx context.Context,id int, employee Employees) (string, Employees, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "UpdateEmployee", "took", time.Since(begin))
	}(time.Now())
	return m.next.UpdateEmployee(ctx ,id, employee)
}
func (m loggingMiddleware) PostLeaves(ctx context.Context,id int) (string, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "PostLeave", "took", time.Since(begin))
	}(time.Now())
	return m.next.PostLeaves(ctx,id)
}
func (m loggingMiddleware) DeleteLeaves(ctx context.Context,id int) (string, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "DeleteLeave", "took", time.Since(begin))
	}(time.Now())
	return m.next.DeleteLeaves(ctx,id)
}
func (m loggingMiddleware) EnterLeaves(ctx context.Context,id int, leave Leaves) ([]Leaves, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "EnterLeave", "took", time.Since(begin))
	}(time.Now())
	return m.next.EnterLeaves(ctx,id, leave)
}
func (m loggingMiddleware) ApproveEmployee(ctx context.Context,id int) (string, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "ApproveEmployee", "took", time.Since(begin))
	}(time.Now())
	return m.next.ApproveEmployee(ctx,id)
}
func (m loggingMiddleware) PostLeaveRequest(ctx context.Context,request Requests) (string, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "PostLEaveRequest", "took", time.Since(begin))
	}(time.Now())
	return m.next.PostLeaveRequest(ctx,request)
}
func (m loggingMiddleware) GetLeaveRequest(ctx context.Context) ([]Requests, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "GetLeaveReuest", "took", time.Since(begin))
	}(time.Now())
	return m.next.GetLeaveRequest(ctx)
}
func (m loggingMiddleware) ApproveLeaveRequest(ctx context.Context,id int) (string, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "ApproveLeaveRequest", "took", time.Since(begin))
	}(time.Now())
	return m.next.ApproveLeaveRequest(ctx,id)
}
func (m loggingMiddleware) DeleteLeaveRequest(ctx context.Context,id int) (string, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "DeleteLeaveRequest", "took", time.Since(begin))
	}(time.Now())
	return m.next.DeleteLeaveRequest(ctx,id)
}
func (m loggingMiddleware) GetLeaves(ctx context.Context) ([]Leaves, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "GetLeaves", "took", time.Since(begin))
	}(time.Now())
	return m.next.GetLeaves(ctx,)
}
func (m loggingMiddleware) GetLeavesById(ctx context.Context,id int) (Leaves, error) {

	defer func(begin time.Time) {
		m.logger.Log("method", "GetLeavesById", "took", time.Since(begin))
	}(time.Now())
	return m.next.GetLeavesById(ctx,id)
}
