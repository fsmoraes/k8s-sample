FROM golang:alpine AS builder

WORKDIR /go/src

COPY /src /go/src

RUN CGO_ENABLED=0 GOOS=linux go build -o ./out/go-app .

FROM scratch 

COPY --from=builder /go/src/out/go-app /app/go-app

ENTRYPOINT ["/app/go-app"]
