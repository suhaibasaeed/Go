package main

import "fmt"

type Cake struct {
	typeOfCake string
	weight int // Weight in grams
}

func main() {

	// Checkpoint 1 code goes here	
  cakes := [3]Cake{{"Chocolate", 1000}, {"Carrot", 500}, {"Ice Cream", 2000}}
	// Checkpoint 2 code goes here
	fmt.Println("Weight of choco cake", cakes[0].weight)
	// Checkpoint 3 code goes here
	cakes[1].weight = 1500
	fmt.Println(cakes)

}
