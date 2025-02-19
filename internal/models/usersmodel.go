package models

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ USERSModel = (*customUSERSModel)(nil)

type (
	// USERSModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUSERSModel.
	USERSModel interface {
		uSERSModel
		withSession(session sqlx.Session) USERSModel
	}

	customUSERSModel struct {
		*defaultUSERSModel
	}
)

// NewUSERSModel returns a model for the database table.
func NewUSERSModel(conn sqlx.SqlConn) USERSModel {
	return &customUSERSModel{
		defaultUSERSModel: newUSERSModel(conn),
	}
}

func (m *customUSERSModel) withSession(session sqlx.Session) USERSModel {
	return NewUSERSModel(sqlx.NewSqlConnFromSession(session))
}
