# Pure Go IPC Lib

用纯 Go 实现的跨平台 IPC 操作库，包含服务端和客户端

## Install

```plaintext
go get github.com/runoneall/pgoipc
```

## Server

```go
server.Serv("example", func(conn net.Conn) {})
```

## Client

```go
conn := client.Connect("example")
```
