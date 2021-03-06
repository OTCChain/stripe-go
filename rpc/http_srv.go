package rpc

import (
	"encoding/json"
	"fmt"
	"github.com/otcChain/chord-go/utils"
	"io"
	"mime"
	"net/http"
)

const (
	maxRequestContentLength = 1024 * 1024 * 5
	contentType             = "application/json"
)

// https://www.jsonrpc.org/historical/json-rpc-over-http.html#id13
var acceptedContentTypes = []string{contentType, "application/json-rpc", "application/jsonrequest"}

type HttpRpcProvider func(*JsonRpcMessageItem) (json.RawMessage, *JsonError)
type HttpApiRouter map[string]HttpRpcProvider

type HttpRpc struct {
	apis *http.ServeMux
}

func (hr *HttpRpc) StartRpc() error {
	endPoint := fmt.Sprintf("%s:%d", _rpcConfig.HttpIP, _rpcConfig.HttpPort)
	server := &http.Server{Addr: endPoint, Handler: hr.apis}

	server.ReadTimeout = _rpcConfig.ReadTimeout
	server.WriteTimeout = _rpcConfig.WriteTimeout
	server.IdleTimeout = _rpcConfig.IdleTimeout

	for id, cb := range HttpRpcApis {
		hr.regService(id, cb)
	}

	return server.ListenAndServe()
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

func (hr *HttpRpc) readRpcMsg(w http.ResponseWriter, r *http.Request) *JsonRpcMessage {
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
	dec := json.NewDecoder(body)
	msg := JsonRpcMessage{}
	if err := dec.Decode(&msg); err != nil {
		if err != io.EOF {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return nil
		}
	}

	if len(msg) == 0 {
		http.Error(w, "no valid rpc request", http.StatusBadRequest)
		return nil
	}

	return &msg
}

func (hr *HttpRpc) processMsg(w http.ResponseWriter, r *http.Request, provider HttpRpcProvider) {

	rpcMsg := hr.readRpcMsg(w, r)
	if rpcMsg == nil {
		return
	}

	answers := make(JsonRpcMessage, 0)

	for _, msg := range *rpcMsg {
		a := &JsonRpcMessageItem{
			Version: msg.Version,
			ID:      msg.ID,
		}
		data, e := provider(msg)
		if e != nil {
			a.Error = e
		} else {
			a.Result = data
		}
		answers = append(answers, a)
	}

	c := json.NewEncoder(w)
	if err := c.Encode(answers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (hr *HttpRpc) regService(name string, provider HttpRpcProvider) {
	hr.apis.HandleFunc(name, func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				utils.LogInst().Trace().Interface("http rpc request", r)
			}
		}()
		hr.processMsg(writer, request, provider)
	})
}

func newHttpRpc() *HttpRpc {
	hr := &HttpRpc{
		apis: http.NewServeMux(),
	}
	return hr
}
