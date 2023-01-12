# Bluetooth checker for Raspberry Pi project

## How does it work?

Based on a bluetooth connection with a Raspberry Pi Pico it triggers mqtt topic that the light is on.

Docker images is based on an images that contains bluetooth connection already set up with trusted device.

To do it manually you have to install `sudo apt-get install bluetooth bluez-utils blueman bluez` and then run:
* `bluetoothctl`
* `scan on` and locate the device
* `connect ${device}`
* `trust ${device}`

Important! Before running docker image, you have to kill bluetooth process on the host: `sudo killall -9 bluetoothd`
