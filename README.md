# jp_prefecture

[![CircleCI](https://circleci.com/gh/diverse-inc/jp_prefecture.svg?style=svg)](https://circleci.com/gh/diverse-inc/jp_prefecture)
[![GoDoc](https://godoc.org/github.com/diverse-inc/jp_prefecture?status.svg)](https://godoc.org/github.com/diverse-inc/jp_prefecture)

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

	pref "github.com/diverse-inc/jp_prefecture"
)

var (
	prefInfo pref.Prefecture
	ok       bool
)

// 都道府県はPrefectureインターフェースとして返されます。
// 取得に失敗するとokにはfalseが設定されます
prefInfo, ok = pref.FindByCode(1)
if !ok {
	log.Print("prefecture not found.")
} else {
	log.Print(prefInfo.Code())       // 都道府県コード
	log.Print(prefInfo.Kanji())      // 都道府県名（漢字：末尾の「都」、「府」、「県」を除外）
	log.Print(prefInfo.KanjiShort()) // 都道府県名（漢字）
	log.Print(prefInfo.Kana())       // 都道府県名（かな）
	log.Print(prefInfo.KanaShort())  // 都道府県名（かな：末尾の「と」、「ふ」、「けん」を除外）
	log.Print(prefInfo.Roma())       // 都道府県名（ローマ字）
	log.Print(prefInfo.RomaShort())  // 都道府県名（ローマ字：末尾の「-to」、「-fu」、「-ken」を除外）
}

// 漢字名検索では末尾の「都」、「府」、「県」は省略して検索できます。
// ※省略しない場合でも問題ありません。
prefInfo, ok = pref.FindByKanji("東京")

// かな検索では末尾の「と」、「ふ」、「けん」は省略して検索できます。
// ※省略しない場合でも問題ありません。
prefInfo, ok = pref.FindByKana("とうきょう")

// ローマ字検索では末尾の「-to」、「-fu」、「-ken」は省略して検索できます。
// ※省略しない場合でも問題ありません。
prefInfo, ok = pref.FindByRoma("tokyo")

// List関数は全ての都道府県リストを返します。
prefs := pref.List()

// RegionList関数は地域のリストを返します。
// 各地域はRegionインターフェースとして返されます。
regions := pref.RegionList()

for _, region := range regions {
	// RegionインターフェースはPrefectureインターフェースと同様に、漢字名、かな名、ローマ字名を取得することが出来ます。
	region.Kanji()
	region.Kana()
	region.Roma()

	// RegionインターフェースはListメソッドを呼ぶことで、その地域に所属する都道府県リストを取得することが出来ます。
	prefs := region.List()
}
```
