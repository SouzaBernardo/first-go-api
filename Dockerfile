FROM golang:1.20

WORKDIR /go/src/SouzaBernardo/first-go-api

COPY *.go ./
COPY go.mod ./

RUN go mod tidy
RUN go build main.go

EXPOSE 8080

CMD ["./main.sh"]
ENTRYPOINT ["/go/src/SouzaBernardo/first-go-api"]