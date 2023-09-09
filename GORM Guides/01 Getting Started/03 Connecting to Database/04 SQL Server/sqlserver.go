package main

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main(){
	dsn := "sqlserver://gorm:LoremIpsum@localhost:9930?database=gorm"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
}