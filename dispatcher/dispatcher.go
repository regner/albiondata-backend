package dispatcher

import (
	"crypto/md5"
	"fmt"
	"os"

	"github.com/go-redis/redis"
	"github.com/nats-io/go-nats"
)

type Work struct {
	Topic   string
	Message []byte
}

var WorkQueue = make(chan Work, 1000)

func StartDispatcher() {
	// Setup NATS
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = nats.DefaultURL
	}

	nc, _ := nats.Connect(natsURL)
	defer nc.Close()

	// Setup Redis
	rc := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	for {
		select {
		case work := <-WorkQueue:
			hash := md5.Sum(work.Message)
			key := fmt.Sprintf("%v-%v", work.Topic, hash)

			if !is_duped_message(key, rc) {
				nc.Publish(work.Topic, work.Message)
			}
		}
	}
}
