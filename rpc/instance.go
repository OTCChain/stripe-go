package rpc

import (
	"sync"
	"time"
)

type ServiceManager struct {
	httpErrChan chan error
	httpRpc     *HttpRpc
	wsErrChan   chan error
	wsRpc       *WsRpc
}

// HTTPTimeouts represents the configuration params for the HTTP RPC server.
type HTTPTimeouts struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

var (
	_instance   *ServiceManager
	once        sync.Once
	HttpRpcApis = make(HttpApiRouter)
)

func Inst() *ServiceManager {
	once.Do(func() {
		_instance = newServiceManager()
	})
	return _instance
}

func newServiceManager() *ServiceManager {

	if _rpcConfig == nil {
		panic("init rpc config first")
	}

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
		sm.httpErrChan = sm.httpRpc.StartRpc()
	}
	if sm.wsRpc != nil {
		sm.wsErrChan = sm.wsRpc.StartRpc()
	}
	return nil
}
