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
	pmsProductFullReductionFieldNames          = builder.RawFieldNames(&PmsProductFullReduction{})
	pmsProductFullReductionRows                = strings.Join(pmsProductFullReductionFieldNames, ",")
	pmsProductFullReductionRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsProductFullReductionFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	pmsProductFullReductionRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsProductFullReductionFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	pmsProductFullReductionModel interface {
		Insert(ctx context.Context, data *PmsProductFullReduction) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PmsProductFullReduction, error)
		Update(ctx context.Context, data *PmsProductFullReduction) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPmsProductFullReductionModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsProductFullReduction struct {
		Id          int64           `db:"id"`
		ProductId   sql.NullInt64   `db:"product_id"`
		FullPrice   sql.NullFloat64 `db:"full_price"`
		ReducePrice sql.NullFloat64 `db:"reduce_price"`
	}
)

func newPmsProductFullReductionModel(conn sqlx.SqlConn) *defaultPmsProductFullReductionModel {
	return &defaultPmsProductFullReductionModel{
		conn:  conn,
		table: "`pms_product_full_reduction`",
	}
}

func (m *defaultPmsProductFullReductionModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPmsProductFullReductionModel) FindOne(ctx context.Context, id int64) (*PmsProductFullReduction, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pmsProductFullReductionRows, m.table)
	var resp PmsProductFullReduction
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

func (m *defaultPmsProductFullReductionModel) Insert(ctx context.Context, data *PmsProductFullReduction) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, pmsProductFullReductionRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ProductId, data.FullPrice, data.ReducePrice)
	return ret, err
}

func (m *defaultPmsProductFullReductionModel) Update(ctx context.Context, data *PmsProductFullReduction) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pmsProductFullReductionRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.ProductId, data.FullPrice, data.ReducePrice, data.Id)
	return err
}

func (m *defaultPmsProductFullReductionModel) tableName() string {
	return m.table
}
