package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/TykTechnologies/storage/temporal/connector"
	tempkv "github.com/TykTechnologies/storage/temporal/keyvalue"
	"github.com/TykTechnologies/storage/temporal/model"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	// Parse the number of records from command line arguments
	recordCount := flag.Int("c", 10, "Number of records to generate")
	recordPrefix := flag.String("p", "record", "Prefix for the record key")
	flag.Parse()

	redisCfg, err := getConfig()
	if err != nil {
		panic(err)
	}

	redisConn, err := connector.NewConnector(model.RedisV9Type, model.WithRedisConfig(redisCfg))
	if err != nil {
		panic("Error connecting to redis:" + err.Error())
	}

	kv, err := tempkv.NewKeyValue(redisConn)
	if err != nil {
		panic(err)
	}

	// Generate and save random data in Redis
	for i := 0; i < *recordCount; i++ {
		key := fmt.Sprintf(*recordPrefix, ":%d", i)
		value := generateRandomData()
		err := kv.Set(context.Background(), key, value, 0)
		if err != nil {
			log.Println("Error saving record:", err)
			continue
		}
		log.Printf("Record %v saved..\n", i)
	}

	log.Printf("%d records generated and saved in Redis.\n", *recordCount)
}

// generateRandomData creates a random string to simulate data
func generateRandomData() string {
	rand.Seed(time.Now().UnixNano())
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := make([]byte, 10) // generate a string of length 10
	for i := range bytes {
		bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(bytes)
}

func getConfig() (*model.RedisOptions, error) {
	cfg := &model.RedisOptions{}
	if err := envconfig.Process("REDIS", cfg); err != nil {
		return nil, err
	}
	log.Println("Config parsed, using the following settings:")
	log.Printf("%+v \n", cfg)

	return cfg, nil
}
