package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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
	//
	// DB Init
	//
	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect DB")
	}
	// Migrate Schema
	db.AutoMigrate(&Order{})

	// Fyne Init
	app := app.New()
	window := app.NewWindow("Coffee Orders")

	var list *widget.List
	// Creating the List Widget
	add := widget.NewButton("Add Order", func() {
		// add order functionality
		log.Fatal("new order")
	})
	exit := widget.NewButton("Quit", func() {
		window.Close()
	})
	window.SetContent(container.NewBorder(nil, container.New(layout.NewVBoxLayout(), add, exit), nil, nil, list))
	window.Resize(fyne.Newsize(800, 600))
	window.SetMaster()
	window.CenterOnScreen()
	window.ShowAndRun()
}

func addOrder(order Order, db Database) {
	db.Create(&order)
}
