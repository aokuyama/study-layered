package admin

import (
	"github.com/aokuyama/circle_scheduler-api/packages/application/create_event/usecase"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
	"github.com/aokuyama/circle_scheduler-api/packages/infra/repository/prisma"
	"github.com/spf13/cobra"
)

var createEventCmd = &cobra.Command{
	Use:  "create_event",
	Long: "create event",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		n, _ := cmd.Flags().GetString("name")
		i := usecase.CreateEventInput{
			CircleID:  args[0],
			EventName: n,
		}

		p, err := prisma.NewPrismaClient()
		if err != nil {
			panic(err)
		}
		defer func() {
			p.Disconnect()
		}()

		f := event.EventFactoryImpl{}
		or := prisma.NewCircleRepositoryPrisma(p)
		cr := prisma.NewEventRepositoryPrisma(p)
		u := usecase.New(f, or, cr)
		out, err := u.Invoke(&i)
		if err != nil {
			panic(err)
		}
		println(out.Event.ID.String())
		println(out.Event.Path.RawValue())
		return nil
	},
}

func CreateEventCmd() *cobra.Command {
	createEventCmd.Flags().String("name", "", "event name")
	return createEventCmd
}
