package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	path := flag.String("file", "../test.log", "target file path")
	flag.Parse()

	checkPath(path)

	r := &ReadFromFile{
		path: *path,
	}

	c := make(chan []byte)

	go r.Read(c)

	for ele := range c {
		log.Println(string(ele))
	}
}

type Reader interface {
	Read(rc chan []byte)
}

type ReadFromFile struct {
	path string // 读取文件的路径
}

func (r *ReadFromFile) Read(rc chan []byte) {
	// 读取模块

	// 打开文件
	f, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file error:%s", err.Error()))
	}

	// 从文件末尾开始逐行读取文件内容
	f.Seek(0, 2)
	rd := bufio.NewReader(f)

	for {
		line, err := rd.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("ReadBytes error:%s", err.Error()))
		}
		rc <- line[:len(line)-1]
	}
}

func checkPath(path *string) {
	if _, err := os.Stat(*path); os.IsNotExist(err) {
		if err != nil {
			panic(fmt.Sprintf("create file err: %s", err.Error()))
		}
	}
}
