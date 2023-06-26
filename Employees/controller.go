package Employee

import (
	"context"
	"errors"
	"lms/common"
)

type RepoService struct {
	controller Controller
}
type RepoController struct {
}

// NewController New Service
func NewController() Controller {
	return &RepoController{}
}

/*
	idh ahn sadhanam used to identify the user type and passed within the

service but here we use aa special file called conmtroller which is accessed before the service
to check the user type and set acces but the working is not correct and need fixes  major fixes
*/
type Controller interface {
	CheckIfAdmin(ctx context.Context) (int, error)
	CheckIfManager(ctx context.Context) (int, error)
	CheckIfEmployee(ctx context.Context) (int, error)
	CheckIfEmployeeOrManager(ctx context.Context) (int, error)
	CheckIfAdminOrManager(ctx context.Context) (int, error)
}

func (r *RepoController) CheckIfAdmin(ctx context.Context) (int, error) {
	usertype := ctx.Value("userType")
	utype, ok := usertype.(float64)
	user := int(utype)
	if !ok {
		return 0, errors.New("cannot Decode type")
	}
	switch user {
	case 1:
		return 1, nil
	default:
		return 0, common.ErrUnauthorized
	}
}

func (r *RepoController) CheckIfManager(ctx context.Context) (int, error) {
	usertype := ctx.Value("userType")
	utype, ok := usertype.(float64)
	user:=int(utype)
	if !ok {
		return 0, errors.New("cannot decode typ")
	}
	switch user {
	case 2:
		return 2, nil
	default:
		return 0, common.ErrUnauthorized
	}
}

func (r *RepoController) CheckIfEmployee(ctx context.Context) (int, error) {
	usertype := ctx.Value("userType")
	utype, ok := usertype.(float64)
	user:=int(utype)
	if !ok {
		return 0, errors.New("cannot decode typ")
	}
	switch user {
	case 3:
		return 3, nil
	default:
		return 0, common.ErrUnauthorized
	}
}

func (r *RepoController) CheckIfEmployeeOrManager(ctx context.Context) (int, error) {
	usertype := ctx.Value("userType")
	utype, ok := usertype.(float64)
	user:=int(utype)
	if !ok {
		return 0, errors.New("cannot decode typ")
	}
	switch user {
	case 2:
		return 2,nil
	case 3:
		return 3, nil
	default:
		return 0, common.ErrUnauthorized
	}
}

func (r *RepoController) CheckIfAdminOrManager(ctx context.Context) (int, error) {
	usertype := ctx.Value("userType")
	utype, ok := usertype.(float64)
	user:=int(utype)
	if !ok {
		return 0, errors.New("cannot decode type")
	}
	switch user {
	case 1:
		return 1,nil
	case 3:
		return 3, nil
	default:
		return 0, common.ErrUnauthorized
	}
}

