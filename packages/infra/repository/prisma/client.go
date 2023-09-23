package prisma

import (
	"github.com/aokuyama/circle_scheduler-api/packages/infra/prisma/db"
)

type PrismaClient struct {
	client *db.PrismaClient
}

func NewPrismaClient() *PrismaClient {
	client := db.NewClient()
	c := PrismaClient{client}
	return &c
}
