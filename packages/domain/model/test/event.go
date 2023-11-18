package test

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/util"
)

func GenEvent(n int) *event.EventEntity {
	return util.PanicOr(event.NewEventEntity(
		util.P("26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"),
		util.P("d833a112-95e8-4042-ab02-ffde48bc874a"),
		util.P("event1"),
		util.P(path.Path{}),
	))
}
