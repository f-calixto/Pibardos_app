FROM arm32v7/golang:alpine3.14  

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

EXPOSE 5003

COPY . .

RUN ["go", "build", "main.go"]

CMD ["./main"]
