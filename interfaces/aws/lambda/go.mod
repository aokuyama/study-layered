module github.com/aokuyama/circle_scheduler-api/interfaces/aws/lambda

go 1.23.4

replace github.com/aokuyama/circle_scheduler-api/packages/domain => ./../../../packages/domain

replace github.com/aokuyama/circle_scheduler-api/packages/application => ./../../../packages/application

replace github.com/aokuyama/circle_scheduler-api/packages/infrastructure => ./../../../packages/infrastructure

require (
	github.com/aokuyama/circle_scheduler-api/packages/infrastructure v0.0.0-20241220233241-849a29516175
	github.com/aws/aws-lambda-go v1.47.0
	github.com/awslabs/aws-lambda-go-api-proxy v0.16.2
	github.com/gin-gonic/gin v1.10.0
)

require (
	github.com/aokuyama/circle_scheduler-api/packages/application v0.0.0-20241220233241-849a29516175 // indirect
	github.com/aokuyama/circle_scheduler-api/packages/domain v0.0.0-20241220233241-849a29516175 // indirect
	github.com/bytedance/sonic v1.12.7 // indirect
	github.com/bytedance/sonic/loader v0.2.2 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/gin-contrib/sse v1.0.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.23.0 // indirect
	github.com/goccy/go-json v0.10.4 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/steebchen/prisma-client-go v0.45.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	go.mongodb.org/mongo-driver/v2 v2.0.0 // indirect
	golang.org/x/arch v0.13.0 // indirect
	golang.org/x/crypto v0.32.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/protobuf v1.36.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
