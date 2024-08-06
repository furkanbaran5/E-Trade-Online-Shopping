package Test

import (
	"module/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginWithPhoneNumber(t *testing.T) {

	db := TestDbConn()
	DeleteTable()

	customer := utils.Customer{
		Id:          0,
		Name:        "Name",
		Surname:     "Surname",
		PhoneNumber: "PhoneNumber",
		Mail:        "Mail",
		Password:    "Password",
	}

	customer.Id, _ = InsertCustomerTable(customer)

	loginInput := utils.LoginInput{
		PhoneNumberOrMail: customer.PhoneNumber,
		Password:          customer.Password,
	}

	expected_id := utils.CustomerID{
		Name: customer.Name,
		Id:   customer.Id,
	}
	expected_control := 1

	var got_id utils.CustomerID
	var got_control int

	got_id, got_control = utils.Login(loginInput, db)

	assert.Equal(t, expected_id, got_id)
	assert.Equal(t, expected_control, got_control)
}

func TestLoginWithMail(t *testing.T) {

	db := TestDbConn()
	DeleteTable()

	customer := utils.Customer{
		Id:          0,
		Name:        "Name",
		Surname:     "Surname",
		PhoneNumber: "PhoneNumber",
		Mail:        "Mail",
		Password:    "Password",
	}

	customer.Id, _ = InsertCustomerTable(customer)

	loginInput := utils.LoginInput{
		PhoneNumberOrMail: customer.Mail,
		Password:          customer.Password,
	}

	expected_id := utils.CustomerID{
		Name: customer.Name,
		Id:   customer.Id,
	}
	expected_control := 1

	var got_id utils.CustomerID
	var got_control int

	got_id, got_control = utils.Login(loginInput, db)

	assert.Equal(t, expected_id, got_id)
	assert.Equal(t, expected_control, got_control)
}

func TestLoginWithWrongPassword(t *testing.T) {

	db := TestDbConn()
	DeleteTable()

	customer := utils.Customer{
		Id:          0,
		Name:        "Name",
		Surname:     "Surname",
		PhoneNumber: "PhoneNumber",
		Mail:        "Mail",
		Password:    "Password",
	}

	customer.Id, _ = InsertCustomerTable(customer)

	loginInput := utils.LoginInput{
		PhoneNumberOrMail: customer.Mail,
		Password:          "WrongPassword",
	}

	expected_id := utils.CustomerID{
		Name: customer.Name,
		Id:   customer.Id,
	}
	expected_control := -1

	var got_id utils.CustomerID
	var got_control int

	got_id, got_control = utils.Login(loginInput, db)

	assert.Equal(t, expected_id, got_id)
	assert.Equal(t, expected_control, got_control)
}

func TestRegister(t *testing.T) {

	db := TestDbConn()
	DeleteTable()

	customer := utils.Customer{
		Id:          1,
		Name:        "Name",
		Surname:     "Surname",
		PhoneNumber: "PhoneNumber",
		Mail:        "Mail",
		Password:    "Password",
	}

	ınfo := utils.Info{
		Ret:  1,
		Text: "Kayıt başarılı",
	}

	expected := ınfo
	got := utils.Register(customer, db)

	assert.Equal(t, expected, got)

}
