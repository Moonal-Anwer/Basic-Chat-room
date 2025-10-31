# Basic-Chat-room
Basic chat room using Go 
Demo Video

📹[Watch the running application here](https://drive.google.com/file/d/1dDYnCAWPvSAsiwmJkrAIvMDkbKsBxaTs/view?usp=sharing)

# How it Works

- The server registers the ChatServer object for RPC.

- Each client connects using rpc.Dial.

- When a message is sent, the server stores it and returns full chat history.

- Clients print the entire chat history each time.


## How to Run

### Start the Server
```bash
go run server.go
```

### Start the Client (each in a new terminal)
```bash
go run client.go
```

### Exit the Client
Type:
```
exit
```
