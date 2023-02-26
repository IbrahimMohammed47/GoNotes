package main

import (
	"encoding/json"
	"fmt"
)

// /// NOTES
// // cpbv types:primitives & arrays & structs containing only primitive types(no maps,no slices, no pointers) are copied by value
// // cpbr types: slices & maps & pointers are copied by reference
// // to deep-copy a cpbv type: just do simple assignment
// // to deep-copy a cpbr type: copy element-wise, pack operators with append, make & copy funcs...etc
type pet struct {
	name string
	age  int
}
type user struct {
	nickname string
	pet      pet
}
type user2 struct {
	nickname string
	pet      *pet
}

func main() {
	fmt.Println("Chao")
	//////////////////////// Arrays
	//// defining and accessing element with index
	// var zeros [5]int
	// zeros2 := [5]int{0,0,0,0,0}
	// fruits := [3]string{"apple", "orange", "banana"}
	// pokemons := [...]string{"pikachu","bulbasor","charmander","squirtle","snorlax"}  // still an array, but the compiler figures out the length of the array value
	// arrWithKeys := [10]int{1, 9: 10, 4: 5} // 1st element is one, 9th index has 10, 4th index has 5  // keys can be used out of order
	// fmt.Println(zeros)
	// fmt.Println(fruits)
	// fmt.Println(pokemons)
	// fmt.Println(zeros == zeros2) // Arrays are compared element-wise
	// zeros2[4] = 1
	// fmt.Println(zeros == zeros2)
	// fmt.Println(arrWithKeys)
	//// Note, array is a value type, once a value type is copied, it's not connected to it's source
	//// copying is deep, it copies the elements
	// p1 := pet{
	// 	name: "Pikachu",
	// }
	// u1 := user{
	// 	nickname: "Tracy",
	// 	pet:      p1,
	// }
	// arr1 := [1]user{u1} // Make this a slice and compare the results
	// arr2 := arr1
	// arr2[0].pet.name = "Bulbasor"
	// fmt.Println(arr1)
	// fmt.Println(arr2)

	//// looping
	// for i := 0; i < len(pokemons); i++ {
	// 	fmt.Print(pokemons[i]+", ")
	// }
	// for _, v := range pokemons {
	// 	fmt.Print(v+", ")
	// }
	// fmt.Println()

	//////////////////////// Slices
	/// slice is essentially an array with a thin layer around it that provides more flexibility
	/// In real-world code, you should be using slices as your go-to for all sorted collections.
	// fruits1 := []string{"apple", "orange", "banana"}
	// fruits2 := append(fruits1, "kiwi") // append 1 element
	// fruits3 := []string{"strawberry", "blueberry"}
	// fruits3 = append(fruits3, "guava", "grapes") // append multiple element
	// fruits3 = append(fruits3, fruits2...)        // append another slice
	// fmt.Println(fruits1)                         // that's functional !
	// fmt.Println(fruits2)
	// fmt.Println(fruits3)
	// fmt.Println("First: ", fruits3[0], fruits3[0:1], fruits3[:1]) // element, slice, prefix-slice
	// fmt.Println("Last: ", fruits3[7], fruits3[7:8], fruits3[7:])  // element, slice, suffix-slice
	// fmt.Println("Middle: ", fruits3[2:5])                         // slice

	//// array is a value type, once a value type is copied, it's not connected to it's source
	//// slice doesn't work like value type, it works like a pointer (but it's NOT a pointer) <--- VERY IMPORTANT, another note: map will act like slice not array
	//// len: size of the slice --- cap: size of the underlying array
	//// appending this element leaded to creating a new underlying array with double size (16), all elements are copied from
	////		the old array to the new array and adds the new element to it too, note that a slice has pointer that points at the current underlying array
	//// Go's built-in make function allows you to set the length and capacity of a slice when creating it. The syntax looks like make(<sliceType>, <length>, <capacity>)
	// fruits3 = append(fruits3, "peach")
	// fmt.Println("length:",len(fruits3),", capacity:", cap(fruits3))
	// fruits3BadCopy := fruits3 // bad: only a pointer is created, both slices are pointing to the same array
	// fmt.Println("copyWithLink   -> Before changing original:", fruits3BadCopy)
	// fruits3[4] = "BOOM"
	// fmt.Println("copyWithLink   -> After  changing original:", fruits3BadCopy)
	// fruits3GoodCopy := make([]string, len(fruits3))
	// copy(fruits3GoodCopy, fruits3) // good: using copy, it requires getting the length, it uses make & copy
	// fmt.Println("copyWithNoLink -> Before changing original:", fruits3GoodCopy)
	// fruits3[5] = "BLA"
	// fmt.Println("copyWithNoLink -> After  changing original:", fruits3GoodCopy)
	// fruits3GoodCopy2 := append([]string{}, fruits3...) // excellent: using append  <-- this is the cleanest and most memory-efficient way to copy a slice.
	// fruits3[6] = "BLOB"
	// fmt.Println("copyWithNoLink2 -> After  changing original:", fruits3GoodCopy2)

	/// Slicing Notes
	// a := [...]int{1, 2, 3} // array
	// b := a[0:1]            // slice out first element from a
	// c := b[0:2]            // slice out 2 elements from b
	// fmt.Println(a, len(a), cap(a))
	// fmt.Println(b, len(b), cap(b))
	// fmt.Println(c, len(c), cap(c))
	// NOTE, in the previous example, the underlying array is a, slices b will have SAME capacity of a, and c will have the same too ,that's how Go works
	// that's why slicing c out of b didn't cause an error
	// to avoid that you can do the folowing x[0:3:3] <-- the 3rd parameter is the capacity
	// a := [...]int{1, 2, 3} // array
	// b := a[0:1:1]          // slice out first element from a
	// fmt.Println(b, len(b), cap(b))
	// c := b[0:2] // error: slice bounds out of range [:2] with capacity 1
	// fmt.Println(c, len(c), cap(c))
	// If we append to b something that exceeds the capacity of a, Go will create a new underlying array with double size (6) and copy all elements from a to it
	// This is forcing reallocation, now appending to b won't affect a

	//// Look at the following example
	// items := [][2]int{{1, 2}, {3, 4}, {5, 6}}
	// a := [][]int{}
	// for _, item := range items {
	// 	fmt.Printf("Address of item=%d:\t%p\n", item, &item)
	// 	a = append(a, item[:]) // the reference is passed, not the value
	// }
	// fmt.Println(a)
	// // The problem is that 'a' kept a reference to a loop variable 'item' that's being reused in each iteration
	// // the solution is to COPY the value of the loop variable 'item' into a new variable 'itemCopy' and then pass the reference to 'itemCopy'
	// a = [][]int{}
	// for _, item := range items {
	// 	itemCopy := item
	// 	fmt.Printf("Address of item=%d:\t%p\n", itemCopy, &itemCopy)
	// 	a = append(a, itemCopy[:])
	// }
	// fmt.Println(a)
	/// Lesson: DON'T PASS REFERENCES TO MUTATING LOOP VARIABLES
	/// note that the reference can be passed explicitly using the '&' operator, but it also can be passed implicitly if the passed variable is a slice for example

	//////////////////////// Maps
	//// Note: unlike arrays, and just like slices, maps are copied by reference not by value, changing the original or the copy will affect the other one
	// phoneNums := map[string]string{
	// 	"ibrahim": "+201521",
	// 	"Saa'd":   "+206152",
	// }
	// fmt.Println(phoneNums["ibrahim"])
	// person1phone, exists1 := phoneNums["Saa'd"] // map read returns 2 things
	// person2phone, exists2 := phoneNums["Khaled"]
	// if exists1 {
	// 	fmt.Println("1", person1phone)
	// }
	// if exists2 {
	// 	fmt.Println("2", person2phone)
	// }
	// delete(phoneNums, "ibrahim")
	// fmt.Println(phoneNums["ibrahim"])
	// var m1 map[string]int      // nil, no storage allocated
	// m2 := make(map[string]int) // non-nil, but empty
	// a, ok1 := m1["a"]
	// b, ok2 := m2["b"]
	// fmt.Println(a, ok1)
	// fmt.Println(b, ok2)
	// m2["one"] = 1
	// m1["one"] = 1 // Panic: assignment to entry in nil map
	////////
	// petStore := map[string]pet{}
	// petStore["dog"] = pet{name: "spike"}
	// fmt.Println(&petStore["dog"]) // NOTE: addresses in maps can't be accessed, BECAUSE are hashtables, so the addresses inside can rearrange
	// petStore["dog"].age += 1 // NOTE: this won't work too, you can't mutate the value of struct, because accessing the value of a map returns a COPY of the value
	// petStore2 := map[string]*pet{}
	// petStore2["dog"] = &pet{name: "spike"}
	// petStore2["dog"].age += 2 // NOTE: this won't work too,
	/// Instead you can do the following
	//////////////////////// Custom types & Structs
	//// custom type id
	// type id string
	// var id1 id
	// var id2 id = "1234-5678"
	// var id3 id
	// id3 = "1234-5678"
	// fmt.Println("id1 == id2    :", id1 == id2)
	// fmt.Println("id2 == id3    :", id2 == id3)
	// fmt.Println("id2 == \"1234-5678\":", id2 == "1234-5678")

	////// structs are value types: copied by value when they are passed to functions, assigned to new vars...etc ===> So if your struct is very large consider passing it by reference with &
	//// struct type user, usually defined at the package scope not the fumction scope
	// type user struct {
	// 	name    string
	// 	age     int
	// 	balance float64
	// 	member  bool
	// }
	// u1 := user{ 	// If you wanna use values only, take care of the order
	// 	balance: 98.43,
	// 	name:  "Tracy",
	// 	member: true,
	// }
	// fmt.Printf("%#v\n", u1) // notice that any unprovided key will have its own type zero value, ex: age:0
	// u2P := new(user)	// pointer to struct object
	// fmt.Printf("%#v\n", u2P)
	// fmt.Printf("%#v\n", *u2P)

	// type point struct{
	// 	x int
	// 	y int
	// }
	/// anonymous struct: they have 2 conscecutive blocks, one for content types, one for content values
	// point1 := struct {
	//   x int
	//   y int
	// }{
	//   x:10,
	//   y:10,
	// }
	// fmt.Printf("%#v", point1); fmt.Println()
	// point2 := struct {
	//   x int
	//   y int
	// }{}	// it is anonymous, it is initialized, but with zero values
	// point2.x = 10
	// point2.y = 5
	// point3 := point{10,5}
	// fmt.Println(point1==point2, point2==point3)

	/// embedded struct
	// type name string
	// type location struct {
	// 	x int
	// 	y int
	// }
	// type size struct {
	// 	width  int
	// 	height int
	// }
	// type dot struct {
	// 	name
	// 	location
	// 	size
	// }
	// dot1 := dot{
	// 	name: "B",
	// 	location: location{
	// 		x: 13,
	// 		y: 27,
	// 	},
	// 	size: size{
	// 		width:  5,
	// 		height: 7,
	// 	},
	// }
	// fmt.Println(dot1.name, dot1.location.x, dot1.size.height)
	// fmt.Println(dot1.x) // notice here that x lies under location not under dot1, however GO is smart enough to find it (it's composotion, but acsts as inheritence)
	////// Notes on structs
	//// 2 structs are compatible if the fields have the same name and type, in the same order, and the same tags
	//// compatible means that you can convert one to the other
	//// 2 structs objects are comparible if the struct fields are comparable

	////// Structs tags
	//// tags are used to add metadata to struct fields, ex: they are used by encoding/json package to encode/decode json
	type user struct {
		Name string `json:"name"`
		Age  int    `json:"age,omitempty"`
	}
	u1 := user{
		Name: "Tracy",
		Age:  10,
	}
	u2 := user{
		Name: "X",
	}
	fmt.Printf("%#v\n", u1)
	fmt.Printf("%#v\n", u2)
	j1, _ := json.Marshal(u1)
	fmt.Println(string(j1))
	j2, _ := json.Marshal(u2)
	fmt.Println(string(j2))

	var v1 user
	json.Unmarshal(j1, &v1)
	fmt.Printf("%#v\n", v1)
}
