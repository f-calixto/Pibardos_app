FROM arm32v7/golang:alpine3.14  

RUN mkdir -p /data/static/80
RUN mkdir /data/static/300
RUN mkdir /data/static/800

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

EXPOSE 5000

COPY . .

RUN ["go", "build", "main.go"]

CMD ["./main"]
