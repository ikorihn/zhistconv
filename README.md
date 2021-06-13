
```shell
# fish_hisoryをzsh_historyに変換して標準出力
$ zhistconv --fish fish_hist
# fish_hisoryをzsh_historyに変換して上書きする
$ zhistconv --fish fish_hist --write
# zsh_historyを読める形式にする
$ zhistconv --parse zsh_hist
# 読める形式にしたzsh_historyを戻す
$ zhistconv --reverse zsh_hist_text
```
