package repository

import (
	"dousheng-demo/config"
	"dousheng-demo/model"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once

func Init() {
	var err error
	once.Do(func() {

		DB = ConnectDB()
	})

	err = InitUser()
	err = InitVideo()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The database is initialized successful.")
	}
	fmt.Println(DB.Find(&model.User{}).RowsAffected)
}

func ConnectDB() (conn *gorm.DB) {
	var err error
	conn, err = gorm.Open(mysql.Open(config.MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	return conn
}

func InitUser() error{
	var err error
	m := DB.Migrator()
	if m.HasTable(&model.User{}) {
		return nil
	}
	err = m.CreateTable(&model.User{})
	return err
}

func InitVideo() error{
	var err error
	m := DB.Migrator()
	if m.HasTable(&model.Video{}) {
		return nil
	}
	err = m.CreateTable(&model.Video{})
	return err
}