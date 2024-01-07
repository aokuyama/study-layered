FROM golang:1.21.3-bookworm as build

WORKDIR /app

COPY ./interfaces/aws/lambda ./interfaces/aws/lambda
COPY ./packages ./packages
COPY ./scripts ./scripts
COPY ./go.work.sum ./

RUN ./scripts/prisma/generate.sh
WORKDIR /app/interfaces/aws/lambda
RUN go build -tags lambda.norpc -o main main.go

FROM golang:1.21.3-bookworm as layer
RUN apt-get update && apt-get install unzip
RUN mkdir -p /opt
COPY ./deploy/aws/lambda/layer.zip ./
RUN unzip layer.zip -d /opt
RUN rm layer.zip

FROM public.ecr.aws/lambda/provided:al2023
COPY --from=layer /opt /opt
COPY --from=build /app/interfaces/aws/lambda/main ./main
ENTRYPOINT [ "./main" ]
