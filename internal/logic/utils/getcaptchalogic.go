package utils

import (
	"context"

	biz "genops-master/internal/biz"
	"genops-master/internal/svc"
	utils "genops-master/internal/svc/utils"
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

func (l *GetCaptchaLogic) GetCaptcha(ctx *svc.ServiceContext, req *types.GetCaptchaReq) (resp *types.Result, err error) {
	// todo: add your logic here and delete this line

	// 1. 生成验证码
	captcha, err := utils.GenerateCaptcha(ctx.RedisClient)
	if err != nil {
		return nil, biz.GenerateCaptchaError
	}

	// 2. 返回验证码
	resp = &types.Result{
		Code:    200,
		Message: "success",
		Data:    captcha,
	}

	return resp, nil
}
