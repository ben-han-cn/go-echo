package main

import (
	"flag"
	"log"
	"net"
	"runtime"

	"github.com/ben-han-cn/cement/signal"
	"github.com/libp2p/go-reuseport"
)

func main() {
	var addr string
	flag.StringVar(&addr, "addr", "127.0.0.1:53", "address to listen")
	flag.Parse()

	recevie_count := runtime.NumCPU()
	for i := 0; i < recevie_count; i++ {
		conn, err := reuseport.ListenPacket("udp", addr)
		if err != nil {
			log.Fatalf("invalid addr:%s", err.Error())
		}
		go echo(conn)
	}
	signal.WaitForInterrupt(func() {
		log.Printf("done\n")
	})
}

func echo(conn net.PacketConn) {
	buf := make([]byte, 512)
	for {
		n, addr, err := conn.ReadFrom(buf)
		if err == nil && n > 0 {
			conn.WriteTo(buf[:n], addr)
		}
	}
}
