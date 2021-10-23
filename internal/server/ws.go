package server

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var clients = make(map[int][]*websocket.Conn)
var lock = sync.Mutex{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin:     func(*http.Request) bool { return true },
}

func (s *Server) wsHandlerFunc(ctx *gin.Context) {
	w, r := ctx.Writer, ctx.Request
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Infoln("error upgrade connection:", err)
		SendErr(ctx, ErrNotCreateWsConnection)
		return
	}
	go handleWs(conn)
}

func handleWs(conn *websocket.Conn) {
	defer conn.Close()
	msg := &wsMessage{}
	maxRetry := 10
	for {
		err := conn.ReadJSON(msg)
		if err != nil {
			logrus.Infoln("error try to read wsMessage:", err)
			maxRetry--
			if maxRetry == 0 {
				return
			}
		}
		auctionId := msg.AuctionID
		if auctionId <= 0 {
			logrus.Infoln("invalid auctionId:", auctionId)
			err = conn.WriteMessage(websocket.TextMessage, []byte("invalid auctionId")) //nolint:errcheck
			if err != nil {
				remove(clients[auctionId], conn)
				conn.Close()
				return
			}
		}
		switch msg.Action {
		case "connect":
			logrus.Info("connect")
			clients[auctionId] = append(clients[auctionId], conn)
			continue
		case "left":
			logrus.Info("left")
			remove(clients[auctionId], conn)
			conn.Close()
			return
		//case "bid":
		//	logrus.Info("bid")
		//	broadcast(auctionId)
		default:
			logrus.Infoln("invalid ws action:", msg.Action)
			err = conn.WriteMessage(websocket.TextMessage, []byte("invalid ws action"))
			if err != nil {
				remove(clients[auctionId], conn)
				conn.Close()
				return
			}
		}
	}
}

func broadcast(msg RespBidMsg) {
	auctionId := msg.AuctionId
	for _, conn := range clients[int(auctionId)] {
		out, _ := json.Marshal(msg)
		err := conn.WriteMessage(websocket.TextMessage, out)
		if err != nil {
			remove(clients[int(auctionId)], conn)
			conn.Close()
		}
	}
}

type wsMessage struct {
	Action    string            `json:"action"`
	AuctionID int               `json:"auctionId"`
	Payload   map[string]string `json:"payload"`
}

func remove(arr []*websocket.Conn, elem *websocket.Conn) {
	lock.Lock()
	defer lock.Unlock()
	for i, e := range arr {
		if e == elem {
			arr = append(arr[:i], arr[i+1:]...)
			return
		}
	}
}
