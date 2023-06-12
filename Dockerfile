FROM golang:1.20

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /src

COPY go.* ./
RUN go mod download

COPY main.go ./
COPY database ./
COPY . .
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
RUN 
# go build -a -installsuffix cgo -o main .
EXPOSE 3000
ENTRYPOINT CompileDaemon -build="go build -o srv main.go" -command="./srv"