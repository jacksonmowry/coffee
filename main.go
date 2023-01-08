package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

type Order struct {
	gorm.Model
	Name        string
	Light       uint
	LightGround uint
	Med         uint
	MedGround   uint
	Dark        uint
	DarkGround  uint
}

func main() {
	log.Fatal("hi")
}
