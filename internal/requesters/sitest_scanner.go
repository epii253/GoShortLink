package requesters

import (
	"context"
	models "project/internal/models"
	"sync"
	"time"
)

func scan(ctx context.Context, wg *sync.WaitGroup, jobs chan models.Task) {
	defer wg.Done()

	for {
		select {
		case _, ok := <-jobs:
			if !ok {
				return
			}

			// TODO do SCAN

		case <-ctx.Done():
			return
		}
	}
}

func SetupScan(workersCup int, tasks []models.Task) {

	paretnCtx := context.Background()

	ctx, cancel := context.WithTimeout(paretnCtx, time.Minute*3)

	defer cancel()

	jobs := make(chan models.Task)

	var wg sync.WaitGroup
	for i := 0; i < workersCup; i++ {
		wg.Add(1)
		go scan(ctx, &wg, jobs)
	}

	for _, task := range tasks {
		select {
		case jobs <- task:

		case <-ctx.Done():
			break
		}
	}

	close(jobs)
	wg.Wait()
}
