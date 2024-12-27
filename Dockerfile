# Goの公式イメージを使用
FROM golang:1.22

# 作業ディレクトリを設定
WORKDIR /app

# 必要なファイルをコンテナにコピー
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# アプリケーションをビルド
RUN go build -o main ./cmd/main.go

# ポートを設定
ENV PORT=8080
EXPOSE 8080

# 実行コマンドを指定
CMD ["./main"]
