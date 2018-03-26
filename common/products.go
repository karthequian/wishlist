package common

type Product struct {
	Name              string  `json:"productName"`
	Price             float32 `json:"price"`
	ID                string  `json:"id"`
	QuantityAvailable int     `json:"quantityAvailable"`
}

var ProductList = []Product{
	Product{
		Name:              "Hammer",
		Price:             9.99,
		ID:                "1",
		QuantityAvailable: 100,
	},
	Product{
		Name:              "Nail",
		Price:             .1,
		ID:                "2",
		QuantityAvailable: 10000,
	},
	Product{
		Name:              "Screw",
		Price:             .12,
		ID:                "3",
		QuantityAvailable: 10000,
	},
	Product{
		Name:              "Corded Drill",
		Price:             98.99,
		ID:                "4",
		QuantityAvailable: 58,
	},
	Product{
		Name:              "Cordless Drill",
		Price:             129.99,
		ID:                "5",
		QuantityAvailable: 62,
	},
	Product{
		Name:              "Screwdriver",
		Price:             12.99,
		ID:                "6",
		QuantityAvailable: 83,
	},
	Product{
		Name:              "Drillbits",
		Price:             21.99,
		ID:                "7",
		QuantityAvailable: 32,
	},
	Product{
		Name:              "Drillbits Jumbo pack",
		Price:             34.99,
		ID:                "8",
		QuantityAvailable: 21,
	},
	Product{
		Name:              "Work Gloves- unisex",
		Price:             8.99,
		ID:                "9",
		QuantityAvailable: 20,
	},
	Product{
		Name:              "Garden gloves- Male",
		Price:             9.99,
		ID:                "10",
		QuantityAvailable: 21,
	},
	Product{
		Name:              "Garden gloves- Female",
		Price:             9.99,
		ID:                "11",
		QuantityAvailable: 21,
	},
}

var ProductMap = make(map[string]Product)

func CreateProductMap() {
	for _, product := range ProductList {
		ProductMap[product.ID] = product
	}

}
