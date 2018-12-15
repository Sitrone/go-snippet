package gotour

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"io/ioutil"
	"log"
	"math"
	"regexp"
	"strings"
)

func Sqrt(x float64) float64 {
	if x == 0 {
		return 0
	}

	step := 0
	z := float64(1)
	for math.Abs(z*z-x) > 0.000001 {
		z -= (z*z - x) / (2 * z)
		step++
		log.Printf("step %d, value is:%v\n", step, z)
	}
	return z
}

func CopyDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	digitRegexp, _ := regexp.Compile("\\d+")
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)

	for i := range ret {
		x := make([]uint8, dx)
		for j := range x {
			switch {
			case j%15 == 0:
				x[j] = 240
			case j%3 == 0:
				x[j] = 120
			case j%5 == 0:
				x[j] = 150
			default:
				x[j] = 100
			}
		}
		ret[i] = x
	}
	return ret
}

func WordCount(s string) map[string]int {
	kvsMap := make(map[string]int)
	fields := strings.Fields(s)

	for _, f := range fields {
		if v, ok := kvsMap[f]; ok {
			kvsMap[f] = v + 1
		} else {
			kvsMap[f] = 1
		}
	}

	return kvsMap
}

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		temp := a
		a, b = b, a+b
		return temp
	}
}

type IPAddr [4]byte

func (ip *IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func SqrtWithErr(x float64) (float64, error) {
	if x < 0 {
		return -1, ErrNegativeSqrt(x)
	}

	if x == 0 {
		return 0, nil
	}

	step := 0
	z := float64(1)
	for math.Abs(z*z-x) > 0.000001 {
		z -= (z*z - x) / (2 * z)
		step++
		log.Printf("step %d, value is:%v\n", step, z)
	}
	return z, nil
}

type MyReader struct{}

func (r MyReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return
}

func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

type Image struct {
	Height, Width int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

func (i Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *Tree, ch chan int) {
	walkImpl(t, ch)
	close(ch)
}

func walkImpl(t *Tree, ch chan int) {
	if t != nil {
		walkImpl(t.Left, ch)
		ch <- t.Value
		walkImpl(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}
