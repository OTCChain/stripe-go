package node

import "github.com/otcChain/chord-go/p2p"

type ChordNodeV1 struct {
}

func newNode() *ChordNodeV1 {
	n := &ChordNodeV1{}
	return n
}

func (fe *ChordNodeV1) Setup() error {
	if err := p2p.Inst().LaunchUp(); err != nil {
		return err
	}

	return nil
}

func (fe *ChordNodeV1) Start() {

}
func (fe *ChordNodeV1) ShutDown() {

}
