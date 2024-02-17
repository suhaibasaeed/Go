package main
import "fmt"

func main() {
    myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}
    changeLastElement(myTutors, "Bobby")
}

func changeLastElement(slice []string, str string){

  slice[len(slice)-1] = str
  fmt.Println(slice)

}