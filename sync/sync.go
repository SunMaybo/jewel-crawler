package sync

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Mutex struct {
	Name    string
	Timeout time.Duration
	redis   *redis.Client
}

func New(redis *redis.Client) *Mutex {
	return &Mutex{
		redis: redis,
	}
}

func (m *Mutex) Lock() {
	for {
		if m.NLock() {
			return
		}
		time.Sleep(1 * time.Millisecond)
	}
}

func (m *Mutex) NLock() bool {
	isOk, err := m.redis.SetNX(context.Background(), m.Name+":lock", 1, m.Timeout).Result()
	if err != nil {
		return false
	}
	if isOk {
		return true
	}
	return false
}
func (m *Mutex) UnLock() {
	m.redis.Del(context.Background(), m.Name+":lock")
}
func (m *Mutex) SetOffset(value string) {
	m.redis.Set(context.Background(), m.Name+":offset", value, 0)
}
func (m *Mutex) GetOffset() (string, error) {
	result, err := m.redis.Get(context.Background(), m.Name+":offset").Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
