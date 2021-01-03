package main

import (
	"common/messagequeue"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

const configFile = "../hotelpublisher/config.yaml"

var mq *messagequeue.RabbitMq

func init() {
	log.SetLevel(log.DebugLevel)
	err := loadConfig(configFile)
	if err != nil {
		panic(err)
	}
	mq = messagequeue.GetNewRabbitMqObject()
	err = mq.Initialise(Config.RMqConfig)
	if err != nil {
		panic(err)
	}
}

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGKILL, syscall.SIGUSR1)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case syscall.SIGKILL:
				done <- true
			case syscall.SIGUSR1:
				readAndPublishHotelInfo()
			}
		}
	}()
	<-done
	log.Debugf("exiting...")
}

func readAndPublishHotelInfo() {
	byteValue, err := ioutil.ReadFile("sampledata.json")
	if err != nil {
		log.Errorf("error in parsing json file - [%s]", err.Error())
		return
	}
	var data interface{}
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Debugf("error in unmarshalling json - [%s]", err.Error())
		return
	}
	log.Debugf("publishing json[%v]", data)
	mq.Publish("hotel", byteValue)
}
