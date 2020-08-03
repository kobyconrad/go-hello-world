package main

// double quotes are important
// fmt is formatting library provided by golang, lets us print to terminal and shit
import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// access the Println function from the fmt library
	// functions from a library start with a capital letter
	fmt.Println("Hello World this is my first go program :)")
	fmt.Println(uuid.New().String())

	// values
	fmt.Println("string " + "concatination")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
