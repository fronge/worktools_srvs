package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"worktools_srvs/model"

	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

func main() {
	dsn := "root:S4SaRs3rUbgJ86kJ@tcp(104.129.60.57:3339)/mxshop_user_srvs?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(

		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold: time.Second * time.Duration(10), // 慢 SQL 阈值
			LogLevel:      logger.Info,                     // 日志级别
			Colorful:      true,                            // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	options := &password.Options{16, 100, 32, sha256.New}
	salt, encodedPwd := password.Encode("admin123", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	for i := 0; i < 10; i++ {
		user := model.User{
			NickName: fmt.Sprintf("Jerry%d", i),
			Mobile:   fmt.Sprintf("1878999988%d", i),
			Password: newPassword,
		}
		db.Save(&user)
	}
	// // 创建表
	// db.AutoMigrate(&model.User{})

	// options := &password.Options{16, 100, 32, sha256.New}
	// salt, encodedPwd := password.Encode("generic password", options)
	// fmt.Println(salt)
	// fmt.Println(encodedPwd)

	// newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	// passwordInfo := strings.Split(newPassword, "$")
	// check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
	// fmt.Println(check)
	// fmt.Println(newPassword)

}
