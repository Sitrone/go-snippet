package cryptor

import (
	"encoding/binary"
	"log"
	"math/rand"

	crand "crypto/rand"
)

const (
	defaultCharacters = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

/**
 * n长度的安全随机字符串，从characters中选择字符
 */
func RandStringWithSelectChs(n int, characters string) string {
	bytes := GenRandomBytes(n)
	for i, v := range bytes {
		bytes[i] = characters[v%byte(len(characters))]
	}
	return string(bytes)
}

//RandString 生成指定长度的secure字符串
func RandString(n int) string {
	bytes := GenRandomBytes(n)
	for i, v := range bytes {
		bytes[i] = defaultCharacters[v%byte(len(defaultCharacters))]
	}
	return string(bytes)
}

func GenRandomBytes(n int) []byte {
	token := make([]byte, n)
	_, err := crand.Read(token)
	if err != nil {
		log.Printf("failed to generate secure rand bytes, err=%s\n", err)
		rand.Read(token) // 非安全rand兜底
	}

	return token
}

// security source ref:https://stackoverflow.com/questions/35203635/golang-cryptographic-shuffle
type CryptoRandSource struct{}

func NewCryptoRandSource() CryptoRandSource {
	return CryptoRandSource{}
}

func (_ CryptoRandSource) Int63() int64 {
	var b [8]byte
	crand.Read(b[:])
	// mask off sign bit to ensure positive number
	return int64(binary.LittleEndian.Uint64(b[:]) & (1<<63 - 1))
}

func (_ CryptoRandSource) Seed(_ int64) {}
