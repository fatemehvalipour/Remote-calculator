package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("connecting...")
	ln, _ := net.Listen("tcp", ":8080")
	con, _ := ln.Accept()
	for {
		message, _ := bufio.NewReader(con).ReadString('\n')
		message = strings.TrimSpace(message)
		if strings.Contains(message, "+") {
			substring := strings.Split(message, "+")
			firstNum, _ := strconv.Atoi(substring[0])
			secondNum, _ := strconv.Atoi(substring[1])
			_, _ = con.Write([]byte(strconv.Itoa(firstNum+secondNum) + "\n"))
		} else if strings.Contains(message, "-") {
			substring := strings.Split(message, "-")
			firstNum, _ := strconv.Atoi(substring[0])
			secondNum, _ := strconv.Atoi(substring[1])
			_, _ = con.Write([]byte(strconv.Itoa(firstNum-secondNum) + "\n"))
		} else if strings.Contains(message, "*") {
			substring := strings.Split(message, "*")
			firstNum, _ := strconv.Atoi(substring[0])
			secondNum, _ := strconv.Atoi(substring[1])
			_, _ = con.Write([]byte(strconv.Itoa(firstNum*secondNum) + "\n"))
		} else if strings.Contains(message, "/") {
			substring := strings.Split(message, "/")
			firstNum, _ := strconv.Atoi(substring[0])
			secondNum, _ := strconv.Atoi(substring[1])
			if secondNum != 0 {
				_, _ = con.Write([]byte(strconv.Itoa(firstNum/secondNum) + "\n"))
			} else {
				_, _ = con.Write([]byte("Division by zero" + "\n"))
			}
		}
	}
}
