package main
import "fmt"

func main() {
  donuts := map[string]int{
    "frosted":   10,
    "chocolate": 15,
    "jelly":     8,
  }
  fmt.Println(donuts)

  // Add your code here
  delete(donuts, "chocolate")
  fmt.Println(donuts)
}