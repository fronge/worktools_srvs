package initialize

import (
	"fmt"
	"log"
	"os"
	"time"
	"worktools_srvs/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitMySQL() {
	c := global.ServerConfig.MysqlInfo
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Name,
		c.Password,
		c.Host,
		c.Port,
		c.DB)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold: time.Second * time.Duration(10), // 慢 SQL 阈值
			LogLevel:      logger.Info,                     // 日志级别
			Colorful:      true,                            // 禁用彩色打印
		},
	)
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}
