package main

import (
	"fmt"
	migratetask "goblock/lib/task/migrate"
	vaulttask "goblock/lib/task/vault"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	for _, command := range vaulttask.Commands {
		cobraCmd := &cobra.Command{
			Use:   command.Use,
			Short: command.Short,
			Run:   command.Run,
		}

		rootCmd.AddCommand(cobraCmd)
	}

	for _, command := range migratetask.Commands {
		cobraCmd := &cobra.Command{
			Use:   command.Use,
			Short: command.Short,
			Run:   command.Run,
		}

		rootCmd.AddCommand(cobraCmd)
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Available tasks:")
			for _, c := range vaulttask.Commands {
				fmt.Printf(" - %s: %s\n", c.Use, c.Short)
			}
			for _, c := range migratetask.Commands {
				fmt.Printf(" - %s: %s\n", c.Use, c.Short)
			}
		},
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
