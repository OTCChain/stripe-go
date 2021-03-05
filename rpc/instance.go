package rpc

import (
	"fmt"
	"sync"
)

type Server interface {
	StartService() error
}

type ServiceManager struct {
	services map[string]Rpc
}

type Rpc interface {
	StartRpc() error
}

var _instance Server
var once sync.Once

func Inst() Server {
	once.Do(func() {
		_instance = newServiceManager()
	})
	return _instance
}

func newServiceManager() Server {
	sm := &ServiceManager{
		services: make(map[string]Rpc),
	}

	if config.CmdEnabled {
		sm.services[SrvNameCmd] = newCmdRpc()
	}
	if config.HttpEnabled {
		sm.services[SrvNameHttp] = newHttpRpc()
	}
	if config.WsEnabled {
		sm.services[SrvNameWs] = newWsRpc()
	}

	return sm
}

func (sm *ServiceManager) StartService() error {
	for name, srv := range sm.services {
		if err := srv.StartRpc(); err != nil {
			return fmt.Errorf("start Rpc[%s] failed:%s", name, err)
		}
	}
	return nil
}
