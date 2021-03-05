package rpc

import (
	"fmt"
	"sync"
)

type Server interface {
	StartService() error
	RegisterHttpSrv(name string, fn HttpRpcProvider) error
}

type ServiceManager struct {
	cmdRpc  *cmdService
	httpRpc *HttpRpc
	wsRpc   *WsRpc
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

	if config.CmdEnabled {
		sm.cmdRpc = newCmdRpc()
	}
	if config.HttpEnabled {
		sm.httpRpc = newHttpRpc()
	}
	if config.WsEnabled {
		sm.wsRpc = newWsRpc()
	}

	return sm
}

func (sm *ServiceManager) StartService() error {
	if sm.cmdRpc != nil {
		if err := sm.cmdRpc.StartRpc(); err != nil {
			return err
		}
	}
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
	sm.httpRpc.regisService(name, fn)
	return nil
}
