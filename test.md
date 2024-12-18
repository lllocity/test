# プロジェクト概要

このプロジェクトは、JSONデータをパースして表示するGoプログラムを含んでいます。

## ファイル構成

- `test.go`: メインのGoプログラムファイル
- `test.md`: このプロジェクトの概要を説明するマークダウンファイル

## 使用方法

1. ターミナルを開き、プロジェクトのディレクトリに移動します。
2. `go`がインストールされていない場合は、[公式サイト](https://golang.org/dl/)からインストールします。
3. インストール後、`go`コマンドが見つからない場合は、以下のコマンドを実行してパスを設定します：
   ```sh
   export PATH=$PATH:/usr/local/go/bin
   ```
4. `go run test.go`コマンドを実行して、JSONデータをパースします。
5. パースされたデータがコンソールに表示されます。

## 例

以下は、プログラムの実行結果の例です：

```
Parsed Record: {RequestID:1 UserID:2 Priority:3 RequestTS:1709168560}
```

# テストの実行手順

1. 必要な依存関係をインストールします。
    ```sh
    go mod tidy
    ```

2. テストを実行します。
    ```sh
    go test -v
    ```

3. すべてのテストがパスすることを確認します。
