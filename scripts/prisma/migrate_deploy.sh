#/bin/bash

cd `dirname $0`/../../packages/infrastructure
go run github.com/steebchen/prisma-client-go migrate deploy
