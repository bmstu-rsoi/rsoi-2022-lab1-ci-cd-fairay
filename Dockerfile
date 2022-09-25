# syntax=docker/dockerfile:1

FROM golang:1.17

WORKDIR /app

COPY src/* .
RUN go mod download

RUN mkdir -p logs
RUN mkdir -p temp

RUN go build -o ./teamdev-backend

EXPOSE 8000

CMD [ "./teamdev-backend ./configs/config.json" ]
