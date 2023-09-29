package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}

	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}

	return p
}

func startNetworkDaemonP() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		connPool := warmServiceConnCache()

		server, err := net.Listen("tcp", "localhost:8000")
		if err != nil {
			log.Fatalf("Cannot listen: %v", err)
		}
		defer server.Close()
		wg.Done()
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("Cannot accept connection %v", err)
				continue
			}
			svcConn := connPool.Get()
			fmt.Fprintln(conn, "")
			connPool.Put(svcConn)
			conn.Close()
		}
	}()
	return &wg
}

func connectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		server, err := net.Listen("tcp", "localhost:8000")
		if err != nil {
			log.Fatalf("Cannot listen: %v", err)
		}
		defer server.Close()

		wg.Done()
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("Cannot accept connection %v", err)
				continue
			}
			connectToService()
			fmt.Fprintln(conn, "")
			conn.Close()
		}
	}()
	return &wg
}

func init() {
	// daemonStarted := startNetworkDaemon()
	daemonStarted := startNetworkDaemonP()
	daemonStarted.Wait()
}

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8000")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}

		conn.Close()
	}
}
