package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})
	// db.Create(&Product{Name: "Laptop", Price: 1000})

	// var product Product
	// db.First(&product, "name = ?", "Laptop")
	//
	// fmt.Println(product)
	//
	// var products []Product
	//
	// db.Find(&products)
	//
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product
	// db.Limit(2).Find(&products)
	//
	// for _, product := range products {
	// 	println(product.Name)
	// }

	// var products []Product
	//
	// db.Where("name LIKE ?", "%laptop%").Find(&products)
	//
	// for _, product := range products {
	// 	println(product.Name)
	// }

	// var product Product
	// db.First(&product)
	//
	// product.Name = "Macbook Pro"
	// db.Save(&product)
	//
	// var product2 Product
	// db.First(&product2)
	//
	// fmt.Println(product2.Name)
	// db.Delete(&product2)

}
