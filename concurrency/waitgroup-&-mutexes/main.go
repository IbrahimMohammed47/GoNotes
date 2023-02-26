package main
import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var mtx = sync.RWMutex{}
/////////////////// race condition on accessing the counter 
// func main()  {
// 	for i := 0; i < 20; i++ {
// 		wg.Add(2)
// 		go sayHello()
// 		go increment()		
// 	}
// 	wg.Wait()
// }

// func sayHello()  {
// 	fmt.Println("hello #", counter)	
// 	wg.Done()
// }

// func increment()  {
// 	counter++
// 	wg.Done()
// }



/////////////////// synchronizing go routines with a mutex (however it accidently became sequential again, look at the mutex, it locks a lot of stuff, it even makes a worse performance that a simple sequential program)
func main()  {
	for i := 0; i < 20; i++ {
		wg.Add(2)
		mtx.RLock()					
		go sayHello()
		mtx.Lock()
		go increment()		
	}
	wg.Wait()
}

func sayHello()  {
	// mtx.RLock()		// moved this to main, to make the mutex locks happen in a SINGLE contex
	fmt.Println("hello #", counter)	
	mtx.RUnlock()
	wg.Done()
}

func increment()  {
	// mtx.Lock()			// moved this to main, to make the mutex locks happen in a SINGLE contex
	counter++
	mtx.Unlock()
	wg.Done()
}