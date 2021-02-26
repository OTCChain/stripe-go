package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/otcChain/chord-go/account"
	"github.com/otcChain/chord-go/consensus"
	"github.com/otcChain/chord-go/node"
	"github.com/otcChain/chord-go/p2p"
	"github.com/otcChain/chord-go/utils"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
)

const ConfFileName = "config.json"
const DefaultBaseDir = "~/.chord"

var param struct {
	password string
	baseDir  string
}

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "init chord node",
	Long:  `TODO::.`,
	Run:   initNode,
	//Args:  cobra.MinimumNArgs(1),
}

func init() {
	InitCmd.Flags().StringVarP(&param.password, "password", "p", "",
		"Password to init chord node wallet")

	InitCmd.Flags().StringVarP(&param.baseDir, "baseDir", "d", "",
		"Password to init chord node wallet")
}

func initNode(_ *cobra.Command, _ []string) {
	dir := param.baseDir
	if dir == "" {
		dir = DefaultBaseDir
	}

	if utils.FileExists(dir) {
		panic("duplicate init operation! please save the old config or use -baseDir for new node config")
	}

	var pwd = param.password
	if pwd == "" {
		fmt.Println("Password=>")
		pw, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			panic(err)
		}
		pwd = string(pw)
	}

	if err := account.NewAccount(param.baseDir, pwd); err != nil {
		panic(err)
	}

	defaultConf := &StoreCfg{
		PCfg: &p2p.Config{},
		CCfg: &consensus.Config{},
		NCfg: &node.Config{},
	}
	bts, err := json.MarshalIndent(defaultConf, "", "\t")
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(param.baseDir+"/"+ConfFileName, bts, 0644); err != nil {
		panic(err)
	}
}

type StoreCfg struct {
	PCfg *p2p.Config       `json:"p2p"`
	CCfg *consensus.Config `json:"consensus"`
	NCfg *node.Config      `json:"node"`
}

func (c StoreCfg) String() string {
	s := fmt.Sprintf("\n===================System config===========================")
	s += c.NCfg.String()
	s += c.PCfg.String()
	s += c.CCfg.String()
	s += fmt.Sprintf("\n===========================================================")
	return s
}
