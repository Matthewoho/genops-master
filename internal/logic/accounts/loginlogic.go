package accounts

import (
	"context"

	"genops-master/internal/biz"
	"genops-master/internal/svc"
	"genops-master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp interface{}, err error) {
	// todo: add your logic here and delete this line

	// 1. 初始化用户模型
	userModel := svc.NewUserService(l.svcCtx.MySQLClient)

	// 2. 检查用户是否存在
	exists, err := userModel.ExistsUser(l.ctx, req.Username)
	if err != nil {
		return nil, biz.DBERROR // 数据库查询错误
	}
	if !exists {
		return nil, biz.InvalidUser // 用户不存在
	}

	// 3. 校验用户名和密码
	verify, err := userModel.VerifyPassword(l.ctx, req.Username, req.Password)
	if err != nil {
		return nil, biz.DBERROR // 数据库查询错误
	}
	if !verify {
		return nil, biz.InvalidUser // 用户名或密码错误
	}

	// 4. 生成token
	tokens, err := svc.GenerateTokens(l.svcCtx, req.Username)
	if err != nil {
		return nil, biz.GenerateTokenError // 生成token错误
	}

	// 5. 获取用户信息
	user, err := userModel.GetUserInfoByUsername(l.ctx, req.Username)
	if err != nil {
		return nil, biz.DBERROR // 数据库查询错误
	}

	// 6. 封装返回结果
	resp = map[string]interface{}{
		"tokens":   tokens,
		"userinfo": user,
	}

	return resp, nil
}
