#!/bin/sh
set -e

# マイグレーション実行
 go run migrations/migration.go

# アプリ本体起動
exec "$@" 