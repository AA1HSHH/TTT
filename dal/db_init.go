package dal

//https://github.com/Moonlight-Zhao/go-project-example/blob/main/repository/db_init.go
import (
	"github.com/AA1HSHH/TTT/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error
	dsn := config.MySQLDefaultDSN
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
