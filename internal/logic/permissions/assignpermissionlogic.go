package permissions

import (
	"context"

	"genops-master/internal/svc"
	"genops-master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssignPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignPermissionLogic {
	return &AssignPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssignPermissionLogic) AssignPermission(req *types.AssignPermissionReq) (resp *types.Result, err error) {
	// todo: add your logic here and delete this line

	return
}
