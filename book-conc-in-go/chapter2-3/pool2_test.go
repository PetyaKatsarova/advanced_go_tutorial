package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

func connectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func startNetworkDeamon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func ()  {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {	log.Fatalf("cant listen: %v", err) }
		defer server.Close()
		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
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
	daemonStarted := startNetworkDeamon() // returns &wg
	daemonStarted.Wait()
}

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil { b.Fatalf("cant dial host: %v", err)}
		if _, err := io.ReadAll(conn); err != nil {
			b.Fatalf("cant read: %v", err)
		}
		conn.Close()
	}
}