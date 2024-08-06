package Test

import (
	"database/sql"
	"log"
	"module/utils"

	_ "github.com/go-sql-driver/mysql"
)

type Test_product struct {
	Type_id  *int
	Model_id *int
	Color_id *int
	Price    *int
}

type Test_stocks struct {
	Quantity int
	Size     string
	Location string
}

type Test_address struct {
	Adres      utils.Address
	CustomerID int
}

type Test_Basket struct {
	Product_id  int
	Order_id    int
	Amount      int
	Size        string
	Location_id int
	Cargo_id    int
}

type Test_order struct {
	address_id  int
	customer_id int
	date        string
	is_guest    int
}

type Test_data struct {
	Brand    string
	Color    string
	Type     string
	Model    string
	Size     string
	Location string
}

func TestDbConn() (db *sql.DB) {

	dbDriver := "sqlname"
	dbUser := "username"
	dbPass := "password"
	dbName := "dbName"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	return db
}
func InsertLocations(db *sql.DB, Locations []utils.Locations) {
	str := "INSERT INTO locations (id, location, capacity, process) VALUES (?,?,?,?)"
	for _, loc := range Locations {
		_, err := db.Exec(str, loc.Id, loc.Location, loc.Capacity, loc.Process)
		if err != nil {
			log.Print(err.Error())
		}
	}
}

func InsertCargoInfos(db *sql.DB, cargoInfos []utils.CargoInfo) {
	str := "INSERT INTO orderpriceinfos (location_id, cargo_id, price_per_distance , discount_per_piece, order_price) VALUES (?,?,?,?,?)"
	for _, cargo := range cargoInfos {
		_, err := db.Exec(str, cargo.Location_id, cargo.Cargo_id, cargo.Price_per_distance, cargo.Discount_per_piece, cargo.Order_price)
		if err != nil {
			log.Print(err.Error())
		}
	}
}

func insert(marka string, color string, typing string, model string, quantity int, size string, locations []string) ([]int, int) {
	var expectedLocationIds []int
	var product Test_product
	_, _ = InsertBrandsTable(marka)
	color_id, _ := InsertColorsTable(color)
	type_id, _ := InsertTypesTable(typing)
	model_id, _ := InsertModelsTable(marka, model)
	product.Color_id = &color_id
	product.Type_id = &type_id
	product.Model_id = &model_id
	product.Price = &quantity
	product_id, _ := InsertProductsTable(product)
	size_id, _ := InsertSizesTable(size)

	for _, loc := range locations {
		location_id := SearchLocationsTestDB(loc)
		expectedLocationIds = append(expectedLocationIds, location_id)
	}
	for _, loc := range expectedLocationIds {
		InsertStocksTable(product_id, quantity, size_id, loc)
	}
	return expectedLocationIds, product_id
}

func SearchLocationsTestDB(location string) int {
	db := TestDbConn()
	defer db.Close()
	var location_id int
	err := db.QueryRow("SELECT id FROM locations WHERE location=?", location).Scan(&location_id)
	if err != nil {
		log.Println("Error querying distances:", err)
		return 0
	}
	return location_id
}

func InsertTypesTable(types string) (int, error) {

	db := TestDbConn()
	defer db.Close()

	var columns = []string{"type"}

	query := "INSERT INTO types ("
	for _, str := range columns {
		query += str
	}
	query += ") VALUES (?)"

	_, err := db.Exec(query, types)

	if err != nil {
		log.Print(err.Error())
	}

	var id int
	query = "SELECT id FROM types WHERE type=?"

	err = db.QueryRow(query, types).Scan(&id)

	if err != nil {
		log.Print(err.Error())
	}

	return id, err
}

func InsertBrandsTable(brand string) (int, error) {

	db := TestDbConn()
	defer db.Close()

	var columns = []string{"brand"}

	query := "INSERT INTO brands ("
	for _, str := range columns {
		query += str
	}
	query += ") VALUES (?)"

	_, _ = db.Exec(query, brand)

	var id int
	query = "SELECT id FROM brands WHERE brand=?"

	err := db.QueryRow(query, brand).Scan(&id)
	return id, err

}

