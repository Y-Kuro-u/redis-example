# 1. 概要
redisのpub/subをテストするためのサンプルレポジトリ

## 2. 使い方
### 2.1 docker-composeの起動
以下コマンドでContainerを起動させる  
```
docker-compose up
```
その後、以下ログが出力されたことを確認する。
```
- go-publish_1    | Starting Publish Server ...
- go-subscribe_1  | Start subscribe ...
```

## 2.2 publish serverへのアクセス
publish serverはhttpサーバを起動しているためcurlコマンド等でアクセスを行う  
```
curl localhost:8080
```
以下ログが出力されていれば成功
```
- go-publish_1    | success Publish
- go-subscribe_1  | mychannel1 payload
```