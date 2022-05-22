package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"myerror/dao/userdata"
)

//调用者获取数据错误，自己判断处理在进行Warp
func GetData(query interface{}) (interface{}, error) {
	userDao := userdata.DBDataHandler{}
	_, err := userDao.GetUserDataErr(query)
	if err != nil {
		if err == userdata.ErrQueryNilError {
			//调用这自己处理
			return nil, errors.Wrap(err, fmt.Sprintf("GetUserDataErr query = %s", query))
		}

		if err == sql.ErrNoRows {
			//调用这自己处理
			return nil, errors.Wrap(err, fmt.Sprintf("GetUserDataErr query = %s", query))
		}

		if err == sql.ErrConnDone {
			//调用这自己处理
			return nil, errors.Wrap(err, fmt.Sprintf("GetUserDataErr query = %s", query))
		}
	}

	return "", nil
}

func main() {

	query := "user = 123456"
	_, err := GetData(query)
	if err != nil {
		log.Println(err)
	}
}
