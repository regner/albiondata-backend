package dispatcher

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

func is_duped_message(key string, rc *redis.Client) bool {
	_, err := rc.Get(key).Result()
	if err == redis.Nil {
		set(key, rc)

		// It didn't exist so not a duped message
		return false
	} else if err != nil {
		fmt.Printf("Error while getting from Redis: %v", err)

		// There was a problem with Redis and since we cannot verify
		// if the message is a dupe lets just say it isn't. Better
		// safe than sorry.
		return false
	} else {
		set(key, rc)

		// There was no problem with Redis and we got a value back.
		// So it was a dupe.
		return true
	}
}

func set(key string, rc *redis.Client) {
	cache_time := time.Duration(os.Getenv("CACHE_TIME")) * time.Second

	_, err := rc.Set(key, 1, cache_time).Result()
	if err != nil {
		fmt.Printf("Something wrong seting redis key: %v", err)
	}
}
