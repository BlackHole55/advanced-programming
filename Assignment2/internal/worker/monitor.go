package worker

import (
	"log"
	"time"

	"github.com/MeirhanSyzdykov/Assignment2/internal/model"
	"github.com/MeirhanSyzdykov/Assignment2/internal/store"
)

func StartMonitoring(store *store.MemoryStore, stop chan struct{}) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			tasks := store.All()
			counts := map[model.TaskStatus]int{
				model.Pending: 0,
				model.Running: 0,
				model.Done:    0,
			}

			for _, t := range tasks {
				counts[t.Status]++
			}

			log.Printf("[Monitor] PENDING=%d RUNNING=%d DONE=%d\n", counts[model.Pending], counts[model.Running], counts[model.Done])

		case <-stop:
			log.Println("[Monitor] Stopping monitoring goroutine")
			return
		}
	}
}
