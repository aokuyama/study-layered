package admin

import (
	"github.com/aokuyama/circle_scheduler-api/packages/infra/repository/prisma"
	usecase "github.com/aokuyama/circle_scheduler-api/packages/usecase/create_circle"
	"github.com/spf13/cobra"
)

var createCircleCmd = &cobra.Command{
	Use:  "create_circle",
	Long: "create circle",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		n, _ := cmd.Flags().GetString("name")
		i := usecase.CreateCircleInput{
			OwnerID:    args[0],
			CircleName: n,
		}

		p, err := prisma.NewPrismaClient()
		if err != nil {
			panic(err)
		}
		defer func() {
			p.Disconnect()
		}()

		or := prisma.NewOwnerRepositoryPrisma(p)
		cr := prisma.NewCircleRepositoryPrisma(p)
		u := usecase.New(or, cr)
		out, err := u.Invoke(&i)
		if err != nil {
			panic(err)
		}
		println(out.Circle.ID.String())
		println(out.Circle.Path.RawValue())
		return nil
	},
}

func CreateCircleCmd() *cobra.Command {
	createCircleCmd.Flags().String("name", "", "circle name")
	return createCircleCmd
}
