package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

var db *sql.DB

var server = "localhost"
var port = 14333
var user = "sa"
var password = "Your_password123"
var database = "DB_NEW_SME"

func main() {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	// Read Products
	count, err := ReadProducts()
	// count, err := GetProductPrice(1012)
	if err != nil {
		log.Fatal("Error reading Products: ", err.Error())
	}
	fmt.Printf("Read %d row(s) successfully.\n", count)

}

// ReadProducts reads all product records
func ReadProducts() (int, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("SELECT [GoodCode],[GoodID], [GoodName1] FROM [dbo].[EMGood];")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	var count int

	// Iterate through the result set.
	for rows.Next() {
		var goodCode, goodName1 string
		var goodId int

		// Get values from row.
		err := rows.Scan(&goodCode, &goodId, &goodName1)
		if err != nil {
			return -1, err
		}

		queryCount, queryErr := GetProductPrice(goodId)
		if queryErr != nil {
			log.Fatal("Error reading Products: ", err.Error())
			break
		}

		if queryCount > 0 {
			fmt.Printf("ID: %d, goodId: %s, Name: %s\n", goodId, goodCode, goodName1)
		}

		count++
	}

	return count, nil
}

func GetProductPrice(goodId int) (int, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("Select [GoodPrice1] FROM [dbo].[EMGoodMultiPrice] WHERE [GoodID] = @id")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql, sql.Named("id", goodId))
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	var count int

	// Iterate through the result set.
	for rows.Next() {
		var goodPrice1 string

		// Get values from row.
		err := rows.Scan(&goodPrice1)
		if err != nil {
			return -1, err
		}

		fmt.Printf("ID: %d, goodPrice1: %s\n", goodId, goodPrice1)
		count++
	}

	// // Execute non-query with named parameters
	// result, err := db.ExecContext(
	// 	ctx,
	// 	tsql,
	// 	sql.Named("Location", location),
	// 	sql.Named("Name", name))
	// if err != nil {
	// 	return -1, err
	// }

	return count, nil
}
