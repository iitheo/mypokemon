package dbconfigs

import (
	"github.com/gomodule/redigo/redis"
)

const REDISCONNSTRING = "redis://:p557474730a5035e5e552638b9c81aaecb57869b3bc71b0bacd1fc14f95fb551b@ec2-52-16-22-73.eu-west-1.compute.amazonaws.com:19519"

func RedisConn() (redis.Conn, error) {
	c, err := redis.DialURL(REDISCONNSTRING, redis.DialTLSSkipVerify(true))
	return c, err
}
