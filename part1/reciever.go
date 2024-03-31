package main

import "fmt"

type reciever struct {
	recieveCh chan int
}

func NewReciever(recieveCh chan int) *reciever {
	return &reciever{
		recieveCh: recieveCh,
	}
}

func (r reciever) StartRecieving() {
	for {
		select {
		case number, ok := <-r.recieveCh:
			if !ok {
				return
			}
			fmt.Printf("Recieved data: %d.\n", number)
		}

	}
}
