package main

import (
	"fmt"
	// "math/rand"
	// "time"
	// "strconv"
)

/////// NOTES
/////// In C++, a function that makes a local variable and returns a pointer to it is UNSAFE, because the stack frame of this function will be thrown in trash as soon as the function finishes execution, therefore that local variable will be cleared as well. as a result the returned pointer is now pointing at ... idk.. nothing? trash?
/////// In Go, it's safe, because the compiler automatically takes that variable out from the stack frame of the function, and places it in the heap memory

var shadow string = "sh"

func main() {
	fmt.Println("Hello")

	/////////////////////////////////////////////////////// Variables
	//// ways to declare vars
	// var i int
	// i = 3
	// var j int = 5
	// k := 7
	// var (
	// 	age  int    = 21
	// 	name string = "ibrahim"
	// )
	// x, y, z := "x", 4, true
	// println(i, j, k)
	// println(name, age)
	// println(x, y, z)

	//// shadowing
	// var shadow string = "shhhhhhhh"
	// println(shadow)

	//// variable scoping
	// package scoping
	// -lower case vars are scoped to the package, any file in the same package can access that variable
	// -upper case vars are globally visible and can be exported by others
	// block scoping: variables declared inside blocks (like function blocks, or any {...} block) are only visible INSIDE the block
	// When accessing a variable, Go looks at the scope the code was defined in. If it can't find a variable with that name, it looks in the parent scope,..., until it gets to the package scope.

	//// printing types
	// a := 5
	// b := 3.14
	// fmt.Printf("%v is of type %T\n", a, a)    // %T prints the type of the variable
	// fmt.Printf("%[1]v is of type %[1]T\n", b) // %v prints the value of the variable, %[1] means the first argument

	//// conversion
	// b := 1.21
	// a := int(b) // converted float64 to int, and lost precision (type as a function)
	// println(a)
	// num := 47
	// println(string(num)) // it converts the unicode 47 to the corresponding character, which is '/'
	// println(strconv.Itoa(num) + " hi")

	//// rand
	// var seed = 1234456789			// seed in int32
	// rand.Seed(int64(seed))		// .Seed(s) accepts int64 only, so we either convert seed to int64, or declare the variable with the type int64
	// rand.Seed(time.Now().UnixNano())
	// println(rand.Intn(100))	  // random number from 1 to 100

	//// constants
	// only numbers, strings, and booleans can be constants,
	// because they are immutable, they can be used in place of literals
	// const s = "hello"
	// fmt.Printf("%T ", s)

	// strings
	// strings are immutable sequence of, you can't change a string, you can only create a new one
	// strings share the underlying storage, so if you have a string variable, and you assign a slice of it to another variable, they will share the same underlying storage
	// logically a sequence of runes (Unicode), int32's, -2^31-2^31-1
	// physically a sequence of bytes (UTF-8 encoded), uint8's, 0-255
	// st := "ثanks"
	// fmt.Printf("%T %[1]v \n", st)
	// fmt.Printf("%T %[1]v \n", []rune(st))             // []rune is a slice of runes, which is an alias for int32
	// fmt.Printf("Rune Length = %v\n", len([]rune(st))) // length of the string in runes, not bytes
	// fmt.Printf("%T %[1]v \n", []byte(st))             // []byte is a slice of bytes, which is an alias for uint8
	// fmt.Printf("Byte Length = %v\n", len(st))         // length of the string in bytes, not runes

	//// By value --VS-- By pointer
	// -primitives are copied by value when passed as arguments to functions
	// -parameters can be marked to be passed using a pointer, passed-by-pointer values are put on the heap not the stack
	// -The heap allows the value to exist until no part of your software has a pointer to it anymore.
	// -unlike C++, heap values and datastructures are collected automatically by the garbage collector,
	// -This garbage collection happens periodically in the background and reclaims heap values allocated memory that no part of your software has a pointer to them anymore
	// -Working out whether a value needs to be put on the heap is called escape analysis.
	// -You have no direct control over whether a value is put on the stack or the heap. Memory management is not part of Go's language specification.
	// - which is better?
	//		- When a value gets copied, Go needs CPU cycles to get that memory and then release it later. Using a pointer avoids this CPU usage when passing it to a function.
	//		- However, having a value on the heap means that it then needs to be managed by the complex garbage collection process. This process can become a CPU bottleneck in certain situations,
	//			for example, if there are lots of values on the heap. When this happens, the garbage collector has to do lots of checking, which uses up CPU cycles. There is no correct answer here,
	// - Best approach is the classic performance optimization one:
	//		Don't prematurely optimize. When you do have a performance problem, measure before you make a change, and then measure after you've made a change.

	//// pointers (extract a pointer with &, dereference a pointer with *)
	/// making pointers
	// var count1 *int			  // Declare a pointer using a var statement:
	// count2 := new(int)		// Create a variable using new:
	// countTemp := 5				// You can't take the address of a literal number (can't do this: count3 = &5 ). Create a temporary variable to hold a number
	// count3 := &countTemp  // Using &, create a pointer from the existing variable:
	// fmt.Printf("count1: %#v\n", count1)  // has value nil because it points at nothing
	// fmt.Printf("count2: %#v\n", count2)  // #v prints (type)(value)
	// fmt.Printf("count3: %#v\n", count3)
	/// dereferencing pointers
	// if count1 != nil {
	// 	fmt.Printf("count1: %#v\n", *count1)
	// }
	// if count2 != nil {
	// 	fmt.Printf("count2: %#v\n", *count2)
	// }
	// if count3 != nil {
	// 	fmt.Printf("count3: %#v\n", *count3)
	// }
	// var count int = 7
	// add5Value(count) 												// pass by value
	// fmt.Println("add5Value post:", count)   // origial unchanged !
	// add5Point(&count)												// pass the pointer
	// fmt.Println("add5Point post:", count)		// original changed  !
	// a, b := 5, 10
	// swap(&a, &b)
	// fmt.Println(a == 10, b == 5)

	//// map, make, constants
	// cache = make(map[string]string)  // make built-in function allocates and initializes an object of type slice, map, or chan (only). Like new, the first argument is a type, not a value. Unlike new, it return the typed value itself not the pointer of it
	// SetBook("1234-5678","AlMajarayat")
	// fmt.Println(GetBook("1234-5678"))

	//// enums
	// const (
	// 	MaxLongitude float32 = 90.0
	// 	MinLongitude float32 = -89.9
	// )
	// const (
	// 	Sunday = (iota + 1) % 7		// iota gives 0
	// 	Monday
	// 	Tuesday
	// 	Wednesday
	// 	Thursday
	// 	Friday
	// 	Saturday
	// )
	// fmt.Println(MaxLongitude, MinLongitude)
	// fmt.Println(Saturday,Sunday,Monday,Tuesday,Wednesday,Thursday,Friday)

}

const GlobalLimit = 100
const MaxCacheSize int = 10 * GlobalLimit

var cache map[string]string // declared in the package scope, initialized in th main function
func cacheGet(key string) string {
	return cache[key]
}
func cacheSet(key, val string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = val
}
func GetBook(isbn string) string {
	return cacheGet("Book-" + isbn)
}
func SetBook(isbn string, name string) {
	cacheSet("Book-"+isbn, name)
}

func add5Value(count int) {
	count += 5
	fmt.Println("add5Value   :", count)
}
func add5Point(count *int) {
	*count += 5
	fmt.Println("add5Point   :", *count)
}
func swap(a *int, b *int) {
	*a += *b
	*b = *a - *b
	*a -= *b
}
