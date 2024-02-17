package main

import "fmt"

type Triangle struct {
	height float32
	base float32
}

// Checkpoint 1 code goes here
func (triangle Triangle) area() float32 {
  return 0.5 * (triangle.height * triangle.base)
}
func main() {

	triangle := Triangle{10, 4}
  fmt.Println(triangle)

	// Call the function here
  fmt.Println("The triangle's area is:", triangle.area())
}