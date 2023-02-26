package main

import (
	"fmt"
	"greet"	// this package directory lives in $GOPATH, which is ~/go/src
)
func init()  {
	fmt.Println("main stuff initialized...")
}

func main()  {
	greet.Greet("ibrahim")
	// greet.helper()	// won't work because `helper` function is lowercased, which means it can't be exported from the greet package
}