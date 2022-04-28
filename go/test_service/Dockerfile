FROM arm32v7/golang:alpine3.14

WORKDIR /app

COPY go.mod .
# COPY go.sum .

RUN go mod download

EXPOSE 5000

COPY . .

RUN ["go", "run", "main.go"]

 CMD ["./main"]