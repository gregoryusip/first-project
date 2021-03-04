FROM golang:1.13

RUN mkdir -p /go/src/first-project
COPY . /go/src/first-project

ADD . /go/src/first-project
WORKDIR /go/src/first-project

RUN go mod download
RUN go build

CMD [ "/first-project/main" ]

EXPOSE 8080