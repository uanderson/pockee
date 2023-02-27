package pocketsmith

import (
	"github.com/robfig/cron/v3"
	"github.com/uanderson/pockee/database"
	exchangedao "github.com/uanderson/pockee/exchange/dao"
	"github.com/uanderson/pockee/setting"
	"log"
)

type Scheduler struct {
	exchangeDao *exchangedao.Queries
}

func Schedule() Scheduler {
	scheduler := Scheduler{
		exchangeDao: exchangedao.New(database.Pool),
	}

	settingService := setting.NewService()
	cronSetting, err := settingService.GetSettingByKey("pocketsmith.cron")
	if err != nil {
		log.Fatal("failed to query 'pocketsmith.cron' setting")
	}

	cron := cron.New()
	cron.AddFunc(cronSetting.Value, scheduler.fetchEvents)
	cron.Start()

	return scheduler
}
