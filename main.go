package main

import (
	"encoding/json"
	"fmt"
	"github.com/otcChain/chord-go/cmd"
	"github.com/otcChain/chord-go/consensus"
	"github.com/otcChain/chord-go/internal"
	"github.com/otcChain/chord-go/node"
	"github.com/otcChain/chord-go/p2p"
	"github.com/otcChain/chord-go/rpc"
	"github.com/otcChain/chord-go/utils"
	"github.com/otcChain/chord-go/utils/fdlimit"
	"github.com/otcChain/chord-go/wallet"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

type SysParam struct {
	version  bool
	baseDir  string
	network  string
	password string
	keyAddr  string
	httpIP   string
	httpPort int16
	wsIP     string
	wsPort   int16
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

	flags := rootCmd.Flags()

	flags.BoolVarP(&param.version, "version",
		"v", false, "chord -v")

	flags.StringVarP(&param.network, "network",
		"n", cmd.TestNet,
		"chord -n|--network ["+cmd.MainNet+"|"+cmd.TestNet+"] default is "+cmd.TestNet+".")

	flags.StringVarP(&param.password, "password",
		"p", "", "chord -p [PASSWORD OF SELECTED KEY]")

	flags.StringVarP(&param.keyAddr, "key",
		"k", "", "chord -k [ADDRESS OF KEY]")

	flags.StringVarP(&param.baseDir, "dir",
		"d", cmd.DefaultBaseDir, "chord -d [BASIC DIRECTORY]")

	flags.StringVar(&param.httpIP, "http.ip", "",
		"chord --http.ip=[IP]")
	flags.Int16Var(&param.httpPort, "http.port", -1,
		"chord --http.port=[Port]")
	flags.StringVar(&param.wsIP, "ws.IP", "",
		"chord --ws.ip=[Port]")
	flags.Int16Var(&param.wsPort, "ws.port", -1,
		"chord --ws.port=[Port]")

	rootCmd.AddCommand(cmd.InitCmd)
	rootCmd.AddCommand(cmd.ShowCmd)
	rootCmd.AddCommand(cmd.WalletCmd)
	rootCmd.AddCommand(cmd.DebugCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func initChordConfig() (err error) {

	conf := make(cmd.StoreCfg)
	dir := utils.BaseUsrDir(param.baseDir)
	confPath := filepath.Join(dir, string(filepath.Separator), cmd.ConfFileName)
	bts, e := os.ReadFile(confPath)
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
	rpc.InitConfig(result.RCfg)

	if param.httpPort != -1 {
		result.RCfg.HttpEnabled = true
		result.RCfg.HttpPort = param.httpPort
		if param.httpIP != "" {
			result.RCfg.HttpIP = param.httpIP
		}
	}

	if param.wsPort != -1 {
		result.RCfg.WsEnabled = true
		result.RCfg.WsPort = param.wsPort
		if param.wsIP != "" {
			result.RCfg.WsIP = param.wsIP
		}
	}

	rpc.InitConfig(result.RCfg)
	return
}

func initWalletKey() error {
	var pwd = param.password
	if pwd == "" {
		fmt.Println("Password=>")
		pw, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return err
		}
		pwd = string(pw)
	}

	if err := wallet.Inst().Active(pwd, param.keyAddr); err != nil {
		return err
	}

	return nil
}

func initSystem() error {

	if err := os.Setenv("GODEBUG", "netdns=go"); err != nil {
		return err
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(int64(time.Now().Nanosecond()))
	limit, err := fdlimit.Maximum()
	if err != nil {
		return fmt.Errorf("failed to retrieve file descriptor allowance:%s", err)
	}
	_, err = fdlimit.Raise(uint64(limit))
	if err != nil {
		return fmt.Errorf("failed to raise file descriptor allowance:%s", err)
	}
	return nil
}

func mainRun(_ *cobra.Command, _ []string) {

	if param.version {
		fmt.Println(Version)
		return
	}

	if err := initSystem(); err != nil {
		panic(err)
	}

	if err := initChordConfig(); err != nil {
		panic(err)
	}

	if err := initWalletKey(); err != nil {
		panic(err)
	}

	if err := rpc.Inst().StartService(); err != nil {
		panic(err)
	}

	go internal.StartRpc()

	if err := node.Inst().Setup(); err != nil {
		panic(err)
	}

	node.Inst().Start()

	waitShutdownSignal()
}

func waitShutdownSignal() {
	sigCh := make(chan os.Signal, 1)

	pid := strconv.Itoa(os.Getpid())
	fmt.Printf("\n>>>>>>>>>>chord node start at pid(%s)<<<<<<<<<<\n", pid)
	path := filepath.Join(utils.BaseUsrDir(param.baseDir), string(filepath.Separator), PidFileName)
	if err := ioutil.WriteFile(path, []byte(pid), 0644); err != nil {
		fmt.Print("failed to write running pid", err)
	}

	signal.Notify(sigCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	sig := <-sigCh
	node.Inst().ShutDown()
	fmt.Printf("\n>>>>>>>>>>process finished(%s)<<<<<<<<<<\n", sig)
}
