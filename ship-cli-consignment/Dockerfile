FROM golang:alpine as builder

RUN apk update && apk upgrade && \
  apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ship-cli-consignment

# Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/consignment.json .
COPY --from=builder /app/ship-cli-consignment .

CMD ["./ship-cli-consignment"]
