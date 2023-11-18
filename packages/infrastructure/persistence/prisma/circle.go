package prisma

import (
	"fmt"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/prisma/db"
)

type circleRepositoryPrisma struct {
	prisma *prisma
}

func NewCircleRepositoryPrisma(client *prisma) *circleRepositoryPrisma {
	r := circleRepositoryPrisma{client}
	return &r
}

func (r *circleRepositoryPrisma) Create(c *circle.Circle) error {
	d := c.Path().Digest()
	en, err := c.Path().Encrypt()
	if err != nil {
		return err
	}

	_, err = r.prisma.client().Circle.CreateOne(
		db.Circle.ID.Set(c.ID().String()),
		db.Circle.PathDigest.Set(d[:]),
		db.Circle.Path.Set(en.Data),
		db.Circle.PathIv.Set(en.Iv),
		db.Circle.Name.Set(c.Name().String()),
		db.Circle.Owner.Link(db.Owner.ID.Set(c.OwnerID().String())),
	).Exec(r.prisma.ctx)

	return err
}

func (r *circleRepositoryPrisma) Find(i *circle.CircleID) (*circle.Circle, error) {
	f, err := r.prisma.client().Circle.FindUnique(db.Circle.ID.Equals(i.String())).Exec(r.prisma.ctx)
	if err != nil {
		return nil, fmt.Errorf("not found\n%w", err)
	}
	en := path.Encrypted{
		Data: f.Path,
		Iv:   f.PathIv,
	}
	path, err := path.DecryptPath(&en)
	if err != nil {
		panic(err)
	}

	c, err := circle.NewCircle(&f.ID, &f.OwnerID, &f.Name, path)
	if err != nil {
		panic(err)
	}
	return c, nil
}

func (r *circleRepositoryPrisma) FindByPath(p *path.Path) (*circle.Circle, error) {
	d := p.Digest()
	f, err := r.prisma.client().Circle.FindUnique(
		db.Circle.PathDigest.Equals(d[:]),
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

	c, err := circle.NewCircle(&f.ID, &f.OwnerID, &f.Name, p2)
	if err != nil {
		panic(err)
	}
	return c, nil
}

func (r *circleRepositoryPrisma) SearchByOwner(i *owner.OwnerID) (*[]circle.CircleID, error) {
	f, err := r.prisma.client().Circle.FindMany(
		db.Circle.OwnerID.Equals(i.String()),
	).Exec(r.prisma.ctx)
	if err != nil {
		return nil, err
	}
	cl := make([]circle.CircleID, len(f))
	for i := range f {
		ci, err := circle.NewCircleID(f[i].ID)
		if err != nil {
			panic(err)
		}
		cl[i] = *ci
	}
	return &cl, nil
}
