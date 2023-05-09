// Code generated by goctl. DO NOT EDIT.

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
	sysDictFieldNames          = builder.RawFieldNames(&SysDict{})
	sysDictRows                = strings.Join(sysDictFieldNames, ",")
	sysDictRowsExpectAutoSet   = strings.Join(stringx.Remove(sysDictFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	sysDictRowsWithPlaceHolder = strings.Join(stringx.Remove(sysDictFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	sysDictModel interface {
		Insert(ctx context.Context, data *SysDict) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysDict, error)
		Update(ctx context.Context, data *SysDict) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysDictModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysDict struct {
		Id          int64          `db:"id"`          // 编号
		Value       string         `db:"value"`       // 数据值
		Label       string         `db:"label"`       // 标签名
		Type        string         `db:"type"`        // 类型
		Description string         `db:"description"` // 描述
		Sort        float64        `db:"sort"`        // 排序（升序）
		CreateBy    string         `db:"create_by"`   // 创建人
		CreateTime  time.Time      `db:"create_time"` // 创建时间
		UpdateBy    sql.NullString `db:"update_by"`   // 更新人
		UpdateTime  time.Time      `db:"update_time"` // 更新时间
		Remarks     sql.NullString `db:"remarks"`     // 备注信息
		DelFlag     int64          `db:"del_flag"`    // 是否删除  -1：已删除  0：正常
	}
)

func newSysDictModel(conn sqlx.SqlConn) *defaultSysDictModel {
	return &defaultSysDictModel{
		conn:  conn,
		table: "`sys_dict`",
	}
}

func (m *defaultSysDictModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysDictModel) FindOne(ctx context.Context, id int64) (*SysDict, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysDictRows, m.table)
	var resp SysDict
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

func (m *defaultSysDictModel) Insert(ctx context.Context, data *SysDict) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysDictRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Value, data.Label, data.Type, data.Description, data.Sort, data.CreateBy, data.UpdateBy, data.Remarks, data.DelFlag)
	return ret, err
}

func (m *defaultSysDictModel) Update(ctx context.Context, data *SysDict) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysDictRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Value, data.Label, data.Type, data.Description, data.Sort, data.CreateBy, data.UpdateBy, data.Remarks, data.DelFlag, data.Id)
	return err
}

func (m *defaultSysDictModel) tableName() string {
	return m.table
}
