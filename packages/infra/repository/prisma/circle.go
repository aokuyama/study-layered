package prisma

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
	"github.com/aokuyama/circle_scheduler-api/packages/infra/prisma/db"
)

type CircleRepositoryPrisma struct {
	prisma *Prisma
}

func NewCircleRepositoryPrisma(client *Prisma) *CircleRepositoryPrisma {
	r := CircleRepositoryPrisma{client}
	return &r
}

func (r *CircleRepositoryPrisma) Save(c *circle.Circle) error {
	d := c.Path.Digest()
	_, err := r.prisma.client().Circle.CreateOne(
		db.Circle.ID.Set(c.ID.String()),
		db.Circle.PathDigest.Set(d[:]),
		db.Circle.Name.Set(c.Name.String()),
		db.Circle.Owner.Link(db.Owner.ID.Set(c.OwnerID.String())),
	).Exec(r.prisma.ctx)

	return err
}

func (r *CircleRepositoryPrisma) Find(i *circle.CircleID) (*circle.Circle, error) {
	f, err := r.prisma.client().Circle.FindUnique(db.Circle.ID.Equals(i.String())).Exec(r.prisma.ctx)
	if err != nil {
		return nil, err
	}
	p := path.Path{}
	return circle.NewCircle(
		(*common.UUID)(&f.ID),
		(*common.UUID)(&f.OwnerID),
		(*circle.Name)(&f.Name),
		&p,
	)
}

func (r *CircleRepositoryPrisma) FindByPath(p *path.Path) (*circle.Circle, error) {
	d := p.Digest()
	f, err := r.prisma.client().Circle.FindUnique(
		db.Circle.PathDigest.Equals(d[:]),
	).Exec(r.prisma.ctx)
	if err != nil {
		return nil, err
	}

	c, err := circle.NewCircle(
		(*common.UUID)(&f.ID),
		(*common.UUID)(&f.OwnerID),
		(*circle.Name)(&f.Name),
		p,
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}
