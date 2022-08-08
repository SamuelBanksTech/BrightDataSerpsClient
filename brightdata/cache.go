package brightdata

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type BDRedisClient struct {
	redisOpts *BrightDataRedisOptions
	RdsClient *redis.Client
}

func NewBrightDataRedis(bdro *BrightDataRedisOptions) *BDRedisClient {

	rdb := redis.NewClient(&redis.Options{
		Addr:         bdro.Addr,
		Password:     bdro.Password,
		DB:           bdro.DB,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	})

	return &BDRedisClient{
		redisOpts: bdro,
		RdsClient: rdb,
	}
}

func (bdrc *BDRedisClient) RdsSet(key string, data interface{}) error {

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = bdrc.RdsClient.Set(context.Background(), key, value, bdrc.redisOpts.CacheExpiry).Err()
	if err != nil {
		return err
	}

	return nil
}

func (bdrc *BDRedisClient) RdsGet(key string, mstruct interface{}) error {

	val, err := bdrc.RdsClient.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return errors.New("key not found")
	} else if err != nil {
		return err
	} else {
		err = json.Unmarshal([]byte(val), &mstruct)
		if err != nil {
			return err
		}
	}

	return nil
}

func (bdrc *BDRedisClient) RdsDel(key string) error {

	_, err := bdrc.RdsClient.Del(context.Background(), key).Result()
	if err == redis.Nil {
		return errors.New("key not found")
	} else if err != nil {
		return err
	}

	return nil
}

func (bdrc *BDRedisClient) CloseConn() error {
	err := bdrc.RdsClient.Close()
	if err != nil {
		return err
	}

	return nil
}
