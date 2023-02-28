package exchange

import (
	"context"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/robfig/cron/v3"
	"github.com/uanderson/pockee/autoid"
	"github.com/uanderson/pockee/database"
	"github.com/uanderson/pockee/exchange/dao"
	"github.com/uanderson/pockee/setting"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Scheduler struct {
	dao        *dao.Queries
	httpClient *httpclient.Client
}

func Schedule() Scheduler {
	scheduler := Scheduler{
		dao: dao.New(database.Pool),
		httpClient: httpclient.NewClient(
			httpclient.WithHTTPTimeout(10*time.Second),
			httpclient.WithRetryCount(3),
		),
	}

	settingService := setting.NewService()
	cronSetting, err := settingService.GetSettingByKey("exchange.cron")
	if err != nil {
		log.Fatal("failed to query 'exchange.cron' setting")
	}

	cron := cron.New()
	cron.AddFunc(cronSetting.Value, scheduler.fetchExchangeRates)
	cron.Start()

	return scheduler
}

func (s *Scheduler) fetchExchangeRates() {
	ctx := context.Background()

	currencies, err := s.dao.GetExchangeCurrencies(ctx)
	if err != nil {
		return
	}

	for _, currency := range currencies {
		err, rate := fetchExchangeRate(currency)
		if err != nil {
			log.Println(err)
			return
		}

		s.updateExchangeRate(currency.Source, currency.Target, rate)
	}
}

func fetchExchangeRate(currency dao.ExchangeCurrency) (error, float64) {
	googleUrl := os.Getenv("GOOGLE_FINANCE_URL")
	url := fmt.Sprintf("%s/%s-%s", googleUrl, currency.Source, currency.Target)

	res, err := http.Get(url)
	if err != nil {
		return err, 0
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("failed to fetch from Google Finance: %d %s", res.StatusCode, res.Status), 0
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err, 0
	}

	lastPriceEl := doc.Find("[data-last-price]").First()
	lastPrice, exists := lastPriceEl.Attr("data-last-price")

	if !exists {
		return errors.New("could not find last price attribute"), 0
	}

	convertedRate, err := strconv.ParseFloat(lastPrice, 64)
	if err != nil {
		return err, 0
	}

	return nil, convertedRate
}

func (s *Scheduler) updateExchangeRate(source string, target string, rate float64) {
	err := s.dao.CreateExchangeRate(context.Background(), dao.CreateExchangeRateParams{
		Date:      time.Now(),
		Id:        autoid.Id(),
		Rate:      rate,
		Source:    source,
		Target:    target,
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Printf("exchange rate %s > %s update failed: %v\n", source, target, err)
		return
	}
}
