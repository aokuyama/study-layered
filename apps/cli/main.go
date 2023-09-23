package main

import (
	"fmt"

	"github.com/aokuyama/circle_scheduler-api/packages/infra/repository/prisma"
	"github.com/aokuyama/circle_scheduler-api/packages/usecase/create_circle"
)

func main() {
	c := prisma.NewPrismaClient()
	r := prisma.NewCircleRepositoryPrisma(c)
	u := create_circle.New(r)
	o, err := u.Invoke(&create_circle.Input{Name: "circle1"})
	if err != nil {
		fmt.Println(err)
		println("error:")
		return
	}
	fmt.Println(o.Circle.Name.String())
}
