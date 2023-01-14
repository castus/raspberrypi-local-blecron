# Bluetooth checker for Raspberry Pi project

## How does it work?

Based on a bluetooth connection with a Raspberry Pi Pico it triggers mqtt topic that the light is on.

### Prerequisites

On your host machine you have to install `sudo apt-get install bluetooth blueman bluez` and then run:
* `bluetoothctl`
* `scan on` and locate the device
* `connect ${device}`
* `trust ${device}`

## Running Docker

This docker image should be run with special options:
- `volume` option to connect to a host bluetooth. Please add it like so: `/var/run/dbus/:/var/run/dbus/:z`
- `privileged` option set to `true`
