# テスト方針

## 単体テスト

- `usecase` レイヤーを中心に記述する
- `repository` はモックに差し替えてテストする
- ファイル名: `xxx_test.go`（テスト対象と同じパッケージ）

## 統合テスト

- `repository` レイヤーは実際の PostgreSQL に対してテストする
- テスト用 DB を使用する（本番 DB とは分離）
- `handler` の E2E テストは HTTP レベルで行う

## 命名規則

```go
func TestXxx_説明(t *testing.T) {}   // 単体テスト
func TestIntegration_説明(t *testing.T) {} // 統合テスト
```

## 統合テストの実行制御

統合テストは `go test -tags=integration ./...` で実行する。
ファイル先頭に以下を記述する:

```go
//go:build integration
```
