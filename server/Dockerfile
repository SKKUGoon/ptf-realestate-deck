FROM golang:1.21.5-alpine AS builder

# Necessary environment because this is alpine
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Go building comile phase
WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o melp-server .

# Use minimal alpine image to run the compiled binary
FROM alpine:latest

WORKDIR /root/

RUN apk update
RUN apk upgrade
RUN apk add curl

# Copy developing environment
COPY --from=builder /app/.env /root/

# Copy compiled program
COPY --from=builder /app/melp-server /root/

CMD [ "./melp-server" ]