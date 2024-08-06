package utils

import (
	"database/sql"
)

func CustomerInfo(id int, db *sql.DB) Customer {
	var customer Customer
	str := "SELECT name, surname, phone_number, mail FROM customers WHERE id=?"
	err := db.QueryRow(str, id).Scan(&customer.Name, &customer.Surname, &customer.PhoneNumber, &customer.Mail)

	if err != nil {
		return customer
	}
	customer.Id = id
	return customer
}

func UpdateCustomerInfo(Customer Customer, db *sql.DB) {
	var str string

	if Customer.Name != "" {
		str = "UPDATE customers SET name =? WHERE id=?"
		uptadeCustomer, err := db.Prepare(str)

		if err != nil {
			panic(err.Error())
		}
		_, err = uptadeCustomer.Exec(Customer.Name, Customer.Id)

		if err != nil {
			panic(err.Error())
		}

	}
	if Customer.Surname != "" {
		str = "UPDATE customers SET surname =? WHERE id=?"
		uptadeCustomer, err := db.Prepare(str)

		if err != nil {
			panic(err.Error())
		}
		_, err = uptadeCustomer.Exec(Customer.Surname, Customer.Id)

		if err != nil {
			panic(err.Error())
		}

	}
	if Customer.PhoneNumber != "" {
		str = "UPDATE customers SET phone_number =? WHERE id=?"
		uptadeCustomer, err := db.Prepare(str)

		if err != nil {
			panic(err.Error())
		}
		_, err = uptadeCustomer.Exec(Customer.PhoneNumber, Customer.Id)

		if err != nil {
			panic(err.Error())
		}

	}
	if Customer.Mail != "" {
		str = "UPDATE customers SET mail =? WHERE id=?"
		uptadeCustomer, err := db.Prepare(str)

		if err != nil {
			panic(err.Error())
		}
		_, err = uptadeCustomer.Exec(Customer.Mail, Customer.Id)

		if err != nil {
			panic(err.Error())
		}

	}

}
