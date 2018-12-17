FROM golang:1.11.2 as builder

WORKDIR /go/src/github.com/kalmeshbhavi/shipper/vessel-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -v -vendor-only
RUN CGO_ENABLED=0 GOOS=linux go build -o vessel-service -a -installsuffix cgo main.go repository.go handler.go datastore.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/kalmeshbhavi/shipper/vessel-service/vessel-service .

CMD ["./vessel-service"]
