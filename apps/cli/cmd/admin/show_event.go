package admin

import (
	"github.com/aokuyama/circle_scheduler-api/packages/infra/repository/prisma"
	usecase "github.com/aokuyama/circle_scheduler-api/packages/usecase/show_event"
	"github.com/spf13/cobra"
)

var showEventCmd = &cobra.Command{
	Use:  "show_event",
	Long: "show event",

	RunE: func(cmd *cobra.Command, args []string) error {
		i := usecase.ShowEventInput{
			Path: args[0],
		}

		p, err := prisma.NewPrismaClient()
		if err != nil {
			panic(err)
		}
		defer func() {
			p.Disconnect()
		}()

		cr := prisma.NewEventRepositoryPrisma(p)
		u := usecase.New(cr)
		out, err := u.Invoke(&i)
		if err != nil {
			panic(err)
		}
		println(out.Event.ID.String())
		println(out.Event.Name.String())
		return nil
	},
}

func ShowEventCmd() *cobra.Command {
	return showEventCmd
}
