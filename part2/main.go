package main

import (
	"flag"
	"sync"

	faker "github.com/bxcodec/faker/v4"
)

const defaultRecieverQuantity = 10
const defaultMessageQuantity = 15

func main() {

	recieverQuantity := flag.Int("recieverQuantity", defaultRecieverQuantity, "setting reciever quantity")
	messageQuantity := flag.Int("messageQuantity", defaultMessageQuantity, "setting message quantity")

	flag.Parse()

	stringPublisher := NewPublisher(*recieverQuantity)

	var tempReciever *stringReciever

	strRecievers := make([]*stringReciever, 0, *recieverQuantity)

	wg := &sync.WaitGroup{}

	for i := 1; i <= *recieverQuantity; i++ {
		wg.Add(1)
		tempReciever = NewReciever(i, *messageQuantity, wg)
		stringPublisher.AddReciever(tempReciever)
		go tempReciever.Recive()
		strRecievers = append(strRecievers, tempReciever)
	}

	go stringPublisher.Start()

	for i := 0; i < *messageQuantity; i++ {
		go stringPublisher.PublishString(faker.Word())
	}

	wg.Wait()

	for _, strReciever := range strRecievers {
		strReciever.PrintAllRecievedValues()
	}
}
