package prisma

import (
	"fmt"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/errs"
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

func (r *eventRepositoryPrisma) Create(e *event.Event) error {
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

func (r *eventRepositoryPrisma) Find(i *event.EventID) (*event.Event, error) {
	return nil, nil
}

func (r *eventRepositoryPrisma) FindByPath(p *path.Path) (*event.Event, error) {
	d := p.Digest()
	f, err := r.prisma.client().Event.FindUnique(
		db.Event.PathDigest.Equals(d[:]),
	).Exec(r.prisma.ctx)
	if err != nil {
		return nil, fmt.Errorf("%w %w", errs.ErrNotFound, err)
	}
	en := path.Encrypted{
		Data: f.Path,
		Iv:   f.PathIv,
	}
	p2, err := path.DecryptPath(&en)
	if err != nil {
		panic(err)
	}

	e, err := event.NewEvent(&event.EventInput{
		ID:       f.ID,
		CircleID: f.CircleID,
		Name:     f.Name,
		Path:     p2.RawValue(), // 一度値オブジェクトにしたものを文字列にして詰め直しているのがイケてないが仕方ない
	})
	if err != nil {
		panic(err)
	}
	return e, nil
}

func (r *eventRepositoryPrisma) Update(after *event.Event, before *event.Event) error {
	return nil
}
