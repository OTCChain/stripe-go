package node

import "sync"

type ChordNode interface {
	Setup() error
	Start()
	ShutDown()
}

var _instance ChordNode
var once sync.Once

func Inst() ChordNode {
	once.Do(func() {
		_instance = newNode()
	})
	return _instance
}
