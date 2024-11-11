FROM golang:1.22-alpine as builder

WORKDIR /app

RUN apk --no-cache add bash git make gcc gettext musl-dev

# dependencies
COPY ["go.mod", "go.sum", "static", "./"]
RUN go mod download

COPY . .
RUN go build -o /bin/main ./cmd/main.go

FROM alpine as runner

WORKDIR /app

# run
COPY --from=builder /bin/main /bin/main
COPY --from=builder /app/static /app/static

CMD ["/bin/main", "serve", "--http=0.0.0.0:8090"]