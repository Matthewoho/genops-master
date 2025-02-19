package roles

import (
	"context"

	"genops-master/internal/svc"
	"genops-master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetachRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetachRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetachRoleLogic {
	return &DetachRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetachRoleLogic) DetachRole(req *types.DetachRoleReq) (resp *types.Result, err error) {
	// todo: add your logic here and delete this line

	return
}
