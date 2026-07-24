package cmd

import (
	"github.com/hanzo-fi/go-libs/v3/service"

	"github.com/hanzo-fi/auth/pkg/storage/sqlstorage"
	"github.com/hanzo-fi/go-libs/v3/bun/bunmigrate"
	"github.com/uptrace/bun"

	"github.com/spf13/cobra"
)

var (
	Version   = "develop"
	BuildDate = "-"
	Commit    = "-"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{}

	cobra.EnableTraverseRunHooks = true

	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cmd.AddCommand(
		newServeCommand(),
		newVersionCommand(),
		bunmigrate.NewDefaultCommand(func(cmd *cobra.Command, args []string, db *bun.DB) error {
			return sqlstorage.Migrate(cmd.Context(), db)
		}))

	return cmd
}

func Execute() {
	service.Execute(NewRootCommand())
}
