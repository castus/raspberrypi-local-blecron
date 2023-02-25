package main

import (
	"encoding/json"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"

	"raspberrypi.local/blecron/mqttHandler"
)

const (
	connectionString   = "#yes#"
	noConnectionString = "#no#"
)

var isConnected = false

const (
	place = "hall"
)

func main() {
	periodicallyCheckForLightTrigger()

	log.Println("type=success msg=\"BLE Cron is up and running\"")

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func periodicallyCheckForLightTrigger() {
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
				log.Printf("type=light-status place=%s is-on=true\n", place)
			}
		} else if strings.Contains(output, noConnectionString) {
			if isConnected {
				isConnected = false
				go mqttHandler.PublishMessage(getMessage(false))
				log.Printf("type=light-status place=%s is-on=false\n", place)
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
		Place:     place,
	}
	m, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	return string(m)
}
