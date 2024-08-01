package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileName := os.Args[1]

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	v, err := io.Copy(os.Stdout, f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
}