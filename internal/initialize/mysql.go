package initialize

import (
	"fmt"
	"go-ecommerce-be/global"
	"go-ecommerce-be/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Database)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		//SkipDefaultTransaction: false,
	})
	if err != nil {
		checkErrorPanic(err, "mysql init failed")
	}
	global.Logger.Info("mysql init success")
	global.Mdb = db

	// set Pool
	SetPool()
	migrateTables()

}

func SetPool() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		global.Logger.Error("mysql init failed", zap.Error(err))
	}
	sqlDb.SetMaxOpenConns(global.Config.Mysql.MaxOpen)
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdle))
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConMaxLifetime))
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{}, &po.Role{},
	)
	if err != nil {
		global.Logger.Error("mysql migrate failed", zap.Error(err))
	}
}
