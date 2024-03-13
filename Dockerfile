FROM golang:alpine as build

# build stage
# create folder app
WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o avenger-clinic

# Final stage
FROM alpine
WORKDIR /app

COPY --from=build /app/avenger-clinic /app/avenger-clinic

ENV DB_HOST=db
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASS=RahasiaBanget123
ENV DB_NAME=avenger_clinic_db
ENV MAX_IDLE=1
ENV MAX_CONN=2
ENV MAX_LIFE_TIME=1h

ENV PORT=8080
ENV LOG_MODE=1

ENTRYPOINT ["/app/avenger-clinic"]