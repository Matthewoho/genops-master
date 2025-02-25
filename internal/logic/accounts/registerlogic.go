package accounts

import (
	"context"

	"genops-master/internal/biz"
	"genops-master/internal/models"
	"genops-master/internal/svc"
	"genops-master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp bool, err error) {
	// todo: add your logic here and delete this line

	// 1. 初始化用户模型
	userModel := svc.NewUserService(l.svcCtx.MySQLClient)

	// 2. 检查用户是否存在
	exists, err := userModel.ExistsUser(l.ctx, req.Username)
	if err != nil {
		return false, biz.DBERROR // 数据库查询错误
	}
	if exists {
		return false, biz.AlredyRegister // 用户已存在
	}

	// 3. 创建用户
	user := &models.USERS{
		USERNAME: req.Username,
		PASSWORD: req.Password,
	}

	err = userModel.CreateUser(l.ctx, user)
	if err != nil {
		return false, biz.DBERROR // 数据库插入错误
	}

	// 4. 返回成功
	return true, nil
}
