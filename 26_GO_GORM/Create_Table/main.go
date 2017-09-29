package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	 _ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type UserModel struct{
	UserId int  `gorm:"primary_key";"AUTO_INCREMENT"`
	Name string `gorm:"size:255"`
	Address string `gorm:"type:varchar(100)"`
}


func main(){
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ormdemo?charset=utf8&parseTime=True")
	if err != nil {
		log.Panic(err)
	}
	log.Println("Connection Established")
	db.DropTableIfExists(&UserModel{})
	db.AutoMigrate(&UserModel{})
}


