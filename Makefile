.PHONY: help docker-up docker-down docker-logs docker-clean migrate-up migrate-down migrate-fresh run-api run-front run-admin build test

# デフォルトターゲット
help:
	@echo "利用可能なコマンド:"
	@echo "  make docker-up        - Dockerコンテナを起動"
	@echo "  make docker-down      - Dockerコンテナを停止"
	@echo "  make docker-logs      - Dockerログを表示"
	@echo "  make docker-clean     - Dockerコンテナとボリュームを削除"
	@echo "  make migrate-up       - マイグレーション実行"
	@echo "  make migrate-down     - マイグレーションロールバック"
	@echo "  make migrate-fresh    - マイグレーションを再実行"
	@echo "  make run-api          - APIサーバーを起動"
	@echo "  make run-front        - Frontサーバーを起動"
	@echo "  make run-admin        - Adminサーバーを起動"
	@echo "  make build            - アプリケーションをビルド"
	@echo "  make test             - テストを実行"

# Docker コマンド
docker-up:
	docker-compose up -d
	@echo "Dockerコンテナが起動しました"
	@echo "MySQL: localhost:3306"
	@echo "Redis: localhost:6379"
	@echo "phpMyAdmin: http://localhost:8888"
	@echo "Redis Commander: http://localhost:8889"

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

docker-clean:
	docker-compose down -v
	@echo "すべてのコンテナとボリュームが削除されました"

# マイグレーション
migrate-up:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down

migrate-fresh:
	go run cmd/migrate/main.go fresh

# サーバー起動
run-api:
	go run cmd/api/main.go

run-front:
	go run cmd/front/main.go

run-admin:
	go run cmd/admin/main.go

# ビルド
build:
	go build -o bin/api cmd/api/main.go
	go build -o bin/front cmd/front/main.go
	go build -o bin/admin cmd/admin/main.go
	go build -o bin/migrate cmd/migrate/main.go
	@echo "ビルドが完了しました: bin/"

# テスト
test:
	go test -v ./...

# 依存関係の更新
deps:
	go mod download
	go mod tidy

# クリーンアップ
clean:
	rm -rf bin/
	@echo "ビルドファイルを削除しました"
