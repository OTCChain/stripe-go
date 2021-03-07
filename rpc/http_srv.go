package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/otcChain/chord-go/common"
	"github.com/otcChain/chord-go/utils"
	"io"
	"mime"
	"net/http"
	"reflect"
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

func (hr *HttpRpc) StartRpc() chan error {
	endPoint := fmt.Sprintf("%s:%d", _rpcConfig.HttpIP, _rpcConfig.HttpPort)
	server := &http.Server{Addr: endPoint, Handler: hr.apis}

	server.ReadTimeout = _rpcConfig.ReadTimeout
	server.WriteTimeout = _rpcConfig.WriteTimeout
	server.IdleTimeout = _rpcConfig.IdleTimeout

	for id, cb := range HttpRpcApis {
		hr.regService(id, cb)
	}
	utils.LogInst().Info().Msgf("http rpc service startup at:%s", endPoint)
	errCh := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()
		errCh <- err
	}()

	return errCh
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
		utils.LogInst().Debug().Msgf("process one of rpc message:%s", msg.String())
		var tps = []reflect.Type{
			common.AddressT,
			reflect.TypeOf(""),
		}
		vas, err := parsePositionalArguments(msg.Params, tps)
		if err != nil {
			panic(err)
		}
		fmt.Println(vas)
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
	hr := &HttpRpc{
		apis: http.NewServeMux(),
	}
	return hr
}

// parsePositionalArguments tries to parse the given args to an array of values with the
// given types. It returns the parsed values or an error when the args could not be
// parsed. Missing optional arguments are returned as reflect.Zero values.
func parsePositionalArguments(rawArgs json.RawMessage, types []reflect.Type) ([]reflect.Value, error) {
	dec := json.NewDecoder(bytes.NewReader(rawArgs))
	var args []reflect.Value
	tok, err := dec.Token()
	switch {
	case err == io.EOF || tok == nil && err == nil:
		// "params" is optional and may be empty. Also allow "params":null even though it's
		// not in the spec because our own client used to send it.
	case err != nil:
		return nil, err
	case tok == json.Delim('['):
		// Read argument array.
		if args, err = parseArgumentArray(dec, types); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("non-array args")
	}
	// Set any missing args to nil.
	for i := len(args); i < len(types); i++ {
		if types[i].Kind() != reflect.Ptr {
			return nil, fmt.Errorf("missing value for required argument %d", i)
		}
		args = append(args, reflect.Zero(types[i]))
	}
	return args, nil
}

func parseArgumentArray(dec *json.Decoder, types []reflect.Type) ([]reflect.Value, error) {
	args := make([]reflect.Value, 0, len(types))
	for i := 0; dec.More(); i++ {
		if i >= len(types) {
			return args, fmt.Errorf("too many arguments, want at most %d", len(types))
		}
		argval := reflect.New(types[i])
		if err := dec.Decode(argval.Interface()); err != nil {
			return args, fmt.Errorf("invalid argument %d: %v", i, err)
		}
		if argval.IsNil() && types[i].Kind() != reflect.Ptr {
			return args, fmt.Errorf("missing value for required argument %d", i)
		}
		args = append(args, argval.Elem())
	}
	// Read end of args array.
	_, err := dec.Token()
	return args, err
}
