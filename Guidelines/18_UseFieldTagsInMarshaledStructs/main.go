package main

type StockBad struct {
	Price int
	Name  string
}

type StockGood struct {
	Price int    `json:"price"`
	Name  string `json:"name"`
	// Safe to rename Name to Symbol.
}

func main() {
	//bytes, err := json.Marshal(StockBad{
	//	Price: 137,
	//	Name:  "UBER",
	//})
	//
	//bytes, err := json.Marshal(StockGood{
	//	Price: 137,
	//	Name:  "UBER",
	//})
}
