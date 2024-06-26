package admin

import (
	"github.com/aokuyama/circle_scheduler-api/packages/application/show_circle/usecase"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/persistence/prisma"
	"github.com/spf13/cobra"
)

var showCircleCmd = &cobra.Command{
	Use:  "show_circle",
	Long: "show circle",

	RunE: func(cmd *cobra.Command, args []string) error {
		i := usecase.ShowCircleInput{
			Path: args[0],
		}

		p, err := prisma.NewPrismaClient()
		if err != nil {
			panic(err)
		}
		defer func() {
			p.Disconnect()
		}()

		cr := prisma.NewCircleRepositoryPrisma(p)
		u := usecase.New(cr)
		out, err := u.Invoke(&i)
		if err != nil {
			panic(err)
		}
		println(out.Circle.ID().String())
		println(out.Circle.Name().String())
		println(out.Circle.Path().RawValue())
		return nil
	},
}

func ShowCircleCmd() *cobra.Command {
	return showCircleCmd
}
