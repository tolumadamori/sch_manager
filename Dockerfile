FROM golang:1.20.1-alpine3.17

COPY . /go/src/sch-manager

WORKDIR /go/src/sch-manager

COPY go.mod ./

RUN go mod download

COPY . .

RUN go mod tidy

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "install", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -polling -log-prefix=false -build="go build ./cmd/main.go" -command="./main"