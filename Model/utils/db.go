package utils

import (
	"database/sql"
	"log"
)

func DBConn() (db *sql.DB) {

	dbDriver := "sqlname"
	dbUser := "username"
	dbPass := "password"
	dbName := "dbname"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func DBUpdate(postId int, amount int, location_id int, size string, db *sql.DB) {

	var size_id int

	err := db.QueryRow("SELECT id FROM sizes WHERE size=?", size).Scan(&size_id)
	if err != nil {
		panic(err.Error())
	}

	updateQuery, err := db.Prepare("UPDATE stocks SET quantity=? WHERE product_id=? AND location_id=? AND size_id=?")
	if err != nil {
		panic(err.Error())
	}
	updateQuery.Exec(amount, postId, location_id, size_id)
}

func DBCollect(db *sql.DB) []Product {

	var model_id int
	var url string

	product := Product{}
	products := []Product{}

	str := "SELECT products.id, products.price, types.type, colors.color, models.model, models.id, brands.brand FROM products LEFT JOIN types ON products.type_id = types.id LEFT JOIN colors ON products.color_id = colors.id LEFT JOIN models ON products.model_id = models.id  LEFT JOIN brands ON models.brand_id = brands.id ORDER BY products.id"

	productDB, err := db.Query(str)

	if err != nil {
		panic(err.Error())
	}

	str2 := "SELECT image_url FROM images WHERE product_id=?"

	for productDB.Next() {

		var img_url []string
		err := productDB.Scan(&product.Id, &product.Price, &product.Type, &product.Color, &product.Model, &model_id, &product.Brand)
		if err != nil {
			log.Fatal(err)
		}
		imageDB, err := db.Query(str2, product.Id)

		if err != nil {
			log.Fatal(err)
		}

		for imageDB.Next() {

			err := imageDB.Scan(&url)
			if err != nil {
				log.Fatal(err)
			}
			img_url = append(img_url, url)
			product.ImageUrl = img_url

		}
		products = append(products, product)
	}
	return products
}

func StockInfo(product_id int, size string, location_id int, db *sql.DB) int {
	var size_id int
	var quantity int

	err := db.QueryRow("SELECT id FROM sizes WHERE size=?", size).Scan(&size_id)
	if err != nil {
		panic(err.Error())
	}

	err = db.QueryRow("SELECT quantity FROM stocks WHERE product_id=? AND size_id=? AND location_id=?", product_id, size_id, location_id).Scan(&quantity)
	if err != nil {
		panic(err.Error())
	}

	return quantity
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func StockSearch(product_id int, db *sql.DB) []string {
	var sizes []string

	stockDB, err := db.Query("SELECT size_id FROM stocks WHERE product_id=? AND quantity!=?", product_id, 0)

	if err != nil {
		panic(err.Error())
	}

	for stockDB.Next() {
		var size_id int
		var size string

		stockDB.Scan(&size_id)
		err1 := db.QueryRow("SELECT size FROM sizes WHERE id=?", size_id).Scan(&size)

		if err1 != nil {
			panic(err.Error())
		}
		if !contains(sizes, size) {
			sizes = append(sizes, size)
		}
	}
	defer stockDB.Close()
	return sizes
}

func DBSearch(product_id int, db *sql.DB) Product {
	var product Product
	var model_id int
	var img_url []string

	str := "SELECT products.id, products.price, colors.color, types.type, models.model, brands.brand, models.id FROM products LEFT JOIN colors ON products.color_id = colors.id LEFT JOIN types ON products.type_id = types.id LEFT JOIN models ON products.model_id = models.id LEFT JOIN brands ON models.brand_id = brands.id WHERE products.id=?"
	err := db.QueryRow(str, product_id).Scan(&product.Id, &product.Price, &product.Color, &product.Type, &product.Model, &product.Brand, &model_id)

	if err != nil {
		panic(err.Error())
	}

	str2 := "SELECT image_url FROM images WHERE product_id=?"
	imageDB, err := db.Query(str2, product.Id)

	if err != nil {
		panic(err.Error())
	}

	for imageDB.Next() {

		var url string
		err := imageDB.Scan(&url)
		if err != nil {
			log.Fatal(err)
		}
		img_url = append(img_url, url)
	}

	product.ImageUrl = img_url
	product.Size = StockSearch(product_id, db)

	if err != nil {
		panic(err.Error())
	}
	return product
}
