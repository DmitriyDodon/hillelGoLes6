package main

type stringPublisher struct {
	strCh     chan string
	recievers []chan string
}

func NewPublisher(bufferSize int) *stringPublisher {
	return &stringPublisher{
		strCh:     make(chan string),
		recievers: make([]chan string, bufferSize),
	}
}

func (s *stringPublisher) AddReciever(r *stringReciever) {
	newRecieverCh := make(chan string)
	s.recievers = append(s.recievers, newRecieverCh)
	r.ch = newRecieverCh
}

func (s *stringPublisher) CloseStrChannel() {
	close(s.strCh)
}

func (s *stringPublisher) PublishString(str string) {
	s.strCh <- str
}

func (s stringPublisher) Start() {
	for {
		str, ok := <-s.strCh

		if !ok {
			return
		}

		for _, reciever := range s.recievers {
			if reciever != nil {
				reciever <- str
			}
		}
	}
}
