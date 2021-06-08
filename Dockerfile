FROM golang:latest

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT go run main.go
