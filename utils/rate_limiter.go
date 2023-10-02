package utils

import (
	"context"
	"strconv"
	"time"
	config "vid/config"
)

func TokenBucket(user_id string, interval_in_seconds int64, maximumRequests int64) bool {
	ctx := context.Background()
	key := user_id
	redisClient := config.ConnectRedis()
	value, _ := redisClient.Get(ctx, key).Result()
	if value == "" {
		redisClient.Set(ctx, key, 1, 0)
	}
	valueInt, _ := strconv.ParseInt(value, 10, 64)
	if valueInt < maximumRequests {
		redisClient.Incr(ctx, key)
		redisClient.Expire(ctx, key, 2*time.Minute)
		return true
	}
	return false
}

func SlidingWindow(user_id string, interval_in_seconds int64, maximumRequests int64) bool {
	redisClient := config.ConnectRedis()
	context := context.Background()
	redisId := user_id + ":last-updated"
	token := redisClient.Get(context, redisId).Val()
	if token == "" {
		token = strconv.FormatInt(maximumRequests, 10)
	}
	currentTime := time.Now().Unix()
	lastUpdatedTime := redisClient.Get(context, redisId+":last-updated").Val()
	if lastUpdatedTime == "" {
		lastUpdatedTime = strconv.FormatInt(currentTime, 10)
	}
	tokenInt, _ := strconv.ParseInt(token, 10, 64)
	lastUpdatedTimeInt, _ := strconv.ParseInt(lastUpdatedTime, 10, 64)
	if tokenInt < maximumRequests {
		timeDifference := currentTime - lastUpdatedTimeInt
		newToken := tokenInt + timeDifference
		if newToken > maximumRequests {
			newToken = maximumRequests
		}
		redisClient.Set(context, redisId, newToken, 0)
		redisClient.Set(context, redisId+":last-updated", currentTime, 0)
		return true
	} else {
		return false
	}
}
