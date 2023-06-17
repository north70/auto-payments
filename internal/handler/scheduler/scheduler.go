package scheduler

import (
	"AutoPayment/internal/service"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog"
	"time"
)

type Scheduler struct {
	schedule *gocron.Scheduler
	service  *service.Service
	log      zerolog.Logger
}

func NewScheduler(service *service.Service, log zerolog.Logger, location *time.Location) *Scheduler {
	schedule := gocron.NewScheduler(location)

	return &Scheduler{service: service, log: log, schedule: schedule}
}

func (s *Scheduler) Start() {
	_, err := s.schedule.Every(1).Day().At("00:00").Do(func() { s.UpdateNextPayDay() })
	if err != nil {
		s.log.Err(err).Msg("error handle schedule command UpdateNextPayDay")
	}

	s.schedule.StartAsync()
}
