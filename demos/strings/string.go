package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

var words = "你好，世界！"

func Rune() {
	b := []byte(words)
	runes := bytes.Runes(b)
	fmt.Println(runes)
}

func Sum(){
	//h := md5.New()
	//h.Write([]byte(words)) // 需要加密的字符串为 123456
	//cipherStr := h.Sum(nil)
	cipherStr := md5.Sum([]byte(words))
	fmt.Printf("%s\n", hex.EncodeToString(cipherStr[:]))
}

func main() {
	Rune()
	Sum()
}
