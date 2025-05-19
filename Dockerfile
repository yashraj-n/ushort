# Distroless img
FROM golang:1.24-alpine AS base

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/does-not-exist" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  go-user

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main ./cmd/main.go

FROM gcr.io/distroless/static-debian11

COPY --from=base /main .
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group

USER go-user:go-user

ENTRYPOINT ["/main"]

