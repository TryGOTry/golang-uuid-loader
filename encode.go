package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)
func Readcode(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("出现错误了。", err)
	}
	return string(data)
}
var (
	codes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	codeLen = len(codes)
)
func RandNewStr(len int) string {  //随机生成deskey
	data := make([]byte, len)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		idx := rand.Intn(codeLen)
		data[i] = byte(codes[idx])
	}
	return string(data)
}

func AESEncrypt(src []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(GenerateKey(key))
	length := (len(src) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, src)
	pad := byte(len(plain) - len(src))
	for i := len(src); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(src); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

func GenerateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("请输入加密的文件名称!")
	} else {
		filename := os.Args[1]
		aeskey := RandNewStr(16)
		code := Readcode(filename)
		newcode := AESEncrypt([]byte(code),[]byte(aeskey))
		ncode := hex.EncodeToString([]byte(newcode)) //hex编码
		fmt.Println("key:",aeskey)
		fmt.Println("code:",ncode)
	}
}
