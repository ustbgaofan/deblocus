package tunnel

import (
	"bufio"
	"fmt"
	"net"
	"net/http"

	log "github.com/Lafeng/deblocus/glog"
	"github.com/gorilla/websocket"
)

type t_httpResponseWriter struct {
	handlerHeader http.Header
	conn          net.Conn
	writer        *bufio.ReadWriter
	written       int64
	wroteHeader   bool
	status        int
}

func newHttpResponseWriter(conn net.Conn) *t_httpResponseWriter {
	return &t_httpResponseWriter{
		handlerHeader: make(http.Header),
		conn:          conn,
		writer:        bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)),
		status:        http.StatusOK,
	}
}

func (w *t_httpResponseWriter) Header() http.Header {
	return w.handlerHeader
}

func (w *t_httpResponseWriter) WriteHeader(code int) {
	if w.wroteHeader {
		log.Errorln("http: multiple response.WriteHeader calls")
		return
	}
	w.wroteHeader = true
	w.status = code
	w.sendHeader()
}

func (w *t_httpResponseWriter) sendHeader() {
	w.writer.WriteString(fmt.Sprintf("HTTP/1.1 %03d OK\r\n", w.status))
	for k, v := range w.handlerHeader {
		w.writer.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	w.writer.WriteString(CRLF)
}

func (w *t_httpResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.conn, w.writer, nil
}

// either dataB or dataS is non-zero.
func (w *t_httpResponseWriter) write(lenData int, dataB []byte, dataS string) (n int, err error) {
	if !w.wroteHeader {
		w.WriteHeader(w.status)
	}
	if lenData == 0 {
		return 0, nil
	}

	defer w.writer.Flush()
	if dataB != nil {
		return w.writer.Write(dataB)
	} else {
		return w.writer.WriteString(dataS)
	}
}

func (w *t_httpResponseWriter) Write(data []byte) (n int, err error) {
	return w.write(len(data), data, "")
}

func (w *t_httpResponseWriter) WriteString(data string) (n int, err error) {
	return w.write(len(data), nil, data)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func wsHandler(conn net.Conn, req *http.Request) {
	var rw = newHttpResponseWriter(conn)
	var c, err = upgrader.Upgrade(rw, req, nil)
	if err != nil {
		log.Errorln("upgrade ws:", err)
		return
	}

	var rErr = make(chan int, 1)
	go func(conn *websocket.Conn, sig chan int) {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Infoln("read:", err)
				sig <- 0xff
				return
			} else {
				log.Infoln("ws msg:", msg)
			}
		}
	}(c, rErr)

	var q = log.LogQueueInstance
	var mn = q.NewMonitor()
	var msg *log.LogEntry

	q.Retrieve(func(m string) bool {
		log.Infoln(m)
		return c.WriteMessage(websocket.TextMessage, []byte(m)) == nil
	})

	for {
		select {
		case <-rErr:
			return

		case <-mn.Changes():
			mn.Next()
			msg = mn.Value().(*log.LogEntry)
			if c.WriteMessage(websocket.TextMessage, []byte(msg.Msg)) != nil {
				return
			}
		}
	}
}
