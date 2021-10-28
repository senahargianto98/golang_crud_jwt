package main

import (
	"skb/Config"
	"skb/Model"
)

func main() {
	db, err := Config.GetDB()
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Model.Artikel{})
	db.AutoMigrate(&Model.User{})
}
