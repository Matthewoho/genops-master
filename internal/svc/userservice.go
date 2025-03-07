package svc

import (
	"context" // 导入 context 包
	"database/sql"
	"fmt"
	"genops-master/internal/models"
	"genops-master/internal/tools"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UserService struct {
	userModel models.USERSModel
}

func NewUserService(conn sqlx.SqlConn) *UserService {
	return &UserService{
		userModel: models.NewUSERSModel(conn),
	}
}

// 获取用户密码盐
func (s *UserService) GetUserSalt(ctx context.Context, username string) (string, error) {
	user, err := s.userModel.FindOneByUSERNAME(ctx, username)
	if err != nil {
		return "", err
	}
	if !user.SALT.Valid {
		return "", sql.ErrNoRows
	}
	return user.SALT.String, nil
}

// 检查用户是否存在
func (s *UserService) ExistsUser(ctx context.Context, username string) (bool, error) {
	_, err := s.userModel.FindOneByUSERNAME(ctx, username) // 传递上下文
	if err == models.ErrNotFound {
		return false, nil // 用户不存在
	}
	if err != nil {
		return false, err // 其他错误
	}
	return true, nil // 用户存在
}

// 创建用户
func (s *UserService) CreateUser(ctx context.Context, user *models.USERS) error {
	// 生成盐
	salt, err := tools.GenerateSalt()
	if err != nil {
		return err
	}

	// 生成哈希密码
	hashedPassword, err := tools.HashPasswordWithSalt(user.PASSWORD, salt)
	if err != nil {
		return err
	}

	// 更新用户密码和盐
	user.PASSWORD = hashedPassword
	user.SALT = sql.NullString{String: salt, Valid: true} // 正确赋值

	// 插入数据库
	_, err = s.userModel.Insert(ctx, user)
	return err
}

// 校验用户名和密码
func (s *UserService) VerifyPassword(ctx context.Context, username, password string) (bool, error) {
	user, err := s.userModel.FindOneByUSERNAME(ctx, username)
	if err != nil {
		return false, fmt.Errorf("error finding user: %w", err) // 处理可能的错误情况
	}
	if !user.SALT.Valid {
		return false, fmt.Errorf("user salt is not valid") // 处理可能的错误情况
	}
	hashedPassword, err := tools.HashPasswordWithSalt(password, user.SALT.String)
	if err != nil {
		return false, fmt.Errorf("error hashing password: %w", err) // 处理可能的错误情况
	}

	// 比较哈希密码是否匹配
	if tools.CheckPassword(hashedPassword, user.PASSWORD, user.SALT.String) {
		return false, fmt.Errorf("密码不匹配")
	}

	return true, nil
}

// 获取用户信息
func (s *UserService) GetUserInfoByUsername(ctx context.Context, username string) (*models.USERS, error) {
	user, err := s.userModel.FindOneByUSERNAME(ctx, username)
	if err != nil {
		return nil, err
	}
	// 清除敏感信息
	user.PASSWORD = ""
	user.SALT = sql.NullString{String: "", Valid: false}

	return user, nil
}
