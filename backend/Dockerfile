# Stage 1: Build the Go application
FROM golang:1.21.5-alpine AS builder

ARG GH_ACCESS_TOKEN
ENV GH_ACCESS_TOKEN=$GH_ACCESS_TOKEN

RUN apk add git make bash
WORKDIR /app
RUN export GOPRIVATE=github.com
RUN git config --global url.https://${GH_ACCESS_TOKEN}@github.com/.insteadOf https://github.com/
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# struct optimizer
RUN go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
RUN fieldalignment -fix ./internal/... || true

RUN make build

# Stage 2: Create the final Docker image
FROM alpine:latest

WORKDIR /app

# Copy the built Go binary from the previous stage
COPY --from=builder /app/bin/app ./app
COPY --from=builder /app/config.yaml ./
COPY --from=builder /app/db ./db
COPY --from=builder /app/i18n ./i18n
COPY --from=builder /app/internal/thirdparty/mailer/templates ./internal/thirdparty/mailer/templates