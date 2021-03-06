package cmd

import (
	"encoding/json"
	"github.com/otcChain/chord-go/consensus"
	"github.com/otcChain/chord-go/node"
	"github.com/otcChain/chord-go/p2p"
	"github.com/otcChain/chord-go/rpc"
	"github.com/otcChain/chord-go/utils"
	"github.com/otcChain/chord-go/wallet"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

const (
	DefaultBaseDir = ".chord"
	ConfFileName   = "config.json"
	MainNet        = "main"
	TestNet        = "test"
)

var param struct {
	baseDir     string
	servicePort *int16
}

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "init chord node",
	Long:  `TODO::.`,
	Run:   initNode,
	//Args:  cobra.MinimumNArgs(1),
}

func init() {
	flags := InitCmd.Flags()
	flags.StringVarP(&param.baseDir, "baseDir", "d", DefaultBaseDir,
		"init -d [Directory]")
}

func initNode(_ *cobra.Command, _ []string) {
	dir := utils.BaseUsrDir(param.baseDir)
	if utils.FileExists(dir) {
		panic("duplicate init operation! please save the old config or use -baseDir for new node config")
	}

	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		panic(err)
	}

	if err := initDefault(dir); err != nil {
		panic(err)
	}
}

func initDefault(baseDir string) error {
	conf := make(StoreCfg)

	mainConf := &CfgPerNetwork{
		Name: MainNet,
		PCfg: p2p.DefaultConfig(true, baseDir),
		CCfg: consensus.DefaultConfig(true, baseDir),
		NCfg: node.DefaultConfig(true, baseDir),
		UCfg: &utils.Config{
			LogLevel: zerolog.ErrorLevel,
		},
		WCfg: wallet.DefaultConfig(true, baseDir),
		RCfg: rpc.DefaultConfig(),
	}
	conf[MainNet] = mainConf

	testConf := &CfgPerNetwork{
		Name: TestNet,
		PCfg: p2p.DefaultConfig(false, baseDir),
		CCfg: consensus.DefaultConfig(false, baseDir),
		NCfg: node.DefaultConfig(false, baseDir),
		UCfg: &utils.Config{
			LogLevel: zerolog.DebugLevel,
		},
		WCfg: wallet.DefaultConfig(false, baseDir),
		RCfg: rpc.DefaultConfig(),
	}
	conf[TestNet] = testConf

	bts, err := json.MarshalIndent(conf, "", "\t")
	if err != nil {
		panic(err)
	}
	path := filepath.Join(baseDir, string(filepath.Separator), ConfFileName)
	if err := os.WriteFile(path, bts, 0644); err != nil {
		panic(err)
	}
	return nil
}
