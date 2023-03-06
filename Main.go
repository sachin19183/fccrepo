package main

import "fmt"

var I int      // declared in uppercase first letter, globally visible to any program that uses this pacakage
var i int = 11 // declared in lowercase first letter, locally visible within the package

//l := 4     this fomrat of variable declaration is not valid for package level declaration
var ( //syntax for declaring a block of variables
	fname string = "Sachin"
	lname string = "Sharma"
	age   int    = 38
)

func main() {
	fmt.Println("Hello GO!!!")
	//var declaration formats
	fmt.Println(" test commit")
	fmt.Println(" Here i declared at pacakage level will be printed:", i)
	var i int // this is block scoped .not visible outside this block
	i = 1
	var j int = 2
	k := 3
	fmt.Print(i, j, k)            //here i declared at the function level will be printed hence preceding the package level declaration. This is called shadowing
	fmt.Printf("\n %v %T ", i, i) //%T shows what is the data type of variable is
}
