package node

import (
	"github.com/otcChain/chord-go/p2p"
	"github.com/otcChain/chord-go/rpc"
)

type ChordNodeV1 struct {
}

func newNode() *ChordNodeV1 {
	n := &ChordNodeV1{}
	return n
}

func (cn *ChordNodeV1) Setup() error {
	if err := p2p.Inst().LaunchUp(); err != nil {
		return err
	}

	for id, cb := range publicHttpRpcApi {
		if err := rpc.Inst().RegisterHttpSrv(id, cb); err != nil {
			return err
		}
	}

	return nil
}

func (cn *ChordNodeV1) Start() {

}
func (cn *ChordNodeV1) ShutDown() {

}
