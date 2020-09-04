package data

import (
	"database/sql"
	"fmt"
	"runtime/debug"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type Test struct {
// 	gorm.Model
// 	Code  string
// 	Price uint
// }

var dbMysql *gorm.DB

// Init : DB 초기화
func Init() bool {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	// db 커넥션 초기화
	dbMysql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("fail to init mysql : %v", err)

		// debug.PrintStack()
		// pc, _, _, ok := runtime.Caller(1)
		// details := runtime.FuncForPC(pc)
		// f, l := details.FileLine(pc)
		// if ok && details != nil {
		// 	fmt.Printf("\ncalled from %s / %s / %d\n", details.Name(), f, l)
		// }
		// _, file, line, ok := runtime.Caller(1)
		// fmt.Printf("\nfail to init mysql : %s , %d, %v\n", file, line, ok)

		return false
	}

	// connection pool 세팅
	var sqlDB *sql.DB
	if sqlDB, err = dbMysql.DB(); err != nil {
		fmt.Printf("fail to init mysql : %v", err)
		debug.PrintStack()
		return false
	}
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(1000)
	sqlDB.SetConnMaxLifetime(time.Hour * 2)

	// 데이터 스키마 마이그레이션
	dbMysql.AutoMigrate(&User{})

	return true
}

// Health : sql연결 체크
func Health() bool {
	var sqlDB *sql.DB
	var err error
	if sqlDB, err = dbMysql.DB(); err != nil {
		fmt.Printf("fail to init mysql : %v", err)
		debug.PrintStack()
		return false
	}
	if err = sqlDB.Ping(); err != nil {
		fmt.Printf("fail to init mysql : %v", err)
		debug.PrintStack()
		return false
	}
	return true
}
