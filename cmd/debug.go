package cmd

import (
	"context"
	"fmt"
	"github.com/otcChain/chord-go/pbs"
	"github.com/otcChain/chord-go/rpc/cmd"
	"github.com/spf13/cobra"
)

var DebugCmd = &cobra.Command{
	Use:   "debug",
	Short: "debug ",
	Long:  `TODO::.`,
	Run:   debug,
}

var P2pCmd = &cobra.Command{
	Use:   "p2p",
	Short: "chord debug p2p -t [TOPIC] -m [MESSAGE]",
	Long:  `TODO::.`,
	Run:   p2pAction,
}

var (
	topic   string
	msgBody string
)

func init() {
	flags := DebugCmd.Flags()
	flags.StringVarP(&topic, "topic", "t", "/global/test/",
		"chord debug p2p -t \"/global/test\" -m \"msg to send\"")
	flags.StringVarP(&msgBody, "message", "m", "",
		"chord debug p2p -t \"/global/test\" -m \"msg to send\"")
	DebugCmd.AddCommand(P2pCmd)
}

func debug(c *cobra.Command, _ []string) {
	_ = c.Usage()
}

func p2pAction(c *cobra.Command, _ []string) {
	if topic == "" || msgBody == "" {
		_ = c.Usage()
		return
	}

	cli := cmd.DialToCmdService()
	rsp, err := cli.P2PSendTopicMsg(context.Background(), &pbs.TopicMsg{
		Topic: topic,
		Msg:   msgBody,
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Msg)
}
