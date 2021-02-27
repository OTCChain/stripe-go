package node

type ChordNodeV1 struct {
}

func newNode() *ChordNodeV1 {
	n := &ChordNodeV1{}
	return n
}

func (fe *ChordNodeV1) Setup(cfg *Config) error {
	return nil
}

func (fe *ChordNodeV1) Start() {

}
func (fe *ChordNodeV1) ShutDown() {

}
