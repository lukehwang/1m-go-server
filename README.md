1M connection server in Go
---------------------------

Copied from https://github.com/eranyanay/1m-go-websockets
for study


## 1. Simple Server

```
cd simple_server
go run server.go
```

Open http://localhost:8000


## 2. Simple Websocket Server

```
cd simple_ws
go run server.go
```

At another terminal

```
go run client.go
```

## 3. Websocket Server - Max rulimit 

```
go run server.go
Set Limit 10496 -> 65536
```
