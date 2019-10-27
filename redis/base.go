package redis

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

var Client *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func GenerateAuthCookie(name, passwd string, expireTime time.Duration) string {
	tuid := uuid.New().String()
	Client.HMSet(tuid, map[string]interface{}{
		"name":   name,
		"passwd": passwd,
	})
	Client.Expire(tuid, expireTime)
	return tuid
}

func GetAuth(tuid string) (string, string) {
	opt := Client.HMGet(tuid, "name", "passwd").Val()
	name, ok := opt[0].(string)
	passwd, ok := opt[1].(string)
	if !ok {
		return "", ""
	}
	return name, passwd
}
