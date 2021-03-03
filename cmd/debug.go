package cmd

import "github.com/spf13/cobra"

var DebugCmd = &cobra.Command{
	Use:   "debug",
	Short: "debug ",
	Long:  `TODO::.`,
	Run:   debug,
}

var P2pCmd = &cobra.Command{
	Use:   "p2p",
	Short: "debug ",
	Long:  `TODO::.`,
	Run:   debug,
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