func InsertModelsTable(brand string, model string) (int, error) {

	db := TestDbConn()
	defer db.Close()

	brand_id, _ := InsertBrandsTable(brand)
	var columns = []string{"brand_id,", "model"}

	query := "INSERT INTO models ("
	for _, str := range columns {
		query += str
	}
	query += ") VALUES (?, ?)"

	_, _ = db.Exec(query, brand_id, model)

	var id int
	query = "SELECT id FROM models WHERE brand_id=? AND model=?"

	err := db.QueryRow(query, brand_id, model).Scan(&id)
	return id, err
}

func InsertColorsTable(color string) (int, error) {

	db := TestDbConn()
	defer db.Close()

	var columns = []string{"color"}

	query := "INSERT INTO colors ("
	for _, str := range columns {
		query += str
	}
	query += ") VALUES (?)"

	_, _ = db.Exec(query, color)

	var id int
	query = "SELECT id FROM colors WHERE color=?"

	err := db.QueryRow(query, color).Scan(&id)

	return id, err
}

func InsertImagesTable(product_id int, image []string) error {

	db := TestDbConn()
	defer db.Close()

	var columns = []string{"product_id,", "image_url"}
	var err error

	for _, img := range image {

		query := "INSERT INTO images ("
		for _, str := range columns {
			query += str
		}
		query += ") VALUES (?, ?)"

		_, err = db.Exec(query, product_id, img)
		if err != nil {
			log.Print(err.Error())
		}

	}

	return err

}

func InsertProductsTable(product Test_product) (int, error) {

	db := TestDbConn()
	defer db.Close()

	var columns = []string{"type_id,", "model_id,", "color_id,", "price"}

	query := "INSERT INTO products ("
	for _, str := range columns {
		query += str
	}
	query += ") VALUES (?, ?, ?, ?)"

	_, _ = db.Exec(query, product.Type_id, product.Model_id, product.Color_id, product.Price)
	var product_id int

	query = "SELECT id FROM products WHERE type_id =? AND model_id =? AND color_id =? AND price =?"
	err := db.QueryRow(query, product.Type_id, product.Model_id, product.Color_id, product.Price).Scan(&product_id)
	return product_id, err

}

func DeleteTable() {

	db := TestDbConn()
	var tables = []string{"basket_product", "orders",
		"address",
		"stocks",
		"sizes",
		"images",
		"products",
		"types",
		"models",
		"brands",
		"colors",
		"customers"}

	for _, tableName := range tables {

		_, err := db.Exec("DELETE FROM " + tableName)

		if err != nil {
			log.Print(err.Error())
		}
	}

}

func InsertSizesTable(size string) (int, error) {

	db := TestDbConn()
	defer db.Close()

	var columns = []string{"size"}

	query := "INSERT INTO sizes ("
	for _, str := range columns {
		query += str
	}
	query += ") VALUES (?)"

	_, _ = db.Exec(query, size)

	var id int
	query = "SELECT id FROM sizes WHERE size=?"

	err := db.QueryRow(query, size).Scan(&id)
	return id, err
}

func InsertLocationsTable(location string) (int, error) {

	db := TestDbConn()
	defer db.Close()

	var columns = []string{"location"}

	query := "INSERT INTO locations ("
	for _, str := range columns {
		query += str
	}
	query += ") VALUES (?)"

	_, _ = db.Exec(query, location)

	var id int
	query = "SELECT id FROM locations WHERE location=?"

	err := db.QueryRow(query, location).Scan(&id)

	return id, err
}

func InsertStocksTable(product_id int, quantity int, size_id int, location_id int) error {

	db := TestDbConn()
	defer db.Close()

	var columns = []string{"product_id,", "quantity,", "size_id,", "location_id"}

	query := "INSERT INTO stocks ("
	for _, str := range columns {
		query += str
	}
	query += ") VALUES (?, ?, ?, ?)"

	_, err := db.Exec(query, product_id, quantity, size_id, location_id)

	return err

}

