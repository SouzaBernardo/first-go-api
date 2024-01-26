FROM golang

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

COPY . .

RUN go build main.go

EXPOSE 8080

ENTRYPOINT ["./main"]