package dispatcher

import (
	"crypto/md5"
	"fmt"
	"os"

	"time"

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

			_, err := rc.Get(key).Result()
			if err == redis.Nil {
				nc.Publish(work.Topic, work.Message)

				_, err := rc.Set(key, 1, 600*time.Second).Result()
				if err != nil {
					fmt.Printf("Something wrong seting redis key: %v", err)
				}
			} else if err != nil {
				fmt.Printf("Error while getting from Redis: %v", err)
			} else {
				// If we saw the value again reset the expiry
				_, err := rc.Set(key, 1, 600*time.Second).Result()
				if err != nil {
					fmt.Printf("Something wrong seting redis key: %v", err)
				}
			}
		}
	}
}
