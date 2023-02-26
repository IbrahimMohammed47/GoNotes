package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

	////// basics
	// ch := make(chan int) //choosing it as the datatype that will flow through the channel
	// ch_signal := make(chan struct{})
	// go func() {
	// 	i := <-ch      // blocks until receiving int from channel ch into a variable i
	// 	fmt.Println(i) // NOTE: order of receiving is not guaranteed because we didn't control order of sender goroutines
	// 	i = <-ch       // blocks until receiving another int from channel ch into a variable i
	// 	fmt.Println(i)
	// 	ch_signal <- struct{}{}
	// }()
	// go func() {
	// 	ch <- 42 // transmitting 42 through  channel ch, this line BLOCKS until somebody receives its msg
	// }()
	// go func() {
	// 	ch <- 41 // transmitting 41 through  channel ch
	// }()
	// <-ch_signal //waits for the goroutine to receive the signal from the channel ch

	///// channel polymorphism to 2 types
	// ch := make(chan int) //choosing it as the datatype that will flow through the channel
	// ch_signal := make(chan struct{})
	// go func(rcvr_channel <-chan int) { // takes a receive-only channel, if you used channel for transmitting, u'll get an error
	// 	i, ok := <-rcvr_channel // ,ok flags whether the channel is open or close
	// 	fmt.Println(i)
	// 	fmt.Println(ok)
	// 	ch_signal <- struct{}{}
	// }(ch)
	// go func(sndr_channel chan<- int) { // takes a send-only channel, if you used channel for receiving, u'll get an error
	// 	sndr_channel <- 7
	// }(ch)
	// <-ch_signal

	///// buffer channels
	// wg = sync.WaitGroup{}
	// wg.Add(2)
	// ch := make(chan int, 2)            // 2 is the buffer size, blocking happends only when the buffer is full
	// go func(rcvr_channel <-chan int) { // takes a receive-only channel, if you used channel for transmitting, u'll get an error
	//// this works, but there is a better way for buffer receivers
	// i, ok := <-rcvr_channel
	// fmt.Println(i, ok)
	// i, ok = <-rcvr_channel
	// fmt.Println(i, ok)
	// i, ok = <-rcvr_channel
	// fmt.Println(i, ok)

	//// this is better
	// for i := range rcvr_channel { // Note the syntax changes, ranges iterates with ONE variable, not 2
	// 	fmt.Println(i)
	// }
	// wg.Done()
	// }(ch)
	// go func(sndr_channel chan<- int) { // takes a send-only channel, if you used channel for receiving, u'll get an error
	// 	sndr_channel <- 7
	// 	println("sndr_channel <- 7")
	// 	sndr_channel <- 3 // this won't block because its msg is buffered
	// 	println("sndr_channel <- 3")
	// 	close(sndr_channel)
	// 	wg.Done()
	// }(ch)
	// wg.Wait()

	///// channel behaviors
	// c := make(chan int, 2) // a buffered channel
	// c <- 3
	// c <- 5
	// // close(c)
	// fmt.Println(len(c), cap(c)) // 2 2
	// x, ok := <-c
	// fmt.Println(x, ok)          // 3 true
	// fmt.Println(len(c), cap(c)) // 1 2
	// x, ok = <-c
	// fmt.Println(x, ok)          // 5 true
	// fmt.Println(len(c), cap(c)) // 0 2
	// go func(c chan int) { fmt.Println("closing the channel"); close(c) }(c)
	// fmt.Println("here")
	// x, ok = <-c // if the channel wasn't closed, this will be blocked until it receives a value, and if there are no one to close it, deadlock happens !
	// fmt.Println("there")
	// fmt.Println(x, ok) // 0 false
	// x, ok = <-c
	// fmt.Println(x, ok)          // 0 false
	// fmt.Println(len(c), cap(c)) // 0 2
	// // close(c)        // panic, because it's already closed!
	// // c <- 7          // panic, can't send to a closed channel!

	// footballgame()

	// /// nil-&-check idiom to break when all receive channels are closed
	// input1 := make(chan int)
	// input2 := make(chan int)

	// go func() {
	// 	<-time.After(time.Second)
	// 	close(input1)
	// 	close(input2)
	// }()

	// for {
	// 	select {
	// 	case _, ok := <-input1:
	// 		if !ok {
	// 			println("nilling 1")
	// 			input1 = nil
	// 		}
	// 	case _, ok := <-input2:
	// 		if !ok {
	// 			println("nilling 2")
	// 			input2 = nil
	// 		}
	// 	}
	// 	println("loop")
	// 	if input1 == nil && input2 == nil {
	// 		break
	// 	}
	// }
	///////////////////// GRACEFUL CLOSING

	////// check if channel if closed without blocking execution
	// c := make(chan int)
	// fmt.Println("1", isClosed(c))
	// close(c)
	// fmt.Println("2", isClosed(c))

	////// M receivers, one sender, the sender says "no more sends" by closing the data channel
	// sndrRecvrs()

	////// One receiver, N senders, the only receiver says "please stop sending more" by CLOSING an additional signal channel
	// sndrsRcvr()

	////// M receivers, N senders, any one of them says "let's end the game" by notifying a moderator to close an additional signal channel
	// sndrsRecvrs()

	// main := make(chan uint64, 1000)
	// kid1 := make(chan uint64)
	// kid2 := make(chan uint64)
	// kid3 := make(chan uint64)
	// for i := 0; i < 1000; i++ {
	// 	main <- uint64(i)
	// }
	// divisor(main, kid1, kid2, kid3)

	// wg.Add(3)
	// go func (){drain(kid1, 1); wg.Done()}()
	// go func (){drain(kid2, 2); wg.Done()}()
	// go func (){drain(kid3, 3); wg.Done()}()
	// close(main)
	// wg.Wait()

}

