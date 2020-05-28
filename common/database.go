package common

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	dbDriver := viper.GetString("db.dbDriver")
	dbName := viper.GetString("db.dbName")
	dbUser := viper.GetString("db.dbUser")
	dbPassword := viper.GetString("db.dbPassword")
	dbHost := viper.GetString("db.dbHost")
	dbPort := viper.GetString("db.dbPort")
	charset := viper.GetString("db.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
		charset) //

	db, err := gorm.Open(dbDriver, args)
	if err != nil {
		panic("faild error in database" + err.Error())
	}
	db.SingularTable(true)
	fmt.Println("数据库连接成功!")
	return db
}

func GetDB() *gorm.DB {
	return DB
}
