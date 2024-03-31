package main

import (
	"fmt"
	"sync"
)

type stringReciever struct {
	name                  string
	ch                    chan string
	recivedValues         []string
	wg                    *sync.WaitGroup
	messageCountToRecieve int
}

func NewReciever(recieverNumber int, messageCountToRecieve int, wg *sync.WaitGroup) *stringReciever {
	return &stringReciever{
		name:                  fmt.Sprintf("Reciever %d", recieverNumber),
		recivedValues:         make([]string, 0, messageCountToRecieve),
		wg:                    wg,
		messageCountToRecieve: messageCountToRecieve,
	}
}

func (r *stringReciever) Recive() {
	for {
		if r.messageCountToRecieve <= len(r.recivedValues) {
			r.wg.Done()
			return
		}
		r.recivedValues = append(r.recivedValues, <-r.ch)
	}
}

func (r stringReciever) PrintAllRecievedValues() {
	fmt.Printf("%s recived:\n", r.name)

	for _, val := range r.recivedValues {
		fmt.Println(val)
	}

	fmt.Print("\n\n")
}
