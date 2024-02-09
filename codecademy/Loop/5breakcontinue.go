package main

import (
	"fmt"
)

func main() {

	for count := 0; count < 20; count++ {
		// WRITE CONTINUE STATEMENT IF COUNT EQUALS 8 BELOW THIS LINE
		if count == 8 {
			continue
		}
		// WRITE BREAK STATEMENT IF COUNT EQUALS 15 BELOW THIS LINE
		if count == 15 {
			break
		}
		fmt.Println(count)
	}

}
