package prisma

import (
	"fmt"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/prisma/db"
)

type ownerRepositoryPrisma struct {
	prisma *prisma
}

func NewOwnerRepositoryPrisma(client *prisma) *ownerRepositoryPrisma {
	r := ownerRepositoryPrisma{client}
	return &r
}

func (r *ownerRepositoryPrisma) Save(o *owner.Owner) error {
	_, err := r.prisma.client().Owner.CreateOne(
		db.Owner.ID.Set(o.ID().String()),
	).Exec(r.prisma.ctx)

	return err
}

func (r *ownerRepositoryPrisma) Find(i *owner.OwnerID) (*owner.Owner, error) {
	var err error
	f, err := r.prisma.client().Owner.FindUnique(db.Owner.ID.Equals(i.String())).Exec(r.prisma.ctx)
	if err != nil {
		return nil, fmt.Errorf("not found\n%w", err)
	}
	i, err = owner.NewOwnerID(f.ID)
	if err != nil {
		panic(err)
	}
	return owner.NewOwner(i)
}
