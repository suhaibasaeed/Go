# Go Notes

## Learn Go - Introduction

### From The Get Go
* Developed by Google
* Sits between low-level & high-level
* Produces code that runs fast and uses **little memory**
    * Has **garbage collection**
* Some **object oriented** features

### Compiling Files
* To make a `.go` file into an executable we need to compile it via the `go build` command
    * E.g. `go build greet.go`
    * After this by doing an `ls` in the directory we can now see that we have an executable called `greet`
```
ls
greet     greet.go
```
* This can be executed via `./greet`

### Running Files
* We can directly run a file instead of the two step process shown above of compilation and then running
    * The `go run` command compiles the code and then runs it
    * But an executable **won't** be created in our directory
    * E.g. `go run greet.go`
* Useful for quick testing of code

### Packages
* Projects can have multiple `.go` files organised into **packages**
  * Essentially like a directory
  * E.g. calculator program would have calculation `.go` files in `calc` package
    *  I/O related go files in `io` package
* Simple example
  * 1st line below `package main` tells compiler **which package** file belongs to
    * It is **package declaration**
    * Specifying `package main` ensures program compiles into **executable**
  * 2nd line of code imports function from another package via `import` command
    * Notice that package name is in **quotes**
    * 
```
package main 

import "fmt" 

func main () {
  fmt.Println("Hello World") 
}

```
### Main Function
* To define a function in Go use the `func` keyword
  * Followed by name of function E.g. main
  * After this we need a set of parentheses i.e. `()`
  * Any code inside the function needs to be within set of **curly braces** i.e. {}
    * Also needs to be **indented**
  * E.g.
```
func main () {
    fmt.Println("Hello World") 
}
```
* Having a main function inside `main.go` is special to Go
  * Created executable when compiled
  * Also starts executing code in here as starting point
    * Same as python
* Defining a function doesn't call it
  * Same as python

### Importing Multiple Packages
* There are two main ways to import multiple palcage
  * USe multiple import statements
  * Or use single statement with **parentheses**
    * E.g.
```
import (
  "package1"
  "package2"
)
```
* We can also give a package an alias
  * Then refer to this instead inside code
  * As we do in python using `as` keyword
  * E.g.
```
import (
  p1 "package1"
  "package2"
)
```

### Comments
* There are two different types of comment in Go
  * Single-line using `//`
    * E.g.
```
// This entire line is ignored by the compiler
// fmt.Println("Does NOT print")
fmt.Println("This gets printed!") // This part gets ignored
```
  * Block comments a.k.a multi-line using `/* */`
    * E.g.
```
/*
This is ignored.
This is also ignored. 
fmt.Println("This WON'T print!")
*/
```

### Go Resources
* Golang includes **go doc** tool for viewing documentation about packages and their functions
  * E.g.
```
$ go doc fmt.Println
package fmt // import "fmt"

func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline is
    appended. It returns the number of bytes written and any write error
    encountered.
```
## Learn Go - Variables & Types
### Literals
* Literals a.k.a values can be anything from number or text
  * Essentially unnamed things
    * i.e. not assigned to variables
* We can do normal arithmetic on integer literals
  * Or add strings together
  * **But we can't multiply string and int like we can in python**
    * i.e. print string x no. of times

### Constants
* One type of named values
* Can't be updated while program running
  * Helps conveys intent of keeping consistent value
* We use `const` keyword to create constant
* E.g.
```
const funFact = "Hummingbirds' wings can beat up to 200 times a second."

fmt.Println("Did you know?")
fmt.Println(funFact)
```
* We must use **camelCase** or **PascalCase** for constants

### Data Types
* Programming languages like Go store data as binary numbers in memory
* Go has **3 basic data types** for **numbers**
  1. **Integers** - `int`
     * Can be positive or negative
     * E.g. `22` and `-1500`
  2. **Floating-point** - `float`
     * Numbers with a decimal point
     * Can also be positive or negative
     * E.g. `1458.2` and `-1900.001`
  3. **Complex Numbers** - `complex`
     * Can also be +ve or -ve
     * Used in 2D coordinates or calculation include sq roots
     * E.g. `3i` and `-14 - -.05i`

