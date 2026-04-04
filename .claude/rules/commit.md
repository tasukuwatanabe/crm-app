# コミットメッセージ規約

Conventional Commits 形式で記述する。

## フォーマット

```
<type>(<scope>): <summary>
```

## type 一覧

| type | 用途 |
|---|---|
| feat | 新機能 |
| fix | バグ修正 |
| docs | ドキュメントのみの変更 |
| refactor | 機能変更を伴わないリファクタリング |
| test | テストの追加・修正 |
| chore | ビルド・設定ファイルの変更 |
| ci | CI/CD の変更 |

## scope 例

`handler`, `usecase`, `repository`, `domain`, `auth`, `customer`, `frontend`

## 例

```
feat(customer): add list endpoint
fix(handler): return 404 when customer not found
test(usecase): add unit tests for customer creation
chore: add CLAUDE.md and rules
```
