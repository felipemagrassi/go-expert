package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// belongs to and has many
type Category struct {
	ID       int `gorm:"primary_key"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

// has one
type SerialNumber struct {
	ID        int `gorm:"primary_key"`
	Number    string
	ProductID int
}

// type Product struct {
// 	ID         int `gorm:"primary_key"`
// 	Name       string
// 	Price      float64
// 	CategoryID int
// 	gorm.Model
// 	Category     Category     `gorm:"foreignKey:CategoryID"`
// 	SerialNumber SerialNumber `gorm:"foreignKey:ProductID"`
// }

type Product struct {
	ID         int `gorm:"primary_key"`
	Name       string
	Price      float64
	CategoryID int
	gorm.Model
	Categories   []Category   `gorm:"many2many:products_categories;"`
	SerialNumber SerialNumber `gorm:"foreignKey:ProductID"`
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&SerialNumber{})

	// category := Category{Name: "Eletrodomésticos"}
	// product := Product{Name: "Geladeira", Price: 1000, Category: category}

	// db.Create(&Category{Name: "Eletrodomésticos"})
	// db.Create(&Product{Name: "Geladeira", Price: 1000, CategoryID: 1})
	// db.Create(&SerialNumber{Number: "123456789", ProductID: 1})

	// var products []Product

	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	//
	// for _, p := range products {
	// 	println(p.Name)
	// 	println(p.Category.Name)
	// 	println(p.SerialNumber.Number)
	// }

	var categories []Category

	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name)
		fmt.Println(category.ID)
		fmt.Println("Produtos:")
		for _, product := range category.Products {
			fmt.Println(product.Name)
			fmt.Println(product.SerialNumber.Number)
		}

	}

}
