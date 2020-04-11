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

// パスワード一致かの判断
func UserSignin(username string, encpwd string) bool {
	stmt, err := mysql.DBConn().Prepare("select * from tbl_user where user_name=? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	rows,err := stmt.Query(username)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}else if rows == nil {
		fmt.Println("username not found:" + username)
		return false
	}

	pRows := mysql.ParseRows(rows)
	if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == encpwd {
		return true
	}
	return false
}

//アクセスtokenの保存と更新
func UpdateToken(username string, token string) bool {
  stmt, err := mysql.DBConn().Prepare("replace into tbl_user_token (`user_name`, `user_token`) values (?, ?)")
  if err != nil {
  	fmt.Println(err.Error())
  	return false
  }
  defer stmt.Close()
  _, err = stmt.Exec(username, token)
	if err != nil {
		fmt.Println(err.Error())
		return false
  }
  return true
}