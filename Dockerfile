FROM golang:1.20-alpine AS builder

# Necessary environment because this is alpine
ENV GO111MODULE=on \ 
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 

# Go Building compile phase
WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o be-melp .

# Use minimal alpine image to run the compiled binary
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/.env.dev /root/.env.dev
COPY --from=builder /app/.env.test /root/.env.test
COPY --from=builder /app/be-market /root/

CMD [ "./be-market" ]