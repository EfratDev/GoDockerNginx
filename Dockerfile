FROM golang:1.12.7-alpine3.10 AS build
WORKDIR /go/src/app
COPY ./main.go ./main.go
RUN go build -o ./bin/webserver ./main.go
FROM alpine:3.10
COPY --from=build /go/src/app/bin /go/bin
EXPOSE 8080
ENTRYPOINT /go/bin/webserver
