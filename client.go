package main

import (
	"bufio" //reads user input line by line i used it instead of scan
	"fmt"
	"net/rpc"
	"os"
	"strings"
)

type Message struct {
	Name string
	Text string
}

func main() {
	client, err := rpc.Dial("tcp", "0.0.0.0:1234")
	defer client.Close()
	read := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your name : ")
	name, _ := read.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Hey  %s!, write a message and engage in active conversations. ", name)
	// a loop until the user enter exit
	for {
		fmt.Print("\n enter a message or 'exit' to get out")
		input, _ := read.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("bye bye...")
			break
		}
		var reply []string
		msg := Message{Name: name, Text: input}

		err = client.Call("ChatroomServer.SendAMessage", msg, &reply)
		if err != nil {
			fmt.Println("something went wrong ooops: ", err)
			break
		}
		fmt.Println("\n ******* chat logs *******")
		for _, m := range reply {
			fmt.Println(m)
		}

		fmt.Println("****************************")
	}
}
