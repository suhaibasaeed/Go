package main

import (
	"fmt"
	"math/rand"
	"time"
)
func randomNumberGen(min float32, max float32) float32 {	
	return min + (max-min)*rand.Float32()
}
// Task implementation goes here
type Stock struct {
  name string
  price float32
}

func (stock *Stock) updateStock() {
  change := randomNumberGen(-10000.0, 10000.0)

  stock.price += change
}

func displayMarket(market []Stock) {
  for _, value := range market {
    fmt.Println("Stock Name:", value.name, "Stock Price:", value.price)
  }
}
func main() {
  rand.Seed(time.Now().UnixNano())
	// Function calls go here
  stockMarket := []Stock{{"GOOG", 2313.50}, {"AAPL", 157.28}, {"FB", 203.77}, {"TWTR", 50.0}}

  displayMarket(stockMarket)
  // Update Google stock
  stockMarket[0].updateStock()

  displayMarket(stockMarket)
}