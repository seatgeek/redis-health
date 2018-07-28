FROM golang:1.10-alpine
WORKDIR /go/src/github.com/seatgeek/redis-health/
COPY . /go/src/github.com/seatgeek/redis-health/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/redis-health  .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/seatgeek/redis-health/build/redis-health .
CMD ["./redis-health"]
