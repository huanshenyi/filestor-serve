package db

import (
	"filestor-serve/db/mysql"
	"fmt"
)

// 新規ユーザーの追加
func UserSignup(username string, passwd string) bool {
	// ignoreがあるとindexが複数存在する場合、データの挿入は行わない
	stmt, err := mysql.DBConn().Prepare("insert ignore into tbl_user(`user_name`, `user_pwd`) value (?,?)")
	if err != nil{
		fmt.Println("Failed to insert, err:"+err.Error())
		return false
	}
	defer stmt.Close()
	ret, err := stmt.Exec(username, passwd)
	if err != nil{
		fmt.Println("Failed to insert, err:"+err.Error())
		return false
	}
	// ignore起動したかどうか
	if  rowsAffected, err := ret.RowsAffected();err==nil&&rowsAffected>0{
       return true
	}
	return false
}
