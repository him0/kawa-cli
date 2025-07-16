# kawa

目黒川ライブカメラの画像をターミナルに表示する CLI ツール

## インストール

### GitHub CLI (gh) を使ったインストール（推奨）

前提条件: [GitHub CLI](https://cli.github.com/) がインストールされている必要があります

```bash
# 1. 最新リリースから自分の環境に合ったバイナリをダウンロード
gh release download -R him0/kawa-cli --pattern "kawa-$(uname -s | tr '[:upper:]' '[:lower:]')-$(uname -m | sed 's/x86_64/amd64/')"

# 2. ダウンロードしたファイルに実行権限を付与
chmod +x kawa-$(uname -s | tr '[:upper:]' '[:lower:]')-$(uname -m | sed 's/x86_64/amd64/')

# 3. ~/.local/bin にインストール（PATHに追加されていることを確認してください）
mkdir -p ~/.local/bin
mv kawa-$(uname -s | tr '[:upper:]' '[:lower:]')-$(uname -m | sed 's/x86_64/amd64/') ~/.local/bin/kawa

# 4. インストール確認
kawa --version
```

#### 対応プラットフォーム

| OS | アーキテクチャ | バイナリ名 |
|---|---|---|
| macOS | Intel (x86_64) | kawa-darwin-amd64 |
| macOS | Apple Silicon (arm64) | kawa-darwin-arm64 |
| Linux | x86_64 | kawa-linux-amd64 |
| Linux | ARM64 | kawa-linux-arm64 |

### ソースからビルド

```bash
git clone https://github.com/him0/kawa-cli.git
cd kawa-cli
make install
```

※ imgcat コマンドが必要です

## 使い方

```bash
# 画像を一度表示
kawa

# ライブモード（60秒ごとに更新）
kawa --live

# 30秒ごとに更新
kawa -l -i 30
```

`Ctrl+C` で終了
