package pocketsmith

import (
	"context"
	"fmt"
	exchangedao "github.com/uanderson/pockee/exchange/dao"
	"gopkg.in/yaml.v3"
	"log"
	"net/url"
	"strings"
	"time"
)

// Event is what we call budget on PocketSmith.
type Event struct {
	Id           string
	Note         string
	CurrencyCode string `json:"currency_code"`
	Date         string
}

// EventMetaData is the note field in an event.
// The contents of the note is in yaml format.
type EventMetaData struct {
	Processed bool
	Rate      float64
	Currency  string
	Tax       float64
	Hours     int
}

func (s *Scheduler) fetchEvents() {
	user, err := fetchUser()

	if err != nil {
		log.Println(err)
		return
	}

	now := time.Now()

	firstOfCurrentMonth := now.AddDate(0, 0, -now.Day()-1)
	lastOfCurrentMonth := firstOfCurrentMonth.AddDate(0, 1, -1)
	firstOfNextMonth := firstOfCurrentMonth.AddDate(0, 1, 0)
	lastOfNextMonth := firstOfNextMonth.AddDate(0, 1, -1)

	s.updateEvents(user, firstOfCurrentMonth, lastOfCurrentMonth)
	s.updateEvents(user, firstOfNextMonth, lastOfNextMonth)
}

func (s *Scheduler) updateEvents(user *User, startDate time.Time, endDate time.Time) {
	params := url.Values{
		"start_date": {startDate.Format("2006-01-02")},
		"end_date":   {endDate.Format("2006-01-02")},
	}

	eventsUrl := fmt.Sprintf("/users/%v/events?%s", user.Id, params.Encode())

	var events []Event
	err := FetchJson(eventsUrl, &events)

	if err != nil {
		log.Println(err)
		return
	}

	for _, event := range events {
		if len(event.Note) > 0 {
			s.updateEvent(event)
		}
	}
}

func (s *Scheduler) updateEvent(event Event) {
	eventMetaData := EventMetaData{}

	err := yaml.Unmarshal([]byte(event.Note), &eventMetaData)
	if err != nil {
		return
	}

	if eventMetaData.Processed {
		return
	}

	exchangeRate, err := s.exchangeDao.GetExchangeRateForConversion(context.TODO(), exchangedao.GetExchangeRateForConversionParams{
		Date:   time.Now(),
		Source: eventMetaData.Currency,
		Target: strings.ToUpper(event.CurrencyCode),
	})

	if err != nil {
		return
	}

	amount := float64(eventMetaData.Hours) * eventMetaData.Rate * exchangeRate.Rate
	amount = amount - amount*eventMetaData.Tax

	eventsUrl := fmt.Sprintf("/events/%s", event.Id)
	eventBody := map[string]interface{}{
		"behaviour": "forward",
		"amount":    amount,
	}

	var rates map[string]interface{}
	err = PutJson(eventsUrl, eventBody, rates)
	if err != nil {
		log.Println(err)
	}
}
