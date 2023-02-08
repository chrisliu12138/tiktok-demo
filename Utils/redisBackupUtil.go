package Utils

import (
	"context"
	"math/rand"
	"strconv"
	"time"
)

func TimeMission() {
	ticker := time.NewTicker(3 * time.Second)
	go func() {
		for {
			<-ticker.C
			te()
		}
	}()
}

func te() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		num := rand.Int63n(30)
		ctx := context.Background()
		RDB.SAdd(ctx, "User_like_2", strconv.Itoa(int(num)))
	}
}
