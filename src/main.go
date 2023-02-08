package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"time"

	"github.com/robfig/cron/v3"

	"raspberrypi.local/blecron/mqttHandler"
)

const (
	connectionString   = "#yes#"
	noConnectionString = "#no#"
)

var isConnected = false

func main() {
	periodicallyCheckForLightTrigger()

	log.Println("BLE Cron is up and running")

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func periodicallyCheckForLightTrigger() {
	l, _ := time.LoadLocation("Europe/Warsaw")
	now := time.Now().In(l)

	c := cron.New()
	_, err := c.AddFunc("@every 2s", func() {
		out, err := exec.Command(os.Getenv("BLUETOOTH_SCRIPT_INTERPRETER"), os.Getenv("BLUETOOTH_SCRIPT_PATH")).Output()
		if err != nil {
			log.Printf("%s", err)
		}

		output := string(out)
		if strings.Contains(output, connectionString) {
			if !isConnected {
				isConnected = true
				go mqttHandler.PublishMessage(getMessage(true))
				log.Println("Send MQTT trigger TURN ON (" + now.String() + ")")
			}
		} else if strings.Contains(output, noConnectionString) {
			if isConnected {
				isConnected = false
				go mqttHandler.PublishMessage(getMessage(false))
				log.Println("Send MQTT trigger TURN OFF (" + now.String() + ")")
			}
		}
	})
	if err != nil {
		panic(err)
	}
	c.Start()
}

func getMessage(isLightOn bool) string {
	message := mqttHandler.Message{
		IsLightOn: isLightOn,
		Place:     "hall",
	}
	m, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	return string(m)
}
