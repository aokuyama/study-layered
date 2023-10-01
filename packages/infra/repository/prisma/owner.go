package prisma

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"

type OwnerRepositoryPrisma struct {
	prisma *Prisma
}

func NewOwnerRepositoryPrisma(client *Prisma) *OwnerRepositoryPrisma {
	c := OwnerRepositoryPrisma{client}
	return &c
}

func (r *OwnerRepositoryPrisma) Save(c *owner.Owner) (*owner.Owner, error) {
	return c, nil
}

func (r *OwnerRepositoryPrisma) LoadByID(*owner.OwnerID) (*owner.Owner, error) {
	c, err := owner.GenerateOwner()
	return c, err
}
