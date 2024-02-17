package main

import "fmt"

type Point struct {
	x int
	y int
}

// Checkpoint 1 goes here
type Circle struct {
  Point
  radius int
}
func main() {
	circle := Circle{Point{4,5}, 2}
	
// Checkpoint 2 code goes here
	fmt.Println(circle.x)
  fmt.Println(circle)
}
