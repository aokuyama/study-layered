package prisma

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
	"github.com/aokuyama/circle_scheduler-api/packages/infra/prisma/db"
)

type OwnerRepositoryPrisma struct {
	prisma *Prisma
}

func NewOwnerRepositoryPrisma(client *Prisma) *OwnerRepositoryPrisma {
	c := OwnerRepositoryPrisma{client}
	return &c
}

func (r *OwnerRepositoryPrisma) Save(o *owner.Owner) (*owner.Owner, error) {
	_, err := r.prisma.client.Owner.CreateOne(
		db.Owner.ID.Set(o.ID.String()),
	).Exec(r.prisma.ctx)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (r *OwnerRepositoryPrisma) Find(*owner.OwnerID) (*owner.Owner, error) {
	c, err := owner.GenerateOwner()
	return c, err
}
