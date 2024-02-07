package main

import (
  "fmt"
)

func main() {

  var name string
  var age int
  
  fmt.Println("What is your name? ")
  fmt.Scan(&name)
  fmt.Println("Hello", name)
  
  fmt.Println("What is your age? ")
  fmt.Scan(&age)
  fmt.Printf("So your name is %s and your age is %d", name, age)

  newMessage := fmt.Sprintf("Oh %s you are only %d years old", name, age)

  fmt.Println(newMessage)

}