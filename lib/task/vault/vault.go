package vaulttask

import (
	"fmt"
	"goblock/config/vaultconfig"

	"github.com/spf13/cobra"
)

type CommandDefinition struct {
	Use   string
	Short string
	Run   func(cmd *cobra.Command, args []string)
}

var Commands = []CommandDefinition{
	{
		Use:   "vault:setup",
		Short: "Run vault setup",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running Setup vault ...")
			vaultconfig.Setup()

			fmt.Println("Done Setup vault")
		},
	},
}
