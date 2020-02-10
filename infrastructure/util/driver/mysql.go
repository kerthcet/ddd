package driver

import (
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	// driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"

	"ddd/infrastructure/config"
)

// DB 数据库链接管理
var DB *gorm.DB

func init() {
	driver := strings.Split(config.ENV.DatabaseURL, "://")

	var err error
	DB, err = gorm.Open(driver[0], driver[1])
	if err != nil {
		log.Error(err)
		panic(err)
	}

	DB.DB().SetMaxIdleConns(10)
	connNumber, _ := strconv.ParseInt("64", 10, 64)
	DB.DB().SetMaxOpenConns(int(connNumber))
	DB.DB().SetMaxIdleConns(int(connNumber))

	DB.LogMode(config.ENV.GormLogMode)
}
