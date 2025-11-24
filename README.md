# Basic-Chat-room
Basic chat room using Go 
Demo Video

[Watch the running application here](https://drive.google.com/file/d/1dDYnCAWPvSAsiwmJkrAIvMDkbKsBxaTs/view?usp=sharing)

# How it Works

- The server registers the ChatRoomServer object for RPC.

- Each client connects using rpc.Dial.

- When a message is sent, the server stores it and returns full chat history.

- Clients print the entire chat history each time.

# project structure 

```text
Basic-Chat-room/
│
├── client/
│     └── main.go
│
├── server/
│     └── main.go
│
├── Dockerfile
├── go.mod
└── README.md
```


## How to Run
## without docker
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

## with docker 
```bash
docker build -t moonal2005/simple-rpc-manal .

```
```bash
docker run -p 1234:1234 moonal2005/simple-rpc-manal

```
[Docker hub image](https://hub.docker.com/repository/docker/manalanwer/simple-rpc-manal/general)

```bash
docker pull moonal2005/simple-rpc-manal

```
