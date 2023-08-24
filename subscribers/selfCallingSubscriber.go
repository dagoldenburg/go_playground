package subscribers

import (
	"fmt"
	"net/http"
)

type SelfCallingSubscriber struct {
	Channel chan interface{}
}

func (scs *SelfCallingSubscriber) GetChannel() chan interface{} {
	return scs.Channel
}

func (scs *SelfCallingSubscriber) Start() {
	for range scs.Channel {
		go func() {
			req, err := http.NewRequest(http.MethodGet, "http://localhost:3333/selfCall", nil)
			if err != nil {
				fmt.Println("could not create request: ", err)
				return
			}
			_, err2 := http.DefaultClient.Do(req)
			if err2 != nil {
				fmt.Println("could not send request: ", err)
				return
			}
		}()
	}
}
