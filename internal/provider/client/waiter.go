package client

import "time"

type Waiter struct {
	start   time.Time
	client  *NotionApiClient
	minWait time.Duration
}

func (w *Waiter) WaitToReserveSpot() {
	w.client.queue <- true
	w.start = time.Now()
}

func (w *Waiter) ReleaseSpot() {
	elapsed := time.Since(w.start)
	if elapsed < w.minWait {
		time.Sleep(w.minWait - elapsed)
	}
	<-w.client.queue
}

func NewWaiter(client *NotionApiClient, minWait time.Duration) *Waiter {
	return &Waiter{
		client:  client,
		minWait: minWait,
	}
}
