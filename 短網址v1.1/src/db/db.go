package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Conn *gorm.DB
)

func Get() *gorm.DB {
	return Conn
}

func Connect(ip, port, database, user, password string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true&parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci",
		user,
		password,
		ip,
		port,
		database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	Conn = db
	sqlDB, err := Conn.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(1024)

	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {

		fmt.Println("無法連接到 MySQL 伺服器")
	} else {
		fmt.Println("已成功連接到 MySQL 伺服器")
	}
	return nil
}
func Disconnect() {
	if Conn != nil {
		db, err := Conn.DB()
		if err != nil {
			fmt.Println("無法獲取資料庫連接:", err)
			return
		}
		db.Close()
		fmt.Println("資料庫連接已關閉")
		Conn = nil
	}
}
