package cmd

import (
	"github.com/aokuyama/circle_scheduler-api/interfaces/cli/cmd/admin"
	"github.com/spf13/cobra"
)

var adminCmd = &cobra.Command{
	Use:  "admin",
	Long: "admin",
}

func AdminCmd() *cobra.Command {
	adminCmd.AddCommand(admin.CreateOwnerCmd())
	adminCmd.AddCommand(admin.CreateCircleCmd())
	adminCmd.AddCommand(admin.ShowCircleCmd())
	adminCmd.AddCommand(admin.CreateEventCmd())
	adminCmd.AddCommand(admin.ShowEventCmd())
	return adminCmd
}
