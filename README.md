# Gin + GORM + MySQL + Redis MVC プロジェクト

Gin、GORM、MySQL、Redisを使用したMVCアーキテクチャのGoプロジェクトです。
API、フロントエンド、管理画面の3つのモジュールに分かれています。

## プロジェクト構造

```
gin-gorm-mvc/
├── cmd/                      # エントリーポイント
│   ├── api/                  # APIサーバー
│   ├── front/                # フロントエンドサーバー
│   ├── admin/                # 管理画面サーバー
│   └── migrate/              # マイグレーションツール
├── internal/                 # 内部パッケージ
│   ├── config/               # 設定
│   ├── database/             # データベース接続
│   ├── redis/                # Redis接続
│   ├── models/               # データモデル
│   ├── repositories/         # リポジトリ層
│   ├── services/             # サービス層
│   ├── controllers/          # コントローラ層
│   ├── middleware/           # ミドルウェア
│   ├── routes/               # ルーティング
│   └── utils/                # ユーティリティ
├── pkg/                      # 公開パッケージ
│   └── response/             # レスポンスヘルパー
├── migrations/               # マイグレーションファイル
├── .env.example              # 環境変数の例
├── go.mod                    # Go モジュール
└── README.md                 # このファイル
```

## 必要な環境

### ローカル開発（Dockerを使用）
- Docker
- Docker Compose

### ローカル開発（Dockerなし）
- Go 1.24+
- MySQL 5.7+
- Redis 6.0+

## セットアップ

### 方法1: Docker を使用（推奨）

最も簡単な方法です。DockerとDocker Composeさえあれば、すぐに開発環境を構築できます。

#### 1. リポジトリのクローン

```bash
git clone <repository-url>
cd gin-gorm-mvc
```

#### 2. 環境変数の設定

```bash
cp .env.example .env
```

`.env`ファイルを編集（Docker環境用の設定例）：

```env
# データベース設定（Dockerコンテナ名を使用）
DB_HOST=mysql
DB_PORT=3306
DB_USER=gin_user
DB_PASSWORD=gin_password
DB_NAME=gin_gorm_mvc

# Redis設定（Dockerコンテナ名を使用）
REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=

# サーバーポート設定
API_PORT=8080
FRONT_PORT=8081
ADMIN_PORT=8082
```

#### 3. Dockerコンテナの起動

```bash
# Makeを使用する場合
make docker-up

# または docker-compose を直接使用
docker-compose up -d
```

これで以下のサービスが起動します：
- MySQL: `localhost:3306`
- Redis: `localhost:6379`
- phpMyAdmin: `http://localhost:8888` (データベース管理ツール)
- Redis Commander: `http://localhost:8889` (Redis管理ツール)

#### 4. マイグレーションの実行

```bash
# Makeを使用する場合
make migrate-up

# または直接実行
go run cmd/migrate/main.go up
```

#### 5. アプリケーションの起動

```bash
# APIサーバー
make run-api

# フロントサーバー
make run-front

# 管理画面サーバー
make run-admin
```

#### Docker環境の停止

```bash
# コンテナを停止
make docker-down

# コンテナとボリュームを削除（データも削除）
make docker-clean
```

### 方法2: ローカル環境（Dockerなし）

#### 1. リポジトリのクローン

```bash
git clone <repository-url>
cd gin-gorm-mvc
```

#### 2. 依存関係のインストール

```bash
go mod download
```

#### 3. 環境変数の設定

```bash
cp .env.example .env
```

`.env`ファイルを編集：

```env
# データベース設定
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=gin_gorm_mvc

# Redis設定
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# サーバーポート設定
API_PORT=8080
FRONT_PORT=8081
ADMIN_PORT=8082
```

#### 4. データベースの作成

MySQLにログインしてデータベースを作成します。

```sql
CREATE DATABASE gin_gorm_mvc CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

#### 5. マイグレーションの実行

```bash
go run cmd/migrate/main.go up
```

## 使い方

### APIサーバーの起動

```bash
go run cmd/api/main.go
```

APIサーバーは `http://localhost:8080` で起動します。

### フロントエンドサーバーの起動

```bash
go run cmd/front/main.go
```

フロントエンドサーバーは `http://localhost:8081` で起動します。

### 管理画面サーバーの起動

```bash
go run cmd/admin/main.go
```

管理画面サーバーは `http://localhost:8082` で起動します。

