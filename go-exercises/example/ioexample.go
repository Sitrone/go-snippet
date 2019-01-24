package example

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func readFields(f string, target int) bool {
	bytes, e := ioutil.ReadFile(f)
	if e != nil {
		panic(e)
	}

	sTarget := strconv.Itoa(target)
	fields := strings.Fields(string(bytes))
	for _, e := range fields[:] {
		if e == sTarget {
			return true
		}
	}
	return false
}
