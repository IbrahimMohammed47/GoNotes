package main

import (
	"fmt"
	"math"
)

// ////////// Go deals with functions as 1st class citizens
// ////////// you can use function decomposition, function factories, currying... and all those functional stuff
// ////////// Go has closures
// ////////// func is a type, so you can create custom types out of it (crtl+f: intGetter)
// ////////// in Go, there are 2 ways to deal with errors:
// //////////  1- error object, make a function returns an error object so the caller can check it and decide what to do with it.
// //////////  2- panic, this is like an exception throwing, it will stop the execution of the current function, call any stacked defers and terminate
// //////////	The 1st way is a common pattern in Go, it prevents the called functions (like std library funcs) from deciding anything about the program flow,
// //////////    it just return a non-nil error object to the caller (our code) and we decide to panic or not for example (this is like Inversion Of Control)
// ////////// One of the main differences compared to try/catch blocks is the way control flows. In a typical try/catch scenario, the code after the catch
// //////////   block will run unless it propagates the error. This is not so with panic/recover. A panic aborts the current function and begins to unwind
// //////////   the stack, running deferred functions (the only place recover does anything) as it encounters them.
func main() {
	fmt.Println("Marhaba ", math.Pi)

	//// functuins have multiple return values
	// ageI, ageS := getAge(1998, 2020)
	// fmt.Println(ageI, ageS)
	// msg, err := message()
	// if err == nil {
	// 	fmt.Println(msg)
	// } else {
	// 	fmt.Println("err:", err)
	// }
	// nums("James", 99, 100)

	//// Anonymous Functions
	// func() { // IIFE
	// 	fmt.Println("Greeting")
	// }()
	// mapF := func(inSlice []float32, f func(float32) float32) (outSlice []float32) {
	// 	outSlice = make([]float32, len(inSlice))
	// 	for idx, e := range inSlice {
	// 		outSlice[idx] = f(e)
	// 	}
	// 	return
	// }
	// area := func(r float32) float32 { return r * r * math.Pi }
	// radii := []float32{1, 2, 3, 4, 5}
	// areas := mapF(radii, area)
	// fmt.Println(areas)

	//// Closures
	// incrementor := getIncrementor(4)
	// fmt.Println(incrementor()) //0
	// fmt.Println(incrementor()) //1
	// fmt.Println(incrementor()) //2
	// fmt.Println(incrementor()) //3
	// fmt.Println(incrementor()) //0

	//// defer: Defers the execution of a function until the surrounding function returns.
	///					The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	///         For a single function block, defers are executed in LIFO order
	///         defers are usually used to claim used resources, that's why defers are executed in LIFO order, ex:  take port -> open connection --> close connection -> leave port
	// defer fmt.Println("world")
	// fmt.Println("hello")
	// x := func(y int) {
	// 	fmt.Println("Deferred: ", y)
	// 	return
	// }
	// y := func(y int) int {
	// 	fmt.Println("Evaluated: ", y)
	// 	return y
	// }
	// defer x(y(5 + 6))
	// fmt.Println("hello")

	//// panic represents throwing exceptions in Go.
	// defer fmt.Println("i will run anyway")	// note any defers are guaranteed to run BEFORE panic panics
	// panic("something is wrong...")

	//// recover allows a program to manage behavior of a panicking goroutine. the return value from recover reports whether the goroutine is panicking or not. 
	//// panic is usually called in a defered execution
	///  It's somehow similar to catching an exception
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

// /// it is idiomatic in Go for the second return type to be of type error
func getAge(birthYear int, currYear int) (int, string) {
	age := currYear - birthYear
	if age <= 18 {
		return age, "adolescent"
	}
	return age, "adult"
}
func message() (message string, err error) { /// naming the returns creates local variables for them that are ready to be assigned with values
	message = "hi"
	// message = "invalid"
	if message == "invalid" {
		err = fmt.Errorf("invalid msg")
		return /// called naked return: if you name the returns, you don't have to explicitly write them in the return statements, you just to assign values to them, otherwise their zero values wil be returned
	}
	return
}

// // A variadic function is a function that accepts a variable number of argument values.
func nums(str string, ns ...int) {
	fmt.Println(str, ns)
}

// // Closure
type intGetter func() int

func getIncrementor(limit int) intGetter {
	i := -1
	return func() int {
		i = (i + 1) % limit
		return i
	}
}

func panicker() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A panicker wanted to say:", err)
			//// we still can re-panic() in here, this will terminate the program and you won't see "end" printed out
		}
	}()
	panic("\"i want to panic now, but before that let's execute the stacked defers\"")
}
