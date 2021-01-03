package messagequeue

// Handler func handles the data received in the queue subscribed
// and returns a boolean indicating whether to ack or nack
type Handler func(data []byte) bool

// MessageQueue defines message queueing interface
type MessageQueue interface {

	// Initialise implements any message queue specific initialization process
	// where config will contain all the configurations to initialize the
	// message queue
	Initialise(config interface{}) error

	// Publish will publish the data onto the queue name provided
	Publish(queueName string, data []byte) error

	// Subscribe subscribes the queue with the name provided as arg
	// and calls the handler func if the message is received onto that queue
	Subscribe(queueName string, handler Handler) error
}
