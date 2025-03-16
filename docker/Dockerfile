# ビルド用のステージ
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 依存関係をコピーしてインストール
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/api ./cmd/api

# 実行用の軽量イメージ
FROM alpine:3.18

WORKDIR /app

# ビルドしたバイナリをコピー
COPY --from=builder /app/api /app/api

# 実行ユーザーを設定
RUN adduser -D -g '' appuser
USER appuser

# ポートを公開
EXPOSE 8080

# アプリケーションを実行
CMD ["/app/api"]