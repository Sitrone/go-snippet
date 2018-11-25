package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	go mockData("../test.log")

	time.Sleep(20 * time.Second)
}

func mockData(path string) {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		// 文件存在则删除
		log.Println(path + "not exist")
		if err := os.Remove(path); err != nil {
			panic(fmt.Sprintf("delete file err:%s", err.Error()))
		}
	}

	cf, err := os.Create(path)
	if err != nil {
		panic(fmt.Sprintf("create file err: %s", err.Error()))
	}

	log.Println("create file success:", cf.Name())

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
