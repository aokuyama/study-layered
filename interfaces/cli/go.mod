module github.com/aokuyama/circle_scheduler-api/interfaces/cli

go 1.21

replace github.com/aokuyama/circle_scheduler-api/packages/domain => ./../../packages/domain

replace github.com/aokuyama/circle_scheduler-api/packages/application => ./../../packages/application

replace github.com/aokuyama/circle_scheduler-api/packages/infrastructure => ./../../packages/infrastructure

require (
	github.com/aokuyama/circle_scheduler-api/packages/application v0.0.0-00010101000000-000000000000
	github.com/aokuyama/circle_scheduler-api/packages/domain v0.0.0-00010101000000-000000000000
	github.com/aokuyama/circle_scheduler-api/packages/infrastructure v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.7.0
)

require (
	github.com/google/uuid v1.3.0 // indirect
	github.com/iancoleman/strcase v0.0.0-20190422225806-e506e3ef7365 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/steebchen/prisma-client-go v0.24.0 // indirect
	github.com/takuoki/gocase v1.0.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)
