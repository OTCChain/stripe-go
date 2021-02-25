package consensus

type FBFTEngine struct {

}

func newEngine() *FBFTEngine {
	fe := &FBFTEngine{}
	return fe
}

func (fe *FBFTEngine) Setup(cfg *Config) error {
	return nil
}
