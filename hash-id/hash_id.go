package main

import (
	"log"
	"reflect"

	"github.com/speps/go-hashids"
)

// https://hashids.org/go/
// https://github.com/speps/go-hashids
// 通过id来生成指定长度的唯一字符串，字符串的字母支持可定制，可逆
func main() {
	hdata := hashids.NewData()
	hdata.Alphabet = "PleasAkMEFoThStx"
	hdata.Salt = "this is my salt"

	hid, _ := hashids.NewWithData(hdata)

	numbers := []int{45, 434, 1313, 99}
	hash, err := hid.Encode(numbers)
	if err != nil {
		log.Fatal(err)
	}
	dec := hid.Decode(hash)

	log.Printf("%v -> %v -> %v", numbers, hash, dec)

	if !reflect.DeepEqual(dec, numbers) {
		log.Fatalf("Decoded numbers `%v` did not match with original `%v`", dec, numbers)
	}
}
