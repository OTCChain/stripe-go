package chord

import "sync"

type Node interface {
	Start() error
	ShutDown()
}

var _instance Node
var once sync.Once

func Inst() Node {
	once.Do(func() {
		_instance = newNode()
	})
	return _instance
}
