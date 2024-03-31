package main

import (
	"flag"
)

const goroutinesCountDefault = 100
const paralelGorutinesRunNumberDefault = 10

func main() {

	goroutinesCount := flag.Int("goroutinesCount", goroutinesCountDefault, "Total number of goroutines!")
	paralelGoroutinesRunNumber := flag.Int("paralelGoroutinesRunNumber", paralelGorutinesRunNumberDefault, "Paralel number of goroutines!")

	flag.Parse()

	dataCh := make(chan int)
	runningGorutinesLimitCh := make(chan struct{}, *paralelGoroutinesRunNumber)

	publisher := NewPublisher(
		runningGorutinesLimitCh,
		dataCh,
		*goroutinesCount,
	)

	reciever := NewReciever(dataCh)

	go publisher.Start()

	reciever.StartRecieving()
}
