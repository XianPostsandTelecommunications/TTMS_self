package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	dao "mognolia/internal/dao/redis"
	"mognolia/internal/global"
	"mognolia/internal/model/automigrate"
	"os"
	"time"
)

func InitMysql() {
	m := global.Settings.Mysql
	fmt.Println(m)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.User, m.Password, m.Port, m.DbName)
	fmt.Println(dsn)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      true,
			LogLevel:      logger.Info,
		})
	//全局模式
	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	if err != nil {
		panic(err)
	}
	dao.Group.DB = DB
	DB.AutoMigrate(&automigrate.Tag{}, &automigrate.ManagerMovie{}, &automigrate.Cinema{}, &automigrate.Plan{}, &automigrate.Seat{}, &automigrate.Ticket{})

}
