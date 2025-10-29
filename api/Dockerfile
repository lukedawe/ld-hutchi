FROM golang:latest AS build

WORKDIR /app

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN swag init

RUN CGO_ENABLED=0 GOOS=linux GIN_MODE=release go build -o api

FROM golang:latest

WORKDIR /app

COPY --from=build /app/api ./

EXPOSE 8080