func drain(kid chan uint64, id int) {
	for e := range kid {
		fmt.Println("kid", id, ":", e)
	}
}

func filter(input <-chan uint64, predicate func(uint64) bool) <-chan uint64 {
	output := make(chan uint64)
	go func() {
		for e := range input {
			if predicate(e) {
				output <- e
			}
		}
		close(output)
	}()
	return output
}

func isClosed(c chan int) bool {
	select {
	case <-c:
		return true
	default:
	}
	return false
}

func sndrRecvrs() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumReceivers = 100

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)

	// the sender
	go func() {
		var value int
		for {
			if value = rand.Intn(Max); value == 0 {
				// The only sender can close the
				// channel at any time safely.
				close(dataCh)
				return
			}
			dataCh <- value

		}
	}()

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			// Receive values until dataCh is
			// closed and the value buffer queue
			// of dataCh becomes empty.
			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}

func sndrsRcvr() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	// ...
	dataCh := make(chan int)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the receiver of channel
	// dataCh, and its receivers are the
	// senders of channel dataCh.

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() { // each sender will send infinite number of rand ints, once a stopCh channel is closed, ALL of them will return
			for {
				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(Max):
				}
			}
		}()
	}

	// the receiver
	go func() {
		defer wgReceivers.Done()

		for value := range dataCh {
			if value == Max-1 {
				// The receiver of channel dataCh is
				// also the sender of stopCh. It is
				// safe to close the stop channel here.
				close(stopCh)
				return
			}

			log.Println(value)
		}
	}()

	// ...
	wgReceivers.Wait()
}

func sndrsRecvrs() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumReceivers = 10
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the moderator goroutine shown
	// below, and its receivers are all senders
	// and receivers of dataCh.
	toStop := make(chan string, 1)
	// The channel toStop is used to notify the
	// moderator to close the additional signal
	// channel (stopCh). Its senders are any senders
	// and receivers of dataCh, and its receiver is
	// the moderator goroutine shown below.
	// It must be a buffered channel.

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					// Here, the try-send operation is
					// to notify the moderator to close
					// the additional signal channel.
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}

				// The try-receive operation here is to
				// try to exit the sender goroutine as
				// early as possible. Try-receive and
				// try-send select blocks are specially
				// optimized by the standard Go
				// compiler, so they are very efficient.
				select {
				case <-stopCh:
					return
				default:
				}

				// Even if stopCh is closed, the first
				// branch in this select block might be
				// still not selected for some loops
				// (and for ever in theory) if the send
				// to dataCh is also non-blocking. If
				// this is unacceptable, then the above
				// try-receive operation is essential.
				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()

			for {
				// Same as the sender goroutine, the
				// try-receive operation here is to
				// try to exit the receiver goroutine
				// as early as possible.
				select {
				case <-stopCh:
					return
				default:
				}

				// Even if stopCh is closed, the first
				// branch in this select block might be
				// still not selected for some loops
				// (and forever in theory) if the receive
				// from dataCh is also non-blocking. If
				// this is not acceptable, then the above
				// try-receive operation is essential.
				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == Max-1 {
						// Here, the same trick is
						// used to notify the moderator
						// to close the additional
						// signal channel.
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}

					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	// ...
	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}

// =========================================================================================== just for fun
func footballgame() {
	var ball = make(chan string)
	kickBall := func(playerName string) {
		for {
			fmt.Println(<-ball, "kicked the ball.")
			time.Sleep(time.Second)
			ball <- playerName
		}
	}
	go kickBall("John")
	go kickBall("Alice")
	go kickBall("Bob")
	go kickBall("Emily")
	ball <- "referee"      // kick off
	var wall chan struct{} // the nil channel can be used as a wall/signal
	<-wall                 // blocking here for ever
}
