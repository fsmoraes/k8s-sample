FROM golang:alpine AS builder

WORKDIR /go/src

COPY /src /go/src

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/go-app

FROM scratch 

COPY --from=builder /go/bin/go-app /go/bin/go-app

ENTRYPOINT ["/go/bin/go-app"]
