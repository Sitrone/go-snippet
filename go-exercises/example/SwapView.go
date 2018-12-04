package example

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

/** go语言实现的
各个语言的SwapView 性能：
https://github.com/lilydjwg/swapview
 */

var (
	nullBytes  = []byte{0x0}
	emptyBytes = []byte(" ")
	units      = []string{"", "K", "M", "G", "T", "P"}
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
)

type SwapView struct {
	Pid  int
	Size int64
	Comm string
}

func GetSwap() chan *SwapView {
	rootFile, e := os.Open("/proc")
	if e != nil {
		panic(fmt.Sprintf("failed to open /proc, msg:%s", e.Error()))
	}
	defer rootFile.Close()
	names, e := rootFile.Readdirnames(0)
	if e != nil {
		panic(fmt.Sprintf("failed to readDirnames /proc, msg:%s", e.Error()))
	}

	ret := make(chan *SwapView)
	var wg sync.WaitGroup
	wg.Add(len(names))
	for _, name := range names {
		pid, err := strconv.Atoi(name)
		if err != nil {
			continue
		}
		go getInfo(ret, wg, pid)
	}

	go func() {
		wg.Wait()
		close(ret)
	}()
	return ret
}

func getInfo(c chan *SwapView, wg sync.WaitGroup, pid int) {
	defer wg.Done()
	swapView := &SwapView{}

	swapView.Pid = pid

	var bs []byte
	bs, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
	if err != nil {
		return
	}
	if bytes.HasSuffix(bs, nullBytes) {
		bs = bs[:len(bs)-1]
	}
	swapView.Comm = string(bytes.Replace(bs, nullBytes, emptyBytes, -1))
	file, err := os.Open(fmt.Sprintf("/proc/%d/smaps", pid))
	if err != nil {
		return
	}

	var total int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		if strings.HasPrefix(line, "Swap:") {
			arr := strings.Split(line, " ")
			size, err := strconv.ParseInt(arr[len(arr)-2], 10, 0)
			if err != nil {
				continue
			}
			total += size
		}
	}

	swapView.Size = total * 1024
	c <- swapView
}

func FormatSize(s int64) string {
	unit := 0

	f := float64(s)
	for f >= KB && unit < len(units) {
		f = f / KB
		unit++
	}
	//switch {
	//case s < KB:
	//	unit = 0
	//case s >= KB && s < MB:
	//	unit = 1
	//case s >= MB && s < GB:
	//	unit = 2
	//case s >= GB && s < TB:
	//	unit = 3
	//case s >= TB:
	//	unit = 4
	//}

	if unit == 0 {
		return fmt.Sprintf("%dB", s)
	} else {
		return fmt.Sprintf("%.2f%sB", f, units[unit])
	}
}