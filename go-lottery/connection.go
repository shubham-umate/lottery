package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	fmt.Println("Connecting to database")
	connection, err := gorm.Open(mysql.Open("root:root@/loot_users"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = connection
	connection.AutoMigrate(&User{},&Lottery{},&Participant{})
}
