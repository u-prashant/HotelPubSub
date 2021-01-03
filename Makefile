all: consumer producer

consumer:
	@echo "\n***Compiling consumer***"
	cd hotelsubscriber && go test -v && go build -o ../bin/subscriber

producer:
	@echo "\n***Compiling producer***"
	cd hotelpublisher && go test -v && go build -o ../bin/producer

.PHONY: all consumer producer