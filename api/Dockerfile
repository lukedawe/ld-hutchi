FROM golang:latest AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/api ./

RUN adduser -D api_user
USER api_user

EXPOSE 8080

CMD [ "./api" ]
