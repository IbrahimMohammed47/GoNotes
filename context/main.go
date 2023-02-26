package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//////// very simple put & get
	parentCtx := context.Background()
	ctx := context.WithValue(parentCtx, "key", "test-value")  // put 
	val := ctx.Value("key").(string) // get and cast to string
	fmt.Println(val, time.Second)


	///////// cancellation option and done
	// ctx, cancel := context.WithCancel(context.Background())
	// go func(){
	// 	<-time.After(time.Second * 2)
	// 	cancel()
	// }()		
	// <- ctx.Done() 
	// fmt.Println("cancel was called, which emitted a Done signal")


	//////// deadlines and Err()
	// ctx1, _ := context.WithCancel(context.Background())
	// ctx2, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second * 2))
	// ctx3, cancelCtx3 := context.WithDeadline(context.Background(), time.Now().Add(time.Second * 2))
	// deadline, ok := ctx1.Deadline()
	// fmt.Println(deadline, ok)					// ok = false because no deadlines were set on this context
	// deadline, ok = ctx2.Deadline()
	// fmt.Println(deadline, ok)
	// deadline, ok = ctx3.Deadline()
	// fmt.Println(deadline, ok)
	
	// fmt.Println(ctx1.Err())			// nil  ==> ctx was cancelled or ctx didn't have deadline or ctx had a deadline that is not reached yet
	// cancelCtx3()
	// fmt.Println(ctx3.Err())     // ctx cancelled
	// <- time.After(time.Second*3)
	// fmt.Println(ctx2.Err())	    // ctx deadline exceeded



	/////// tree of contexts: when parent ctx is cancelled, all children ctxs are cancelled 
	// parentCtx, cancel := context.WithCancel(context.Background())
	// childCtx1, _ := context.WithCancel(parentCtx)
	// childCtx2, _ := context.WithCancel(parentCtx)
	// grandChildCtx1, _ := context.WithCancel(childCtx1)
	// fmt.Println(childCtx1.Err(), time.Second)
	// cancel()
	// fmt.Println(parentCtx.Err())
	// fmt.Println(childCtx1.Err())
	// fmt.Println(childCtx2.Err())
	// fmt.Println(grandChildCtx1.Err())
}
