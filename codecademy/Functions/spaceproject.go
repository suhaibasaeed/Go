package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
  fmt.Println(fuel)
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
  var fuel int

  if planet == "Venus" {
    fuel = 300000
  } else if planet == "Mercury" {
    fuel = 500000
  } else if planet == "Mars" {
    fuel = 700000
  }
  return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
  fmt.Println("welcome to planet", planet)
}

// Create the function cantFly() here
func cantFly() {
  fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
  var fuelRemaining, fuelCost int

  fuelRemaining = fuel
  fuelCost = calculateFuel(planet)

  if fuelRemaining >= fuelCost {
    greetPlanet(planet)
    fuelRemaining -= fuelCost
  } else {
    cantFly()
  }

  return fuelRemaining
}

func main() {
  // Test your functions!
  
  // Create `planetChoice` and `fuel`
  fuel := 1000000
  planetChoice := "Venus"
  // And then liftoff!
  fuel = flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)
}