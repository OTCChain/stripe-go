package p2p

import (
	"fmt"
	"testing"
)

func TestInst(t *testing.T) {
	config := DefaultConfig(false)
	InitConfig(config)
	_inst := Inst()
	fmt.Printf("%x\n", _inst)
	_instS := Inst()
	if _inst != _instS {
		t.Fatal("single instance failed")
	}

}
