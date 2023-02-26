package main

import (
	"fmt"
)

func init() {
	fmt.Println("init() is called before main")
}
func main() {
	fmt.Println("Hola")

	////////////////////// if conditions
	// if 1 == 1 {
	// 	fmt.Println("1 is equal to 1")
	// } else if 2 == 2 {
	// 	fmt.Println("2 is equal to 2")
	// } else {
	// 	fmt.Println("1 is not equal to 1")
	// }

	// if ok := true; ok {
	// 	fmt.Println("ok is:", ok)
	// }

	////////////////////// for loops
	//// for <initializer>;<test>;<incrementer> {}		// simple for loop (every part is optional)
	//// for <test> {}									// while loop
	//// for {}											// infinite loop, we can excape it with a break statement
	//// for <index>, <value> := range <collection> {}	// range loop (NOTE: second value is a copy of the element, not a pointer to it, so use it only if copying is cheap)
	//// exiting early: break, continue, labels

	////////////////////// switch
	// i := 10
	// switch {
	// case i <= 10, i < 11:
	// 	fmt.Println(i, "is less-or-equal 10")
	// 	fallthrough // switch statements break once the have a match, to continue after that WITHOUT matching use fallthrough
	// case i <= 20:
	// 	fmt.Println(i, "is less-or-equal 20")
	// default:
	// 	fmt.Println(i, "is less than infinity ")
	// }

	//// type switch
	// 	var i interface{} = 1.4		/// interface{} is a type that can hold any type of data: primitive, pointer, complex..etc
	// 	switch i.(type) {
	// 		case int: fmt.Println(i, "is an int")
	// 		case float64: fmt.Println(i, "is a float64")
	// 		case string: fmt.Println(i, "is a string")
	// 		default: fmt.Println(i, "is another type")
	// 	}

	////////////////////// for range
	// myArr := [3]string{"a","b","c"}
	// mySlice := []string{"x", "y", "z"}
	// myMap := map[string]int{"x":1,"y":2}	// NOTE, key,value pairs order is undeterministic
	// for _, elem := range myArr {
	// 	fmt.Println(elem)
	// }
	// for idx, elem := range mySlice {
	// 	fmt.Println(idx, elem)
	// }
	// for key, value := range myMap {
	// 	fmt.Println(key, value)
	// }
}
