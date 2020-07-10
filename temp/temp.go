package temp

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type Temp map[string]interface{}

func (t *Temp) Merge(tinyExtras map[string]interface{}) {
	if tinyExtras == nil {
		return
	}
	for k, v := range tinyExtras {
		(*t)[k] = v
	}
}

type TempStorage struct {
	rc *redis.Client
}

func NewTempStorage(rc *redis.Client) *TempStorage {
	return &TempStorage{rc: rc}
}
func (ts *TempStorage) Set(ctx context.Context, key string, value Temp, expiration time.Duration) error {
	buff, _ := json.Marshal(value)
	return ts.rc.Set(ctx, key, string(buff), expiration).Err()
}
func (ts *TempStorage) Clear(ctx context.Context, key string) error {
	return ts.rc.Del(ctx, key).Err()
}
func (ts *TempStorage) Get(ctx context.Context, key string) (Temp, error) {
	result, err := ts.rc.Get(ctx, key).Result()
	if err == nil {
		var temp Temp
		err = json.Unmarshal([]byte(result), &temp)
		return temp, err
	}
	return nil, err
}
