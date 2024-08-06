package utils

import (
	"database/sql"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(login LoginInput, db *sql.DB) (CustomerID, int) {
	var customerID CustomerID

	var phoneNumber string
	var mail string
	var password string

	str := "SELECT name, id, phone_number, mail, password FROM customers WHERE phone_number=? OR mail=?"
	err := db.QueryRow(str, login.PhoneNumberOrMail, login.PhoneNumberOrMail).Scan(&customerID.Name, &customerID.Id, &phoneNumber, &mail, &password)

	if err != nil {
		if err == sql.ErrNoRows {
			return customerID, -1
		}
	} else if VerifyPassword(login.Password, password) {
		return customerID, 1
	}
	return customerID, -1

}

func Register(customer Customer, db *sql.DB) Info {
	customerDB, err := db.Prepare("INSERT INTO customers (name, surname, phone_number, mail, password) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	customer.Password, err = HashPassword(customer.Password)
	if err != nil {
		panic(err.Error())
	}
	_, err = customerDB.Exec(customer.Name, customer.Surname, customer.PhoneNumber, customer.Mail, customer.Password)
	if err != nil {
		if strings.Contains(err.Error(), "phone_number") {
			//log.Println("Syntax error in SQL query:", err.Error())
			info := Info{
				Ret:  -1,
				Text: "Bu telefon numarası daha önce kullanılmış",
			}
			return info
		}
		if strings.Contains(err.Error(), "mail") {
			info := Info{
				Ret:  -1,
				Text: "Bu mail adresi daha önce kullanılmış",
			}
			return info
		}
	}
	info := Info{
		Ret:  1,
		Text: "Kayıt başarılı",
	}
	defer customerDB.Close()
	return info
}
