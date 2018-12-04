package example

import (
	"math"
)

type vertex struct {
	X, Y float64
}

func Abs(v vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//func main() {
//	v := vertex{3, 4}
//	Scale(&v, 10)
//	fmt.Println(Abs(v))
//}

