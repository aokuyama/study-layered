module github.com/aokuyama/circle_scheduler-api/packages/usecase

go 1.21

replace github.com/aokuyama/circle_scheduler-api/packages/domain => ./../../packages/domain

require github.com/aokuyama/circle_scheduler-api/packages/domain v0.0.0-00010101000000-000000000000

require github.com/google/uuid v1.3.0 // indirect
