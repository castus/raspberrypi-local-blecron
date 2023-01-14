#!/bin/bash

service dbus start
bluetoothd &

go mod init raspberrypi.local/blecron
go get github.com/eclipse/paho.mqtt.golang
go get github.com/gorilla/websocket
go get golang.org/x/net/proxy
go get github.com/robfig/cron/v3

go run raspberrypi.local/blecron
