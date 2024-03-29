// Code generated by goctl. DO NOT EDIT.

package favorites

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
	favoritesFieldNames          = builder.RawFieldNames(&Favorites{})
	favoritesRows                = strings.Join(favoritesFieldNames, ",")
	favoritesRowsExpectAutoSet   = strings.Join(stringx.Remove(favoritesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	favoritesRowsWithPlaceHolder = strings.Join(stringx.Remove(favoritesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	favoritesModel interface {
		Insert(ctx context.Context, data *Favorites) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Favorites, error)
		Update(ctx context.Context, data *Favorites) error
		Delete(ctx context.Context, id int64) error
	}

	defaultFavoritesModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Favorites struct {
		Id       int64         `db:"id"`        // 自增主键
		CreateAt time.Time     `db:"create_at"` // 创建时间
		UserId   sql.NullInt64 `db:"user_id"`   // 点赞用户ID
		VideoId  int64         `db:"video_id"`  // 被点赞视频ID
	}
)

func newFavoritesModel(conn sqlx.SqlConn) *defaultFavoritesModel {
	return &defaultFavoritesModel{
		conn:  conn,
		table: "`favorites`",
	}
}

func (m *defaultFavoritesModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultFavoritesModel) FindOne(ctx context.Context, id int64) (*Favorites, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", favoritesRows, m.table)
	var resp Favorites
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

func (m *defaultFavoritesModel) Insert(ctx context.Context, data *Favorites) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, favoritesRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.VideoId)
	return ret, err
}

func (m *defaultFavoritesModel) Update(ctx context.Context, data *Favorites) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, favoritesRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.VideoId, data.Id)
	return err
}

func (m *defaultFavoritesModel) tableName() string {
	return m.table
}
