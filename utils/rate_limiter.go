package utils

import (
	"context"
	"strconv"
	"time"
	config "vid/config"
)

func LimitRate(user_id string, interval_in_seconds int64, maximumRequests int64) bool {
	now := time.Now().Unix()
	ctx := context.Background()
	currentWindow := strconv.FormatInt(now/interval_in_seconds, 10)
	key := user_id + ":" + currentWindow
	redisClient := config.ConnectRedis()
	value, _ := redisClient.Get(ctx, key).Result()
	requestCountCurrentWindow, _ := strconv.ParseInt(value, 10, 64)
	if requestCountCurrentWindow >= maximumRequests {
		// drop request
		return false
	}
	lastWindow := strconv.FormatInt(now/interval_in_seconds-1, 10)
	key = user_id + ":" + lastWindow // user userID + last time window
	// get last window count
	value, _ = redisClient.Get(ctx, key).Result()
	requestCountlastWindow, _ := strconv.ParseInt(value, 10, 64)

	elapsedTimePercentage := float64(now%interval_in_seconds) / float64(interval_in_seconds)
	if (float64(requestCountlastWindow)*(1-elapsedTimePercentage))+float64(requestCountCurrentWindow) >= float64(maximumRequests) {
		return false
	}
	redisClient.Incr(ctx, user_id+":"+currentWindow)
	return true
}
