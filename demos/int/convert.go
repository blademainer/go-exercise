package main

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("string cast: ", string(199))
	fmt.Println("string cast: ", string(1))
	fmt.Println("strconv: ", strconv.FormatInt(199, 10))
	fmt.Println(binary.BigEndian.Uint16([]byte{1, 255}))
	fmt.Println(binary.LittleEndian.Uint16([]byte{1, 255}))
}
