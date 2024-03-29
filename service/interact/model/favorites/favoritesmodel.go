package favorites

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ FavoritesModel = (*customFavoritesModel)(nil)

type (
	// FavoritesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFavoritesModel.
	FavoritesModel interface {
		favoritesModel
	}

	customFavoritesModel struct {
		*defaultFavoritesModel
	}
)

// NewFavoritesModel returns a model for the database table.
func NewFavoritesModel(conn sqlx.SqlConn) FavoritesModel {
	return &customFavoritesModel{
		defaultFavoritesModel: newFavoritesModel(conn),
	}
}
