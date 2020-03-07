package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var data = make([]byte, 4096)

var (
	inputcmd string
)

func readMsg(conn net.Conn) {
	for {
		count, err := conn.Read(data)
		fmt.Printf(string(data[:count]))
		if err != nil {
			break
		}
	}
}
func main() {

	if len((os.Args)) > 1 {
		clientAddr := os.Args[1] + ":" + os.Args[2]
		conn, err := net.Dial("tcp", clientAddr)
		if err != nil {
			fmt.Println("conn error")
			return
		}
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Shell:")
			input, _ := reader.ReadString('\n')
			cmd := strings.Trim(input, "\n")
			if cmd == "exit" {
				break
			}
			conn.Write([]byte(cmd))
			count, err := conn.Read(data)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf(string(data[:count]) + "\n")
		}
	} else {
		fmt.Println("[+]Usage: client[.exe] ip serverport")
	}

}
