# クリーンアーキテクチャ ルール

## 依存方向

```
handler → usecase → domain ← repository
```

- `domain` は他のレイヤーに依存しない
- `usecase` は `domain` のインターフェースのみに依存する（`repository` の実装には依存しない）
- `handler` はHTTPの関心事のみを持つ（ビジネスロジックを書かない）
- `repository` はDBアクセスの実装を持つ（`domain` のインターフェースを実装する）

## 各レイヤーの責務

### domain/
- エンティティの構造体定義
- repository インターフェースの定義
- usecase インターフェースの定義

### usecase/
- ビジネスロジックの実装
- repository インターフェースを通じてデータにアクセス

### handler/
- リクエストのバリデーション
- usecase の呼び出し
- レスポンスの整形

### repository/
- SQL の実行（sqlc 生成コードを使用）
- domain インターフェースの実装

## 禁止事項

- handler 内でのSQL直接実行
- usecase 内での `http` パッケージのインポート
- domain 内での外部パッケージへの依存
