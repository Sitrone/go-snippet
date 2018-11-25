package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

//利用go程和管道，实现Linux的tailf功能

func main() {
	path := "../test.log"

	c := make(chan string)
	go tailfile(path, c)
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	time.Sleep(20 * time.Second)
}

func tailfile(path string, c chan string) {
	f, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("open file error:%s", err.Error()))
	}

	defer f.Close()

	f.Seek(0, 2)
	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("ReadBytes error:%s", err.Error()))
		}

		//fmt.Println(string(line))
		c <- string(line[:len(line)-1])
	}
}
