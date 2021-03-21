package chord

import (
	"github.com/otcChain/chord-go/chord/txpool"
	"github.com/otcChain/chord-go/p2p"
	"github.com/otcChain/chord-go/rpc"
)

type NodeV1 struct {
	microTxPool *txpool.MicTxPool
}

func newNode() *NodeV1 {
	if _nodeConfig == nil {
		panic("please init instance config first")
	}

	n := &NodeV1{
		microTxPool: txpool.NewMicroTxPool(),
	}

	n.initRpcApis()
	return n
}

func (cn *NodeV1) Start() error {
	if err := p2p.Inst().LaunchUp(); err != nil {
		return err
	}

	if err := rpc.Inst().StartService(); err != nil {
		return err
	}

	return nil
}

func (cn *NodeV1) ShutDown() {

}
