package utils

import (
	"context"

	"genops-master/internal/svc"
	"genops-master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysStatusLogic {
	return &SysStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysStatusLogic) SysStatus(req *types.SysStatusReq) (resp *types.Result, err error) {
	// todo: add your logic here and delete this line

	return
}
