package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var r io.Reader

	tty, err := os.OpenFile("text.txt", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	b1 := make([]byte, 11)
	n1, err := tty.Read(b1)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))
	tty.Seek(0, 0) // set offset for next Read to 0

	r = tty
	b2 := make([]byte, 11)
	n2, err := r.Read(b2)
	fmt.Printf("%d bytes: %s\n", n2, string(b2))

	var w io.Writer
	w = r.(io.Writer)

	d1 := []byte("Hello")
	w.Write(d1)

	var empty interface{}
	empty = w // empty cannot call any method
	fmt.Printf("empty: %v\n", empty)
}
