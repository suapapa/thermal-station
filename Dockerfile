# build stage
FROM golang:1.19 as builder

ENV CGO_ENABLED=0

RUN apt-get -qq update && \
    apt-get install -yqq upx

COPY . /build
WORKDIR /build

ARG BUILD_TAG=dev

RUN go build \
    -ldflags "-X main.programVer=${BUILD_TAG}"
RUN strip /build/thermal-station
RUN upx -q -9 /build/thermal-station

# ---
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/thermal-station .

ENV MQTT_USERNAME=secret
ENV MQTT_PASSWORD=secret

EXPOSE 8080

ENTRYPOINT ["./thermal-station"]
