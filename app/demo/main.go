package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {

	w := md5.New()
	w.Write([]byte("123456"))

	bytes := w.Sum(nil)
	fmt.Println(bytes)

	fmt.Println(hex.EncodeToString(bytes), len(bytes))

	t := float32(101) / float32(100)

	fmt.Println(t)

}
