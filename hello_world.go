package main

// double quotes are important
// fmt is formatting library provided by golang, lets us print to terminal and shit
import (
	"fmt"
	"math"

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

	//variables
	var a = "hello variable"
	fmt.Println(a)

	var b int = 2
	fmt.Println(b)

	var c, d int = 3, 4
	fmt.Println(c + d)

	var e = true
	fmt.Println(e)

	var f int
	fmt.Println(f)

	g := "initialize"
	fmt.Println(g)

	g = "reassigning variable"
	fmt.Println(g)

	//constants
	const s string = "constant"
	const n = 500000000
	const z = 3e20 / n
	fmt.Println(z)

	fmt.Println(int64(z))

	fmt.Println(math.Sin(n))

}
