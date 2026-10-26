package main

import "fmt"

var d int;
var e string="how you doing"
func main(){
	// var abc string="hello there"
	// var abc2 = 123
	// abc3:=45

	// fmt.Println(abc)
	// fmt.Println(abc2)
	// fmt.Println(abc3)

	// just intialising the variables
	var a string
	var b int
	var c bool
	d=1000
	fmt.Println(a," ",b," ",c," ",d," ",e)

	// mutliple declaring
	var f,g,h,e int= 1,2,3,4
	ab,cd:= "hello",23
	fmt.Println(f,g,h,e,ab,cd)

	var (
		a1 int
		a2 string="happy"
		a3 int=4326
	)
	fmt.Println(a1,a2,a3)
}