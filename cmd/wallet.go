package cmd

import (
	"fmt"
	"github.com/otcChain/chord-go/utils"
	"github.com/otcChain/chord-go/wallet"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"path/filepath"
)

var WalletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "chord wallet",
	Long:  `TODO::.`,
	Run:   walletAction,
}

var walletCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create chord wallet",
	Long:  `TODO::.`,
	Run:   createAcc,
	//Args:  cobra.MinimumNArgs(1),
}
var (
	password string
	baseDir  string
	forTest  bool
)

func init() {
	flags := walletCreateCmd.Flags()
	flags.StringVarP(&password, "password", "p", "",
		"chord wallet create -p[--password] [PASSWORD]")
	flags.StringVarP(&baseDir, "base", "b", DefaultBaseDir,
		"chord wallet create -b[--base] [DIRECTORY OF CHORD NODE]")
	flags.BoolVarP(&forTest, "test", "t", false, "chord wallet create -b")

	WalletCmd.AddCommand(walletCreateCmd)
}

func walletAction(c *cobra.Command, _ []string) {
	_ = c.Usage()
}

func createAcc(_ *cobra.Command, _ []string) {
	if password == "" {
		fmt.Println("Password=>")
		pw, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			panic(err)
		}
		password = string(pw)
	}

	dir := utils.BaseUsrDir(baseDir)
	var walletDir = ""
	if forTest {
		walletDir = filepath.Join(dir, string(filepath.Separator), wallet.TestKeyStoreScheme)
	} else {
		walletDir = filepath.Join(dir, string(filepath.Separator), wallet.KeyStoreScheme)
	}

	wallet.InitConfig(&wallet.Config{Dir: walletDir})
	if err := wallet.Inst().CreateNewKey(password); err != nil {
		panic(err)
	}
	fmt.Println("create success!")
}
