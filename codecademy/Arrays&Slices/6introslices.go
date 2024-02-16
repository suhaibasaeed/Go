package main

import "fmt"

func main() {
    myTutors := [4]string{"Kirsty", "Mishell", "Jose", "Neil"}
    fmt.Println(myTutors)

    a_slice := myTutors[:]
    // Create new slice
    subjects := []string{"Go", "Java", "Python"}
    
    fmt.Println(a_slice, subjects)
}