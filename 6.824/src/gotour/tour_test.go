package gotour

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestSqrt(t *testing.T) {
	fmt.Println(Sqrt(2))
}

func TestPic(t *testing.T) {
	Show(Pic)
}

func TestWordCount(t *testing.T) {
	Test(WordCount)
}
func TestFibonacci(t *testing.T) {
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func TestInterface(t *testing.T) {
	testInterface()
}

func TestIPAddr_String(t *testing.T) {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, &ip)
	}
}

func TestSqrtWithErr(t *testing.T) {
	fmt.Println(SqrtWithErr(2))
	fmt.Println(SqrtWithErr(-2))
}

func TestMyReader_Read(t *testing.T) {
	Validate(MyReader{})
}

func TestRot13Reader_Read(t *testing.T) {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func TestImage_At(t *testing.T) {
	m := Image{256, 256}
	ShowImage(m)
}

func TestWalk(t *testing.T) {
	ch := make(chan int)
	go Walk(New(1), ch)
	for e := range ch {
		fmt.Println(e)
	}
}

func TestSame(t *testing.T) {
	if Same(New(1), New(1)) {
		fmt.Println("same PASSED")
	} else {
		fmt.Println("same FAILED")
	}
}

func TestCrawl(t *testing.T) {
	Crawl("https://golang.org/", 4, fetcher)
}
