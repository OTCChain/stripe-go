package main

import (
	"fmt"
	"github.com/otcChain/chord-go/cmd"
	"github.com/otcChain/chord-go/node"
	rpcCmd "github.com/otcChain/chord-go/rpc/cmd"
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
	rootCmd.AddCommand(cmd.WalletCmd)
	rootCmd.AddCommand(cmd.DebugCmd)
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

	if err := cmd.InitConfig(param.baseDir, param.network); err != nil {
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

	go rpcCmd.StartCmdService()
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
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	sig := <-sigCh
	node.Inst().ShutDown()
	fmt.Printf("\n>>>>>>>>>>process finished(%s)<<<<<<<<<<\n", sig)
}
