# golang
## deferについて
```
func DeferFunc() {
  defer 関数
}
```
とすることで、DeferFuncの終了後にdeferの処理が実行される。
- 複数の処理を実行したい場合は無名関数を使用する
```
defer func() {
  ...
}
```
## panic & recover
- 例外処理として使用する。

## goroutin
- goの並行処理のことを言う。
