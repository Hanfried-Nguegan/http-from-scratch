package main

import (
	"bufio"
	_ "bytes"
	"fmt"
	"httpServer/internal/http1"
	_ "io"
	"log"
	"net"
)

// func getLinesChannel(f io.ReadCloser) <-chan string {
// 	out := make(chan string, 1)

// 	go func() {
// 		defer close(out)
// 		defer f.Close()

// 		str := ""
// 		for {
// 			data := make([]byte, 8)
// 			n, err := f.Read(data)
// 			if err != nil {
// 				break
// 			}

// 			data = data[:n]
// 			if i := bytes.IndexByte(data, '\n'); i != -1 {
// 				str += string(data[:i])
// 				data = data[i+1:]
// 				out <- str
// 				str = ""
// 			}

// 			str += string(data)
// 		}
// 		if len(str) != 0 {
// 			out <- str
// 		}
// 	}()

// 	return out
// }

func main() {
	listener, err := net.Listen("tcp", ":8080")
	//f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error", "error", err)
		}
		r, err := http1.ParseRequest(bufio.NewReader(conn))
		if err != nil {
			log.Fatal("error", "error", err)
		}
		fmt.Printf("Request line:\n")
		fmt.Printf("- Method: %s\n", r.RequestLine.Method)
		fmt.Printf("- Target: %s\n", r.RequestLine.RequestTarget)
		fmt.Printf("- Version: %s\n", r.RequestLine.HttpVersion)
	}
}
