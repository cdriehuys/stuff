FROM golang:1.23 AS builder

WORKDIR /app

# These will be uncommented later when the application is more complicated than
# "Hello, World!"
# COPY go.mod go.sum ./
# RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /stuff ./cmd/stuff


FROM gcr.io/distroless/base-debian12 AS final

WORKDIR /

COPY --from=builder /stuff /stuff

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/stuff"]
