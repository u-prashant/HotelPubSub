all: consumer producer

consumer:
	@echo "\n***Compiling consumer***"
	cd hotelsubscriber && go build -o ../bin/subscriber

producer:
	@echo "\n***Compiling producer***"
	cd hotelpublisher && go build -o ../bin/producer

.PHONY: all consumer producer