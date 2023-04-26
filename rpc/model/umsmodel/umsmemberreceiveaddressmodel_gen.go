// Code generated by goctl. DO NOT EDIT!

package umsmodel

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
	umsMemberReceiveAddressFieldNames          = builder.RawFieldNames(&UmsMemberReceiveAddress{})
	umsMemberReceiveAddressRows                = strings.Join(umsMemberReceiveAddressFieldNames, ",")
	umsMemberReceiveAddressRowsExpectAutoSet   = strings.Join(stringx.Remove(umsMemberReceiveAddressFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	umsMemberReceiveAddressRowsWithPlaceHolder = strings.Join(stringx.Remove(umsMemberReceiveAddressFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	umsMemberReceiveAddressModel interface {
		Insert(ctx context.Context, data *UmsMemberReceiveAddress) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UmsMemberReceiveAddress, error)
		Update(ctx context.Context, data *UmsMemberReceiveAddress) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUmsMemberReceiveAddressModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UmsMemberReceiveAddress struct {
		Id            int64          `db:"id"`
		MemberId      sql.NullInt64  `db:"member_id"`
		Name          sql.NullString `db:"name"` // 收货人名称
		PhoneNumber   sql.NullString `db:"phone_number"`
		DefaultStatus sql.NullInt64  `db:"default_status"` // 是否为默认
		PostCode      sql.NullString `db:"post_code"`      // 邮政编码
		Province      sql.NullString `db:"province"`       // 省份/直辖市
		City          sql.NullString `db:"city"`           // 城市
		Region        sql.NullString `db:"region"`         // 区
		DetailAddress sql.NullString `db:"detail_address"` // 详细地址(街道)
	}
)

func newUmsMemberReceiveAddressModel(conn sqlx.SqlConn) *defaultUmsMemberReceiveAddressModel {
	return &defaultUmsMemberReceiveAddressModel{
		conn:  conn,
		table: "`ums_member_receive_address`",
	}
}

func (m *defaultUmsMemberReceiveAddressModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUmsMemberReceiveAddressModel) FindOne(ctx context.Context, id int64) (*UmsMemberReceiveAddress, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", umsMemberReceiveAddressRows, m.table)
	var resp UmsMemberReceiveAddress
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

func (m *defaultUmsMemberReceiveAddressModel) Insert(ctx context.Context, data *UmsMemberReceiveAddress) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, umsMemberReceiveAddressRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.MemberId, data.Name, data.PhoneNumber, data.DefaultStatus, data.PostCode, data.Province, data.City, data.Region, data.DetailAddress)
	return ret, err
}

func (m *defaultUmsMemberReceiveAddressModel) Update(ctx context.Context, data *UmsMemberReceiveAddress) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, umsMemberReceiveAddressRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.MemberId, data.Name, data.PhoneNumber, data.DefaultStatus, data.PostCode, data.Province, data.City, data.Region, data.DetailAddress, data.Id)
	return err
}

func (m *defaultUmsMemberReceiveAddressModel) tableName() string {
	return m.table
}