### Basic Numeric Types in Go
* There are 15 ways to describe a number in Go
  * 2 types for floating-point
  * 11 types for integer
  * 2 types for complex 
* The different types dictate **how much memory** number takes to store
  * As well as **how many binary digits** it uses to store it
    * Less bits means fewer possible values i.e. min/max values
  * Better to use types with **smaller range** if possible
* Integers can be **signed** or **unsigned**
  * Unsigned can **only be positive**
    * Minimum value is 0
  * Signed can be either -ve or +ve
    * Max value lower than unsigned for same no. of bits vs unsigned
* Boolean type only uses **1 bit**
* Numeric Data Types Table

![Numeric Types](codecademy/images/numeric_types.png)

* Floats and complex no. **don't** have max. or min. numbers
* Floats can be `float32` or `float64`
  * Difference is how much data used to **ensure value's precision**

### What is a Variable
* Variables in Go are defined using **var keyword**
  * They also need their **type defined**
  * E.g. `var neighbourUp bool`
* Should also be in **camelCase**
* Other examples
  * `var lengthOfSong uint16` - Unsigned 16-bit integer
  * `var songRating float32` - 32-bit float

### Reading Go Errors
* When Go compiler raises error code cannot be turned into binary
  * Thus code can't be run
* An example of an error is when we define variable but **don't use it**
  * E.g. `./main.go:4:7: numberWheels declared and not used`
    * Unused variables are waste of space

### Assigning Variables
* After defining a variable we can assign it a value
  * E.g. 
```
var ipAddress
ipAddress = "192.168.1.2"
```
* We could also **define** and **assign** a value in one line
  * E.g. `var kilometersToMars int32 = 62100000`
* **Strings can't have single quotes**
  * USed for specific character called rune

### Zero Values
* Variables of different types have **default values**
  * i.e. String defaults to "", int to 0 and boolean to false
  * E.g.
```
var classTime uint32
var averageGrade float32
var teacherName string
var isPassFail bool

fmt.Println(classTime) // Prints 0
fmt.Println(averageGrade) // Prints 0
fmt.Println(teacherName) // Doesn't print anything
fmt.Println(isPassFail) // Prints false
```

### Inferring Values
* In Go we **don't** have to **specify variable type** if we use `:=` operator
  * Called short decleration operator
  * We can use this if we know what variable should hold when creating it
* E.g.
```
nuclearMeltdownOccurring := true
radiumInGroundWater := 4.521 //Type float64
daysSinceLastWorkplaceCatastrophe := 0 //Type int32/int64
externalMessage := "Everything is normal. Keep calm and carry on."
```
* Notice we don't need to use `var` keyword 
* The long way to do the above is:
```
var nuclearMeltdownOccurring = true
var radiumInGroundWater = 4.521
var daysSinceLastWorkplaceCatastrophe = 0
var externalMessage = "Everything is normal. Keep calm and carry on."
```

### Default int Type
* We can assign an integer `uint` or `int` type
* If computer architecture is 32-bit it will use int32/uint32
  * If 64-bit then int64/uint64
* Usually recommended to use either `uint` or `int`
  * Instead of specifying no. of bits
  * E.g.
```
var timesWeWereFooled int
var foolishGamesPlayed uint
```
* When we use **value inferring** with integers it defaults to `int`
  * E.g. `consolationPrizes := 2`

### Updating Variables
* As with Python we can update variables in Go
  * In below e.g. `basketTotal` var defaults to 0
    * Wouldn't be the case in Python
```
var basketTotal float64
carrotPrice := 0.75

basketTotal = basketTotal + carrotPrice
fmt.Println(basketTotal) // Prints: 0.75
```
* Same as Python we can also use `+=` above instead
  * Also works on strings
  * Others operators are: `*=`, `/=` and `-=`

### Multiple Variable Declaration
* We can declare multiple variables in **single line**
  * E.g.
```
var part1, part2 string
part1 = "To be..."
part2 = "Not to be..."
```
* Same goes for inferring also 
  * E.g. `quote, fact := "Bears, Beets, Battlestar Galactica", true`

## Learn Go: FMT Package

### The fmt Package
* fmt is one of Go's **core packages**
  * Also has utilities that help us **format** data
    * Apart from printing things
