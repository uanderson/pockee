package exchange

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/robfig/cron/v3"
	"github.com/uanderson/pockee/autoid"
	"github.com/uanderson/pockee/database"
	"github.com/uanderson/pockee/exchange/dao"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
	"time"
)

// Values are temporary fixed here
var exchanges = []string{"USD_BRL", "EUR_BRL"}

type Scheduler struct {
	dao        *dao.Queries
	httpClient *httpclient.Client
}

func (s *Scheduler) fetchExchangeRates() {
	apiKey := os.Getenv("EXCHANGE_API_KEY")
	apiUri := os.Getenv("EXCHANGE_API_URI")

	params := url.Values{
		"q":       {strings.Join(exchanges[:], ",")},
		"compact": {"ultra"},
		"apiKey":  {apiKey},
	}

	apiUri = fmt.Sprintf("%s?%s", apiUri, params.Encode())

	response, err := s.httpClient.Get(apiUri, nil)
	if err != nil {
		log.Printf("could not fetch the exchange rates: %v\n", err)
		return
	}

	body, err := ioutil.ReadAll(response.Body)

	defer response.Body.Close()

	var rates map[string]interface{}
	err = json.Unmarshal(body, &rates)
	if err != nil {
		log.Println("could not parse exchange rates response")
		return
	}

	for _, exchange := range exchanges {
		if _, ok := rates[exchange]; ok {
			split := strings.Split(exchange, "_")
			rate := rates[exchange].(float64)

			s.updateExchangeRate(split[0], split[1], rate)
		}
	}
}

func (s *Scheduler) updateExchangeRate(source string, target string, rate float64) {
	err := s.dao.UpdateExchangeRate(context.TODO(), dao.UpdateExchangeRateParams{
		Date:   time.Now(),
		Id:     autoid.Id(),
		Rate:   rate,
		Source: source,
		Target: target,
	})

	if err != nil {
		log.Printf("exchange rate %s > %s update failed: %v\n", source, target, err)
		return
	}
}

func Schedule() Scheduler {
	scheduler := Scheduler{
		dao: dao.New(database.Pool),
		httpClient: httpclient.NewClient(
			httpclient.WithHTTPTimeout(10*time.Second),
			httpclient.WithRetryCount(3),
		),
	}

	cron := cron.New()
	cron.AddFunc("0 20-23 * * 1-5", scheduler.fetchExchangeRates)
	cron.Start()

	return scheduler
}
