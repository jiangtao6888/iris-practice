package config

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris/v12"
)

func InitDB(c *DbConfig) (err error) {
	if Log == nil || Log.Logger == nil {
		return errors.New("logger uninitialized")
	}
	DB, err = gorm.Open(c.Drive, c.User+":"+c.Password+"@/"+c.Database+"?charset=utf8&parseTime=True&loc=Local") //protocol("+c.Host+":"+c.Port+")
	if err != nil {
		Log.LogInfo("orm failed to initialized: ", err)
	}
	iris.RegisterOnInterrupt(func() {
		_ = DB.Close()
	})
	return err
}

func GetDB() *DbConfig {
	return config.DB
}
