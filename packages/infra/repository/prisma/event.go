package prisma

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
	"github.com/aokuyama/circle_scheduler-api/packages/infra/prisma/db"
)

type EventRepositoryPrisma struct {
	prisma *Prisma
}

func NewEventRepositoryPrisma(client *Prisma) *EventRepositoryPrisma {
	r := EventRepositoryPrisma{client}
	return &r
}

func (r *EventRepositoryPrisma) Create(e *event.RegisterEvent) error {
	d := e.Path.Digest()
	_, err := r.prisma.client().Event.CreateOne(
		db.Event.ID.Set(e.ID.String()),
		db.Event.PathDigest.Set(d[:]),
		db.Event.Name.Set(e.Name.String()),
		db.Event.Circle.Link(db.Circle.ID.Set(e.CircleID.String())),
	).Exec(r.prisma.ctx)

	return err
}

func (r *EventRepositoryPrisma) FindByPath(p *path.Path) (*event.EventEntity, error) {
	d := p.Digest()
	f, err := r.prisma.client().Event.FindUnique(
		db.Event.PathDigest.Equals(d[:]),
	).Exec(r.prisma.ctx)
	if err != nil {
		return nil, err
	}

	e, err := event.NewEventEntity(
		(*common.UUID)(&f.ID),
		(*common.UUID)(&f.CircleID),
		(*event.Name)(&f.Name),
	)
	if err != nil {
		return nil, err
	}
	return e, nil
}
