package main
import (
  "fmt"
  "math/rand"
  "time"
)
func getRandomElement(slice []string) string {

  randElement := rand.Intn(len(slice))
  return slice[randElement]
}
func main() {
  
  people := [5]string{"David B", "Ramesh", "Lele", "Pepe", "Nelo"}

  objects := [3]string{"Chest", "Crate", "Box"}

  fmt.Print("People: ", people)
  fmt.Println()
  fmt.Print("Objects: ", objects)

  // Use seeding using time to generate random numbers
  rand.Seed(time.Now().UnixNano())

  objSlice := objects[:]
  peopleSlice := people[:]

  randObj := getRandomElement(objSlice)
  randPerson := getRandomElement(peopleSlice)

  fmt.Println("Random Object:", randObj, "Random Person:", randPerson)
}
