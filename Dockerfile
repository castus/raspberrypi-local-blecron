FROM --platform=linux/arm/v7 c4stus/raspberrypi:blecron-base-image AS builder
ARG TARGETPLATFORM
ARG BUILDPLATFORM
RUN echo "I am running on $BUILDPLATFORM, building for $TARGETPLATFORM"

WORKDIR /data
COPY ./src /data
RUN sh go-init.sh
RUN go build -o blecron

FROM builder
WORKDIR /root/
COPY --from=builder /data/blecron ./
COPY --from=builder /data/isDeviceConnected.py ./
CMD ["./blecron"]
