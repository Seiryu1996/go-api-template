# Gin REST API サンプル

このプロジェクトは [Gin](https://github.com/gin-gonic/gin) を使ったGo製のサンプルRESTful APIです。ユーザー認証（JWT）、アイテム管理（CRUD）、Docker対応、自動テスト機能を備えています。

## 主な機能
- ユーザー登録・ログイン（JWT認証）
- アイテムのCRUD（作成・取得・更新・削除）
- ユーザーごとのアイテム所有権管理
- RESTfulなAPI設計
- Docker・docker compose対応
- Goのテストフレームワークによる自動テスト

## はじめかた

### 1. リポジトリをクローン
```sh
git clone <リポジトリURL>
cd <クローンしたディレクトリ>/app
```

### 2. 環境変数ファイルの準備
- `.env.example` を `.env` にコピーし、必要に応じて編集してください（DB情報やJWTシークレットなど）
- **本物の `.env` ファイルは絶対にGitにコミットしないでください！**

### 3. Dockerで起動
```sh
docker compose up --build
```
- APIは `http://localhost:8080` で利用できます

### 4. ローカルで直接起動（Dockerを使わない場合）
```sh
go mod tidy
go run main.go
```

## APIエンドポイント一覧

### 認証系
- `POST /auth/signup` — ユーザー登録
- `POST /auth/login` — ログイン（JWTトークンを返却）

### アイテム系（GET /items 以外はJWT必須）
- `GET /items` — アイテム一覧取得
- `GET /items/:id` — アイテム詳細取得（認証必要）
- `POST /items` — アイテム作成（認証必要）
- `PUT /items/:id` — アイテム更新（認証必要）
- `DELETE /items/:id` — アイテム削除（認証必要）

## 環境変数（.env）
- `SECRET_KEY` — JWT署名キー
- `DB_USER`, `DB_PASS`, `DB_HOST`, `DB_PORT`, `DB_NAME` — データベース接続情報

## テストの実行方法

```sh
cd app/tests
go test -v
```
- テスト実行時はDBが自動的に初期化され、主要なAPI機能が検証されます。

## 注意事項
- `.env` は `.gitignore` に追加済みです。
- 本番運用時はDB情報やJWTキーの管理に十分ご注意ください。

## ライセンス
MIT 