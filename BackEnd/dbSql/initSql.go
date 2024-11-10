package dbsql

import (
	"fmt"

	"github.com/lyx1213812138/BilibiliCleanPlan/data"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitSql() error {
	var err error // LEARN： 保证db是全局变量，所以这里不用 :=
	Db, err = gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	Db.AutoMigrate(&data.Up{})
	Db.AutoMigrate(&data.Video{})
	Db.AutoMigrate(&data.Season{})
	return nil
}
