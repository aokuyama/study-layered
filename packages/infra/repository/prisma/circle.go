package prisma

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"

type CircleRepositoryPrisma struct {
	prisma *Prisma
}

func NewCircleRepositoryPrisma(client *Prisma) *CircleRepositoryPrisma {
	c := CircleRepositoryPrisma{client}
	return &c
}

func (r *CircleRepositoryPrisma) Save(c *circle.Circle) (*circle.Circle, error) {
	return c, nil
}

func (r *CircleRepositoryPrisma) LoadByID(*circle.CircleID) (*circle.Circle, error) {
	n := "dummy"
	c, err := circle.GenerateCircle(&n)
	return c, err
}
