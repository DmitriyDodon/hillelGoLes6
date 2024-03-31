package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type publisher struct {
	limiterCh       chan struct{}
	publishingCh    chan int
	goroutinesCount int
	wg              *sync.WaitGroup
}

func NewPublisher(limiterCh chan struct{}, publishingCh chan int, goroutinesCount int) *publisher {
	return &publisher{
		limiterCh:       limiterCh,
		publishingCh:    publishingCh,
		goroutinesCount: goroutinesCount,
		wg:              &sync.WaitGroup{},
	}
}

func (p publisher) Start() {
	go func() {
		p.wg.Wait()
		close(p.publishingCh)
	}()
	for i := 0; i < p.goroutinesCount; i++ {
		go p.putRandIntIntoPublishingChannel()
	}
}

func (p publisher) putRandIntIntoPublishingChannel() {
	p.limiterCh <- struct{}{}
	fmt.Printf("New goroutin started. Current number of goroutines are %d.\n", len(p.limiterCh))
	p.wg.Add(1)
	for i := 0; i < 10; i++ {
		p.publishingCh <- rand.Int()
		time.Sleep(time.Millisecond * time.Duration(getRandomInt64InRange(10, 1000)))
	}
	<-p.limiterCh
	fmt.Printf("Goroutin Stoped. Current number of goroutines are %d.\n", len(p.limiterCh))
	p.wg.Done()
}

func getRandomInt64InRange(min int, max int) int64 {
	return int64(rand.Intn(max-min+1) + min)
}
