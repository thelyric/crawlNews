package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func ConnectMysql() (*gorm.DB, error) {
	count := 1
	user := getEnv("DB_USER", "root")
	pass := getEnv("DB_PASS", "123456")
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "3307")
	name := getEnv("DB_NAME", "foodDelivery")
	params := getEnv("DB_PARAMS", "charset=utf8mb4&parseTime=True&loc=Local")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, pass, host, port, name, params)

	for {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("Connected to database successfully")
			return db, nil
		}
		if count >= 5 {
			fmt.Printf("Failed to connect to database, retrying %d times\n", count)
			return nil, err
		}
		count++
	}
}

func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Error getting raw database object:", err)
		return err
	}
	return sqlDB.Close()
}
