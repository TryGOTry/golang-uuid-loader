package main

import (
	"demo2/root"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("请输入加密的文件名称!")
	} else {
		filename := os.Args[1]
		aeskey :=root.RandNewStr(16)
		code := root.Readcode(filename)
		newcode := root.AESEncrypt([]byte(code),[]byte(aeskey))
		ncode := hex.EncodeToString([]byte(newcode)) //hex编码
		fmt.Println("key:",aeskey)
		fmt.Println("code:",ncode)
	}
}
