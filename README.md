# jp_prefecture

`jp_prefecture`は _JIS X 0401_ の規格に基づいた都道府県情報の取得、及び検索機能を提供するライブラリです。

参考：[全国地方公共団体コード](https://ja.wikipedia.org/wiki/%E5%85%A8%E5%9B%BD%E5%9C%B0%E6%96%B9%E5%85%AC%E5%85%B1%E5%9B%A3%E4%BD%93%E3%82%B3%E3%83%BC%E3%83%89)

## インストール

```
go get github.com/diverse-inc/jp_prefecture
```

## 使い方

```go
package main

import (
	"log"

	"github.com/diverse-inc/jp_prefecture"
)

var (
	prefecture jp_prefecture.Prefecture
	ok         bool
)

// 取得に失敗するとokにはfalseが設定されます
prefecture, ok = jp_prefecture.FindByCode(1)
if !ok {
	log.Print("prefecture not found.")
} else {
	log.Print(prefecture.Code())  // 都道府県コード
	log.Print(prefecture.Kanji()) // 都道府県名（漢字）
	log.Print(prefecture.Kana())  // 都道府県名（かな）
	log.Print(prefecture.Roma())  // 都道府県名（ローマ字）
}

// 漢字名検索では末尾の「都」、「府」、「県」は省略して検索できます。
// ※省略しない場合でも問題ありません。
prefecture, ok = jp_prefecture.FindByKanji("東京")

// かな検索では末尾の「と」、「ふ」、「けん」は省略して検索できます。
// ※省略しない場合でも問題ありません。
prefecture, ok = jp_prefecture.FindByKana("とうきょう")

// ローマ字検索では末尾の「-to」、「-fu」、「-ken」は省略して検索できます。
// ※省略しない場合でも問題ありません。
prefecture, ok = jp_prefecture.FindByRoma("tokyo")

// List関数は全ての都道府県リストを返します。
// 各都道府県はPrefectureインターフェースとして返されます。
prefectures := jp_prefecture.List()
```
