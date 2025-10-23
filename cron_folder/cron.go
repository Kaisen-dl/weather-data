package cron_folder

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func DoCron(s gocron.Scheduler) ([]gocron.Job, error) {
	j, err := s.NewJob(
		gocron.DurationJob(
			10*time.Second,
		),
		gocron.NewTask(
			func() {
				fmt.Println("Hello!")
			},
		),
	)
	if err != nil {
		log.Println(err)
	}
	return []gocron.Job{j}, nil
}
