package main

import "fmt"

type Triangle struct {
	height float32
	base float32
}

// Checkpoint 1 code goes here
func (triangle *Triangle) updateBase (newBase float32) {
  triangle.base = newBase
}
func main() {

  triangle := Triangle{10, 4}
  
  // Call the function here
  triangle.updateBase(13)
  fmt.Println(triangle)

}
