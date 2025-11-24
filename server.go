package main

import (
	"errors"  // so we can handel errors uniquely
	"fmt"     // for prints
	"net"     // to establish tpc connection
	"net/rpc" // for the rpc server
)

// define a new type that stores all messages
type ChatroomServer struct {
	messages []string
}

// messages will be a data type that carries name and the text
type Message struct {
	Name string
	Text string
}

// define a method {SendAMessage} that client can call remotely
// it takes a msg and send it back as a chat history to all clients
func (h *ChatroomServer) SendAMessage(msg Message,
	reply *[]string) error {
	if msg.Text == "" { // handling empty messages
		return errors.New("we received no message :(")
	}
	// formate how messages will look like > name : text
	formate := fmt.Sprintf("> %s: %s",
		msg.Name, msg.Text)
	h.messages = append(h.messages, formate)
	// print the messages on server side as well
	fmt.Println(formate)
	*reply = h.messages // send history back to the client
	return nil
}

func main() {

	chatServer := new(ChatroomServer)
	rpc.Register(chatServer)

	listener, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		fmt.Println("something went wrong: ", err)
		return
	}
	fmt.Println("server running on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("something went wrong: ", err)
			continue
		}
		// handle multiple clients at once
		go rpc.ServeConn(conn)
	}
}
