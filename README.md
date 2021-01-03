# HotelPubSub

HotelPubSub is an example of rabbitMQ producer and consumer. The producer listens for a signal and on receiving that signal, it reads hotel offer json from the file and publishes it onto the message queue with the name hotel. The consumer is listening onto the same queue and on receiving the hotel json offer, it stores into the mysql db.

![](compilation_tutorial.gif)

## Setup

Prerequisites

1. RabbitMq server must be running. Update the rabbitmq server configuration address and port in producer config yaml (hotelpublisher/config.yaml) and consumer config yaml (hotelsubscriber/config.yaml) to appropriate. In case, rabbit mq server is not running .. it can be run locally using `docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management` 
2. MySql Server should be running for storing the hotel offers. Update the database configurations in hotelsubscriber/config.yaml to appropriate.

Producer Consumer
```
git clone https://github.com/u-prashant/HotelPubSub.git 
cd HotelPubSub 
make
cd bin
./subscriber #to run the consumer
./producer   #to run the producer
```
