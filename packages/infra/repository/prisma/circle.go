package prisma

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/infra/prisma/db"
)

type CircleRepositoryPrisma struct {
	prisma *Prisma
}

func NewCircleRepositoryPrisma(client *Prisma) *CircleRepositoryPrisma {
	c := CircleRepositoryPrisma{client}
	return &c
}

func (r *CircleRepositoryPrisma) Save(c *circle.Circle) error {
	_, err := r.prisma.client.Circle.CreateOne(
		db.Circle.ID.Set(c.ID.String()),
		db.Circle.PathDigest.Set(c.Path.Digest()),
		db.Circle.Name.Set(c.Name.String()),
		db.Circle.Owner.Link(db.Owner.ID.Set(c.OwnerID.String())),
	).Exec(r.prisma.ctx)

	return err
}

func (r *CircleRepositoryPrisma) Find(*circle.CircleID) (*circle.Circle, error) {
	panic("not implemented")
}
