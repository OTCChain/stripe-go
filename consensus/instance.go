package consensus

import "sync"

type Engine interface {
	Setup(*Config) error
}

var _instance Engine
var once sync.Once

func Inst() Engine {
	once.Do(func() {
		_instance = newEngine()
	})
	return _instance
}
