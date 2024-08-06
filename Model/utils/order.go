package utils

import (
	"database/sql"
	"strconv"
)

func WriteAddress(order Order, db *sql.DB) int {
	str := "INSERT INTO address (adres, city, ilce, name, is_register, customer_id, surname, phone_number, title) VALUES (?,?,?,?,?,?,?,?,?)"
	_, err := db.Exec(str, order.Address.Adres, order.Address.City, order.Address.Ilce, order.Address.Name, order.Address.IsRegister, order.CustomerId, order.Address.Surname, order.Address.PhoneNumber, order.Address.Title)

	if err != nil {
		panic(err.Error())
	}

	var adres_id int

	str = "SELECT id FROM address WHERE adres=? AND city=?  AND ilce=? AND name=? AND is_register=? AND customer_id=? AND surname=? AND phone_number=? AND title=?"
	err = db.QueryRow(str, order.Address.Adres, order.Address.City, order.Address.Ilce, order.Address.Name, order.Address.IsRegister, order.CustomerId, order.Address.Surname, order.Address.PhoneNumber, order.Address.Title).Scan(&adres_id)

	if err != nil {
		panic(err.Error())
	}

	return adres_id
}

func WriteOrder(order Order, db *sql.DB) {
	var adres_id int

	str := "SELECT id FROM address WHERE adres=? AND city=?  AND ilce=? AND name=? AND is_register=? AND customer_id=? AND surname=? AND phone_number=? AND title=?"
	err := db.QueryRow(str, order.Address.Adres, order.Address.City, order.Address.Ilce, order.Address.Name, order.Address.IsRegister, order.CustomerId, order.Address.Surname, order.Address.PhoneNumber, order.Address.Title).Scan(&adres_id)

	if err != nil {
		adres_id = WriteAddress(order, db)
	}

	str = "INSERT INTO orders (address_id, customer_id, date,is_guest) VALUES (?,?,?,?)"
	_, err = db.Exec(str, adres_id, order.CustomerId, order.Date, order.IsGuest)
	if err != nil {
		panic(err.Error())
	}

	var order_id int

	str = "SELECT id FROM orders WHERE address_id=? AND customer_id=? AND date=?"
	err = db.QueryRow(str, adres_id, order.CustomerId, order.Date).Scan(&order_id)

	if err != nil {
		panic(err.Error())
	}

	for _, basket := range order.Baskets {

		str := "INSERT INTO basket_product (product_id, order_id, amount, size, location_id, cargo_id) VALUES (?,?,?,?,?,?)"

		amount, _ := strconv.Atoi(basket.Amount)

		for i := 0; i < amount; i++ {
			_, err := db.Exec(str, basket.Id, order_id, 1, basket.Size, basket.Location_id, basket.Cargo_id)
			if err != nil {
				panic(err.Error())
			}
		}
	}

}

func GetOrder(customerId int, db *sql.DB) []Order {

	var orders []Order

	str := "SELECT id, address_id, date FROM orders WHERE customer_id=?"

	orderDB, err := db.Query(str, customerId)

	if err != nil {
		panic(err.Error())
	}

	for orderDB.Next() {

		var adres_id int
		var order_id int
		var order Order

		err := orderDB.Scan(&order_id, &adres_id, &order.Date)

		if err != nil {
			panic(err.Error())
		}

		str := "SELECT adres, city, ilce, name, surname, phone_number, title FROM address WHERE id=?"

		err = db.QueryRow(str, adres_id).Scan(&order.Address.Adres, &order.Address.City, &order.Address.Ilce, &order.Address.Name, &order.Address.Surname, &order.Address.PhoneNumber, &order.Address.Title)
		if err != nil {
			panic(err.Error())
		}

		str = "SELECT product_id, amount, size FROM basket_product WHERE order_id=?"
		basketDB, err := db.Query(str, order_id)

		if err != nil {
			panic(err.Error())
		}

		for basketDB.Next() {

			var basket Basket

			err := basketDB.Scan(&basket.Id, &basket.Amount, &basket.Size)

			if err != nil {
				panic(err.Error())
			}

			order.Baskets = append(order.Baskets, basket)
		}

		orders = append(orders, order)

	}
	defer orderDB.Close()
	return orders
}

func GetAdress(customer_id int, db *sql.DB) []Address {
	var address []Address
	var adres Address

	query := "SELECT adres, city, ilce, name, surname, phone_number, title, is_register FROM address WHERE customer_id=? AND is_register=?"

	adreesDB, err := db.Query(query, customer_id, "1")
	if err != nil {
		panic(err.Error())
	}

	for adreesDB.Next() {

		err = adreesDB.Scan(&adres.Adres, &adres.City, &adres.Ilce, &adres.Name, &adres.Surname, &adres.PhoneNumber, &adres.Title, &adres.IsRegister)
		if err != nil {
			panic(err.Error())
		}

		address = append(address, adres)

	}

	defer adreesDB.Close()
	return address
}
