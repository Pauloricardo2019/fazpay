FROM golang:1.21-alpine as builder

WORKDIR /app

COPY . .
COPY .env .

RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.8
RUN swag init --parseInternal --parseDepth 1 -g ./cmd/api/main.go

# Set environment variables for build
RUN go env -w CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Finally, do the build
RUN go build -a -tags netgo \
    -ldflags '-w -extldflags "-static"'\
    -o /app/api.bin ./cmd/api/main.go

FROM alpine:3.15 as release

RUN apk add --no-cache bash \
    && adduser --disabled-password --gecos "" --no-create-home app

COPY --from=builder --chown=app /app/*.bin /app/
COPY --from=builder --chown=app /app/.env /app/


WORKDIR /app

USER app

CMD ./api.bin
