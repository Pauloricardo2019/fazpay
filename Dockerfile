FROM golang:1.21-alpine as builder

WORKDIR /app

COPY . .
COPY .env .

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
