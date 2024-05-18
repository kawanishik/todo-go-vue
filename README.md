# GoとVue.jsを使用したTodoアプリ
## settings
### go
1. docker-compose up -d --build
   - dbの構築
2. go run main.go
   - ホットリロードを使用する場合は`air -c .air.toml`

### Vue
1. cd app
2. yarn install
3. yarn run dev

## 色々
- golangci-lint run main.go
  - goのlint
- mysql -h 127.0.0.1 -P 3306 -u root -p
  - dockerへのアクセス方法

## curlコマンド
- Get
  - curl http://localhost:3000/api/index
- POST
  - curl -X POST http://localhost:3000/api/create -H "Content-Type: application/json" -d '{"Title": "テスト", "Description": "curlコマンドからの登録"}'
- Delete
  - curl -X DELETE http://localhost:3000/api/delete -H "Content-Type: application/json" -d '{"Id":2}'
- Put
  - 

## 参考
- `zsh: command not found: air`が発生した場合
  - https://zenn.dev/awonosuke/articles/47336619a4f039
- データを返す部分（ginを使用しない場合）
  - https://qiita.com/hrk_ym/items/c73c5ad41c92688c3b94
- mysqlとの接続
  - https://zenn.dev/ajapa/articles/03dcf8fd12d086