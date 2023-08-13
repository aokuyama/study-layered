package main

import (
	"fmt"

	"github.com/aokuyama/circle_scheduler-api/apps/cli/registry"
	"github.com/aokuyama/circle_scheduler-api/packages/usecase/create_circle"
)

func main() {
	r := registry.Dummy{}
	u := create_circle.New(&r)
	o, err := u.Invoke(&create_circle.Input{})
	if err != nil {
		fmt.Println(err)
		println("error:")
		fmt.Println(err)
		return
	}
	fmt.Println(o.Circle.Name.String())
}
