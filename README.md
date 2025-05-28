# 🛡️ Go Auth API

JWT認証つきのシンプルなGo製APIサーバー。  
個人学習・検証用に最小構成で構築。Echoフレームワーク + Docker環境。

---

## 🚀 使用技術

- **Go**: 1.23
- **Echo**: v4.13.4
- **MySQL**: 8.0.36
- **Docker / Docker Compose**

---

## 📦 セットアップ手順

### 1. リポジトリをクローン

```bash
git clone git@github.com:daihase/go-auth-api.git
cd go-auth-api
```

### 2. .env ファイルを作成

```env
PORT=8080
JWT_SECRET=your_jwt_secret
```

※ `your_jwt_secret` は任意の安全な文字列に変更してください。

### 3. Docker コンテナ起動

```bash
docker-compose up --build
```

起動後、ブラウザで http://localhost:8080 にアクセス可能。

---

## 🔌 API エンドポイント一覧

### ✅ POST /register

ユーザー登録。

**リクエストボディ:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

### 🔐 POST /login

ログインし、JWTトークンを取得。

**リクエストボディ:**
```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

**レスポンス例:**
```json
{
  "token": "xxxxx.yyyyy.zzzzz"
}
```

### 🙋‍♂️ GET /me

認証されたユーザー情報を取得。

**ヘッダー:**
```
Authorization: Bearer {JWTトークン}
```

---

## 📁 ディレクトリ構成

```
go-auth-api/
├── db/                 # 初期化用SQL (init.sql)
├── handlers/           # 各APIのビジネスロジック
├── middleware/         # JWTミドルウェア
├── models/             # DB接続・ユーザーモデル
├── main.go             # エントリーポイント
├── Dockerfile          # アプリ用Docker設定
├── docker-compose.yml  # 全体構成
└── .env                # 環境変数ファイル
```

---

## 🧪 動作確認コマンド例

```bash
# 登録
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Taro","email":"taro@example.com","password":"pass123"}'

# ログイン
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"taro@example.com","password":"pass123"}'

# 認証付きユーザー情報取得
curl http://localhost:8080/me \
  -H "Authorization: Bearer xxxxx.yyyyy.zzzzz"
```

---

## 📓 備考

- `.env` は `.gitignore` 推奨（秘密情報を含むため）
- JWT付きエンドポイントは `/me` のみ
- MySQLは `3307:3306` でポート分離
- コンテナ起動時に `http server started` が出たら成功

---
