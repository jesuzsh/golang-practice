package main

import (
	"os"
)

func main() {
	f, _ := os.Create("./file.txt")
	b := []byte{115, 111, 109}
	f.Write(b)
	f.Close()
}
