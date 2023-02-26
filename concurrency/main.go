package main

import (
	"fmt"
	"log"
	"sync"
	"time"
	// "runtime"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Print("Salut ", time.Now().Weekday(), " ")
	log.Print()
	//// Doing it wrong
	// msg := "hello"
	// go func() {
	// 	fmt.Println(msg) // main goroutines finished and terminated before this goroutine runs
	// }()

	//// Doing it right but stupid
	// msg:= "hello"
	// go func() {
	//   fmt.Println(msg)
	// }()
	// time.Sleep(100 * time.Millisecond)	// sending the main goroutine to sleep allows the schedueler to start executing the other goroutine

	//// passing variable by closure, stupid
	// msg := "hello"
	// go func() {
	// 	fmt.Println(msg)
	// }()
	// msg = "bye"
	// time.Sleep(100 * time.Millisecond)

	//// passing variable by parameter, fine, but sleeping is still stupid
	// msg := "hello"
	// go func(m string) {
	// 	fmt.Println(m)
	// }(msg)
	// msg = "bye"
	// time.Sleep(100 * time.Millisecond)

	//// Using a waitgroup for synchronization, finally something fair
	// msg := "hello"
	// wg.Add(1)
	// go func(m string) {
	// 	fmt.Println(m)
	// 	wg.Done() // it decrements wg by 1
	// }(msg)
	// msg = "bye"
	// wg.Wait() // this is like a counting semaphore, it blocks main gouroutine until the wg counter drops back to 0

	//// maxProcs
	// fmt.Println("OS Threads available:", runtime.GOMAXPROCS(-1)) // by default:  # of threads = # of cores on a machine, -1 means return
	// runtime.GOMAXPROCS(1)                                        // making your program a single-threaded
	// fmt.Println("OS Threads available:", runtime.GOMAXPROCS(-1))
	// runtime.GOMAXPROCS(100)
	// fmt.Println("OS Threads available:", runtime.GOMAXPROCS(-1))

}
