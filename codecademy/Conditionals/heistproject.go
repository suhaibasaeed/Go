package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Generate unqiue seed
	rand.Seed(time.Now().UnixNano())

	isHeistOn := true
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}

	openedVault := rand.Intn(100)
	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("The vault can't be opened")
	}

	leftSafely := rand.Intn(5)

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Heist Failed")
		case 1:
			isHeistOn = false
			fmt.Println("Vault doors won't open from inside")
		default:
			fmt.Println("Start the gataway car!")
		}
	}
	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println(amtStolen, "was stolen")
	}

	fmt.Println("The heist is still on yes?", isHeistOn)
}
