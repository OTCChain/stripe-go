package cmd

import (
	"fmt"
	"github.com/otcChain/chord-go/chain"
	"github.com/otcChain/chord-go/consensus"
	"github.com/otcChain/chord-go/p2p"
	"github.com/otcChain/chord-go/rpc"
	"github.com/otcChain/chord-go/utils"
	"github.com/otcChain/chord-go/wallet"
)

type StoreCfg map[string]*CfgPerNetwork

type CfgPerNetwork struct {
	Name string            `json:"name"`
	PCfg *p2p.Config       `json:"p2p"`
	CCfg *consensus.Config `json:"consensus"`
	NCfg *chain.Config     `json:"node"`
	UCfg *utils.Config     `json:"utils"`
	WCfg *wallet.Config    `json:"wallet"`
	RCfg *rpc.Config       `json:"rpc"`
}

func (sc StoreCfg) DebugPrint() {
	for _, c := range sc {
		fmt.Println(c)
	}
}

func (c CfgPerNetwork) String() string {
	s := fmt.Sprintf("\n<<<===================System[%s] Config===========================", c.Name)
	s += c.NCfg.String()
	s += c.PCfg.String()
	s += c.CCfg.String()
	s += c.UCfg.String()
	s += c.WCfg.String()
	s += fmt.Sprintf("\n======================================================================>>>")
	return s
}
