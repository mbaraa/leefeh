FROM alpine:latest as build

RUN apk add go git

WORKDIR /app
COPY . .

RUN go get
RUN go build -ldflags="-w -s"

FROM alpine:latest as run

WORKDIR /app

COPY --from=build /app/salsa ./run
COPY --from=build /app/config.json ./

EXPOSE 8080

CMD ["./run"]
