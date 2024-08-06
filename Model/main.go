package main

import (
	"encoding/json"
	"fmt"
	"log"
	"module/utils"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	db := utils.DBConn()
	defer db.Close()

	app.Get("/request", func(c fiber.Ctx) error {
		return c.JSON(utils.DBCollect(db))
	})

	app.Post("/id", func(c fiber.Ctx) error {
		body := c.Body()
		var data string
		err := json.Unmarshal(body, &data)
		if err != nil {
			fmt.Printf("err was %v\n", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		id, err := strconv.Atoi(data)
		if err != nil {
			fmt.Printf("err was %v\n", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id format"})
		}

		return c.JSON(utils.DBSearch(id, db))

	})

	app.Post("/baskets", func(c fiber.Ctx) error {
		var Order utils.Order

		body := c.Body()
		var val []byte = []byte(body)
		jsonInput, err := strconv.Unquote(string(val))
		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal([]byte(jsonInput), &Order)
		if err != nil {
			fmt.Printf("err was %v", err)
		}
		info := utils.Allocator(Order, db)
		if info.Ret == 1 {
			utils.WriteOrder(Order, db)
		}

		utils.GetPdf(Order, db)

		return c.JSON(info)
	})

	app.Post("/getOrder", func(c fiber.Ctx) error {
		var Id int

		body := c.Body()
		var val []byte = []byte(body)
		jsonInput, err := strconv.Unquote(string(val))
		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal([]byte(jsonInput), &Id)
		if err != nil {
			fmt.Printf("err was %v", err)
		}
		Order := utils.GetOrder(Id, db)
		if Order != nil {
			return c.JSON(Order)
		} else {
			return c.JSON(-1)
		}

	})

	app.Post("/register", func(c fiber.Ctx) error {
		var Customer utils.Customer

		body := c.Body()
		var val []byte = []byte(body)
		jsonInput, err := strconv.Unquote(string(val))
		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal([]byte(jsonInput), &Customer)
		if err != nil {
			fmt.Printf("err was %v", err)
		}

		info := utils.Register(Customer, db)
		return c.JSON(info)
	})

	app.Post("/login", func(c fiber.Ctx) error {
		var LoginInput utils.LoginInput

		body := c.Body()
		var val []byte = []byte(body)
		jsonInput, err := strconv.Unquote(string(val))
		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal([]byte(jsonInput), &LoginInput)
		if err != nil {
			fmt.Printf("err was %v", err)
		}

		Customer, control := utils.Login(LoginInput, db)
		if control == -1 {
			return nil
		}
		return c.JSON(Customer)
	})

	app.Post("/customerInfo", func(c fiber.Ctx) error {
		var Id int

		body := c.Body()
		var val []byte = []byte(body)
		jsonInput, err := strconv.Unquote(string(val))
		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal([]byte(jsonInput), &Id)
		if err != nil {
			fmt.Printf("err was %v", err)
		}
		Customer := utils.CustomerInfo(Id, db)

		return c.JSON(Customer)
	})

	app.Post("/updateCustomerInfo", func(c fiber.Ctx) error {
		var Customer utils.Customer

		body := c.Body()
		var val []byte = []byte(body)
		jsonInput, err := strconv.Unquote(string(val))
		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal([]byte(jsonInput), &Customer)
		if err != nil {
			fmt.Printf("err was %v", err)
		}
		utils.UpdateCustomerInfo(Customer, db)

		return c.JSON(Customer)
	})

	app.Get("/GetCities", func(c fiber.Ctx) error {
		return c.JSON(utils.CityCollect(db))
	})

	app.Post("/district", func(c fiber.Ctx) error {
		var District string

		body := c.Body()

		err := json.Unmarshal([]byte(body), &District)
		if err != nil {
			fmt.Printf("err was %v", err)
		}

		return c.JSON(utils.GetDistrict(District, db))
	})

	app.Post("/address", func(c fiber.Ctx) error {
		var Id int

		body := c.Body()
		var val []byte = []byte(body)
		jsonInput, err := strconv.Unquote(string(val))
		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal([]byte(jsonInput), &Id)
		if err != nil {
			fmt.Printf("err was %v", err)
		}
		Address := utils.GetAdress(Id, db)

		return c.JSON(Address)
	})

	// Sunucuyu port 3000'de başlatın
	log.Fatal(app.Listen(":3000"))
}
