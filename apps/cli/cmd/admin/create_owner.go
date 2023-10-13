package admin

import (
	"github.com/aokuyama/circle_scheduler-api/packages/infra/repository/prisma"
	"github.com/aokuyama/circle_scheduler-api/packages/usecase/admin_create_owner"
	"github.com/spf13/cobra"
)

var CreateOwnerCmd = &cobra.Command{
	Use:  "create_owner",
	Long: "create owner",

	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := prisma.NewPrismaClient()
		if err != nil {
			panic(err)
		}
		defer func() {
			p.Disconnect()
		}()

		o := prisma.NewOwnerRepositoryPrisma(p)
		u := admin_create_owner.New(o)
		out, err := u.Invoke(&admin_create_owner.Input{})
		if err != nil {
			panic(err)
		}
		println(out.Owner.ID.String())
		return nil
	},
}
