package subscribers

type Subscriber interface {
	GetChannel() chan interface{}
	Start()
}
