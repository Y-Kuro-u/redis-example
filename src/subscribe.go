package main
import (
	"fmt"
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func subscribe(){
    rdb := redis.NewClient(&redis.Options{
        Addr:     "redis-container:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

	// There is no error because go-redis automatically reconnects on error.
	pubsub := rdb.Subscribe(ctx, "mychannel1")
	// Close the subscription when we are done.
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
	
		fmt.Println(msg.Channel, msg.Payload)
	}
}

func main() {
	subscribe()
}