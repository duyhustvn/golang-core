package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sync"
)

func processChunk(buf []byte) {
	// TODO handle chunk here
}

func main() {
	fileName := ""
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file %s", fileName)
	}

	bufioReader := bufio.NewReader(f)

	var wg sync.WaitGroup
	eof := false
	for !eof {
		// Read 4KB
		buf := make([]byte, 4*1024)
		n, err := bufioReader.Read(buf)
		if err != nil {
			switch err {
			case io.EOF:
				eof = true
			default:
				continue
			}
		}

		if n == 0 {
			continue
		}

		// Keep reading until \n characters
		line, err := bufioReader.ReadBytes('\n')
		if err != io.EOF {
			buf = append(buf, line...)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			processChunk(buf)
		}()
	}

	wg.Wait()
}
