FROM golang:1.22-alpine as builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

# dependencies
COPY ["go.mod", "go.sum", "static", "./"]
RUN go mod download

COPY . .
RUN go build -o /bin/main ./cmd/main.go

FROM alpine as runner

# run
COPY --from=builder /bin/main /bin/main

CMD ["/bin/main"]