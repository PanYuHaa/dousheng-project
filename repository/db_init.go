package repository

import (
	"dousheng-demo/config"
	"dousheng-demo/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var DB *gorm.DB
var once sync.Once
var mu sync.Mutex

func Init() {
	var err error
	once.Do(func() {
		DB = ConnectDB()
	})

	err = InitUser()
	err = InitVideo()
	err = InitFavorite()
	err = InitFollow()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The database is initialized successful.")
	}
}

func ConnectDB() (conn *gorm.DB) {
	var err error
	conn, err = gorm.Open(mysql.Open(config.MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	return conn
}

func InitUser() error {
	var err error
	m := DB.Migrator()
	if m.HasTable(&model.User{}) {
		return nil
	}
	err = m.CreateTable(&model.User{})
	return err
}

func InitVideo() error {
	var err error
	m := DB.Migrator()
	if m.HasTable(&model.Video{}) {
		return nil
	}
	err = m.CreateTable(&model.Video{})
	return err
}

func InitFavorite() error {
	var err error
	m := DB.Migrator()
	if m.HasTable(&model.Favorite{}) {
		return nil
	}
	err = m.CreateTable(&model.Favorite{})
	return err
}

func InitFollow() error {
	var err error
	m := DB.Migrator()
	if m.HasTable(&model.Follow{}) {
		return nil
	}
	err = m.CreateTable(&model.Follow{})
	return err
}
