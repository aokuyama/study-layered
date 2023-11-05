package prisma

import (
	"fmt"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/prisma/db"
)

type userRepositoryPrisma struct {
	prisma *prisma
}

func NewUserRepositoryPrisma(client *prisma) *userRepositoryPrisma {
	r := userRepositoryPrisma{client}
	return &r
}

func (r *userRepositoryPrisma) Create(u *user.UserWithPassword) error {
	d := u.Password.Digest()

	_, err := r.prisma.client().User.CreateOne(
		db.User.ID.Set(u.User.ID().String()),
		db.User.PasswordDigest.Set(d[:]),
	).Exec(r.prisma.ctx)
	return err
}

func (r *userRepositoryPrisma) Find(i *user.UserID) (*user.User, error) {
	f, err := r.prisma.client().User.FindUnique(db.User.ID.Equals(i.String())).Exec(r.prisma.ctx)
	if err != nil {
		return nil, fmt.Errorf("not found\n%w", err)
	}

	i, err = user.NewUserID(f.ID)
	if err != nil {
		panic(err)
	}
	return user.NewUser(i)
}
