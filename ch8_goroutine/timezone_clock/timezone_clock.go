// Clock1 is a TCP server that periodically writes the time.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var p = flag.String("port", "", "port")

// killall timezone_clock to terminate all processes
// killall timezone_clock && go build ./ch8_goroutine/timezone_clock && TZ=US/Eastern ./timezone_clock -port 8010 & TZ=Asia/Tokyo ./timezone_clock -port 8020 & TZ=Europe/London ./timezone_clock -port 8030 &
func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", *p))

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}

		go handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		tz := os.Getenv("TZ")
		timezone, err := time.LoadLocation(tz)

		if err != nil {
			panic(fmt.Sprintf("Can't found timezone %s", tz))
		}

		_, err = io.WriteString(c, time.Now().In(timezone).Format("15:04:05\n"))
		if err != nil {
			log.Print(err)
			return // e.g., client disconnected
		}

		time.Sleep(1 * time.Second)
	}
}
