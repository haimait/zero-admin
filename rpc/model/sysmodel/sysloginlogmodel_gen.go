// Code generated by goctl. DO NOT EDIT!

package sysmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysLoginLogFieldNames          = builder.RawFieldNames(&SysLoginLog{})
	sysLoginLogRows                = strings.Join(sysLoginLogFieldNames, ",")
	sysLoginLogRowsExpectAutoSet   = strings.Join(stringx.Remove(sysLoginLogFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	sysLoginLogRowsWithPlaceHolder = strings.Join(stringx.Remove(sysLoginLogFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	sysLoginLogModel interface {
		Insert(ctx context.Context, data *SysLoginLog) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysLoginLog, error)
		Update(ctx context.Context, data *SysLoginLog) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysLoginLogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysLoginLog struct {
		Id             int64          `db:"id"`               // 编号
		UserName       sql.NullString `db:"user_name"`        // 用户名
		Status         sql.NullString `db:"status"`           // 登录状态（online:在线，登录初始状态，方便统计在线人数；login:退出登录后将online置为login；logout:退出登录）
		Ip             sql.NullString `db:"ip"`               // IP地址
		CreateBy       sql.NullString `db:"create_by"`        // 创建人
		CreateTime     time.Time      `db:"create_time"`      // 创建时间
		LastUpdateBy   sql.NullString `db:"last_update_by"`   // 更新人
		LastUpdateTime sql.NullTime   `db:"last_update_time"` // 更新时间
	}
)

func newSysLoginLogModel(conn sqlx.SqlConn) *defaultSysLoginLogModel {
	return &defaultSysLoginLogModel{
		conn:  conn,
		table: "`sys_login_log`",
	}
}

func (m *defaultSysLoginLogModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysLoginLogModel) FindOne(ctx context.Context, id int64) (*SysLoginLog, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysLoginLogRows, m.table)
	var resp SysLoginLog
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysLoginLogModel) Insert(ctx context.Context, data *SysLoginLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, sysLoginLogRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserName, data.Status, data.Ip, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime)
	return ret, err
}

func (m *defaultSysLoginLogModel) Update(ctx context.Context, data *SysLoginLog) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysLoginLogRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserName, data.Status, data.Ip, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime, data.Id)
	return err
}

func (m *defaultSysLoginLogModel) tableName() string {
	return m.table
}