func setupTestData(data Test_data) (int, int, int) {
	_, err := InsertBrandsTable(data.Brand)
	if err != nil {
		log.Fatalf("Failed to insert brand: %v", err)
	}

	color_id, err := InsertColorsTable(data.Color)
	if err != nil {
		log.Fatalf("Failed to insert color: %v", err)
	}

	type_id, err := InsertTypesTable(data.Type)
	if err != nil {
		log.Fatalf("Failed to insert type: %v", err)
	}

	model_id, err := InsertModelsTable(data.Brand, data.Model)
	if err != nil {
		log.Fatalf("Failed to insert model: %v", err)
	}

	price := 1800
	product := Test_product{
		Color_id: &color_id,
		Type_id:  &type_id,
		Model_id: &model_id,
		Price:    &price,
	}

	product_id, err := InsertProductsTable(product)
	if err != nil {
		log.Fatalf("Failed to insert product: %v", err)
	}

	size_id, err := InsertSizesTable(data.Size)
	if err != nil {
		log.Fatalf("Failed to insert size: %v", err)
	}

	location_id, err := InsertLocationsTable(data.Location)
	if err != nil {
		log.Fatalf("Failed to insert location: %v", err)
	}

	err = InsertStocksTable(product_id, 1800, size_id, location_id)
	if err != nil {
		log.Fatalf("Failed to insert stock: %v", err)
	}

	return product_id, size_id, location_id
}

func InsertCustomerTable(customer utils.Customer) (int, error) {

	db := TestDbConn()
	var columns = []string{"name,", "surname,", "phone_number,", "mail,", "password"}

	query := "INSERT INTO customers ("
	for _, str := range columns {
		query += str
	}
	query += ") VALUES (?,?,?,?,?)"
	HashPassword, _ := utils.HashPassword(customer.Password)

	_, _ = db.Exec(query, customer.Name, customer.Surname, customer.PhoneNumber, customer.Mail, HashPassword)

	str := "SELECT id FROM customers WHERE name=? AND surname=? AND phone_number=? AND mail=? "
	err := db.QueryRow(str, customer.Name, customer.Surname, customer.PhoneNumber, customer.Mail).Scan(&customer.Id)

	return customer.Id, err
}

func InsertAdressTable(data Test_address) int {

	db := TestDbConn()
	var columns = []string{"adres,", "city,", "ilce,", "name,", "surname,", "phone_number,", "title,", "is_register,", "customer_id"}

	query := "INSERT INTO address ("
	for _, str := range columns {
		query += str
	}
	query += ") VALUES (?,?,?,?,?,?,?,?,?)"

	_, err := db.Exec(query,
		data.Adres.Adres,
		data.Adres.City,
		data.Adres.Ilce,
		data.Adres.Name,
		data.Adres.Surname,
		data.Adres.PhoneNumber,
		data.Adres.Title,
		data.Adres.IsRegister,
		data.CustomerID)

	if err != nil {
		panic(err.Error())
	}

	var adres_id int
	str := "SELECT id FROM address WHERE adres=? AND city=? AND ilce=? AND name=? AND surname=? AND phone_number=? AND title=? AND is_register=? AND customer_id=?"
	_ = db.QueryRow(str,
		data.Adres.Adres,
		data.Adres.City,
		data.Adres.Ilce,
		data.Adres.Name,
		data.Adres.Surname,
		data.Adres.PhoneNumber,
		data.Adres.Title,
		data.Adres.IsRegister,
		data.CustomerID).Scan(&adres_id)
	return adres_id
}

