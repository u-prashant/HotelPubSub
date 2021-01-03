package main

import (
	"common/database"
	"common/messagequeue"
	"encoding/json"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

var (
	mq     *messagequeue.RabbitMq
	dbCtxt *database.DbCtxt
)

func init() {
	log.SetLevel(log.DebugLevel)
	err := loadConfig()
	if err != nil {
		panic(err)
	}
	mq = messagequeue.GetNewRabbitMqObject()
	err = mq.Initialise(Config.RMqConfig)
	if err != nil {
		panic(err)
	}
	dbCtxt = database.GetNewDbCtxt()
	err = dbCtxt.ConnectToDb(Config.DbConfig)
	if err != nil {
		panic(err)
	}
	err = dbCtxt.InitDatabase()
	if err != nil {
		panic(err)
	}
}

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGKILL)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case syscall.SIGKILL:
				done <- true
			}
		}
	}()
	go mq.Subscribe("hotel", offerHandler)
	<-done
	log.Debugf("exiting...")
}

func offerHandler(data []byte) (ack bool) {
	offers := &database.HotelOffers{}
	err := json.Unmarshal(data, &offers)
	if err != nil {
		log.Errorf("error in unmarshalling the offers [%s]", err.Error())
		return
	}
	log.Debugf("received message [%+v]", offers)
	dbCtxt.StoreOfferInDB(offers)
	ack = true
	return
}
