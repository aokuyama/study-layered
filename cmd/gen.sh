#/bin/bash

# generate mock
cd `dirname $0`/../packages/domain
go generate ./...

# generate prisma
cd `dirname $0`/../packages/infra
go run github.com/steebchen/prisma-client-go generate
