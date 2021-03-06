package chain

import (
	"github.com/otcChain/chord-go/p2p"
)

type ChordNodeV1 struct {
}

func newNode() *ChordNodeV1 {
	if _nodeConfig == nil {
		panic("please init instance config first")
	}
	n := &ChordNodeV1{}
	n.initRpcApis()
	return n
}

func (cn *ChordNodeV1) Setup() error {
	if err := p2p.Inst().LaunchUp(); err != nil {
		return err
	}
	return nil
}

func (cn *ChordNodeV1) Start() {

}
func (cn *ChordNodeV1) ShutDown() {

}