* Other supported methods
  * `fmt.Print()`
  * `fmt.Printf()`
  * The below format but doesn't print anything to the console
  * `fmt.Sprint()`
  * `fmt.Sprintln()`
  * `fmt.Sprintf()`
  * To get user input
    * * `fmt.Scan()`

### The Print Method
* The `fmt.Println()` method does some formatting for us by default
  * It adds a **space** to the arguments inside it and adds **line break** at the end
* Sometimes we may not want this behaviour
  * i.e don't want the extra space or the line break
* E.g.
```
fmt.Print("The answer is", ": ")
fmt.Print("12")
// Prints: The answer is: 12
```

### The Printf Method
* The `fmt.Printf()` method allows us to do **string interpolation**
  * Like f-strings
* Example
  * `%v` part is called **verb**
    * Specific letter after `%` tells it what exactly to put in placeholder
```
selection1 := "soup"
selection2 := "salad"
fmt.Printf("Do I want %v or %v?", selection1, selection2)
// Prints: Do I want soup or salad?
```

### Different Verbs
* Another verb is `%T` which prints out **type** of 2nd argument
  * E.g.
```
specialNum := 42
fmt.Printf("This value's type is %T.", specialNum)
// Prints: This value's type is int.

quote := "To do or not to do"
fmt.Printf("This value's type is %T.", quote)
// Prints: This value's type is string.
```
* The `%d` verb prints out a number in string format
  * E.g.
```
votingAge := 18
fmt.Printf("You must be %d years old to vote.", votingAge)
// Prints: You must be 18 years old to vote. 
```
* The `%f` verb to print out a float in string format
  * E.g.
```
gpa := 3.8
fmt.Printf("You're averaging: %f.", gpa)
// Prints: You're averaging 3.800000.
```
* We can also control **number of decimals** in float
  * E.g.
```
gpa := 3.8
fmt.Printf("You're averaging: %.2f.", gpa)
// Prints: You're averaging 3.80.
```
  
### Sprint and Sprintln
* These methods allow us to format strings **without printing** them
  * E.g.
```
grade := "100"
compliment := "Great job!"
teacherSays := fmt.Sprint("You scored a ", grade, " on the test! ", compliment)

fmt.Print(teacherSays)
// Prints: You scored a 100 on the test! Great job!
```
* As before `Sprintln` adds spacing and a new line
  * E.g.
```
quote = fmt.Sprintln("Look ma,", "no spaces!")
fmt.Print(quote) // Prints Look ma, no spaces!
```

### Sprintf Method
* Method a bit like using python f-string without print statement
  * Let's us do interpolation
* Uses same verbs as Printf
```
correctAns := "A"
answer := fmt.Sprintf("And the correct answer is… %v!", correctAns)

fmt.Print(answer) // Prints: And the correct answer is… A!
```

### Getting User Input
* We can use the `fmt.Scan` method to get input from user
  * E.g. below would take in single word from user
    * If we entered 2 words separated by space only 1st one taken
```
fmt.Println("How are you doing?") 

var response string 
fmt.Scan(&response)

fmt.Printf("I'm %v.", response) 
```
* Below example takes 2 arguments i.e. 2 words
```
fmt.Println("How are you doing?") 

var response1 string 
var response2 string 
fmt.Scan(&response1)
fmt.Scan(&response2)

fmt.Printf("I'm %v %v", response1, response2)
```
* The `&` character used references **addresses** as `fmt.Scan()` expects these
  * Without it the program will not actually ask for any input

## Conditionals

### If Statement
* Our condition can be within parenthesis but not mandatory
* Code to be executed if condition is true needs to be within **curly braces**
  * No colon needed like Python
* Booleans in Go are **all lowercase**. E.g. `true`
* E.g.
```
alarmRinging := true
if alarmRinging {
  fmt.Println("Turn off the alarm!!") 
}
```

### The Else Statement
* Else block will also need to have it's code in curly braces
  * E.g.
```
isHungry := false
if isHungry {
  fmt.Println("Eat the cookie") 
} else {
  fmt.Println("Step away from the cookie...")
}
```

### Comparison Operators
* Comparison operators work in the exact same way as Python

### Logical Operators
* Different to Python
  * Not is `!`
  * And is `&&`
  * Or is `||`
