package p2p

import "sync"

type ChordNetwork interface {
	LaunchUp() error
	Destroy() error
}

var _instance ChordNetwork
var once sync.Once

func Inst() ChordNetwork {
	once.Do(func() {
		_instance = newNetwork()
	})
	return _instance
}
