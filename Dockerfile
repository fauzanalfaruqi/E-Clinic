FROM golang:alpine as build

# build stage
# create folder app
WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o avengers-clinic

# Final stage
FROM alpine
WORKDIR /app

COPY --from=build /app/avengers-clinic /app/avengers-clinic

ENTRYPOINT ["/app/avengers-clinic"]