package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysqlのドライバー
	"os"
)

var DB *gorm.DB

// GetDBEnv - DBの接続先を環境変数から取得
func GetDBEnv() interface{} {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	dbName := os.Getenv("MYSQL_DBNAME")
	connect := user + ":" + pass + "@" + "tcp(127.0.0.1:3306)" + "/" + dbName
	return connect
}

// GetConnection - コネクションの確立
func CreateConnection(connect interface{}) {
	db, err := gorm.Open("mysql", connect)

	if err != nil {
		panic(err.Error())
	}
	DB = db
}
