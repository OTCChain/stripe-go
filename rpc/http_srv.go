package rpc

import (
	"bytes"
	"fmt"
	pbs "github.com/otcChain/chord-go/pbs/rpc"
	"github.com/otcChain/chord-go/utils"
	"github.com/otcChain/chord-go/utils/thread"
	"google.golang.org/protobuf/proto"
	"io"
	"mime"
	"net"
	"net/http"
)

const (
	maxRequestContentLength = 1024 * 1024 * 5
	contentType             = "application/json"
	httpThreadName          = "http rpc thread"
	websocketThreadName     = "websocket rpc thread"
)

// https://www.jsonrpc.org/historical/json-rpc-over-http.html#id13
var acceptedContentTypes = []string{contentType, "application/json-rpc", "application/jsonrequest"}

type HttpRpcProvider func(*pbs.RpcMsgItem) *pbs.RpcResponse
type HttpApiRouter map[string]HttpRpcProvider

type HttpRpc struct {
	apis *http.ServeMux
	srv  *http.Server
}

func (hr *HttpRpc) StartRpc() error {
	endPoint := fmt.Sprintf("%s:%d", _rpcConfig.HttpIP, _rpcConfig.HttpPort)
	ln, err := net.Listen("tcp4", endPoint)
	if err != nil {
		return err
	}
	thread.NewThreadWithName(httpThreadName, func(_ chan struct{}) {
		utils.LogInst().Info().Msgf("http rpc service startup at:%s", endPoint)
		err = hr.srv.Serve(ln)
		utils.LogInst().Err(err).Str("http rpc", "Exit").Send()
		hr.ShutDown()
	}).Run()
	return nil
}

func (hr *HttpRpc) ShutDown() {
	if hr.srv == nil {
		return
	}
	_ = hr.srv.Close()
	hr.srv = nil
}

func validateRequest(r *http.Request) (int, error) {
	if r.Method == http.MethodPut || r.Method == http.MethodDelete {
		return http.StatusMethodNotAllowed, fmt.Errorf("method not allowed")
	}
	if r.ContentLength > maxRequestContentLength {
		err := fmt.Errorf("content length too large (%d>%d)", r.ContentLength, maxRequestContentLength)
		return http.StatusRequestEntityTooLarge, err
	}
	// Allow OPTIONS (regardless of content-type)
	if r.Method == http.MethodOptions {
		return 0, nil
	}
	// Check content-type
	if mt, _, err := mime.ParseMediaType(r.Header.Get("content-type")); err == nil {
		for _, accepted := range acceptedContentTypes {
			if accepted == mt {
				return 0, nil
			}
		}
	}
	// Invalid content-type
	err := fmt.Errorf("invalid content type, only %s is supported", contentType)
	return http.StatusUnsupportedMediaType, err
}

func (hr *HttpRpc) readRpcMsg(w http.ResponseWriter, r *http.Request) []*pbs.RpcMsgItem {
	if r.Method == http.MethodGet && r.ContentLength == 0 && r.URL.RawQuery == "" {
		w.WriteHeader(http.StatusOK)
		return nil
	}
	if code, err := validateRequest(r); err != nil {
		http.Error(w, err.Error(), code)
		return nil
	}
	w.Header().Set("content-type", contentType)

	body := io.LimitReader(r.Body, maxRequestContentLength)
	var b bytes.Buffer
	n, err := b.ReadFrom(body)
	if err != nil && n != 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	requests := &pbs.RpcRequest{}
	if err := proto.Unmarshal(b.Bytes(), requests); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	if len(requests.Request) == 0 {
		http.Error(w, "no valid rpc request", http.StatusBadRequest)
		return nil
	}

	return requests.Request
}

func (hr *HttpRpc) processMsg(w http.ResponseWriter, r *http.Request, provider HttpRpcProvider) {

	rpcMsg := hr.readRpcMsg(w, r)
	if rpcMsg == nil {
		return
	}

	aws := make([]*pbs.RpcResponse, 0)
	for _, msg := range rpcMsg {
		utils.LogInst().Debug().Msgf("process one of rpc message:%s", msg.String())
		res := provider(msg)
		aws = append(aws, res)
	}
	answers := &pbs.RpcAnswer{
		Answer: aws,
	}
	protoData, err := proto.Marshal(answers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	n, err := w.Write(protoData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.LogInst().Debug().Int("http rpc response", n)
}

func (hr *HttpRpc) regService(name string, provider HttpRpcProvider) {
	utils.LogInst().Info().Msgf("api path[%s] register success", name)
	hr.apis.HandleFunc(name, func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				utils.LogInst().Error().Interface("http rpc request fatal:", r).Send()
			}
		}()
		hr.processMsg(writer, request, provider)
	})
}

func newHttpRpc() *HttpRpc {
	apis := http.NewServeMux()
	server := &http.Server{
		Handler:      apis,
		ReadTimeout:  _rpcConfig.ReadTimeout,
		WriteTimeout: _rpcConfig.WriteTimeout,
		IdleTimeout:  _rpcConfig.IdleTimeout,
	}

	hr := &HttpRpc{
		apis: apis,
		srv:  server,
	}
	for id, cb := range HttpRpcApis {
		hr.regService(id, cb)
	}
	return hr
}
