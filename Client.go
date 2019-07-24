package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	/*
		enter your mathematical phrase and press enter to see the result
		The Server performs the calculation
	*/
	con, _ := net.Dial("tcp", "127.0.0.1:8080")
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		_, _ = fmt.Fprintf(con, text+"\n")
		result, _ := bufio.NewReader(con).ReadString('\n')
		fmt.Print("result: " + result)

	}
}
