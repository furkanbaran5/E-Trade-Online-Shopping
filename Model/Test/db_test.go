package Test

import (
	"fmt"
	"module/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDbCollet(t *testing.T) {

	DeleteTable()
	db := TestDbConn()

	_, err := InsertBrandsTable("Brand1")

	if err != nil {
		fmt.Print(err.Error())
	}
	var price = 1800
	var product Test_product

	color_id, _ := InsertColorsTable("Color1")
	type_id, _ := InsertTypesTable("Shoes")
	model_id, _ := InsertModelsTable("Brand1", "Model1")

	product.Color_id = &color_id
	product.Type_id = &type_id
	product.Model_id = &model_id
	product.Price = &price

	product_id, _ := InsertProductsTable(product)

	var img = []string{"asdas", "asdasd"}

	_ = InsertImagesTable(product_id, img)

	var size []string

	expected := []utils.Product{
		{
			Type:     "Shoes",
			Brand:    "Brand1",
			Model:    "Model1",
			Color:    "Color1",
			Price:    1800,
			Id:       product_id,
			ImageUrl: img,
			Size:     size,
		},
	}

	got := utils.DBCollect(db)

	assert.Equal(t, expected, got)

}

func TestDbCollet_Empty(t *testing.T) {

	DeleteTable()
	db := TestDbConn()

	expected := []utils.Product{}

	got := utils.DBCollect(db)

	assert.Equal(t, expected, got)

}

func TestDbsearch(t *testing.T) {

	DeleteTable()
	db := TestDbConn()

	_, _ = InsertBrandsTable("Brand1")

	var price = 1800
	var product Test_product

	color_id, _ := InsertColorsTable("Color1")
	type_id, _ := InsertTypesTable("Shoes")
	model_id, _ := InsertModelsTable("Brand1", "Model1")

	product.Color_id = &color_id
	product.Type_id = &type_id
	product.Model_id = &model_id
	product.Price = &price

	product_id, _ := InsertProductsTable(product)

	var img = []string{"asdas", "asdasd"}

	_ = InsertImagesTable(product_id, img)

	var size []string

	expected := utils.Product{

		Type:     "Shoes",
		Brand:    "Brand1",
		Model:    "Model1",
		Color:    "Color1",
		Price:    1800,
		Id:       product_id,
		ImageUrl: img,
		Size:     size,
	}

	got := utils.DBSearch(product_id, db)

	assert.Equal(t, expected, got)
}

func TestDbsearch_With_Wrong_id(t *testing.T) {
	db := TestDbConn()
	defer db.Close()

	defer func() {
		if r := recover(); r != nil {
			t.Logf("Panic occurred: %v", r)
		}
	}()

	// Geçersiz ID
	product := utils.DBSearch(-1, db)

	if product.Id != 0 {
		t.Errorf("Expected product ID to be 0 for invalid input, got %d", product.Id)
	}
}

func TestDbUpdate(t *testing.T) {

	DeleteTable()

	db := TestDbConn()

	_, _ = InsertBrandsTable("Brand1")

	var price = 1800
	var product Test_product

	color_id, _ := InsertColorsTable("Color1")
	type_id, _ := InsertTypesTable("Shoes")
	model_id, _ := InsertModelsTable("Brand1", "Model1")

	product.Color_id = &color_id
	product.Type_id = &type_id
	product.Model_id = &model_id
	product.Price = &price

	product_id, _ := InsertProductsTable(product)

	size_id, _ := InsertSizesTable("S")
	location_id, _ := InsertLocationsTable("Rize")
	fmt.Println(size_id, location_id)

	InsertStocksTable(product_id, 1800, size_id, location_id)
	utils.DBUpdate(product_id, 999, location_id, "S", db)

	got := 0
	expected := 999

	query := "SELECT quantity FROM stocks WHERE product_id=? AND location_id=? AND size_id=?"

	_ = db.QueryRow(query, product_id, location_id, size_id).Scan(&got)

	assert.Equal(t, expected, got)
}

func TestStockInfo(t *testing.T) {

	DeleteTable()

	db := TestDbConn()

	data := Test_data{
		Brand:    "Brand1",
		Color:    "Color1",
		Type:     "Type1",
		Model:    "Model1",
		Size:     "Size1",
		Location: "Location1",
	}

	product_id, _, location_id := setupTestData(data)

	got := utils.StockInfo(product_id, "Size1", location_id, db)
	expected := 1800

	query := "SELECT quantity FROM stocks WHERE product_id=? AND location_id=? AND size_id=?"
	_ = db.QueryRow(query, product_id, location_id, "S").Scan(&got)
	assert.Equal(t, expected, got)
}

func TestStockInfoWithWrongProductID(t *testing.T) {
	db := TestDbConn()
	defer db.Close()

	defer func() {
		if r := recover(); r != nil {
			t.Logf("Panic occurred: %v", r)
		}
	}()

	//Yanlış product_id
	id := utils.StockInfo(-1, "S", 5, db)

	if id != 0 {
		t.Errorf("Expected product ID to be 0 for invalid input, got %d", id)
	}
}

func TestStockInfoWithWrongSizeId(t *testing.T) {

	DeleteTable()
	db := TestDbConn()
	defer db.Close()

	defer func() {
		if r := recover(); r != nil {
			t.Logf("Panic occurred: %v", r)
		}
	}()

	data := Test_data{
		Brand:    "Brand1",
		Color:    "Color1",
		Type:     "Type1",
		Model:    "Model1",
		Size:     "Size1",
		Location: "Location1",
	}

	product_id, _, location_id := setupTestData(data)
	//Yanlış Size
	got := utils.StockInfo(product_id, "M", location_id, db)
	if got != 0 {
		t.Errorf("Expected product ID to be 0 for invalid input, got %d", got)
	}
}

func TestStockSearch(t *testing.T) {

	DeleteTable()
	//DeleteProductsTable()

	db := TestDbConn()

	data := Test_data{
		Brand:    "Brand1",
		Color:    "Color1",
		Type:     "Type1",
		Model:    "Model1",
		Size:     "Size1",
		Location: "Location1",
	}

	product_id, _, location_id := setupTestData(data)

	got := utils.StockSearch(product_id, db)
	expected := []string{"Size1"}

	query := "SELECT quantity FROM stocks WHERE product_id=? AND location_id=? AND size_id=?"
	_ = db.QueryRow(query, product_id, location_id, "S").Scan(&got)
	assert.Equal(t, expected, got)
}

func TestStockSearchWithWrongId(t *testing.T) {

	DeleteTable()

	db := TestDbConn()
	defer db.Close()

	data := Test_data{
		Brand:    "Brand1",
		Color:    "Color1",
		Type:     "Type1",
		Model:    "Model1",
		Size:     "Size1",
		Location: "Location1",
	}

	_, _, _ = setupTestData(data)

	got := utils.StockSearch(-1, db)
	var expected []string

	assert.Equal(t, expected, got)
}
