// Code generated by goctl. DO NOT EDIT!

package pmsmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	pmsProductCollectFieldNames          = builder.RawFieldNames(&PmsProductCollect{})
	pmsProductCollectRows                = strings.Join(pmsProductCollectFieldNames, ",")
	pmsProductCollectRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsProductCollectFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	pmsProductCollectRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsProductCollectFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	pmsProductCollectModel interface {
		Insert(ctx context.Context, data *PmsProductCollect) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PmsProductCollect, error)
		Update(ctx context.Context, data *PmsProductCollect) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPmsProductCollectModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsProductCollect struct {
		Id          int64        `db:"id"`
		UserId      int64        `db:"user_id"`      // 用户表的用户ID
		ValueId     int64        `db:"value_id"`     // 如果type=0，则是商品ID；如果type=1，则是专题ID
		CollectType int64        `db:"collect_type"` // 收藏类型，如果type=0，则是商品ID；如果type=1，则是专题ID
		AddTime     sql.NullTime `db:"add_time"`     // 创建时间
		UpdateTime  sql.NullTime `db:"update_time"`  // 更新时间
		Deleted     int64        `db:"deleted"`      // 逻辑删除
	}
)

func newPmsProductCollectModel(conn sqlx.SqlConn) *defaultPmsProductCollectModel {
	return &defaultPmsProductCollectModel{
		conn:  conn,
		table: "`pms_product_collect`",
	}
}

func (m *defaultPmsProductCollectModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPmsProductCollectModel) FindOne(ctx context.Context, id int64) (*PmsProductCollect, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pmsProductCollectRows, m.table)
	var resp PmsProductCollect
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

func (m *defaultPmsProductCollectModel) Insert(ctx context.Context, data *PmsProductCollect) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, pmsProductCollectRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.ValueId, data.CollectType, data.AddTime, data.Deleted)
	return ret, err
}

func (m *defaultPmsProductCollectModel) Update(ctx context.Context, data *PmsProductCollect) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pmsProductCollectRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.ValueId, data.CollectType, data.AddTime, data.Deleted, data.Id)
	return err
}

func (m *defaultPmsProductCollectModel) tableName() string {
	return m.table
}
