FROM golang:1.18-alpine3.16 as build-stage

COPY ./ /go/src/as-booking-consumer
WORKDIR /go/src/as-booking-consumer

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o as-booking-consumer

FROM alpine:latest as production-stage

RUN apk --no-cache add ca-certificates

COPY --from=build-stage /go/src/as-booking-consumer /as-booking-consumer
WORKDIR /as-booking-consumer

CMD ["./as-booking-consumer"]