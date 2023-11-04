module github.com/aokuyama/circle_scheduler-api/packages/infrastructure

go 1.21

replace github.com/aokuyama/circle_scheduler-api/packages/domain => ./../../packages/domain

require (
	github.com/aokuyama/circle_scheduler-api/packages/domain v0.0.0-00010101000000-000000000000
	github.com/golang-jwt/jwt/v5 v5.0.0
	github.com/iancoleman/strcase v0.0.0-20190422225806-e506e3ef7365
	github.com/joho/godotenv v1.5.1
	github.com/shopspring/decimal v1.3.1
	github.com/steebchen/prisma-client-go v0.24.0
	github.com/takuoki/gocase v1.0.0
	golang.org/x/text v0.13.0
)

require github.com/google/uuid v1.3.0 // indirect
