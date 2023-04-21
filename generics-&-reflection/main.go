package main

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"strings"

	"golang.org/x/exp/constraints"
)

// Generic func
func SumIntsOrFloats[V int | float64](slice []V) V { // V is a type parameter, (int | float64) is a type constraint
	var s V
	for _, v := range slice {
		s += v
	}
	return s
}

// Generic func
func Min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// Generic func
func Push[S ~[]E, E any](slice *S, elem E) {
	*slice = append(*slice, elem)
}

// Generic Type
type Tree[T constraints.Ordered] struct {
	left, right *Tree[T]
	data        T
}

// Generic Type Method
func (t *Tree[T]) Contains(x T) bool {
	if t == nil {
		return false
	}
	return t.data == x || t.left.Contains(x) || t.right.Contains(x)
}

// Bad Generic Func -> WHY?
// ex: type Point = []int32
// func (p Point) String() string
// the resulting slice won't have the String method(won't be of type Point, it will be a slice)
func Scale[E constraints.Integer](slice []E, factor E) []E {
	r := make([]E, len(slice))
	for i, v := range slice {
		r[i] = v * factor
	}
	return r
}

// Good Generic Func
// ex: type Point = []int32
// func (p Point) String() string
// the resulting slice will have the String method -> because ~[]E means substitute all types with underlying type []E
func Scale2[E constraints.Integer, S ~[]E](slice S, factor E) S {
	r := make(S, len(slice))
	for i, v := range slice {
		r[i] = v * factor
	}
	return r
}

func main() {

	//////// Generics
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}
	println(SumIntsOrFloats(s1))
	println(SumIntsOrFloats(s2))

	println(Min[int](5, 7)) // here we pass type arguments
	println(Min(5, 7))      // here type arguments got implicitly inferred <- So we will use this syntax
	println(Min("asdasgf", "asd"))
	println(Min(2.4, 5.11))

	intMin := Min[int] // instantiating Min function without calling it, intMin is non-generic min function
	println(intMin(5, 7))
	// println(intMin(1.5, 7.4)) // error

	s3 := []int{}
	println(len(s3))
	Push(&s3, 6)
	println(len(s3))

	rightTree := Tree[string]{nil, nil, "foo"}
	leftTree := Tree[string]{nil, nil, "boo"}
	parentTree := Tree[string]{left: &leftTree, right: &rightTree, data: "zoo"}
	println(rightTree.Contains("boo"))
	println(leftTree.Contains("boo"))
	println(parentTree.Contains("boo"))

	//////// Reflection
	var w io.Writer = os.Stdout
	// f := w.(*os.File)      // successful downcasting
	// println(f.Name())
	// c := w.(*bytes.Buffer) // panic
	/// Another way is using 2 variable to not panic
	c, ok := w.(*bytes.Buffer) // panic
	if !ok {
		println(c == nil)
		println("w is not bytes.Buffer")
	}

	//// Reflection package
	a := struct {
		x string
		y []int
	}{
		x: "foo",
		y: []int{1, 2, 3},
	}
	b := struct {
		x string
		y []int
	}{
		x: "foo",
		y: []int{1, 2, 3},
	}
	println(reflect.DeepEqual(a, b))

	//// Reflection switch (like pattern matching)
	var x interface{} = 5 // try int and string
	switch v := x.(type) {
	case string:
		println(strings.ToUpper(v))
	case int:
		println(v + 5)
	}
}
