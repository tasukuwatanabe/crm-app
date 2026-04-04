# crm-app

顧客管理システム。Go + PostgreSQL + React で構築するフルスタックアプリケーション。学習を目的とし、機能を段階的に拡張していく。

## 技術スタック

- **バックエンド**: Go, Echo, sqlc
- **データベース**: PostgreSQL
- **フロントエンド**: React 19, Vite, TypeScript, TanStack Query（予定）
- **CI/CD**: GitHub Actions

## ディレクトリ構成

モノレポ。バックエンドはクリーンアーキテクチャで実装。

```
crm-app/
├── backend/
│   ├── cmd/api/       # エントリーポイント
│   └── internal/
│       ├── domain/    # エンティティ・インターフェース定義
│       ├── usecase/   # ビジネスロジック
│       ├── handler/   # HTTP ハンドラー
│       └── repository/ # DB アクセス実装
└── frontend/          # React アプリ（今後追加）
```

## 開発コマンド

```bash
cd backend && go run ./cmd/api   # サーバー起動
cd backend && go test ./...      # テスト全実行
```

## 規約

- コミット: Conventional Commits 形式（詳細 → `.claude/rules/commit.md`）
- テスト: 単体テスト・統合テストを記述（詳細 → `.claude/rules/testing.md`）
- アーキテクチャ: クリーンアーキテクチャの依存方向を厳守（詳細 → `.claude/rules/architecture.md`）
- 学習スタイル: コマンドはユーザーが実行するので提示のみ行う（詳細 → `.claude/rules/teaching.md`）
