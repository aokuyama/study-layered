package main

import (
	"os"

	"github.com/aokuyama/circle_scheduler-api/apps/cli/cmd"
	"github.com/aokuyama/circle_scheduler-api/apps/cli/cmd/admin"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cmd.AdminCmd.AddCommand(admin.CreateOwnerCmd)
	cmd.AdminCmd.AddCommand(admin.CreateCircleCmd())
	rootCmd.AddCommand(cmd.AdminCmd)
}
