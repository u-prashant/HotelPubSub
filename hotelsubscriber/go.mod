module hotelsubscriber

go 1.14

replace common => ../common

require (
	common v0.0.0-00010101000000-000000000000
	github.com/go-delve/delve v1.5.1
	github.com/sirupsen/logrus v1.7.0
	gopkg.in/yaml.v2 v2.4.0
)
