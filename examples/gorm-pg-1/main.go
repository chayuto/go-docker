package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string
	Code  string
	Price uint
}

func main() {
	fmt.Printf("hello, Go\n")

	// dsn := "host=localhost user=postgres password=changeme dbname=local_pgdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "host=localhost user=postgres password=changeme dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})
	db.Create(&Product{Code: "C43", Price: 200})

	// Read
	var product Product
	// db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	// db.Delete(&product, 2)

}
