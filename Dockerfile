FROM golang:alpine as build

WORKDIR /app


COPY . .


RUN go mod download

RUN go build -o booking-room


FROM alpine:latest
WORKDIR /app


COPY --from=build /app/avengers-clinic /app/avengers-clinic

ENTRYPOINT ["/app/avengers-clinic"]