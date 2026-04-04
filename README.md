# crm-app

顧客管理システム。Go + PostgreSQL + React で構築するフルスタックアプリケーション。

## 技術スタック

- **バックエンド**: Go, Echo, sqlc
- **データベース**: PostgreSQL
- **フロントエンド**: React 19, Vite, TypeScript（予定）
- **CI/CD**: GitHub Actions
- **デプロイ**: Cloudflare Pages（フロント）, Fly.io（バックエンド）

## 開発環境のセットアップ

### 必要なツール

- Go 1.26+
- PostgreSQL 17+
- Node.js（フロントエンド追加時）

### バックエンド起動

```bash
cd backend
go run ./cmd/api
```

## ディレクトリ構成

```
crm-app/
├── backend/
│   ├── cmd/api/         # エントリーポイント
│   └── internal/
│       ├── domain/      # エンティティ・インターフェース定義
│       ├── usecase/     # ビジネスロジック
│       ├── handler/     # HTTP ハンドラー
│       └── repository/  # DB アクセス実装
└── frontend/            # React アプリ（今後追加）
```
