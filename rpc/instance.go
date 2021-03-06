package rpc

import (
	"sync"
	"time"
)

type Server interface {
	StartService() error
}

type ServiceManager struct {
	httpRpc *HttpRpc
	wsRpc   *WsRpc
}

// HTTPTimeouts represents the configuration params for the HTTP RPC server.
type HTTPTimeouts struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

var (
	_instance   Server
	once        sync.Once
	HttpRpcApis = make(HttpApiRouter)
)

func Inst() Server {
	once.Do(func() {
		_instance = newServiceManager()
	})
	return _instance
}

func newServiceManager() Server {
	sm := &ServiceManager{}

	if _rpcConfig.HttpEnabled {
		sm.httpRpc = newHttpRpc()
	}
	if _rpcConfig.WsEnabled {
		sm.wsRpc = newWsRpc()
	}

	return sm
}

func (sm *ServiceManager) StartService() error {
	if sm.httpRpc != nil {
		if err := sm.httpRpc.StartRpc(); err != nil {
			return err
		}
	}
	if sm.wsRpc != nil {
		if err := sm.wsRpc.StartRpc(); err != nil {
			return err
		}
	}
	return nil
}
