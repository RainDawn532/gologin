package internal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var DB *gorm.DB
var err error

const admin = "root"
const Password = "root"
const DBName = "bootexplore"
const IP = "127.0.0.1"
const Port = "3306"

func InitDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), //output to stdout
		logger.Config{
			SlowThreshold:             time.Second, //slow sql
			LogLevel:                  logger.Info, //log level
			IgnoreRecordNotFoundError: true,        // ignore record not found Err
			Colorful:                  true,        //colorful
		},
	)
	// connect to mysql,no use config
	conn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", admin, Password,
		IP, Port, DBName)
	DB, err = gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //table name singular
		},
	})
	if err != nil {
		panic("connect_to_mysql_failed" + err.Error())
	}
	//自动生成数据库表  前提在model中配置了相关 grom:......   具体看相关
	//err = DB.AutoMigrate(&model.User{})
	//if err != nil {
	//	fmt.Println("auto_migrate_failed", err)
	//}
}
