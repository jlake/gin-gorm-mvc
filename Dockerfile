# ビルドステージ
FROM golang:1.24-alpine AS builder

# 作業ディレクトリを設定
WORKDIR /app

# 依存関係をコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# バイナリをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/api ./cmd/api/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/front ./cmd/front/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/admin ./cmd/admin/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/migrate ./cmd/migrate/main.go

# 実行ステージ
FROM alpine:latest

# 必要なパッケージをインストール
RUN apk --no-cache add ca-certificates tzdata

# タイムゾーンを設定
ENV TZ=Asia/Tokyo

# 作業ディレクトリを設定
WORKDIR /app

# ビルドステージからバイナリをコピー
COPY --from=builder /app/bin/api /app/api
COPY --from=builder /app/bin/front /app/front
COPY --from=builder /app/bin/admin /app/admin
COPY --from=builder /app/bin/migrate /app/migrate

# .envファイルをコピー（オプション）
COPY .env.example /app/.env.example

# 実行権限を付与
RUN chmod +x /app/api /app/front /app/admin /app/migrate

# デフォルトでAPIサーバーを起動
CMD ["/app/api"]
