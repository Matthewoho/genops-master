package models

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ServiceFunctionModel = (*customServiceFunctionModel)(nil)

type (
	// ServiceFunctionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customServiceFunctionModel.
	ServiceFunctionModel interface {
		serviceFunctionModel
		withSession(session sqlx.Session) ServiceFunctionModel
	}

	customServiceFunctionModel struct {
		*defaultServiceFunctionModel
	}
)

// NewServiceFunctionModel returns a model for the database table.
func NewServiceFunctionModel(conn sqlx.SqlConn) ServiceFunctionModel {
	return &customServiceFunctionModel{
		defaultServiceFunctionModel: newServiceFunctionModel(conn),
	}
}

func (m *customServiceFunctionModel) withSession(session sqlx.Session) ServiceFunctionModel {
	return NewServiceFunctionModel(sqlx.NewSqlConnFromSession(session))
}
