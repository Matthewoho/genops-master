package permissions

import (
	"context"

	"genops-master/internal/svc"
	"genops-master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetachPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetachPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetachPermissionLogic {
	return &DetachPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetachPermissionLogic) DetachPermission(req *types.DetachPermissionReq) (resp *types.Result, err error) {
	// todo: add your logic here and delete this line

	return
}
