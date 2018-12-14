#FROM alpine:latest
#
#RUN mkdir -p /app
#WORKDIR /app
#
#ADD consignment.json /app/consignment.json
#ADD consignment-cli /app/consignment-cli
#
#CMD ["./consignment-cli"]

#FROM alpine:latest
#
#RUN mkdir /app
#WORKDIR /app
#ADD consignment-service /app/consignment-service
#
#CMD ["./consignment-service"]
#


FROM golang:1.11.2 as builder


WORKDIR /go/src/github.com/kalmeshbhavi/shipper/consignment-cli



COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

ADD consignment.json /app/consignment.json
COPY --from=builder /go/src/github.com/kalmeshbhavi/shipper/consignment-cli/consignment-cli .

CMD ["./consignment-cli"]