package main

import (
	"encoding/json"
	"fmt"
	"github.com/otcChain/chord-go/cmd"
	"github.com/otcChain/chord-go/consensus"
	"github.com/otcChain/chord-go/node"
	"github.com/otcChain/chord-go/p2p"
	"github.com/otcChain/chord-go/utils"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type SysParam struct {
	version  bool
	baseDir  string
	network  string
	password string
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
		"n", "jack", "chord -p [PASSWORD]")

	rootCmd.Flags().StringVarP(&param.password, "password",
		"p", "", "chord -p [PASSWORD]")

	rootCmd.Flags().StringVarP(&param.baseDir, "base directory",
		"d", cmd.DefaultBaseDir, "chord -d [base directory]")

	rootCmd.AddCommand(cmd.InitCmd)
	rootCmd.AddCommand(cmd.ShowCmd)
	rootCmd.AddCommand(cmd.AccCmd)
}

func InitConfig() (err error) {
	conf := make(cmd.StoreCfg)
	bts, e := os.ReadFile(param.baseDir + "/" + cmd.ConfFileName)
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

	cfg := &node.Config{}
	if err := node.Inst().Setup(cfg); err != nil {
		panic(err)
	}

	node.Inst().Start()
	sigCh := make(chan os.Signal, 1)
	waitSignal(sigCh)
}

func waitSignal(sigCh chan os.Signal) {

	pid := strconv.Itoa(os.Getpid())
	fmt.Printf("\n>>>>>>>>>>chord node start at pid(%s)<<<<<<<<<<\n", pid)
	if err := ioutil.WriteFile(param.baseDir+"/"+PidFileName, []byte(pid), 0644); err != nil {
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
