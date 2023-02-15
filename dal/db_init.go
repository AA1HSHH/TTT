package dal

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
	TimeOut    string
	db         *gorm.DB
)

func Init() error {
	var err error
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", DbUser, DbPassWord, DbHost, DbPort, DbName, TimeOut)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic("fail to connect database ,err = " + err.Error())
	}
	return err
}
func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}
