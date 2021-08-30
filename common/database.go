package common

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB

func InitDb() {
	// dbDriver := viper.GetString("db.dbDriver")
	dbName := viper.GetString("db.dbName")
	dbUser := viper.GetString("db.dbUser")
	dbPassword := viper.GetString("db.dbPassword")
	dbHost := viper.GetString("db.dbHost")
	dbPort := viper.GetString("db.dbPort")
	charset := viper.GetString("db.charset")
	maxConn := viper.GetInt("db.maxConn") // 最大连接数
	maxOpen := viper.GetInt("db.maxOpen") // 最大空闲数
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
		charset,
	)
	fmt.Println(args)
	var err error
	_db, err = gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败： " + err.Error())
	}

	// _db.Callback().Create().Replace("gorm:create_time", CreateTimeCallBack)

	sqlDB, err := _db.DB()
	if err != nil {
		panic("数据错误： " + err.Error())
	}
	sqlDB.SetMaxIdleConns(maxOpen) // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(maxConn) // 设置打开数据库连接的最大数量。
	fmt.Println("数据库连接成功!")
}

func GetDB() *gorm.DB {
	return _db
}
