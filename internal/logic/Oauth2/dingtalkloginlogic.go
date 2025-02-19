package Oauth2

import (
	"context"

	"genops-master/internal/svc"
	"genops-master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DingTalkLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDingTalkLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DingTalkLoginLogic {
	return &DingTalkLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DingTalkLoginLogic) DingTalkLogin(req *types.DingTalkLoginReq) (resp *types.Result, err error) {
	// todo: add your logic here and delete this line

	return
}
