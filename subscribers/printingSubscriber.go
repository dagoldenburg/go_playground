package subscribers

import "fmt"

type PrintingSubscriber struct {
	Channel chan interface{}
}

func (ps *PrintingSubscriber) GetChannel() chan interface{} {
	return ps.Channel
}

func (ps *PrintingSubscriber) Start() {
	for eventChannel := range ps.Channel {
		go fmt.Println(eventChannel)
	}
}
