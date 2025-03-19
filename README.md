# Private_memo
## 概要
https://github.com/ken7python/GinAuth
このユーザ認証システムを用いて作ったシンプルなメモアプリです

## 使用技術
 - Golang 1.24.1
 - Gin (Webフレームワーク)
 - GORM (ORM)
 - MySQL (データベース)
 - JWT (認証)

## 環境変数の設定
.envを以下のようにしてください
```
DB_USER=username
DB_PASSWORD=password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=yourdatabase
SECRET_KEY=your_secret_key
```

## 実行方法
```sh
go run main.go database.go user.go memo.go
```

## Memo API使い方
### メモ一覧取得
`GET api/memos` （JWTトークンが必要）

**ヘッダー**
```
Authorization: Bearer <your-jwt-token>
```
**リクエスト**
```
{
    "title": "title",
    "content": "message"
}
```
**レスポンス**
```
"memos": [
        {
            "ID": 1,
            "Title": "title",
            "Content": "message",
            "CreatedAt": "2025-03-19T17:33:05.773+09:00",
            "UpdatedAt": "2025-03-19T17:33:05.773+09:00",
            "UserUUID": "uuid"
        }
    ]
```

### メモ追加
`POST api/add_memo` （JWTトークンが必要）

**ヘッダー**
```
Authorization: Bearer <your-jwt-token>
```
**リクエスト**
```
{
    "title": "title",
    "content": "message"
}
```
**レスポンス**
```
{
    "message": "メモを作成しました"
}
```