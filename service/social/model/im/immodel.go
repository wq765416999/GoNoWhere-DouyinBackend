package im

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ImModel = (*customImModel)(nil)

type (
	// ImModel is an interface to be customized, add more methods here,
	// and implement the added methods in customImModel.
	ImModel interface {
		imModel
	}

	customImModel struct {
		*defaultImModel
	}
)

// NewImModel returns a model for the database table.
func NewImModel(conn sqlx.SqlConn) ImModel {
	return &customImModel{
		defaultImModel: newImModel(conn),
	}
}
