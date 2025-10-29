FROM golang:latest AS build

WORKDIR /app

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN swag init

RUN CGO_ENABLED=0 GOOS=linux go build -o api

FROM golang:latest

WORKDIR /app

COPY --from=build /app/api ./

COPY .env .

ENV GIN_MODE=release

CMD [ "./api" ]

EXPOSE 8080