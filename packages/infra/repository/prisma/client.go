package prisma

import (
	"context"
	"log/slog"

	"github.com/aokuyama/circle_scheduler-api/packages/infra/prisma/db"
)

type Prisma struct {
	client *db.PrismaClient
	ctx    context.Context
}

func NewPrismaClient() (*Prisma, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}
	slog.Debug("connect.")

	ctx := context.Background()
	c := Prisma{client, ctx}
	return &c, nil
}

func (p *Prisma) Disconnect() {
	if err := p.client.Prisma.Disconnect(); err != nil {
		panic(err)
	}
	slog.Debug("disconnect.")
}
