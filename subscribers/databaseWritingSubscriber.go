package subscribers

import "fmt"

type DataBaseWritingSubscriber struct {
	Channel chan interface{}
}

func (ps *DataBaseWritingSubscriber) GetChannel() chan interface{} {
	return ps.Channel
}

func (ps *DataBaseWritingSubscriber) Start() {
	for eventChannel := range ps.Channel {
		go fmt.Println(eventChannel)
	}
}
