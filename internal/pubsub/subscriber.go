package pubsub

import (
	"context"
	"encoding/json"
	"log"
	"moneytrx/internal/model"
	"moneytrx/internal/repository"

	"github.com/redis/go-redis/v9"
)

type Subscriber struct {
	Db    repository.PgRepo
	Redis *redis.Client
}

func (s *Subscriber) Subscribe() {
	pubsub := s.Redis.Subscribe(context.TODO(), "trx")
	defer pubsub.Close()
	_, err := pubsub.Receive(context.TODO())
	if err != nil {
		panic(err)
	}

	ch := pubsub.Channel()

	for msg := range ch {
		var req model.TrxRecord
		json.Unmarshal([]byte(msg.Payload), &req)

		if req.Success {
			err := s.Db.SuccessTrx(context.TODO(), req.Id)
			if err != nil {
				log.Println(err.Error())
			}
			return
		}

		err := s.Db.FailedTrx(context.TODO(), req.Id, 1, req.Amount)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
