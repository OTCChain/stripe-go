package main

import (
	"encoding/json"
	"fmt"
	"github.com/otcChain/chord-go/cmd"
	"github.com/otcChain/chord-go/consensus"
	"github.com/otcChain/chord-go/node"
	"github.com/otcChain/chord-go/p2p"
	"github.com/otcChain/chord-go/utils"
	"github.com/otcChain/chord-go/wallet"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
)

type SysParam struct {
	version  bool
	baseDir  string
	network  string
	password string
	keyAddr  string
}

const (
	PidFileName = "pid"
	Version     = "0.1.0"
)

var (
	param = &SysParam{}
)

var rootCmd = &cobra.Command{
	Use: "chord",

	Short: "chord is a decentralized OTC system",

	Long: `usage description::TODO::`,

	Run: mainRun,
}

func init() {

	rootCmd.Flags().BoolVarP(&param.version, "version",
		"v", false, "chord -v")

	rootCmd.Flags().StringVarP(&param.network, "network",
		"n", cmd.TestNet,
		"chord -n|--network ["+cmd.MainNet+"|"+cmd.TestNet+"] default is "+cmd.TestNet+".")

	rootCmd.Flags().StringVarP(&param.password, "password",
		"p", "", "chord -p [PASSWORD OF SELECTED KEY]")

	rootCmd.Flags().StringVarP(&param.keyAddr, "key",
		"k", "", "chord -k [ADDRESS OF KEY]")

	rootCmd.Flags().StringVarP(&param.baseDir, "dir",
		"d", cmd.DefaultBaseDir, "chord -d [BASIC DIRECTORY]")

	rootCmd.AddCommand(cmd.InitCmd)
	rootCmd.AddCommand(cmd.ShowCmd)
	rootCmd.AddCommand(cmd.AccCmd)
}

func InitConfig() (err error) {
	conf := make(cmd.StoreCfg)
	dir := utils.BaseUsrDir(param.baseDir)
	bts, e := os.ReadFile(dir + "/" + cmd.ConfFileName)
	if e != nil {
		return e
	}

	if err = json.Unmarshal(bts, &conf); err != nil {
		return
	}

	result, ok := conf[param.network]
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

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func mainRun(_ *cobra.Command, _ []string) {
	if param.version {
		fmt.Println(Version)
		return
	}

	if err := InitConfig(); err != nil {
		panic(err)
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

	if err := wallet.Inst().Active(pwd, param.keyAddr); err != nil {
		panic(err)
	}

	if err := node.Inst().Setup(); err != nil {
		panic(err)
	}

	node.Inst().Start()
	sigCh := make(chan os.Signal, 1)
	waitSignal(sigCh)
}

func waitSignal(sigCh chan os.Signal) {

	pid := strconv.Itoa(os.Getpid())
	fmt.Printf("\n>>>>>>>>>>chord node start at pid(%s)<<<<<<<<<<\n", pid)
	path := filepath.Join(utils.BaseUsrDir(param.baseDir), string(filepath.Separator), PidFileName)
	if err := ioutil.WriteFile(path, []byte(pid), 0644); err != nil {
		fmt.Print("failed to write running pid", err)
	}

	signal.Notify(sigCh,
		//syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	sig := <-sigCh
	node.Inst().ShutDown()
	fmt.Printf("\n>>>>>>>>>>process finished(%s)<<<<<<<<<<\n", sig)
}
