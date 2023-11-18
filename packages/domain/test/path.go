package test

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"

func GenPathString() string {
	p, err := path.GeneratePath()
	if err != nil {
		panic(err)
	}
	return p.RawValue()
}
