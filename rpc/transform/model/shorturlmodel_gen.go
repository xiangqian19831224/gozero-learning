// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.5

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
	shorturlRowsExpectAutoSet   = strings.Join(stringx.Remove(shorturlFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	shorturlRowsWithPlaceHolder = strings.Join(stringx.Remove(shorturlFieldNames, "`shorten`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheShorturlShortenPrefix = "cache:shorturl:shorten:"
)

type (
	// 定义了短网址模型需要实现的方法，包括插入、查询、更新和删除操作
	shorturlModel interface {
		Insert(ctx context.Context, data *Shorturl) (sql.Result, error)
		FindOne(ctx context.Context, shorten string) (*Shorturl, error)
		Update(ctx context.Context, data *Shorturl) error
		Delete(ctx context.Context, shorten string) error
	}

	// defaultShortUrlMode是shorturlModel的实现类：
	// 1. 在Go语言中，接口是一种类型，它定义了一组方法签名，但不实现这些方法。
	//    任何实现了接口中所有方法的类型都隐式地实现了该接口
	// 2. shorturlmodel_gen中defaultShortUrlModel实现了shorturlModel的方法
	defaultShorturlModel struct {
		sqlc.CachedConn
		table string
	}

	Shorturl struct {
		Shorten string `db:"shorten"` // shorten key
		Url     string `db:"url"`     // original url
	}
)

func newShorturlModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultShorturlModel {
	return &defaultShorturlModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`shorturl`",
	}
}

func (m *defaultShorturlModel) Delete(ctx context.Context, shorten string) error {
	shorturlShortenKey := fmt.Sprintf("%s%v", cacheShorturlShortenPrefix, shorten)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `shorten` = ?", m.table)
		return conn.ExecCtx(ctx, query, shorten)
	}, shorturlShortenKey)
	return err
}

func (m *defaultShorturlModel) FindOne(ctx context.Context, shorten string) (*Shorturl, error) {
	shorturlShortenKey := fmt.Sprintf("%s%v", cacheShorturlShortenPrefix, shorten)
	var resp Shorturl
	err := m.QueryRowCtx(ctx, &resp, shorturlShortenKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `shorten` = ? limit 1", shorturlRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, shorten)
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
	shorturlShortenKey := fmt.Sprintf("%s%v", cacheShorturlShortenPrefix, data.Shorten)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, shorturlRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Shorten, data.Url)
	}, shorturlShortenKey)
	return ret, err
}

func (m *defaultShorturlModel) Update(ctx context.Context, data *Shorturl) error {
	shorturlShortenKey := fmt.Sprintf("%s%v", cacheShorturlShortenPrefix, data.Shorten)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `shorten` = ?", m.table, shorturlRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Url, data.Shorten)
	}, shorturlShortenKey)
	return err
}

func (m *defaultShorturlModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheShorturlShortenPrefix, primary)
}

func (m *defaultShorturlModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `shorten` = ? limit 1", shorturlRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultShorturlModel) tableName() string {
	return m.table
}
