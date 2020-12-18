package data

import "github.com/google/wire"

var Provider = wire.NewSet(New)

type DataMessage struct {
	message string
}

func New() *DataMessage {
	return &DataMessage{message: "I am an dataMessager!"}
}

func (m *DataMessage) Get() string {
	return m.message
}
