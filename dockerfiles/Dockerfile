FROM golang:1.15.5-alpine as builder
RUN apk update && apk add git make gcc g++
WORKDIR /go/src/app
ENV GO111MODULE=on
COPY ../go.mod .
COPY ../go.sum .
RUN go mod download
COPY ../ .
WORKDIR /go/src/app/cmd/api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app -ldflags="-s -w" .

FROM gcr.io/distroless/static:latest
COPY --from=builder /go/bin/app /go/bin/app
CMD while ! nc -z ${PSQL_HOST} ${PSQL_PORT}; do sleep 1; done && ./go/bin/app
