#/bin/bash

cd `dirname $0`/../packages/domain
go generate ./...
