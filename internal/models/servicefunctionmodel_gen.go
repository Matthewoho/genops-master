// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.6

package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	serviceFunctionFieldNames          = builder.RawFieldNames(&ServiceFunction{})
	serviceFunctionRows                = strings.Join(serviceFunctionFieldNames, ",")
	serviceFunctionRowsExpectAutoSet   = strings.Join(stringx.Remove(serviceFunctionFieldNames, "`ID`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	serviceFunctionRowsWithPlaceHolder = strings.Join(stringx.Remove(serviceFunctionFieldNames, "`ID`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	serviceFunctionModel interface {
		Insert(ctx context.Context, data *ServiceFunction) (sql.Result, error)
		FindOne(ctx context.Context, iD uint64) (*ServiceFunction, error)
		Update(ctx context.Context, data *ServiceFunction) error
		Delete(ctx context.Context, iD uint64) error
	}

	defaultServiceFunctionModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ServiceFunction struct {
		ID          uint64         `db:"ID"`          // 主键ID
		Name        string         `db:"Name"`        // 名称
		Description sql.NullString `db:"Description"` // 描述
		Permission  sql.NullString `db:"Permission"`  // 权限
		Visibility  int64          `db:"Visibility"`  // 是否可见
		CreatedAt   time.Time      `db:"CreatedAt"`   // 创建时间
		UpdatedAt   sql.NullTime   `db:"UpdatedAt"`   // 更新时间
	}
)

func newServiceFunctionModel(conn sqlx.SqlConn) *defaultServiceFunctionModel {
	return &defaultServiceFunctionModel{
		conn:  conn,
		table: "`ServiceFunction`",
	}
}

func (m *defaultServiceFunctionModel) Delete(ctx context.Context, iD uint64) error {
	query := fmt.Sprintf("delete from %s where `ID` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, iD)
	return err
}

func (m *defaultServiceFunctionModel) FindOne(ctx context.Context, iD uint64) (*ServiceFunction, error) {
	query := fmt.Sprintf("select %s from %s where `ID` = ? limit 1", serviceFunctionRows, m.table)
	var resp ServiceFunction
	err := m.conn.QueryRowCtx(ctx, &resp, query, iD)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultServiceFunctionModel) Insert(ctx context.Context, data *ServiceFunction) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, serviceFunctionRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.Description, data.Permission, data.Visibility, data.CreatedAt, data.UpdatedAt)
	return ret, err
}

func (m *defaultServiceFunctionModel) Update(ctx context.Context, data *ServiceFunction) error {
	query := fmt.Sprintf("update %s set %s where `ID` = ?", m.table, serviceFunctionRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Description, data.Permission, data.Visibility, data.CreatedAt, data.UpdatedAt, data.ID)
	return err
}

func (m *defaultServiceFunctionModel) tableName() string {
	return m.table
}
