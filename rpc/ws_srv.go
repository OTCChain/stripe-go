package rpc

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/otcChain/chord-go/chord/transaction"
	"github.com/otcChain/chord-go/utils"
	"github.com/otcChain/chord-go/utils/thread"
	"net"
	"net/http"
	"time"
)

const (
	ChordEvent = "/chord/event"
)

type WsRpc struct {
	upGrader  *websocket.Upgrader
	server    *http.Server
	kaTimer   *time.Ticker
	eventChan chan *transaction.Event
}

func (wr *WsRpc) StartRpc() error {
	endPoint := fmt.Sprintf("%s:%d", _rpcConfig.WsIP, _rpcConfig.WsPort)
	ln, err := net.Listen("tcp4", endPoint)
	if err != nil {
		utils.LogInst().Error().Str("end", endPoint).
			Str("websocket rpc", err.Error()).
			Send()
		return err
	}
	thread.NewThreadWithName(websocketThreadName, func(_ chan struct{}) {
		utils.LogInst().Info().Msgf("websocket rpc service startup at:%s", endPoint)
		err = wr.server.Serve(ln)
		utils.LogInst().Err(err).Str("websocket rpc", "Exit").Send()
		wr.ShutDown()
	})
	return nil
}

func (wr *WsRpc) ShutDown() {
	if wr.server == nil {
		return
	}

	_ = wr.server.Close()
	wr.server = nil
}

func newWsRpc() *WsRpc {
	apis := http.NewServeMux()
	server := &http.Server{
		Handler: apis,
	}
	wr := &WsRpc{
		server:    server,
		upGrader:  _rpcConfig.newUpGrader(),
		kaTimer:   time.NewTicker(_rpcConfig.PingPeriod),
		eventChan: make(chan *transaction.Event, _rpcConfig.WsEventBufSize),
	}
	apis.HandleFunc(ChordEvent, wr.eventWatcher)
	return wr
}

func (wr *WsRpc) eventWatcher(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if r := recover(); r != nil {
			utils.LogInst().Warn().Msgf("websocket service panic by one server :=>%s", r)
		}
	}()

	webSocket, err := wr.upGrader.Upgrade(w, r, nil)
	if err != nil {
		utils.LogInst().Err(err).Send()
		return
	}

	webSocket.SetReadLimit(int64(_rpcConfig.WsIOBufferSize))
	_ = webSocket.SetReadDeadline(time.Now().Add(_rpcConfig.PongWait))
	webSocket.SetPongHandler(func(string) error {
		return webSocket.SetReadDeadline(time.Now().Add(_rpcConfig.PongWait))
	})

	thread.NewThread(func(stop chan struct{}) {
		wr.pingPong(stop, webSocket)
	})
	thread.NewThread(func(stop chan struct{}) {
		wr.eventWriter(stop, webSocket)
	})
}

func (wr *WsRpc) pingPong(stop chan struct{}, conn *websocket.Conn) {
	for {
		select {
		case <-wr.kaTimer.C:
			utils.LogInst().Debug().Str("WS ping pong", "sent").Send()
			if err := conn.SetWriteDeadline(time.Now().Add(_rpcConfig.WsWriteTimeout)); err != nil {
				utils.LogInst().Warn().Str("WS write deadline", err.Error()).Send()
				return
			}
			if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				utils.LogInst().Warn().Str("WS write ping", err.Error()).Send()
				return
			}
		case <-stop:
			utils.LogInst().Debug().Str("websocket rpc ping pong thread", "exit").Send()
			return
		}
	}
}

func (wr *WsRpc) eventWriter(stop chan struct{}, conn *websocket.Conn) {
	for {
		select {
		case <-stop:
			utils.LogInst().Debug().Str("websocket rpc write thread", "exit").Send()
			return
		case event := <-wr.eventChan:
			if event == nil {
				utils.LogInst().Info().Str("websocket rpc write message chan", " closed").Send()
				_ = conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := conn.SetWriteDeadline(time.Now().Add(_rpcConfig.WsWriteTimeout)); err != nil {
				utils.LogInst().Warn().Str("websocket rpc write set write timeout ", err.Error()).Send()
				return
			}

			w, err := conn.NextWriter(websocket.TextMessage)
			if err != nil {
				utils.LogInst().Warn().Str("websocket rpc write get next writer ", err.Error()).Send()
				return
			}

			_, err = w.Write(event.Data())
			if err := w.Close(); err != nil {
				utils.LogInst().Warn().Str("websocket rpc write ", err.Error()).Send()
				return
			}
		}
	}
}
