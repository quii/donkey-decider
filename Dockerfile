FROM golang:1.14.0 AS builder
ENV CGO_ENABLED=0

WORKDIR /working
COPY main.go main.go

RUN go build -o donkey main.go

FROM scratch
COPY --from=builder /working/donkey .

ENTRYPOINT ["./donkey"]
