module github.com/aokuyama/circle_scheduler-api/apps/cli

go 1.21

replace github.com/aokuyama/circle_scheduler-api/packages/domain => ./../../packages/domain

replace github.com/aokuyama/circle_scheduler-api/packages/usecase => ./../../packages/usecase

replace github.com/aokuyama/circle_scheduler-api/packages/infra => ./../../packages/infra

require (
	github.com/aokuyama/circle_scheduler-api/packages/infra v0.0.0-00010101000000-000000000000
	github.com/aokuyama/circle_scheduler-api/packages/usecase v0.0.0-00010101000000-000000000000
)

require (
	github.com/aokuyama/circle_scheduler-api/packages/domain v0.0.0-00010101000000-000000000000 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/iancoleman/strcase v0.0.0-20190422225806-e506e3ef7365 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/steebchen/prisma-client-go v0.24.0 // indirect
	github.com/takuoki/gocase v1.0.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)
