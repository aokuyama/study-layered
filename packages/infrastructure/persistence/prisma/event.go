package prisma

import (
	"fmt"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/prisma/db"
)

type eventRepositoryPrisma struct {
	prisma *prisma
}

func NewEventRepositoryPrisma(client *prisma) *eventRepositoryPrisma {
	r := eventRepositoryPrisma{client}
	return &r
}

func (r *eventRepositoryPrisma) Create(e *event.EventEntity) error {
	d := e.Path().Digest()
	en, err := e.Path().Encrypt()
	if err != nil {
		return err
	}

	_, err = r.prisma.client().Event.CreateOne(
		db.Event.ID.Set(e.ID().String()),
		db.Event.PathDigest.Set(d[:]),
		db.Event.Path.Set(en.Data),
		db.Event.PathIv.Set(en.Iv),
		db.Event.Name.Set(e.Name().String()),
		db.Event.Circle.Link(db.Circle.ID.Set(e.CircleID().String())),
	).Exec(r.prisma.ctx)

	return err
}

func (r *eventRepositoryPrisma) FindByPath(p *path.Path) (*event.EventEntity, error) {
	d := p.Digest()
	f, err := r.prisma.client().Event.FindUnique(
		db.Event.PathDigest.Equals(d[:]),
	).Exec(r.prisma.ctx)
	if err != nil {
		return nil, fmt.Errorf("not found\n%w", err)
	}
	en := path.Encrypted{
		Data: f.Path,
		Iv:   f.PathIv,
	}
	p2, err := path.DecryptPath(&en)
	if err != nil {
		panic(err)
	}

	e, err := event.NewEventEntity(
		&f.ID,
		&f.CircleID,
		&f.Name,
		p2,
	)
	if err != nil {
		panic(err)
	}
	return e, nil
}
