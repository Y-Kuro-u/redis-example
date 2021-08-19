package main
import (
    "context"
    "fmt"

    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ExampleClient() {
	// redisとの接続
    rdb := redis.NewClient(&redis.Options{
        Addr:     "redis-container:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

	// mychannellの名前を付けたgroupにpayloadという文字列を送信する
    err := rdb.Publish(ctx, "mychannel1", "payload").Err()
    if err != nil {
        panic(err)
    } else {
        fmt.Println("success Publish")
    }
}

func main() {
    ExampleClient()
}