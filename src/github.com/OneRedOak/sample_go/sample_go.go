package main

import "github.com/OneRedOak/newmath"

var result float64

func main() {
	result = newmath.Sqrt(4)
}

func getResult() float64 {
	return result
}