## マイグレーションコマンド

### マイグレーション実行

```bash
# Makeを使用
make migrate-up

# または直接実行
go run cmd/migrate/main.go up
```

### マイグレーションロールバック

```bash
# Makeを使用
make migrate-down

# または直接実行
go run cmd/migrate/main.go down
```

### すべてのテーブルを削除して再作成

```bash
# Makeを使用
make migrate-fresh

# または直接実行
go run cmd/migrate/main.go fresh
```

## Makeコマンド一覧

プロジェクトには便利なMakeコマンドが用意されています：

```bash
make help           # ヘルプを表示
make docker-up      # Dockerコンテナを起動
make docker-down    # Dockerコンテナを停止
make docker-logs    # Dockerログを表示
make docker-clean   # Dockerコンテナとボリュームを削除
make migrate-up     # マイグレーション実行
make migrate-down   # マイグレーションロールバック
make migrate-fresh  # マイグレーションを再実行
make run-api        # APIサーバーを起動
make run-front      # Frontサーバーを起動
make run-admin      # Adminサーバーを起動
make build          # アプリケーションをビルド
make test           # テストを実行
make deps           # 依存関係の更新
make clean          # ビルドファイルを削除
```

## API エンドポイント

### APIモジュール (`/api/v1`)

#### ユーザー

- `POST /api/v1/users` - ユーザー作成
- `GET /api/v1/users` - ユーザー一覧
- `GET /api/v1/users/:id` - ユーザー詳細
- `PUT /api/v1/users/:id` - ユーザー更新
- `DELETE /api/v1/users/:id` - ユーザー削除

#### 記事

- `POST /api/v1/articles` - 記事作成
- `GET /api/v1/articles` - 記事一覧
- `GET /api/v1/articles/:id` - 記事詳細
- `GET /api/v1/articles/author/:author_id` - 著者別記事一覧
- `PUT /api/v1/articles/:id` - 記事更新
- `DELETE /api/v1/articles/:id` - 記事削除

### フロントエンドモジュール (`/front`)

- `GET /front/articles` - 記事一覧
- `GET /front/articles/:id` - 記事詳細
- `GET /front/articles/author/:author_id` - 著者別記事一覧
- `GET /front/users/:id` - ユーザー詳細

### 管理画面モジュール (`/admin`)

- APIモジュールと同じエンドポイント（`/admin`プレフィックス付き）

## アーキテクチャ

このプロジェクトはMVCパターンに基づいています：

- **Models** (`internal/models`): データベーステーブルの構造を定義
- **Repositories** (`internal/repositories`): データベース操作を抽象化
- **Services** (`internal/services`): ビジネスロジックを実装
- **Controllers** (`internal/controllers`): HTTPリクエストを処理
- **Routes** (`internal/routes`): URLルーティングを定義

## 機能

- ✅ RESTful API
- ✅ MVCアーキテクチャ
- ✅ GORM（ORM）
- ✅ MySQL接続
- ✅ Redis接続
- ✅ マイグレーション
- ✅ ミドルウェア（CORS、ロガー、リカバリー）
- ✅ 環境変数管理
- ✅ 依存性注入
- ✅ エラーハンドリング
- ✅ ページネーション
- ✅ Docker対応
- ✅ Docker Compose設定
- ✅ phpMyAdmin / Redis Commander付属

## Docker環境の詳細

### 起動するコンテナ

| コンテナ | 説明 | ポート |
|---------|------|--------|
| mysql | MySQL 8.0 データベース | 3306 |
| redis | Redis 7 キャッシュ | 6379 |
| phpmyadmin | MySQL管理ツール | 8888 |
| redis-commander | Redis管理ツール | 8889 |

### 管理ツールへのアクセス

- **phpMyAdmin**: http://localhost:8888
  - サーバー: `mysql`
  - ユーザー名: `root`
  - パスワード: `root`

- **Redis Commander**: http://localhost:8889
  - 自動的にRedisに接続されます

## トラブルシューティング

### Dockerコンテナが起動しない

```bash
# コンテナのログを確認
make docker-logs

# すべてのコンテナを削除して再起動
make docker-clean
make docker-up
```

### マイグレーションが失敗する

データベース接続情報が正しいか `.env` ファイルを確認してください。

### ポートが既に使用されている

`docker-compose.yml` のポート番号を変更してください。

## ライセンス

MIT License
