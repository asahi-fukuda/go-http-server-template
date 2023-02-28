# GO HTTP SERVER TEMPLATE

## 開発環境構築

### 1. DB 作成とマイグレーションの実行

以下のコマンドでマイグレーションを実行します。

```
./scripts/sql-migrate up
```

以下のコマンドで dry run すると実行予定の SQL が出力されます。

```
./scripts/sql-migrate up
```

### 2. アプリケーションのビルドとサーバー起動

以下のコマンドを実行します。

```
# 初回のみ
make build
```

```
# サーバー起動
make serve
```

### OpenAPI Stub, API Doc の生成

Makefile の以下のコマンドを実行すると生成できます。

```
make all
```

別々に生成したい場合は Makefile に記載されたコマンドを適宜実行してください。
