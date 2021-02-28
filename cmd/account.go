package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

var accParam struct {
	password string
}

var AccCmd = &cobra.Command{
	Use:   "account",
	Short: "create chord account",
	Long:  `TODO::.`,
	Run:   showAccounts,
	//Args:  cobra.MinimumNArgs(1),
}

var AccCreateCmd = &cobra.Command{
	Use:   "account",
	Short: "create chord account",
	Long:  `TODO::.`,
	Run:   createAcc,
	//Args:  cobra.MinimumNArgs(1),
}

func init() {
	AccCreateCmd.Flags().StringVarP(&accParam.password, "password", "p", "",
		"Password to create chord node account")
	AccCmd.AddCommand(AccCreateCmd)
}

func showAccounts(_ *cobra.Command, _ []string) {
}

func createAcc(_ *cobra.Command, _ []string) {
	var pwd = accParam.password
	if pwd == "" {
		fmt.Println("Password=>")
		pw, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			panic(err)
		}
		pwd = string(pw)
	}
}
