package prisma

import (
	"context"
	"log/slog"

	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/prisma/db"
)

type Prisma struct {
	_client *db.PrismaClient
	ctx     context.Context
}

func NewPrismaClient() (*Prisma, error) {
	ctx := context.Background()
	c := Prisma{nil, ctx}
	return &c, nil
}

func (p *Prisma) client() *db.PrismaClient {
	if p._client != nil {
		return p._client
	}
	p._client = db.NewClient()
	if err := p._client.Prisma.Connect(); err != nil {
		panic(err)
	}
	slog.Debug("connect.")
	return p._client
}

func (p *Prisma) Disconnect() {
	if p._client == nil {
		return
	}
	if err := p._client.Prisma.Disconnect(); err != nil {
		panic(err)
	}
	slog.Debug("disconnect.")
}
