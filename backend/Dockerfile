FROM golang:1.23-alpine AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./
RUN go mod tidy

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=build /app/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]
