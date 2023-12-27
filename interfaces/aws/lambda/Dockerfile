FROM golang:1.21.3-bookworm as build

WORKDIR /app

COPY ./ ./
RUN ./scripts/prisma.sh

WORKDIR /app/interfaces/aws/lambda
RUN go build -tags lambda.norpc -o main main.go

FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /app/interfaces/aws/lambda/main ./main
ENTRYPOINT [ "./main" ]
