package main

import (
	"fmt"
	"strconv"
)

func forFunc() {
	for i := 0; i < 10; i++ {
		fmt.Print(strconv.Itoa(i) + " ")
	}
}

const (
	E = 0.0001
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for z*z-x > E || z*z-x < -E {
		z = (z + x/z) / 2
	}
	return z
}

func main() {
	forFunc()
}