func InsertOrdersTable(order Test_order) int {

	db := TestDbConn()

	var order_id int
	var columns = []string{"address_id,", "customer_id,", "date,", "is_guest"}

	query := "INSERT INTO orders ("
	for _, str := range columns {
		query += str
	}
	query += ") VALUES (?,?,?,?)"

	_, _ = db.Exec(query, order.address_id, order.customer_id, order.date, order.is_guest)

	str := "SELECT id FROM orders WHERE address_id=? AND customer_id=? AND date=? AND is_guest=?"
	_ = db.QueryRow(str, order.address_id, order.customer_id, order.date, order.is_guest).Scan(&order_id)

	return order_id
}

func InsertBasketProductTable(basket Test_Basket) int {

	db := TestDbConn()
	defer db.Close()

	var columns = []string{"product_id,", "order_id,", "amount,", "size,", "location_id,", "cargo_id"}

	query := "INSERT INTO basket_product ("
	for _, str := range columns {
		query += str
	}
	query += ") VALUES (?,?,?,?,?,?)"

	_, _ = db.Exec(query, basket.Product_id, basket.Order_id, basket.Amount, basket.Size, basket.Location_id, basket.Cargo_id)

	return 1
}

func InsertCargo(db *sql.DB, names []string) {
	str := "INSERT INTO cargos (id,name) VALUES (?,?)"

	for i, cargo := range names {
		_, err := db.Exec(str, i+1, cargo)
		if err != nil {
			log.Print(err.Error())
		}
	}
}

func GetCargoInfo() []utils.CargoInfo {

	var cargoInfo = []utils.CargoInfo{
		{Location_id: 1, Cargo_id: 1, Price_per_distance: 0.04, Discount_per_piece: 0.5, Order_price: 50},
		{Location_id: 1, Cargo_id: 2, Price_per_distance: 0.03, Discount_per_piece: 0.6, Order_price: 48},
		{Location_id: 1, Cargo_id: 3, Price_per_distance: 0.028, Discount_per_piece: 0.4, Order_price: 52},
		{Location_id: 2, Cargo_id: 1, Price_per_distance: 0.032, Discount_per_piece: 0.4, Order_price: 48},
		{Location_id: 2, Cargo_id: 2, Price_per_distance: 0.026, Discount_per_piece: 0.5, Order_price: 47},
		{Location_id: 2, Cargo_id: 3, Price_per_distance: 0.026, Discount_per_piece: 0.6, Order_price: 50},
		{Location_id: 3, Cargo_id: 1, Price_per_distance: 0.038, Discount_per_piece: 0.5, Order_price: 51},
		{Location_id: 3, Cargo_id: 2, Price_per_distance: 0.03, Discount_per_piece: 0.5, Order_price: 49},
		{Location_id: 3, Cargo_id: 3, Price_per_distance: 0.026, Discount_per_piece: 0.5, Order_price: 53},
		{Location_id: 4, Cargo_id: 1, Price_per_distance: 0.024, Discount_per_piece: 0.6, Order_price: 45},
		{Location_id: 4, Cargo_id: 2, Price_per_distance: 0.022, Discount_per_piece: 0.4, Order_price: 45},
		{Location_id: 4, Cargo_id: 3, Price_per_distance: 0.024, Discount_per_piece: 0.6, Order_price: 42},
		{Location_id: 5, Cargo_id: 1, Price_per_distance: 0.016, Discount_per_piece: 0.4, Order_price: 40},
		{Location_id: 5, Cargo_id: 2, Price_per_distance: 0.018, Discount_per_piece: 0.6, Order_price: 42},
		{Location_id: 5, Cargo_id: 3, Price_per_distance: 0.022, Discount_per_piece: 0.4, Order_price: 45},
	}
	return cargoInfo
}

func GetLocatonsInfo() []utils.Locations {

	var Locations = []utils.Locations{
		{Id: 1, Location: "İstanbul", Capacity: 1000, Process: 500},
		{Id: 2, Location: "Ankara", Capacity: 800, Process: 200},
		{Id: 3, Location: "İzmir", Capacity: 800, Process: 400},
		{Id: 4, Location: "Diyarbakır", Capacity: 600, Process: 250},
		{Id: 5, Location: "Rize", Capacity: 450, Process: 150},
	}
	return Locations
}
