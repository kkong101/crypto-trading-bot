package tools

import (
	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/service"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
	"time"
)

type OrderCondition struct {
	Condition func(df *ninjabot.Dataframe) bool
	Size      float64
	Side      ninjabot.SideType
}

type Scheduler struct {
	pair            string
	intervalTime    time.Duration
	orderConditions []OrderCondition
}

func NewScheduler(pair string) *Scheduler {
	return &Scheduler{pair: pair}
}

func (s *Scheduler) AddSellCondition(size float64, condition func(df *ninjabot.Dataframe) bool) {
	s.orderConditions = append(
		s.orderConditions,
		OrderCondition{Condition: condition, Size: size, Side: ninjabot.SideTypeSell},
	)
}

func (s *Scheduler) AddBuyCondition(size float64, condition func(df *ninjabot.Dataframe) bool) {
	s.orderConditions = append(
		s.orderConditions,
		OrderCondition{Condition: condition, Size: size, Side: ninjabot.SideTypeBuy},
	)
}

func (s *Scheduler) CheckCondition(df *ninjabot.Dataframe, broker service.Broker) {
	s.orderConditions = lo.Filter[OrderCondition](s.orderConditions, func(oc OrderCondition, _ int) bool {
		if oc.Condition(df) {
			_, err := broker.CreateOrderMarket(oc.Side, s.pair, oc.Size)
			if err != nil {
				log.Error(err)
				return true
			}
			return false
		}
		return true
	})
}
