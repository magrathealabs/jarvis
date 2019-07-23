FROM golang:1.12

ADD . /go/src/github.com/magrathealabs/jarvis

RUN go install github.com/magrathealabs/jarvis/services/device_service
