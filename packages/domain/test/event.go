package test

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
)

func GenEvent(n int) *event.Event {
	if n == 1 {
		return PanicOr(event.NewEvent(
			P("26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"),
			P("d833a112-95e8-4042-ab02-ffde48bc874a"),
			P("event1"),
			P(path.Path{}),
		))
	}
	if n == 11 {
		return PanicOr(event.NewEvent(
			P("26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"),
			P("d833a112-95e8-4042-ab02-ffde48bc874a"),
			P("event1_1"),
			P(path.Path{}),
		))
	}
	panic("undefined factory")
}
