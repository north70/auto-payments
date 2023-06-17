package scheduler

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"time"
)

func (s *Scheduler) UpdateNextPayDay() {
	limit := 100
	offset := 0
	countUpdated := 0

	s.log.Info().Msg("start updating next pay day for payments")
	now := time.Now()
	for {
		payments, err := s.service.IndexByTime(limit, offset, now)
		if err != nil {
			log.Err(err).Msg(fmt.Sprintf("error get payments with limit = %d and offset = %d", limit, offset))
			break
		}

		if len(payments) == 0 {
			break
		}

		for _, payment := range payments {
			err = s.service.UpdateNextPayDay(payment.Id)
			if err != nil {
				log.Err(err).Msg(fmt.Sprintf("not update next pay day for payment with id = %d", payment.Id))
				continue
			}
			countUpdated++
		}

		offset += limit
	}
	s.log.Info().Msg(fmt.Sprintf("updating next pay day for payments ended. Updated %d records", countUpdated))
}
