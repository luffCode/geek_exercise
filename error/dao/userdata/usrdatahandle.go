package userdata

import (
	"database/sql"
	"github.com/pkg/errors"
)

var ErrQueryNilError = errors.New("query is nil")

type DBDataHandler struct {
}

//我的理解是调用数据库驱动，直接返回低层代码的错误信息
//GetUserDataErr 查询用户数据
func (d *DBDataHandler) GetUserDataErr(query interface{}) (interface{}, error) {
	if query == nil {
		return nil, ErrQueryNilError
	}

	//模拟查询.....
	err := sql.ErrNoRows
	if err == sql.ErrNoRows {
		return nil, err
	}

	return "ok", err
}
