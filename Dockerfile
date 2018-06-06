FROM golang:alpine3.6 AS builder

RUN apk update; apk add git gcc musl-dev
RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/muchrm/science-syllabus
COPY . .
RUN dep ensure
RUN go build -o app

FROM alpine:3.6
WORKDIR /app
COPY --from=builder /go/src/github.com/muchrm/science-syllabus/app .
CMD ["/app/app"]