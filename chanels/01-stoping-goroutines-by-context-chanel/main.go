package main

import (
	"context"
	"log"
)

func main() {

	ctx, finish := context.WithCancel(context.Background())
	workersCount := 0
	for i := 0; i < 100; i++ {
		select {
		case <-ctx.Done():
			break
		default:
			workersCount = i
			go worker(ctx, finish, i)
		}
	}

Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		default:
		}
	}

	log.Printf("%v workers was started", workersCount)
	log.Println("Exit")
}

func worker(ctx context.Context, finish context.CancelFunc, workerNumber int) {

	for i := 0; i < 100; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			if i == 20 {
				log.Printf("Worker #%v acheived iteration #%v", workerNumber, i)
				finish()
				return
			}
		}
	}

}
