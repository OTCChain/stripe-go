package rpc

import (
	"fmt"
	"sync"
	"time"
)

type Server interface {
	StartService() error
	RegisterHttpSrv(name string, fn HttpRpcProvider) error
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
	HttpRpcInvalid = fmt.Errorf("http rpc serivice is not valaible")
	_instance      Server
	once           sync.Once
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

func (sm *ServiceManager) RegisterHttpSrv(name string, fn HttpRpcProvider) error {
	if sm.httpRpc == nil {
		return HttpRpcInvalid
	}
	sm.httpRpc.regService(name, fn)
	return nil
}
