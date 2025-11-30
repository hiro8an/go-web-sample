# go-web-sample

シンプルな認証付き Go Web アプリケーションのサンプルです。セッション管理、ログイン／ログアウト、簡易的なユーザ管理、ミドルウェア（認証・CSRF・ロギング）を実装しています。

**主な特徴**
- **セッション管理**: `gorilla/sessions` を利用
- **ユーザ管理**: ローカル SQLite データベースにユーザ情報を保存（bcrypt ハッシュ）
- **ミドルウェア**: ログイン保護、CSRF 相当処理、リクエストロギング

**動作ポート**: `:8080`

**プロジェクト構成（主要部分）**

```
go-web-sample/
├── main.go                 # アプリ起動（DB 初期化、シード、サーバ起動）
├── go.mod
├── README.md
├── web/
│   ├── routes.go           # ルート登録
│   ├── auth/
│   │   └── session.go      # セッション管理（SESSION_KEY 環境変数を使用可）
│   ├── handler/            # ハンドラとテンプレート読み込み
│   │   ├── home.go
│   │   ├── login.go
│   │   └── template.go
│   ├── database/           # SQLite3 を使った簡易ユーザ DB
│   │   ├── db.go
│   │   └── user.go
│   └── middleware/         # auth, csrf, logging
│       ├── auth.go
│       ├── csrf.go
│       └── logging.go
└── web/templates/
	├── home.html
	└── login.html
```

**主要依存**
- `github.com/gorilla/sessions` (セッション管理)
- `github.com/ncruces/go-sqlite3` (組み込み SQLite ドライバ)
- `golang.org/x/crypto/bcrypt` (パスワードハッシュ)

（依存は `go.mod` に記載されています）

**前提**
- Go 1.25 以上（`go.mod` を参照）

## セットアップと実行

1. リポジトリをクローンして依存を取得します。

```pwsh
git clone <repository-url>
cd go-web-sample
go mod download
```

2. セッション署名キー（任意）を設定します。設定しない場合はソース内の固定キーが使用されます（開発用のみ）。PowerShell の例:

```pwsh
$env:SESSION_KEY = 'your-32-or-more-byte-secret'
```

3. サーバを起動します。

```pwsh
go run main.go
```

起動後、`http://localhost:8080` にアクセスします。

## データベースと初期ユーザ

- プロジェクトルートに `app.db`（SQLite）が作成されます。
- `web/database.Seed()` は初期ユーザ `demo` を登録します。ソース上の登録時パスワードは `Password123` になっています（もしログに `demo/demo` と出力されている場合は出力とシード実装の不一致に注意してください）。

## ルート一覧（ソースに基づく）

- `GET /login` : ログインフォーム表示
- `POST /login`: ログイン送信（認証してセッション作成）
- `GET /logout`: ログアウト（セッション削除）
- `GET /`       : ログイン必須のホーム（ミドルウェアで保護）

## 環境変数

- `SESSION_KEY` : セッション Cookie の署名に使うキー（任意だが推奨）
- `TRUSTED_ORIGIN`: `web/middleware/csrf.go` で CSRF 相当の処理に信頼オリジンを追加するために使用

## 実装メモ / 注意点

- セッション Cookie 名はソース内で `session_id` に設定され、TTL は 24 時間です。
- パスワードは `bcrypt` でハッシュ化して保存します。
- 現在のルーティング実装は `web/routes.go` にまとめられています。ハンドラやミドルウェアの実装を変更するとルート動作が変わります。

## 追加作業提案

- 本番運用向けには `SESSION_KEY` を安全に管理し、HTTPs を必須にしてください。
- CSRF 保護の実装やエラーメッセージの国際化、入力検証の強化を検討してください。

ご希望であれば、この README をさらに詳しく（Docker サポート、DB マイグレーション手順、テストスクリプト追加など）拡張します。
