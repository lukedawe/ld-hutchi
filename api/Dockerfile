FROM golang:latest AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api

FROM golang:latest

WORKDIR /app

COPY --from=build /app/api ./

ENV GIN_MODE=release

CMD [ "./api" ]

EXPOSE 8080