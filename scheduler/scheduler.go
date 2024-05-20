package scheduler

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type Scheduler struct {
	Controller gocron.Scheduler
}

func NewScheduler() Scheduler {
	s, err := gocron.NewScheduler()
	if err != nil {
		fmt.Println(err.Error())
	}

	return Scheduler{
		Controller: s,
	}
}

func (s *Scheduler) CreateJob(duration time.Duration, task any) error {
	_, err := s.Controller.NewJob(
		gocron.DurationJob(duration),
		gocron.NewTask(
			task,
		),
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Scheduler) Start() {
	s.Controller.Start()
}
