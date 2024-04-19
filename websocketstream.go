package websocketstream

import (
	"io"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketStream struct {
	conn    *websocket.Conn
	msgType int
	reader  io.Reader
}
type Dialer struct {
	Dialer *websocket.Dialer
}
type Upgrader struct {
	ReadBufferSize  int
	WriteBufferSize int
	CheckOrigin     func(r *http.Request) bool
	Error           func(w http.ResponseWriter, r *http.Request, status int, reason error)
}

func NewWebSocketStream(conn *websocket.Conn) *WebSocketStream {
	return &WebSocketStream{conn: conn}
}

func (u *Upgrader) Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*WebSocketStream, error) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  u.ReadBufferSize,
		WriteBufferSize: u.WriteBufferSize,
		CheckOrigin:     u.CheckOrigin,
		Error:           u.Error,
	}
	c, err := upgrader.Upgrade(w, r, responseHeader)
	if err != nil {
		return nil, err
	}

	return NewWebSocketStream(c), nil
}

func IsWebSocketUpgrade(r *http.Request) bool {
	return websocket.IsWebSocketUpgrade(r)
}

func (d *Dialer) Dial(urlStr string, requestHeader http.Header) (*WebSocketStream, *http.Response, error) {
	c, resp, err := d.Dialer.Dial(urlStr, requestHeader)
	if err != nil {
		return nil, nil, err
	}

	return NewWebSocketStream(c), resp, nil
}

var DefaultDialer = &Dialer{Dialer: websocket.DefaultDialer}

func (w *WebSocketStream) Close() error {
	return w.conn.Close()
}

func (w *WebSocketStream) SetReadDeadline(t time.Time) error {
	return w.conn.SetReadDeadline(t)
}

func (w *WebSocketStream) SetWriteDeadline(t time.Time) error {
	return w.conn.SetWriteDeadline(t)
}
func (ws *WebSocketStream) WriteMessage(messageType int, data []byte) error {
	return ws.conn.WriteMessage(messageType, data)
}

func (ws *WebSocketStream) ReadMessage() (messageType int, p []byte, err error) {
	return ws.conn.ReadMessage()
}
func (w *WebSocketStream) Read(p []byte) (n int, err error) {
	for {
		if w.reader == nil {
			// 读取下一个消息
			w.msgType, w.reader, err = w.conn.NextReader()
			if err != nil {
				return 0, err
			}
		}

		n, err = w.reader.Read(p)
		if err == io.EOF {
			// 读取完毕，准备读取下一个消息
			w.reader = nil
			if n > 0 {
				return n, nil
			} else {
				// 如果没有数据可读，继续读取下一个消息
				continue
			}
		}
		return n, err
	}
}

func (w *WebSocketStream) Write(p []byte) (n int, err error) {
	err = w.conn.WriteMessage(websocket.BinaryMessage, p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
