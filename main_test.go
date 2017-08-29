package main

import (
	"github.com/antonyho/go-kart/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"testing"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&models.Group{},
		&models.OrderedGroup{},
		&models.Image{},
		&models.Item{},
		&models.OrderedItem{},
	)
}

func TestMain(m *testing.M) {
	defer db.Close()
	returnCode := m.Run()
	os.Exit(returnCode)
}
