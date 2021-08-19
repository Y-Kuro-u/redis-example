package main
import (
    "context"
	"fmt"
	"net/http"

    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func Publish(w http.ResponseWriter, r *http.Request ) {
	// redisとの接続
    rdb := redis.NewClient(&redis.Options{
        Addr:     "redis-container:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

	// mychannellの名前を付けたgroupにPublishという文字列を送信する
    err := rdb.Publish(ctx, "mychannel1", "Publish").Err()
    if err != nil {
        panic(err)
    } else {
        fmt.Println("success Publish")
    }
}

func main() {
	// http サーバにハンドラの追加
	http.HandleFunc("/", Publish)

    // http サーバのListen開始
    fmt.Println("Starting Publish Server ...")
	http.ListenAndServe(":8080", nil)
}