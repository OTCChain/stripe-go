package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	pbs "github.com/otcChain/chord-go/pbs/rpc"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync/atomic"
)

type HttpClient struct {
	idCounter uint32
	rootCtx   context.Context
	client    *http.Client
	headers   http.Header
	url       string
}

func DialHTTP(endpoint string) (Client, error) {
	return DialHTTPWithClient(endpoint, new(http.Client))
}

func DialHTTPWithClient(endpoint string, client *http.Client) (*HttpClient, error) {
	_, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	headers := make(http.Header, 2)
	headers.Set("accept", contentType)
	headers.Set("content-type", contentType)

	return &HttpClient{
		client:  client,
		headers: headers,
		url:     endpoint,
		rootCtx: context.Background(),
	}, nil
}

func (h *HttpClient) Close() {
	return
}

func (h *HttpClient) Call(url string, args []byte) ([]byte, error) {
	ctx, cancel := context.WithTimeout(h.rootCtx, DefaultReadTimeout+DefaultWriteTimeout)
	defer cancel()
	return h.CallContext(ctx, url, args)
}

func (h *HttpClient) buildMsg(args []byte) (json.RawMessage, error) {
	id := atomic.AddUint32(&h.idCounter, 1)
	//msg := &JsonRpcMessageItem{Version: vsn, ID: id, Params: args}

	msg := &pbs.RpcMsgItem{
		VersionID: vsn,
		ID:        id,
		Parameter: args,
	}
	msgs := make([]*pbs.RpcMsgItem, 1)
	msgs[0] = msg
	ret := &pbs.RpcRequest{
		Request: msgs,
	}
	body, err := proto.Marshal(ret)
	if err != nil {
		fmt.Println("-------->")
		fmt.Println(err.Error())
		fmt.Println("-------->")
		return nil, err
	}
	return body, nil
}

func (h *HttpClient) buildRequest(ctx context.Context, path string, body []byte) (*http.Request, error) {

	req, err := http.NewRequestWithContext(ctx,
		"POST",
		h.url+path,
		ioutil.NopCloser(bytes.NewReader(body)))
	if err != nil {
		return nil, err
	}

	req.ContentLength = int64(len(body))
	req.Header = h.headers.Clone()
	return req, nil
}

func (h *HttpClient) CallContext(ctx context.Context, path string, args []byte) ([]byte, error) {

	body, err := h.buildMsg(args)
	if err != nil {
		return nil, err
	}

	req, err := h.buildRequest(ctx, path, body)
	if err != nil {
		return nil, err
	}
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	out, err := ioutil.ReadAll(resp.Body)

	aws := &pbs.RpcAnswer{}
	if err := proto.Unmarshal(out, aws); err != nil {
		return nil, err
	}

	fmt.Println("context call body:=>", aws.String())

	code := aws.Answer[0].Code
	data := aws.Answer[0].Data
	if code == int32(pbs.ApiRet_Error) {
		return nil, fmt.Errorf(string(data))
	}
	return data, nil
}
