package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

var (
	mockFile string
	duration int
)

func main() {
	runtime.GOMAXPROCS(1)
	flag.StringVar(&mockFile, "path", "../test.log", "log file path")
	flag.IntVar(&duration, "duration", 20, "generate data duration")
	flag.Parse()

	go GenerateMockData(mockFile)

	time.Sleep(20 * time.Second)
}

func GenerateMockData(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		cf, err := os.Create(path)
		if err != nil {
			panic(fmt.Sprintf("create file err: %s", err.Error()))
		}
		log.Println("create file success:", cf.Name())

	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("Open file err: %s", err.Error()))
	}

	defer file.Close()

	for {
		line := fmt.Sprintf("%d\n", time.Now().UnixNano())
		file.Write([]byte(line))
		log.Println(line)
		time.Sleep(100 * time.Millisecond)
	}
}
