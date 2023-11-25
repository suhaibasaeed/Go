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
