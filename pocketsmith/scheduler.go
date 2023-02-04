package pocketsmith

import (
	"github.com/robfig/cron/v3"
	"github.com/uanderson/pockee/database"
	exchangedao "github.com/uanderson/pockee/exchange/dao"
)

type Scheduler struct {
	exchangeDao *exchangedao.Queries
}

func Schedule() Scheduler {
	scheduler := Scheduler{
		exchangeDao: exchangedao.New(database.Pool),
	}

	cron := cron.New()
	cron.AddFunc("30 * * * *", scheduler.fetchEvents)
	cron.Start()

	return scheduler
}
