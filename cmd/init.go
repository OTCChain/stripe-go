package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/otcChain/chord-go/consensus"
	"github.com/otcChain/chord-go/node"
	"github.com/otcChain/chord-go/p2p"
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
		"base directory for chord node")
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

type StoreCfg map[string]*CfgPerNetwork

type CfgPerNetwork struct {
	Name string            `json:"name"`
	PCfg *p2p.Config       `json:"p2p"`
	CCfg *consensus.Config `json:"consensus"`
	NCfg *node.Config      `json:"node"`
	UCfg *utils.Config     `json:"utils"`
	WCfg *wallet.Config    `json:"wallet"`
}

func initDefault(baseDir string) error {
	conf := make(StoreCfg)

	mainConf := &CfgPerNetwork{
		Name: MainNet,
		PCfg: p2p.DefaultConfig(true, baseDir),
		CCfg: consensus.InitDefaultConfig(),
		NCfg: node.InitDefaultConfig(),
		UCfg: &utils.Config{
			LogLevel: zerolog.ErrorLevel,
		},
		WCfg: wallet.DefaultConfig(true, baseDir),
	}
	conf[MainNet] = mainConf

	testConf := &CfgPerNetwork{
		Name: TestNet,
		PCfg: p2p.DefaultConfig(false, baseDir),
		CCfg: consensus.InitDefaultConfig(),
		NCfg: node.InitDefaultConfig(),
		UCfg: &utils.Config{
			LogLevel: zerolog.DebugLevel,
		},
		WCfg: wallet.DefaultConfig(false, baseDir),
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

func InitChordConfig(baseDir, network string) (err error) {
	conf := make(StoreCfg)
	dir := utils.BaseUsrDir(baseDir)
	confPath := filepath.Join(dir, string(filepath.Separator), ConfFileName)
	bts, e := os.ReadFile(confPath)
	if e != nil {
		return e
	}

	if err = json.Unmarshal(bts, &conf); err != nil {
		return
	}

	result, ok := conf[network]
	if !ok {
		err = fmt.Errorf("failed to find node config")
		return
	}

	fmt.Println(result.String())

	wallet.InitConfig(result.WCfg)
	node.InitConfig(result.NCfg)
	p2p.InitConfig(result.PCfg)
	consensus.InitConfig(result.CCfg)
	utils.InitConfig(result.UCfg)
	return
}
