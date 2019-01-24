package main

import (
	"fmt"
	"time"
)

func main() {

	// Strings, which can be added together with `+`.
	fmt.Println("go" + "lang")

	// Integers and floats.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println("0.1+0.2=", 0.1+0.2)

	// Booleans, with boolean operators as you'd expect.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	s := make([]string, 3)
	fmt.Println("emp:", len(s), cap(s), s[0] == "", s[1], s[2])

	// We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` returns the length of the slice as expected.
	fmt.Println("len:", len(s))

	// In addition to these basic operations, slices
	// support several more that make them richer than
	// arrays. One is the builtin `append`, which
	// returns a slice containing one or more new values.
	// Note that we need to accept a return value from
	// `append` as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can also be `copy`'d. Here we create an
	// empty slice `c` of the same length as `s` and copy
	// into `c` from `s`.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	sliceFromLiteral()
}

func sliceFromLiteral() {
	fmt.Printf("\n\n** working as expected: **\n")
	// this doesnt work
	i := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("initial slice: ", i)
	j := append(i, 100)
	g := append(i, 101)
	h := append(i, 102)
	fmt.Printf("i: %v\nj: %v\ng: %v\nh: %v\n", i, j, g, h)
}
