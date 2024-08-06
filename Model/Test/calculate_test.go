package Test

import (
	"module/utils"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cargoInfo = []utils.CargoInfo{
	{Location_id: 1, Cargo_id: 1, Price_per_distance: 0.04, Discount_per_piece: 0.5, Order_price: 50},
	{Location_id: 1, Cargo_id: 2, Price_per_distance: 0.03, Discount_per_piece: 0.6, Order_price: 48},
	{Location_id: 1, Cargo_id: 3, Price_per_distance: 0.028, Discount_per_piece: 0.4, Order_price: 52},
	{Location_id: 2, Cargo_id: 1, Price_per_distance: 0.32, Discount_per_piece: 0.4, Order_price: 48},
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
var Locations = []utils.Locations{
	{Id: 1, Location: "İstanbul", Capacity: 1000, Process: 500},
	{Id: 2, Location: "Ankara", Capacity: 800, Process: 200},
	{Id: 3, Location: "İzmir", Capacity: 800, Process: 400},
	{Id: 4, Location: "Diyarbakır", Capacity: 600, Process: 250},
	{Id: 5, Location: "Rize", Capacity: 450, Process: 150},
}
var db = TestDbConn()

func TestWriteInOrder(t *testing.T) {
	tests := []struct {
		name            string
		baskets         []utils.Basket
		address         utils.Address
		order           utils.Order
		bestCombination utils.BestCombination
		allCombination  [][]int
		expectedBaskets []utils.Basket
	}{
		{
			name: "Test Case 1",
			baskets: []utils.Basket{
				{Id: "305", Size: "M", Amount: "1", Location_id: 0, Cargo_id: 0},
				{Id: "388", Size: "S", Amount: "1", Location_id: 0, Cargo_id: 0},
			},
			address: utils.Address{
				Adres: "asd", City: "İstanbul", Ilce: "Bakırköy", Name: "Furkan", PhoneNumber: "5321212125", Surname: "Baran", Title: "Ev", IsRegister: "0",
			},
			order: utils.Order{
				Address: utils.Address{Adres: "asd", City: "İstanbul", Ilce: "Bakırköy", Name: "Furkan", PhoneNumber: "5321212125", Surname: "Baran", Title: "Ev", IsRegister: "0"},
				Baskets: []utils.Basket{
					{Id: "305", Size: "M", Amount: "1", Location_id: 0, Cargo_id: 0},
					{Id: "388", Size: "S", Amount: "1", Location_id: 0, Cargo_id: 0},
				},
				CustomerId: 35, Date: "safas", IsGuest: 0,
			},
			bestCombination: utils.BestCombination{
				Combination: []int{2, 1}, Point: 4.02, Ways: []utils.Way{
					{Location_id: 1, Cargo_id: 2, Amount: 75.9},
					{Location_id: 2, Cargo_id: 2, Amount: 59.97},
				},
			},
			allCombination: [][]int{
				{305, 388}, {2, 1}, {4, 1},
			},
			expectedBaskets: []utils.Basket{
				{Id: "305", Size: "M", Amount: "1", Location_id: 2, Cargo_id: 2},
				{Id: "388", Size: "S", Amount: "1", Location_id: 1, Cargo_id: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.WriteInOrder(&tt.order, tt.bestCombination, tt.allCombination)
			assert.Equal(t, tt.expectedBaskets, tt.order.Baskets)
		})
	}
}

func TestCalculatePointToCargo(t *testing.T) {
	tests := []struct {
		name         string
		distance     utils.Distance
		cargoInfo    []utils.CargoInfo
		adet         int
		ExpectedWays utils.Way
	}{
		{
			name: "Test Case 1",
			distance: utils.Distance{
				Location: 4, City: 1, Distance: 526, DistanceKey: 2,
			},
			cargoInfo: []utils.CargoInfo{
				{Location_id: 4, Cargo_id: 1, Price_per_distance: 0.024, Discount_per_piece: 0.6, Order_price: 45},
				{Location_id: 4, Cargo_id: 2, Price_per_distance: 0.022, Discount_per_piece: 0.4, Order_price: 45},
				{Location_id: 4, Cargo_id: 3, Price_per_distance: 0.024, Discount_per_piece: 0.6, Order_price: 42},
			},
			adet: 1,
			ExpectedWays: utils.Way{
				Location_id: 4, Cargo_id: 3, Amount: 54.624,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.CalculatePointToCargo(tt.distance, tt.cargoInfo, float32(tt.adet))
			assert.Equal(t, tt.ExpectedWays, got)
		})
	}
}

func TestCalculatePointToCapacity(t *testing.T) {
	tests := []struct {
		name          string
		Loc           int
		ExpectedPoint float32
	}{
		{
			name:          "Test Case 1",
			Loc:           1,
			ExpectedPoint: 2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.CalculatePointToCapacity(Locations, tt.Loc)
			assert.Equal(t, float32(tt.ExpectedPoint), got)
		})
	}
}
func TestCalculateAllCombinations(t *testing.T) {

	tests := []struct {
		name                string
		matrixDistances     [][]utils.Distance
		allCombination      [][]int
		len                 int
		expectedCombination utils.BestCombination
	}{
		{
			name: "Test Case 1",
			matrixDistances: [][]utils.Distance{
				{
					{Location: 2, City: 1, Distance: 499, DistanceKey: 2},
					{Location: 4, City: 1, Distance: 526, DistanceKey: 2},
				},
				{
					{Location: 1, City: 1, Distance: 930, DistanceKey: 4},
				},
			},
			allCombination: [][]int{{305, 388}, {2, 1}, {4, 1}},
			len:            2,
			expectedCombination: utils.BestCombination{
				Combination: []int{2, 1},
				Point:       4.021723,
				Ways: []utils.Way{
					{Location_id: 1, Cargo_id: 2, Amount: 75.9},
					{Location_id: 2, Cargo_id: 2, Amount: 59.974},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.CalculateAllCombinations(tt.allCombination, tt.matrixDistances, cargoInfo, Locations, tt.len)
			assert.Equal(t, tt.expectedCombination, got)
		})
	}
}

func TestFilterCargosByLocationOfSingle(t *testing.T) {
	tests := []struct {
		name              string
		Loc               int
		ExpectedCargoInfo []utils.CargoInfo
	}{
		{
			name: "Test Case 1",
			Loc:  1,
			ExpectedCargoInfo: []utils.CargoInfo{
				{Location_id: 1, Cargo_id: 1, Price_per_distance: 0.04, Discount_per_piece: 0.5, Order_price: 50},
				{Location_id: 1, Cargo_id: 2, Price_per_distance: 0.03, Discount_per_piece: 0.6, Order_price: 48},
				{Location_id: 1, Cargo_id: 3, Price_per_distance: 0.028, Discount_per_piece: 0.4, Order_price: 52},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.FilterCargosByLocationOfSingle(cargoInfo, tt.Loc)
			assert.Equal(t, tt.ExpectedCargoInfo, got)
		})
	}
}
func TestFindCombinations(t *testing.T) {
	tests := []struct {
		name                 string
		locations            []utils.LocationsForCombination
		allCombination       [][]int
		ExpectAllCombination [][]int
	}{
		{
			name: "Test Case 1",
			locations: []utils.LocationsForCombination{
				{
					Location_ids: []int{2, 4},
					Product_id:   305,
				},
				{
					Location_ids: []int{1},
					Product_id:   388,
				},
			},
			allCombination:       [][]int{{305, 388}},
			ExpectAllCombination: [][]int{{305, 388}, {2, 1}, {4, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.FindCombinations(tt.locations, []int{}, 0, len(tt.locations), &tt.allCombination)
			assert.Equal(t, tt.ExpectAllCombination, tt.allCombination)
		})
	}
}

func TestFilterDistancesByLocations(t *testing.T) {
	tests := []struct {
		name             string
		distance         []utils.Distance
		ExpectedDistance []utils.Distance
	}{
		{
			name: "Test Case 1",
			distance: []utils.Distance{
				{Location: 2, City: 1, Distance: 499, DistanceKey: 2},
				{Location: 3, City: 1, Distance: 902, DistanceKey: 4},
				{Location: 4, City: 1, Distance: 526, DistanceKey: 2},
				{Location: 5, City: 1, Distance: 955, DistanceKey: 4},
			},
			ExpectedDistance: []utils.Distance{
				{Location: 2, City: 1, Distance: 499, DistanceKey: 2},
				{Location: 4, City: 1, Distance: 526, DistanceKey: 2},
			},
		},
		{
			name: "Test Case 2",
			distance: []utils.Distance{
				{Location: 1, City: 1, Distance: 930, DistanceKey: 4},
			},
			ExpectedDistance: []utils.Distance{
				{Location: 1, City: 1, Distance: 930, DistanceKey: 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.FilterDistancesByLocations(tt.distance)
			assert.Equal(t, tt.ExpectedDistance, got)
		})
	}
}

func TestFilterDistances(t *testing.T) {
	tests := []struct {
		name             string
		distances        []utils.Distance
		location_ids     []int
		ExpectedDistance []utils.Distance
	}{
		{
			name: "Test Case 1",
			distances: []utils.Distance{
				{Location: 1, City: 1, Distance: 930, DistanceKey: 4},
				{Location: 2, City: 1, Distance: 499, DistanceKey: 2},
				{Location: 3, City: 1, Distance: 902, DistanceKey: 4},
				{Location: 4, City: 1, Distance: 526, DistanceKey: 2},
				{Location: 5, City: 1, Distance: 955, DistanceKey: 4},
			},
			location_ids: []int{2, 3, 4, 5},
			ExpectedDistance: []utils.Distance{
				{Location: 2, City: 1, Distance: 499, DistanceKey: 2},
				{Location: 3, City: 1, Distance: 902, DistanceKey: 4},
				{Location: 4, City: 1, Distance: 526, DistanceKey: 2},
				{Location: 5, City: 1, Distance: 955, DistanceKey: 4},
			},
		},
		{
			name: "Test Case 2",
			distances: []utils.Distance{
				{Location: 1, City: 1, Distance: 930, DistanceKey: 4},
				{Location: 2, City: 1, Distance: 499, DistanceKey: 2},
				{Location: 3, City: 1, Distance: 902, DistanceKey: 4},
				{Location: 4, City: 1, Distance: 526, DistanceKey: 2},
				{Location: 5, City: 1, Distance: 955, DistanceKey: 4},
			},
			location_ids: []int{1},
			ExpectedDistance: []utils.Distance{
				{Location: 1, City: 1, Distance: 930, DistanceKey: 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.FilterDistances(tt.distances, tt.location_ids)
			assert.Equal(t, tt.ExpectedDistance, got)
		})
	}
}

func TestSearchCapacity(t *testing.T) {

	tests := []struct {
		name      string
		Locations []utils.Locations
	}{
		{
			name: "Test Case 2",
			Locations: []utils.Locations{
				{Id: 1, Location: "İstanbul", Capacity: 1000, Process: 500},
				{Id: 2, Location: "Ankara", Capacity: 800, Process: 200},
				{Id: 3, Location: "İzmir", Capacity: 800, Process: 400},
				{Id: 4, Location: "Diyarbakır", Capacity: 600, Process: 250},
				{Id: 5, Location: "Rize", Capacity: 450, Process: 150},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteTable()
			InsertLocations(db, tt.Locations)
			got := utils.SearchCapacity(db)
			assert.Equal(t, tt.Locations, got)
		})
	}
}

func TestSearchCargoInfo(t *testing.T) {
	DeleteTable()
	got := utils.SearchCargoInfo(db)
	assert.Equal(t, cargoInfo, got)
}

func TestSearchLocations(t *testing.T) {
	tests := []struct {
		name      string
		marka     string
		quantity  int
		color     string
		typing    string
		model     string
		size      string
		locations []string
	}{
		{
			name:      "Test Case 1",
			marka:     "brand1",
			quantity:  300,
			color:     "color",
			typing:    "hat",
			model:     "air",
			size:      "S",
			locations: []string{"İstanbul", "Ankara", "Diyarbakır"},
		},
		{
			name:      "Test Case 2",
			marka:     "brand1",
			quantity:  300,
			color:     "color",
			typing:    "hat",
			model:     "air",
			size:      "S",
			locations: []string{"İzmir"},
		},
		{
			name: "Test Case 3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteTable()
			expectedLocationIds, product_id := insert(tt.marka, tt.color, tt.typing, tt.model, tt.quantity, tt.size, tt.locations)
			got := utils.SearchLocations(product_id, tt.size, 10, db)
			assert.Equal(t, expectedLocationIds, got)
		})
	}
}

func TestCalculateProductStock(t *testing.T) {
	tests := []struct {
		name      string
		order     utils.Order
		locations []string
		expID     string
	}{
		{
			name: "Test Case 1",
			order: utils.Order{
				Address: utils.Address{City: "Adana"},
				Baskets: []utils.Basket{
					{Id: "305", Size: "M", Amount: "1", Location_id: 0, Cargo_id: 0},
					{Id: "310", Size: "S", Amount: "1", Location_id: 0, Cargo_id: 0},
				},
				CustomerId: 35, Date: "safas", IsGuest: 0,
			},
			locations: []string{"Ankara", "Diyarbakır"},
			expID:     "",
		},
		{
			name: "Test Case 2",
			order: utils.Order{
				Address: utils.Address{City: "Muş"},
				Baskets: []utils.Basket{
					{Id: "305", Size: "M", Amount: "1", Location_id: 0, Cargo_id: 0},
					{Id: "310", Size: "S", Amount: "1", Location_id: 0, Cargo_id: 0},
				},
				CustomerId: 35, Date: "safas", IsGuest: 0,
			},
			locations: []string{"Diyarbakır"},
			expID:     "",
		},
		{
			name: "Test Case 3",
			order: utils.Order{
				Address: utils.Address{City: "Diyarbakır"},
				Baskets: []utils.Basket{
					{Id: "305", Size: "M", Amount: "1", Location_id: 0, Cargo_id: 0},
				},
				CustomerId: 35, Date: "safas", IsGuest: 0,
			},
			locations: []string{"İstanbul", "İzmir"},
			expID:     "",
		},
		{
			name: "Test Case 3",
			order: utils.Order{
				Address: utils.Address{City: "Diyarbakır"},
				Baskets: []utils.Basket{
					{Id: "305", Size: "M", Amount: "1", Location_id: 0, Cargo_id: 0},
					{Id: "310", Size: "L", Amount: "1", Location_id: 0, Cargo_id: 0},
				},
				CustomerId: 35, Date: "safas", IsGuest: 0,
			},
			locations: []string{"İstanbul", "İzmir"},
			expID:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteTable()
			var expectedIDS []int
			length := len(tt.order.Baskets)
			expectedLocations := make([]utils.LocationsForCombination, length)
			for i := 0; i < length; i++ {
				x := tt.order.Baskets[i].Id
				location_ids, product_id := insert(x, x, x, x, 10, tt.order.Baskets[i].Size, tt.locations)
				tt.order.Baskets[i].Id = strconv.Itoa(product_id)
				expectedLocations[i].Location_ids = location_ids
				expectedLocations[i].Product_id = product_id
				expectedIDS = append(expectedIDS, product_id)
			}
			cust := utils.Customer{
				Id:          0,
				Name:        "",
				Surname:     "",
				PhoneNumber: "",
				Mail:        "",
				Password:    "",
			}
			tt.order.CustomerId, _ = InsertCustomerTable(cust)
			_, locations, Ids, cargo, ID := utils.CalculateProductStock(tt.order, db)
			assert.Equal(t, expectedLocations, locations)
			assert.Equal(t, expectedIDS, Ids)
			assert.Equal(t, cargoInfo, cargo)
			assert.Equal(t, tt.expID, ID)
		})
	}
}
