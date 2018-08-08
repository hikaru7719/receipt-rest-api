package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

func GetConnection() *gorm.DB {

	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	dbname := ps.Getenv("MYSQL_DBNAME")
	connect := user + ":" + pass + "@" + "tcp(127.0.0.1:3306)" + "/" + dbname
	db, err := gorm.Open("mysql", connect)

	if err != nil {
		panic(err.Error())
	}

	return db
}
