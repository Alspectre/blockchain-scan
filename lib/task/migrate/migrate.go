package migratetask

import (
	"fmt"
	"goblock/db/migrate"

	"github.com/spf13/cobra"
)

type CommandDefinition struct {
	Use   string
	Short string
	Run   func(cmd *cobra.Command, args []string)
}

var Commands = []CommandDefinition{
	{
		Use:   "db:create",
		Short: "Run create database",
		Run: func(cmd *cobra.Command, args []string) {
			migrate.Create()
			fmt.Println("Database created")
		},
	},
	{
		Use:   "db:migrate",
		Short: "Run migrate table",
		Run: func(cmd *cobra.Command, args []string) {
			migrate.Migrate()
			fmt.Println("Table succefullty migrate")
		},
	},
	{
		Use:   "db:drop",
		Short: "Run vault setup",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Database dropped")
		},
	},
}
