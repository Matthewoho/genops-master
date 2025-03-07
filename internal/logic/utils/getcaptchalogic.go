package utils

import (
	"context"

	biz "genops-master/internal/biz"
	"genops-master/internal/svc"
	"genops-master/internal/tools"
	"genops-master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha(ctx *svc.ServiceContext, req *types.GetCaptchaReq) (resp interface{}, err error) {
	// todo: add your logic here and delete this line

	// 1. 生成验证码
	captcha, err := tools.GenerateCaptcha(ctx.RedisClient)
	if err != nil {
		return nil, biz.GenerateCaptchaError
	}

	return captcha, nil
}
