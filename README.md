# Bluetooth checker for Raspberry Pi project

## How does it work?

It checks the Bluetooth connection every and trigger MQTT topic if hall light is on.

### Prerequisites

This image requires Bluetooth to be installed on Linux's machine with `sudo apt-get install bluetooth blueman bluez`. It's intended to work only on Linux.

#### Preparing base image

Base image is prepared and available here c4stus/raspberrypi:blecron-base-image, but if you want to create a new base image:

1. Prepare image from `Dockerfile.sourceimage` and run it
2. Go into it with `docker exec`
3. Run `bluetoothctl`
4. `scan on` and locate the device "26:1A:DC:1C:6D:75"
5. `connect ${device}`
6. `trust ${device}`
7. Exit container and commit the image with `docker commit -m "Source image for BLE cron" blecron-source-image c4stus/raspberrypi:blecron-base-image`
8. Push the image to docker registry with `docker push c4stus/raspberrypi:blecron-base-image` (PAT with write permission is necessary)

## Running Docker

This docker image should be run with special options:
- `volume` option to connect to a host bluetooth. Please add it like so: `/var/run/dbus/:/var/run/dbus/:z`
- `privileged` option set to `true`
