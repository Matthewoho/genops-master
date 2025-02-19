package models

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PermissionModel = (*customPermissionModel)(nil)

type (
	// PermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPermissionModel.
	PermissionModel interface {
		permissionModel
		withSession(session sqlx.Session) PermissionModel
	}

	customPermissionModel struct {
		*defaultPermissionModel
	}
)

// NewPermissionModel returns a model for the database table.
func NewPermissionModel(conn sqlx.SqlConn) PermissionModel {
	return &customPermissionModel{
		defaultPermissionModel: newPermissionModel(conn),
	}
}

func (m *customPermissionModel) withSession(session sqlx.Session) PermissionModel {
	return NewPermissionModel(sqlx.NewSqlConnFromSession(session))
}
