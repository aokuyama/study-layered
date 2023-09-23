package prisma

import (
	"github.com/aokuyama/circle_scheduler-api/packages/infra/prisma/db"
)

type Prisma struct {
	client *db.PrismaClient
}

func NewPrismaClient() *Prisma {
	client := db.NewClient()
	c := Prisma{client}
	return &c
}
