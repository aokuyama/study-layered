module github.com/aokuyama/circle_scheduler-api/apps/cli

go 1.20

replace github.com/aokuyama/circle_scheduler-api/packages/domain => ./../../packages/domain

replace github.com/aokuyama/circle_scheduler-api/packages/usecase => ./../../packages/usecase

require (
	github.com/aokuyama/circle_scheduler-api/packages/domain v0.0.0-00010101000000-000000000000
	github.com/aokuyama/circle_scheduler-api/packages/usecase v0.0.0-00010101000000-000000000000
)

require github.com/google/uuid v1.3.0 // indirect
