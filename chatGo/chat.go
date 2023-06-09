package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	host = flag.String("host", "localhost", "Host of the server")
	port = flag.Int("port", 3090, "Port of the server")
)

func main() {

	flag.Parse()

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()

	CopyContent(conn, os.Stdin)
	conn.Close()
	<-done

}

func CopyContent(destination io.Writer, source io.Reader) {
	_, err := io.Copy(destination, source)
	if err != nil {
		log.Fatal(err)
	}
}
