FROM golang:1.14
RUN mkdir -p /go/src/first-project
COPY . /go/src/first-project
ADD . /go/src/first-project
WORKDIR /go/src/first-project
RUN go mod download
RUN go build
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
CMD CompileDaemon -log-prefix=false -build="go build" -command="./first-project"
EXPOSE 8080