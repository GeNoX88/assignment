FROM golang:1.17.3-alpine
WORKDIR /

COPY main.go index.js index.html go.sum go.mod favicon.ico ./

RUN go mod download

RUN go build main.go
ENTRYPOINT [ "go", "run", "main.go" ]
