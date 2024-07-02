package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"ocr-service-dev/internal/models"
	"os"

	"github.com/redis/rueidis"
)

func InitializeRedisClient() (rueidis.Client, error) {
	key := os.Getenv("REDIS_CLIENT_PASSWORD")
	if key == "" {
		return nil, fmt.Errorf("REDIS_CLIENT_PASSWORD environment variable not set")
	}
	rclient, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress:  []string{"redis-10547.c246.us-east-1-4.ec2.redns.redis-cloud.com:10547"},
		Username:     "default",
		Password:     key,
		DisableCache: true, // Disable client-side caching
	})
	if err != nil {
		return nil, err
	}
	return rclient, nil
}

func InitializeSearchIndex(ctx context.Context) {

	rclient, err := InitializeRedisClient()
	if err != nil {
		log.Println(err)
	}
	defer rclient.Close()
	cmd := rclient.B().FtList().Build()
	resp, err := rclient.Do(ctx, cmd).ToArray()

	if err != nil {
		return
	}

	var indexExists bool
	var index models.SearchIndex

	for _, idx := range resp {

		if err := json.Unmarshal([]byte(idx.String()), &index); err != nil {
			log.Println(err)
			continue
		}
		if index.Value == "FILE_NAME" {
			indexExists = true
			break
		}
	}

	if indexExists {
		log.Println("Index 'FILE_NAME' already exists. Skipping creation.")
		return
	}

	// Create the index
	cmdCreate := rclient.B().FtCreate().Index("FILE_NAME").OnJson().Schema().FieldName("$.FILE_NAME.text").As("FILE_NAME").Text().Build()
	_, err = rclient.Do(ctx, cmdCreate).ToString()
	if err != nil {
		log.Println("Error creating index:", err)
		return
	}

	fmt.Println("Index initialized.")
}