* E.g.
```
if storeLights == "on" && doorsOpen {
  fmt.Println("You can enter the store!")
} 
```

### Not Operator
* Essentially reverses value of a boolean
  * E.g.
```
bored := true
fmt.Println(!bored) // Prints false

tired := false;
fmt.Println(!tired) // Prints true
```
* We can also add them to if statements
  * E.g.
```
if !readyToGo {
		fmt.Println("Start the car!")
}
```

### Else if Statement
* Instead of elif in Python in Go we use `else if`
  * E.g.
```
position := 2

if position == 1 {
  fmt.Println("You won the gold!")
} else if position == 2 {
  fmt.Println("You got the silver medal.")
} else if position == 3 {
  fmt.Println("Great job on bronze.")
} else {
  fmt.Println("Sorry, better luck next time?")
}
```

### Switch Statement
* Instead of the verbosity of writing multiple else if statements we can use `switch` and `match`
* `switch` is followed by **variable** the conditional is looking at 
* `case` is followed by **value** it should be equal to and colon
  * Then code that should be run for condition is **indented below**
  * Can also use comparsion operators
* Finally the `default` keyword is used at the end like an **else**
* E.g.
  * Using else if
```
clothingChoice := "sweater"

if clothingChoice == "shirt" {
  fmt.Println("We have shirts in S and M only.")
} else if clothingChoice == "polos" {
  fmt.Println("We have polos in M, L, and XL.")
} else if clothingChoice == "sweater" {
  fmt.Println("We have sweaters in S, M, L, and XL.")
} else {
  fmt.Println("Sorry, we don't carry that.")
}
```
* Same example now using switch
```
clothingChoice := "sweater"

switch clothingChoice {
case "shirt":
  fmt.Println("We have shirts in S and M only.")
case "polos":
  fmt.Println("We have polos in M, L, and XL.")
case "sweater":
  fmt.Println("We have sweaters in S, M, L, and XL.")
case "jackets":
  fmt.Println("We have jackets in all sizes.")
default:
  fmt.Println("Sorry, we don't carry that")
}
// Prints: We have sweaters in S, M, L, and XL.
```

### Scoped Short Decleration Statement
* We can **declare variables** within an if or switch statement
  * Done using `:=` straight after `if` keyword
  * But before the comparison
  * Separated by `;`
  * E.g.
```
x := 8
y := 9
if product := x * y; product > 60 {
  fmt.Println(product, "  is greater than 60")
}
```
* E.g. below shows how to do it within switch statement
```
switch season := "summer" ; season {
case "summer":
  fmt.Println("Go out and enjoy the sun!")
}
```
* **Variable defined like this are only valid within the block**
  * i.e. Trying to use these after outside of the block will throw an error

### Randomising
* Go has a `math/rand` library which can generate random numbers for us
  * E.g. of usage
    * below should print a random number from 0 to 99 but it will always print 81
```
import (
  "math/rand"
  "fmt"
)

func main() {
  fmt.Println(rand.Intn(100))
}
```

### Seeding
* The reason the above code did not give us random number was due to how Go **chooses starting no.**
  * i.e. starting point for generating random number  
  * Known as seed number which defaults to 1
* To actually get random number we need to pass in **unique seed number**
  * To get seed no. we use `rand.Seed()` method
  * One way to do this is to use the **time** via `time` library
    * As it's always different
    * E.g.
```
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  // Gives us time difference since 1 Jan 1970 UTC in us thus different seed no.
  rand.Seed(time.Now().UnixNano())
  fmt.Println(rand.Intn(100))
}
```
* Essentially this line will give us random number: `rand.seed(time.Now().UnixNano())`

## Functions

### What is a function
* In Go we use `func` keyword to define a function
  * Type hints are **mandatory** unlike python
  * E.g. below shows function with input and output being ints 
    * Code also needs to be in curly braces
```
func doubleNum(num int) int {
  return num * 2
}
```
* `main` function **doesn't** have to be called as compiler already knows to do it
  * Unlike Python
* A function is called in the exact same way as python
  * E.g. `fmt.Println(doubleNum(x)) // Prints: 10`

### Scope
* There are 3 scopes for variables inside Go as with Python
  * Global
  * Within function
  * Within main function 