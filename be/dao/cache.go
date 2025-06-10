package dao

import (
	"time"
)

func SetResumeCache(taskID string, data string, ttl time.Duration) error {
	return Rdb.Set(Ctx, "resume:"+taskID, data, ttl).Err()
}

func GetResumeCache(taskID string) (string, error) {
	return Rdb.Get(Ctx, "resume:"+taskID).Result()
}
