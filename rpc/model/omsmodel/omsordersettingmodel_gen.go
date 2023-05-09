// Code generated by goctl. DO NOT EDIT.

package omsmodel

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
	omsOrderSettingFieldNames          = builder.RawFieldNames(&OmsOrderSetting{})
	omsOrderSettingRows                = strings.Join(omsOrderSettingFieldNames, ",")
	omsOrderSettingRowsExpectAutoSet   = strings.Join(stringx.Remove(omsOrderSettingFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	omsOrderSettingRowsWithPlaceHolder = strings.Join(stringx.Remove(omsOrderSettingFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	omsOrderSettingModel interface {
		Insert(ctx context.Context, data *OmsOrderSetting) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*OmsOrderSetting, error)
		Update(ctx context.Context, data *OmsOrderSetting) error
		Delete(ctx context.Context, id int64) error
	}

	defaultOmsOrderSettingModel struct {
		conn  sqlx.SqlConn
		table string
	}

	OmsOrderSetting struct {
		Id                  int64 `db:"id"`
		FlashOrderOvertime  int64 `db:"flash_order_overtime"`  // 秒杀订单超时关闭时间(分)
		NormalOrderOvertime int64 `db:"normal_order_overtime"` // 正常订单超时时间(分)
		ConfirmOvertime     int64 `db:"confirm_overtime"`      // 发货后自动确认收货时间（天）
		FinishOvertime      int64 `db:"finish_overtime"`       // 自动完成交易时间，不能申请售后（天）
		CommentOvertime     int64 `db:"comment_overtime"`      // 订单完成后自动好评时间（天）
	}
)

func newOmsOrderSettingModel(conn sqlx.SqlConn) *defaultOmsOrderSettingModel {
	return &defaultOmsOrderSettingModel{
		conn:  conn,
		table: "`oms_order_setting`",
	}
}

func (m *defaultOmsOrderSettingModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultOmsOrderSettingModel) FindOne(ctx context.Context, id int64) (*OmsOrderSetting, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", omsOrderSettingRows, m.table)
	var resp OmsOrderSetting
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

func (m *defaultOmsOrderSettingModel) Insert(ctx context.Context, data *OmsOrderSetting) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, omsOrderSettingRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.FlashOrderOvertime, data.NormalOrderOvertime, data.ConfirmOvertime, data.FinishOvertime, data.CommentOvertime)
	return ret, err
}

func (m *defaultOmsOrderSettingModel) Update(ctx context.Context, data *OmsOrderSetting) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, omsOrderSettingRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.FlashOrderOvertime, data.NormalOrderOvertime, data.ConfirmOvertime, data.FinishOvertime, data.CommentOvertime, data.Id)
	return err
}

func (m *defaultOmsOrderSettingModel) tableName() string {
	return m.table
}
