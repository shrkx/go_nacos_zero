// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	shorturlFieldNames          = builder.RawFieldNames(&Shorturl{})
	shorturlRows                = strings.Join(shorturlFieldNames, ",")
	shorturlRowsExpectAutoSet   = strings.Join(stringx.Remove(shorturlFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	shorturlRowsWithPlaceHolder = strings.Join(stringx.Remove(shorturlFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheShorturlIdPrefix = "cache:shorturl:id:"
)

type (
	shorturlModel interface {
		Insert(ctx context.Context, data *Shorturl) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Shorturl, error)
		Update(ctx context.Context, data *Shorturl) error
		Delete(ctx context.Context, id int64) error
		FindOneByShortUrl(ctx context.Context, shortUrl string) (*Shorturl, error) //新增方法
	}

	defaultShorturlModel struct {
		sqlc.CachedConn
		table string
	}

	Shorturl struct {
		Id       int64  `db:"id"`
		ShortUrl string `db:"short_url"` // 生成的key
		Url      string `db:"url"`       // 原始的url
	}
)

func newShorturlModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultShorturlModel {
	return &defaultShorturlModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`shorturl`",
	}
}

func (m *defaultShorturlModel) Delete(ctx context.Context, id int64) error {
	shorturlIdKey := fmt.Sprintf("%s%v", cacheShorturlIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, shorturlIdKey)
	return err
}

func (m *defaultShorturlModel) FindOne(ctx context.Context, id int64) (*Shorturl, error) {
	shorturlIdKey := fmt.Sprintf("%s%v", cacheShorturlIdPrefix, id)
	var resp Shorturl
	err := m.QueryRowCtx(ctx, &resp, shorturlIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", shorturlRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultShorturlModel) Insert(ctx context.Context, data *Shorturl) (sql.Result, error) {
	shorturlIdKey := fmt.Sprintf("%s%v", cacheShorturlIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, shorturlRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ShortUrl, data.Url)
	}, shorturlIdKey)
	return ret, err
}

func (m *defaultShorturlModel) Update(ctx context.Context, data *Shorturl) error {
	shorturlIdKey := fmt.Sprintf("%s%v", cacheShorturlIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, shorturlRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ShortUrl, data.Url, data.Id)
	}, shorturlIdKey)
	return err
}

func (m *defaultShorturlModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheShorturlIdPrefix, primary)
}

func (m *defaultShorturlModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", shorturlRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultShorturlModel) tableName() string {
	return m.table
}

// 新增方法
func (m *defaultShorturlModel) FindOneByShortUrl(ctx context.Context, shortUrl string) (*Shorturl, error) {
	shorturlIdKey := fmt.Sprintf("%s%v", cacheShorturlIdPrefix, shortUrl)
	var resp Shorturl
	err := m.QueryRowCtx(ctx, &resp, shorturlIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `short_url` = ? limit 1", shorturlRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, shortUrl)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
