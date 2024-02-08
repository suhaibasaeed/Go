package main

import "fmt"

func main() {
	star := "Polaris"
	
	starAddress := &star
	
	// Dereference starAddress and assign new value
  *starAddress = "Sirius"
  
  fmt.Println("The actual value of star is", star)
}
