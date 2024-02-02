# WebSocketStream

WebSocketStream 是一个 Go 语言库，它提供了一个简单且一致的接口来处理 WebSocket 连接。它基于 [gorilla/websocket](https://github.com/gorilla/websocket) 库，并提供了一些额外的功能，如连接升级和消息读写。

## 特性

- **WebSocket 连接升级**：`Upgrader` 结构体提供了 `Upgrade` 方法，可以将 HTTP 连接升级为 WebSocket 连接。
- **WebSocket 连接拨号**：`Dialer` 结构体提供了 `Dial` 方法，可以创建一个新的 WebSocket 连接。
- **消息读写**：`WebSocketStream` 结构体提供了 `ReadMessage` 和 `WriteMessage` 方法，可以方便地读取和发送 WebSocket 消息。
- **流式读写**：`WebSocketStream` 结构体实现了 `io.Reader` 和 `io.Writer` 接口，可以像处理普通的流一样处理 WebSocket 连接。

## 安装

你可以使用 `go get` 命令来安装 WebSocketStream：

```bash
go get github.com/1370339317/websocketstream
