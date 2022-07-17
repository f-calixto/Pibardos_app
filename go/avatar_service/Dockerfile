FROM arm64v8/alpine:3.14

RUN mkdir -p /data/static/80
RUN mkdir /data/static/300
RUN mkdir /data/static/800

COPY ./bin .

EXPOSE 5001

CMD ["./main"]
