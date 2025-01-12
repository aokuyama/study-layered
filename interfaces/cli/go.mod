module github.com/aokuyama/circle_scheduler-api/interfaces/cli

go 1.23.4

replace github.com/aokuyama/circle_scheduler-api/packages/domain => ./../../packages/domain

replace github.com/aokuyama/circle_scheduler-api/packages/application => ./../../packages/application

replace github.com/aokuyama/circle_scheduler-api/packages/infrastructure => ./../../packages/infrastructure

require (
	github.com/aokuyama/circle_scheduler-api/packages/application v0.0.0-20241220233241-849a29516175
	github.com/aokuyama/circle_scheduler-api/packages/domain v0.0.0-20241220233241-849a29516175
	github.com/aokuyama/circle_scheduler-api/packages/infrastructure v0.0.0-20241220233241-849a29516175
	github.com/spf13/cobra v1.8.1
)

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/steebchen/prisma-client-go v0.45.0 // indirect
	go.mongodb.org/mongo-driver/v2 v2.0.0 // indirect
)
