package main

import (
	"fmt"
	"github.com/otcChain/chord-go/node"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type GlobalConf struct {
	PidPath string
}
var globalConf *GlobalConf = &GlobalConf{}

func initConfig(){
	globalConf.PidPath = ".pid"
}

func main()  {

	initConfig()

	cfg := &node.Config{}
	if err := node.Inst().Setup(cfg); err != nil{
		panic(err)
	}

	node.Inst().Start()
	sigCh := make(chan os.Signal, 1)
	waitSignal(sigCh)
}

func waitSignal(sigCh chan os.Signal) {

	pid := strconv.Itoa(os.Getpid())
	fmt.Printf("\n>>>>>>>>>>chord node start at pid(%s)<<<<<<<<<<\n", pid)
	if err := ioutil.WriteFile(globalConf.PidPath, []byte(pid), 0644); err != nil {
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