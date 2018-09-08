package datastore

import (
	"fmt"
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysqlのドライバー
	"os"
)

// DB - コネクション
var DB = CreateConnection(GetDBEnv())

// GetDBEnv - DBの接続先を環境変数から取得
func GetDBEnv() interface{} {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	dbName := os.Getenv("MYSQL_DBNAME")
	dbIp := os.Getenv("MYSQL_IP")
	connect := user + ":" + pass + "@" + "tcp(" + dbIp + ")" + "/" + dbName
	return connect
}

// CreateConnection - コネクションの確立
func CreateConnection(connect interface{}) *gorm.DB {
	db, err := gorm.Open("mysql", connect)

	if err != nil {
		fmt.Println(connect)
		panic(err.Error())
	}
	db.AutoMigrate(&model.Credit{})
	db.AutoMigrate(&model.Receipt{})
	return db
}
