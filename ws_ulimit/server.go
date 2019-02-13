package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"syscall"

	"github.com/gorilla/websocket"
)

var count int64

func printConnections(cnt int64) {
	log.Printf("Total number of connections: %v", cnt)
}
func ws(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	new := atomic.AddInt64(&count, 1)
	if new%100 == 0 {
		printConnections(new)
	}
	defer func() {
		new := atomic.AddInt64(&count, -1)
		if new%100 == 0 {
			printConnections(new)
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		log.Printf("msg: %s", string(msg))
	}
}

func main() {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	fmt.Printf("Set Limit %d -> %d\n", rLimit.Cur, rLimit.Max)
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}

	go func() {
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			log.Fatalf("Pprof failed: %v", err)
		}
	}()

	http.HandleFunc("/", ws)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
