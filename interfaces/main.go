package main

import (
	"fmt"
	"strconv"
)
	/////////////// Notes and Best practices
	//// Accepting Interfaces (make functions that take interfaces parameters(abstract), and return structs(concrete))
	////	- liberal with what we accept
	////	-	not forcing a user to use a specific concrete type
	////	- increase flexibility
	////  - increase number of user implementations
	//// Best Practices
	////  - Use many, small interfaces
	////  - Don't export interfaces for types that will be consumed
	////  - Do    export interfaces for types that will be used by package
	////  - Design functions and methods to receive interfaces whenever possible
	//// interfaces promote behavoirs over entities, because entities may share behaviors

func main()  {
	fmt.Println("Konnichiwa :" +strconv.Itoa(3))

	/////////////// Basics
	// p := person{name: "Cailyn", age: 44, isMarried: false}
	// c := cat{}
	// d := duck{}
	// fmt.Println(p)	// uses String() method
	// SpeakStation(p)		// for this to work ; you must provide implementations for the whole interface method set, not just the methods used inside SpeakStation function
	// SpeakStation(c)
	// SpeakStation(d)
	// fmt.Println("------------------------------------")
	/////////////// Empty Interface
	//// all types implement the empty interface{}.
	// emptyDetails(p)
	// emptyDetails(c)
	// emptyDetails(5)
	// emptyDetails("Any thing, Any type")


	///////////////// Type Conversion
	//// v := s.(T) :- asserts that the interface value s is of type T and assigns the underlying value of v --> if s is of type T, assign it to v with the type T, the 2nd returned value is a flag that is false if the conversion fails
	// var i interface{} = 7
	/// strconv.Itoa(i) // we cannot use type conversion because the types are not compatible with type conversion
	// iInt, isValid := i.(int)	// return 2 things, the value converted and conversion success flag
	// fmt.Println("did the conversion work? ", isValid)
	// strconv.Itoa(iInt)		
	// switch v:= i.(type){			// switching on the type and putting the value in v to use directly in the cases
	// 	case int: fmt.Println(v + 7)     // try "7" and it will fail, because inside this case v is an int, and string "7" can't be added to int v
	// 	case string: fmt.Println(v + "7")// try 7 and it will fail, because inside this case v is a string, and num 7 can't be concatenated to string v
	// }

}

type Speaker interface { 
	Speak() string
	// can have more methods signatures
} 

type person struct {
  name string
  age int
	isMarried bool
}
type cat struct {}
type duck struct {}

func (p person) String() string {
  return fmt.Sprintf("%v (%v years old).\nMarried status: %v ", p.name, p.age, p.isMarried)
}
func (p person) Speak()string{
	return p.name + ": I am a person... I am a speaker"
}	
func (c cat) Speak()string{	// can just put the type without the c variable identifier, because we won't use it anyway
	return "cat: Purr... Meow"
}	
func (duck) Speak()string{
	return "duck: Quack... Quack"
}

func SpeakStation(spkr Speaker) {
	fmt.Println("")
	fmt.Println("a speaker will speak now ...")
	fmt.Println(spkr.Speak())
}

func emptyDetails(s interface{}) { 
	fmt.Printf("(%#v)\n", s) 
}