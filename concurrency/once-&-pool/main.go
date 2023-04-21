package main

import (
	"bytes"
	"sync"
)

type Singleton struct{}

var sing *Singleton
var once sync.Once

func main() {
	// ONCE
	// for i := 0; i < 1000; i++ {
	// 	go func() {
	// 		// once.DO guarantees that the passed function is executed only once no matter how many concuurent goroutines try to call it
	// 		once.Do(initialize)
	// 		// (checking if sing == nil, and if yes call it) is NOT safe
	// 	}()
	// }

	// POOL
	//// a pool provides fpr efficient & safe reuse of objects, but it's a container of interfaces
	var bufPool = sync.Pool{
		New: func() any {
			return new(bytes.Buffer)
		},
	}
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	// write stuff to it, then put it back
	bufPool.Put(b)
}

func initialize() {
	println("once")
	sing = new(Singleton)
}
