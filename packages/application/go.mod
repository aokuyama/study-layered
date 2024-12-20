module github.com/aokuyama/circle_scheduler-api/packages/application

go 1.23.4

replace github.com/aokuyama/circle_scheduler-api/packages/domain => ./../../packages/domain

require (
	github.com/aokuyama/circle_scheduler-api/packages/domain v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.8.4
	go.uber.org/mock v0.3.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
