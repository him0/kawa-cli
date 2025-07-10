# kawa

目黒川ライブカメラの画像をターミナルに表示するCLIツール

## インストール

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