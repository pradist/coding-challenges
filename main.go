package main

import (
	"github.com/pradist/coding-challenges/ccwc"
)

func main() {
	i, err := ccwc.CountLines("testdata/example.txt")
	if err != nil {
		panic(err)
	}
	println("Number of lines:", i)
}
