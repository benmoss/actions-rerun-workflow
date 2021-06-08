FROM golang:latest

WORKDIR /src
COPY . .

RUN go build -o /bin/action

ENTRYPOINT ["/bin/action"]
