FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /go-api ./cmd

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build /go-api /usr/local/bin/go-api
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/go-api"]
