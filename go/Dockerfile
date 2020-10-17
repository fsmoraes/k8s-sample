FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY /src /go/src/app

RUN go build -o /go/bin/hello

FROM scratch

COPY --from=builder /go/bin/hello /go/bin/hello

ENTRYPOINT ["/go/bin/hello"]
