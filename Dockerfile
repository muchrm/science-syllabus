FROM golang:stretch

RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/src/github.com/muchrm/science-syllabus
COPY . .
RUN dep ensure