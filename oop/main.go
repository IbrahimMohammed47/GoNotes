package main

type ponger interface {
	pong()
}

type english struct{}
type french struct{}

// pong is a method of english
func (p english) pong() { // p is called a "receiver"
	println("ha?")
}

// pong is a method of french too
func (p french) pong() {
	println("eu?")
}

func ping(p ponger) {
	p.pong()
}

type myInt int

func (i myInt) add(j int) int {
	return int(i) + j
}

type Point struct {
	x, y float64
}

// passes pointer of receiver
func (p *Point) Move(x, y float64) {
	p.x += x
	p.y += y
}

// passes copy of the receiver
func (p Point) Move2(x, y float64) Point {
	return Point{p.x + x, p.y + y}
}

type Eater interface {
	Eat(food string)
}
type Breather interface {
	Breathe(air string)
}
type LivingThing interface {
	Eater
	Breather
}
type Person struct{}

func (p *Person) Eat(food string) {
	println("yum")
}
func (p *Person) Breathe(air string) {
	println("hmm")
}

type Human struct {
	name string
}

type Hero struct {
	Human
	power string
}

func (h *Human) Eat(food string) {
	println("yummy")
}

func (h *Human) Drink(drink string) {
	println("glug")
}

func (h *Hero) Drink(drink string) {
	println("glug like a hero")
}

func Call(h Human) {
	println("yes sir")
}

func (h Human) CallMethod() {
	println("yes sir 2")
}

type Callable interface {
	CallMethod()
}

func Call2(c Callable) {
	c.CallMethod()
}

func main() {
	// 1: polymorphism
	// var e ponger = english{}
	// f := french{}
	// ping(e)
	// ping(f)

	// 2: methods on non-struct types
	// var i myInt = 1
	// println(i.add(2))

	// 3: receiver is either a value or a pointer
	// p := Point{1, 2}
	// p.Move(3, 4)
	// p2 := Point{1, 2}
	// p2 = p2.Move2(3, 4)
	// fmt.Println(p)
	// fmt.Println(p2)
	//// Note: Nothing in Go prevents calling a method on a nil receiver

	// 4: interfaces composition
	// p1 := Person{}
	// var l LivingThing = &p1
	// l.Eat("pizza")
	// l.Breathe("air")

	// 5: struct composition and fields/methods promotion
	/// fields of embedded struct(Human) are promoted to the embedding struct(Hero)
	// man := Human{"Peter Parker"}
	// spiderman := Hero{man, "shoots webs"} // Note: man is copied
	// man.name = "Bruce Wayne"
	// println(spiderman.Human.name) // ok
	// println(spiderman.name)       // ok too
	// println(spiderman.power)
	// spiderman.Eat("pizza")

	// 6: if method is not on struct then go looks for it on the embedded struct (promoted method)
	// spiderman := Hero{Human{"Peter Parker"}, "shoots webs"}
	// spiderman.Drink("water")

	// 7: Composition doesn't give polymorphism (it's not inheritence, and no interfaces are involved)
	// man := Human{"Peter Parker"}
	// spiderman := Hero{man, "shoots webs"}
	// Call(man)       // ok
	// Call(spiderman) // error: because this composition doesn't give polymorphism, interface do

	// 8: How to get polymorphism with composition -> use interfaces
	// man := Human{"Peter Parker"}
	// spiderman := Hero{man, "shoots webs"}
	// Call2(man)       // ok, because man has method CallMethod()
	// Call2(spiderman) // ok, because spiderman has method CallMethod() by PROMOTION

	// 9: Nil interface values
	//// Interfaces have 2 parts:
	//// 1. a value or a pointer of some type
	//// 2. a pointer to actual object
	//// interface -> (type, ptr) -> object
	//// If both are nil, then the interface is nil
	// var r io.Reader     // nil interface value
	// println(r == nil)   // true, because r -> (nil, nil) ->
	// var b *bytes.Buffer // nil pointer
	// r = b               // r now is NOT nil, but it has a nil pointer to a Buffer
	// println(r == nil)   // false, because r -> (Buffer, nil) ->
	// var r Callable
	// var b Human // empty struct is nil value for struct
	// println(r == nil)
	// r = b
	// println(r == nil)
	// r.CallMethod()
	//// Note: when making a method that returns error, make 'error' (the interface) as the return type

	// 10: Empty interfaces
	//// An empty interface is an interface with no methods
	//// so it is satisfied by anything
	//// interface{} = any type, like "any" in TypeScript or *void in C/C++
	//// Reflection would be used to determine the concrete type of any object that was passed via an empty interface parameter

	// 11: Method values
	//// A method value is a function that has a receiver, it can stored in a variable so it closes over the receiver
	//// just like cyrrying a function closes over a parameter
	// x := Point{1, 2}
	// moveFromX := x.Move2 // the receiver value is copied here
	// fmt.Println(moveFromX(1, 1))
	// x = Point{6, 7}
	// fmt.Println(moveFromX(1, 1)) // if the receiver was a pointer, then the value would be changed

}
