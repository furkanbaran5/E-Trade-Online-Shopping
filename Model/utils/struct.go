package utils

type Product struct {
	Type     string   `json:"Type"`
	Brand    string   `json:"Brand"`
	Model    string   `json:"Model"`
	Color    string   `json:"Color"`
	Price    float64  `json:"Price"`
	Id       int      `json:"Id"`
	ImageUrl []string `json:"ImageUrl"`
	Size     []string `json:"Size"`
}

type Basket struct {
	Id          string `json:"Id"`
	Size        string `json:"Size"`
	Amount      string `json:"Amount"`
	Location_id int
	Cargo_id    int
}

type Address struct {
	Adres       string `json:"adres"`
	City        string `json:"city"`
	Ilce        string `json:"ilce"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Surname     string `json:"surname"`
	Title       string `json:"title"`
	IsRegister  string `json:"isRegister"`
}

type Order struct {
	Address    Address  `json:"Address"`
	Baskets    []Basket `json:"Baskets"`
	CustomerId int      `json:"CustomerId"`
	Date       string   `json:"Date"`
	IsGuest    int      `json:"IsGuest"`
}

type CustomerID struct {
	Name string `json:"name"`
	Id   int    `json:"Id"`
}

type Info struct {
	Ret  int    `json:"Ret"`
	Text string `json:"Text"`
}

type Customer struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phoneNumber"`
	Mail        string `json:"mail"`
	Password    string `json:"password"`
}

type LoginInput struct {
	PhoneNumberOrMail string `json:"phoneNumberOrMail"`
	Password          string `json:"password"`
}

type LocationsForCombination struct {
	Location_ids []int
	Product_id   int
}

type Distance struct {
	Location    int
	City        int
	Distance    float32
	DistanceKey float32
}

type CargoInfo struct {
	Location_id        int
	Cargo_id           int
	Price_per_distance float32
	Discount_per_piece float32
	Order_price        float32
}

type Way struct {
	Location_id int
	Cargo_id    int
	Amount      float32
}

type Locations struct {
	Id       int
	Location string
	Capacity float32
	Process  float32
}

type BestCombination struct {
	Combination []int
	Point       float32
	Ways        []Way
}
type EmailMessage struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Date    string `json:"date"`
}
