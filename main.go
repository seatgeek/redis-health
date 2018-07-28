package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	// create client
	client := redisClient()

	// check if its alive
	if err := client.Ping().Err(); err != nil {
		critical("Could not ping redis: %s", err)
	}

	// get the client info
	raw, err := client.Info().Result()
	if err != nil {
		critical("Could not get client info: %s", err)
	}

	// parse out
	info := parseKeyValue(raw)

	// check if redis is loading data from disk
	if loading, ok := info["loading"]; ok && loading == "1" {
		critical("Redis is currently loading data from disk")
	}

	// check if redis is syncing data from master
	if masterSyncInProgress, ok := info["master_sync_in_progress"]; ok && masterSyncInProgress == "1" {
		critical("Redis slave is currently syncing data from its master instance")
	}

	// check if the master link is up
	if masterLinkUpString, ok := info["master_link_status"]; ok && masterLinkUpString != "up" {
		critical("Redis slave do not have active connection to its master instance")
	}

	fmt.Printf("OK!")
}

func redisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         os.Getenv("REDIS_ADDR"),
		DialTimeout:  1 * time.Second,
		Password:     os.Getenv("REDIS_PASS"),
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	})
}

func critical(str string, args ...interface{}) {
	fmt.Printf(str, args...)
	os.Exit(2)
}

func parseKeyValue(str string) map[string]string {
	res := make(map[string]string)

	lines := strings.Split(str, "\r\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			continue
		}

		pair := strings.Split(line, ":")
		if len(pair) != 2 {
			continue
		}

		res[pair[0]] = pair[1]
	}

	return res
}
