FROM golang:1.24.4-alpine AS builder
WORKDIR /app
COPY . .

RUN apk add --no-cache upx ca-certificates

RUN go mod download
RUN go mod verify
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build \
    -o backend -a -ldflags '-s -w' -installsuffix cgo

RUN upx -9 backend

FROM alpine:latest
WORKDIR /app

RUN apk add --no-cache ca-certificates
COPY --from=builder /app/backend /app/

CMD ["/app/backend"]