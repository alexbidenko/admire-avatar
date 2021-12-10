FROM node:16-alpine as build-client

WORKDIR /app
COPY client .
RUN yarn
RUN yarn build

FROM golang:1.17-alpine AS build-server

ARG VERSION
ARG POSTGRES_PASSWORD
ARG ACCESS_TOKEN_SECRET_KEY
ARG REFRESH_TOKEN_SECRET_KEY

ENV GOPATH="/go/src"

WORKDIR /go/src/application

RUN apk --no-cache add gcc g++ make openssh-client
RUN apk add git
COPY server .

RUN GOOS=linux go build -ldflags="-X admire-avatar/config.Version=v${VERSION} -X admire-avatar/config.PostgresPassword=${POSTGRES_PASSWORD} -X admire-avatar/config.AccessSecret=${ACCESS_TOKEN_SECRET_KEY} -X admire-avatar/config.RefreshSecret=${REFRESH_TOKEN_SECRET_KEY}" -o main .

FROM docker/compose:latest

RUN apk add openssh-client git ca-certificates

WORKDIR /app

COPY --from=build-client /app/dist dist
COPY --from=build-server /go/src/application/entrypoint.sh .
COPY --from=build-server /go/src/application/main .

RUN chmod +x ./entrypoint.sh

ENTRYPOINT  ["./entrypoint.sh"]
