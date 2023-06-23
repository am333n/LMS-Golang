package Employee

import (
	"context"
	"lms/common"
)

type RepoService struct {
}
type RepoController struct{

}
/* idh ahn sadhanam used to identify the user type and passed within the 
service but here we use aa special file called conmtroller which is accessed before the service 
to check the user type and set acces but the working is not correct and need fixes  major fixes*/
type Controller interface {
	CheckUserType(ctx context.Context) (int, error)
}

func (r *RepoController) CheckUserType(ctx context.Context) (int, error) {
	usertype:=ctx.Value("usertype")
	switch usertype {
	case 1:
		return 1, nil
	default:
		return 0, common.ErrUnauthorized
	}
}



