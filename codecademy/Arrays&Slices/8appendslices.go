package main

import "fmt"

func main() {
    myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}
    fmt.Println(myTutors)
    myTutors = append(myTutors, "Josh")
    fmt.Println(myTutors)
}