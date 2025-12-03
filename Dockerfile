# build stage
FROM golang:1.20-alpine AS build
WORKDIR /src
# install git for module fetches if needed
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app .

# final stage
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build /src/app /app/app
EXPOSE 9595
ENTRYPOINT ["/app/app"]