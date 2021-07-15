package main

type Core_theme struct {
	Id        int    `gorm:"id" json:"id"`
	Directory string `gorm:"directory" json:"directory"`
	Name      string `gorm:"name" json:"name"`
	Data      string `gorm:"data" json:"data"`
}
