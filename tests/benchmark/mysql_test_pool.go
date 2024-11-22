package benchmark

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

type User struct {
	gorm.Model
	Name string
}

func insertRecord(b *testing.B, db *gorm.DB) {
	user := User{Name: "test"}
	if err := db.Create(&user).Error; err != nil {
		b.Fatal(err)
	}
}

func BenchMarkMaxOpenConns1(b *testing.B) {
	dsn := "root:1@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// check if table exist
	if db.Migrator().HasTable(&User{}) == false {
		// Drop tb if exist
		if err := db.Migrator().CreateTable(&User{}); err != nil {
			log.Fatal(err)
			// Handle error
		}
	}
	// Create new tb
	err = db.AutoMigrate(&User{})
	if err != nil {
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// setting number connection
	sqlDB.SetMaxOpenConns(1)
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

// go test -bench=. -benchmem

func BenchMarkMaxOpenConns10(b *testing.B) {
	dsn := "root:1@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// check if table exist
	if db.Migrator().HasTable(&User{}) == false {
		// Drop tb if exist
		if err := db.Migrator().CreateTable(&User{}); err != nil {
			log.Fatal(err)
			// Handle error
		}
	}
	// Create new tb
	err = db.AutoMigrate(&User{})
	if err != nil {
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// setting number connection
	sqlDB.SetMaxOpenConns(10)
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}
