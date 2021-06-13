# zhistconv

マルチバイト文字を含むzsh_historyを変換するコマンド

zsh_historyは `0x80-0xA2` のバイトがあるとき、直前に `0x83` を挿入してから6bit目を反転するという仕様になっているため、
これを変換して読める形式にしたり、もとに戻したりするために使用する

## できること

- マルチバイト文字を含むzsh_historyファイルを読める形式に変換
- プレーンなテキストをzsh_historyファイル形式に変換
- fish_historyをzsh_historyファイル形式に変換

## 使い方

```shell
# zsh_historyを読める形式にする
$ zhistconv parse zsh_hist
# 読める形式にしたzsh_historyを戻す
$ zhistconv reverse zsh_hist_text
# fish_hisoryをzsh_historyに変換して標準出力
$ zhistconv fish fish_hist
```
