package prisma

import (
	"fmt"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/errs"
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
	s := user.GenerateSalt()
	d := u.Password.Digest(s)

	_, err := r.prisma.client().User.CreateOne(
		db.User.ID.Set(u.User.ID().String()),
		db.User.PasswordDigest.Set(d[:]),
		db.User.PasswordSalt.Set(s[:]),
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

func (r *userRepositoryPrisma) FindWithPasswordAuth(i *user.UserID, p *user.Password) (*user.User, error) {
	f, err := r.prisma.client().User.FindUnique(db.User.ID.Equals(i.String())).Exec(r.prisma.ctx)
	if err != nil {
		return nil, fmt.Errorf("not found\n%w", err)
	}

	i, err = user.NewUserID(f.ID)
	if err != nil {
		panic(err)
	}
	s := user.PasswordSalt(f.PasswordSalt)
	d := p.Digest(&s)
	if string(f.PasswordDigest) != string(d[:]) {
		return nil, errs.ErrUnauthorized
	}

	return user.NewUser(i)
}
