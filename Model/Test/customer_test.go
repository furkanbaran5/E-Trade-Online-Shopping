package Test

import (
	"module/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomerInfo(t *testing.T) {

	db := TestDbConn()
	DeleteTable()

	customer := utils.Customer{
		Id:          0,
		Name:        "Name",
		Surname:     "Surname",
		PhoneNumber: "PhoneNumber",
		Mail:        "Mail",
		Password:    "",
	}

	customer.Id, _ = InsertCustomerTable(customer)
	expected := customer
	got := utils.CustomerInfo(customer.Id, db)

	assert.Equal(t, expected, got)

}

func TestCustomerInfoWithWrongİd(t *testing.T) {

	db := TestDbConn()
	DeleteTable()

	customer := utils.Customer{}

	_, _ = InsertCustomerTable(customer)
	expected := customer
	got := utils.CustomerInfo(-1, db)

	assert.Equal(t, expected, got)

}
