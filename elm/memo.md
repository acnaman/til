## 部分適用
Elmは関数の引数を部分適用できる。
例えば3つの引数の合計値を返すような関数を考えたとする。
これを`elm repl`上で入力すると、関数の型は以下のように表示される

``` elm
> sum3num num1 num2 num3 = num1 + num2 + num3
<function> : number -> number -> number -> number
```

複数の引数をとる関数は、内部的には1つずつ処理される。
↑の関数は例えば `100`, `20` ,`3` の3引数を渡したとき、以下の順で処理されるイメージになる

``` elm
> f1 = sum3num 100
<function> : number -> number -> number
> f2 = f1 20
<function> : number -> number
> f2 3
123 : number
```

(f1とf2は順を追うために便宜上使用している中間の匿名関数)

逆に言うと、3つの引数が必要な関数に対して、1つの引数のみを渡すこともできる。このように複数の引数のうちの一部の引数のみを適用することを**部分適用**という。

部分適用を行うことによって得られる中間の関数を別の関数の引数に渡す、みたいなことも出てくる。混乱してきたら関数の型定義がどうなっているのかをよく見て、関数の構造を理解したほうが良い。

## コマンド
Elm RuntimeとApplicationの橋渡しを行う
Cmd Msgは `Msg` 型のメッセージでを送るコマンド的な感じ
`Browser.application` の場合、コマンドを受け取るとそのコマンドを引数にしたupdate関数を呼んで実行する

```
> Browser.application
<function>
    : { init : flags -> Url.Url -> Browser.Navigation.Key -> ( model, Cmd msg )
      , onUrlChange : Url.Url -> msg
      , onUrlRequest : Browser.UrlRequest -> msg
      , subscriptions : model -> Sub msg
      , update : msg -> model -> ( model, Cmd msg )
      , view : model -> Browser.Document msg
      }
      -> Program flags model msg
```

「更新必要なし」の状態を `Cmd.none` で表現できる

```
> Cmd.none
<internals> : Cmd msg
```

## Result
以下のような形で定義する
```
type Result error value
  = Ok value
  | Err error
```
