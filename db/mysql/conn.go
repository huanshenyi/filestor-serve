package mysql

import (
	"database/sql"
	"errors"
	"filestor-serve/db/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"os"
)

var db *sql.DB

var DbEngine *xorm.Engine

func init() {
	// テーブルの作成
	//CreatTable()
	db, _ = sql.Open("mysql", "root:root@tcp(192.168.99.100:3306)/fileserver?charset=utf8")
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql, err:" + err.Error())
		os.Exit(1)
	}

}

// DBConn:データベースのコネクトを返す
func DBConn() *sql.DB {
	return db
}

// テーブルの初期化用(windowsのdockerがゴミ)
func CreatTable(){
	driverName := "mysql"
	DsName := "root:root@tcp(192.168.99.100:3306)/fileserver?charset=utf8"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName,DsName)
	if err != nil && err.Error() != ""{
		log.Fatal(err.Error())
	}
	defer DbEngine.Close()
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(1)
	DbEngine.Sync2(new(model.TblFile))
	fmt.Println("init data base ok")
}