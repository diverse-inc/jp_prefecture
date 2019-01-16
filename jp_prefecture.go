package jp_prefecture

import (
	"strings"
)

// prefectureMap は都道府県の情報を”コード：名称のMap（漢字、かな、ローマ字）”形式で定義した値です。
var prefectureMap = map[int]map[string]string{
	1:  {"kanji": "北海道", "kana": "ほっかいどう", "roma": "hokkaido"},
	2:  {"kanji": "青森県", "kana": "あおもりけん", "roma": "aomori-ken"},
	3:  {"kanji": "岩手県", "kana": "いわてけん", "roma": "iwate-ken"},
	4:  {"kanji": "宮城県", "kana": "みやぎけん", "roma": "miyagi-ken"},
	5:  {"kanji": "秋田県", "kana": "あきたけん", "roma": "akita-ken"},
	6:  {"kanji": "山形県", "kana": "やまがたけん", "roma": "yamagata-ken"},
	7:  {"kanji": "福島県", "kana": "ふくしまけん", "roma": "fukushima-ken"},
	8:  {"kanji": "茨城県", "kana": "いばらきけん", "roma": "ibaraki-ken"},
	9:  {"kanji": "栃木県", "kana": "とちぎけん", "roma": "tochigi-ken"},
	10: {"kanji": "群馬県", "kana": "ぐんまけん", "roma": "gunma-ken"},
	11: {"kanji": "埼玉県", "kana": "さいたまけん", "roma": "saitama-ken"},
	12: {"kanji": "千葉県", "kana": "ちばけん", "roma": "chiba-ken"},
	13: {"kanji": "東京都", "kana": "とうきょうと", "roma": "tokyo-to"},
	14: {"kanji": "神奈川県", "kana": "かながわけん", "roma": "kanagawa-ken"},
	15: {"kanji": "新潟県", "kana": "にいがたけん", "roma": "niigata-ken"},
	16: {"kanji": "富山県", "kana": "とやまけん", "roma": "toyama-ken"},
	17: {"kanji": "石川県", "kana": "いしかわけん", "roma": "ishikawa-ken"},
	18: {"kanji": "福井県", "kana": "ふくいけん", "roma": "fukui-ken"},
	19: {"kanji": "山梨県", "kana": "やまなしけん", "roma": "yamanashi-ken"},
	20: {"kanji": "長野県", "kana": "ながのけん", "roma": "nagano-ken"},
	21: {"kanji": "岐阜県", "kana": "ぎふけん", "roma": "gifu-ken"},
	22: {"kanji": "静岡県", "kana": "しずおかけん", "roma": "shizuoka-ken"},
	23: {"kanji": "愛知県", "kana": "あいちけん", "roma": "aichi-ken"},
	24: {"kanji": "三重県", "kana": "みえけん", "roma": "mie-ken"},
	25: {"kanji": "滋賀県", "kana": "しがけん", "roma": "shiga-ken"},
	26: {"kanji": "京都府", "kana": "きょうとふ", "roma": "kyoto-fu"},
	27: {"kanji": "大阪府", "kana": "おおさかふ", "roma": "osaka-fu"},
	28: {"kanji": "兵庫県", "kana": "ひょうごけん", "roma": "hyogo-ken"},
	29: {"kanji": "奈良県", "kana": "ならけん", "roma": "nara-ken"},
	30: {"kanji": "和歌山県", "kana": "わかやまけん", "roma": "wakayama-ken"},
	31: {"kanji": "鳥取県", "kana": "とっとりけん", "roma": "tottori-ken"},
	32: {"kanji": "島根県", "kana": "しまねけん", "roma": "shimane-ken"},
	33: {"kanji": "岡山県", "kana": "おかやまけん", "roma": "okayama-ken"},
	34: {"kanji": "広島県", "kana": "ひろしまけん", "roma": "hiroshima-ken"},
	35: {"kanji": "山口県", "kana": "やまぐちけん", "roma": "yamaguchi-ken"},
	36: {"kanji": "徳島県", "kana": "とくしまけん", "roma": "tokushima-ken"},
	37: {"kanji": "香川県", "kana": "かがわけん", "roma": "kagawa-ken"},
	38: {"kanji": "愛媛県", "kana": "えひめけん", "roma": "ehime-ken"},
	39: {"kanji": "高知県", "kana": "こうちけん", "roma": "kochi-ken"},
	40: {"kanji": "福岡県", "kana": "ふくおかけん", "roma": "fukuoka-ken"},
	41: {"kanji": "佐賀県", "kana": "さがけん", "roma": "saga-ken"},
	42: {"kanji": "長崎県", "kana": "ながさきけん", "roma": "nagasaki-ken"},
	43: {"kanji": "熊本県", "kana": "くまもとけん", "roma": "oita-ken"},
	44: {"kanji": "大分県", "kana": "おおいたけん", "roma": "kumamoto-ken"},
	45: {"kanji": "宮崎県", "kana": "みやざきけん", "roma": "miyazaki-ken"},
	46: {"kanji": "鹿児島県", "kana": "かごしまけん", "roma": "kagoshima-ken"},
	47: {"kanji": "沖縄県", "kana": "おきなわけん", "roma": "okinawa-ken"},
}

const (
	prefectureSize = 47

	prefectureCodeOfHokkaido = 1
	prefectureCodeOfTokyo    = 13
	prefectureCodeOfKyoto    = 26
	prefectureCodeOfOsaka    = 27
)

// nameFindMap は漢字、かな、ローマ字で都道府県情報を検索しやすくするためのインデックスマップです。
var nameFindMap = func() map[string]map[string]int {
	findMap := map[string]map[string]int{
		"kanji": make(map[string]int, prefectureSize),
		"kana":  make(map[string]int, prefectureSize),
		"roma":  make(map[string]int, prefectureSize),
	}

	for code, texts := range prefectureMap {
		findMap["kanji"][texts["kanji"]] = code
		findMap["kana"][texts["kana"]] = code
		findMap["roma"][texts["roma"]] = code

		switch code {
		case prefectureCodeOfHokkaido:
			findMap["kanji"][texts["kanji"]] = code
			findMap["kana"][texts["kana"]] = code
			findMap["roma"][texts["roma"]] = code

		case prefectureCodeOfTokyo:
			kanjiIndex := strings.TrimSuffix(texts["kanji"], "都")
			findMap["kanji"][kanjiIndex] = code

			kanaIndex := strings.TrimSuffix(texts["kana"], "と")
			findMap["kana"][kanaIndex] = code

			romaIndex := strings.TrimSuffix(texts["roma"], "-to")
			findMap["roma"][romaIndex] = code

		case prefectureCodeOfKyoto, prefectureCodeOfOsaka:
			kanjiIndex := strings.TrimSuffix(texts["kanji"], "府")
			findMap["kanji"][kanjiIndex] = code

			kanaIndex := strings.TrimSuffix(texts["kana"], "ふ")
			findMap["kana"][kanaIndex] = code

			romaIndex := strings.TrimSuffix(texts["roma"], "-fu")
			findMap["roma"][romaIndex] = code

		default:
			kanjiIndex := strings.TrimSuffix(texts["kanji"], "県")
			findMap["kanji"][kanjiIndex] = code

			kanaIndex := strings.TrimSuffix(texts["kana"], "けん")
			findMap["kana"][kanaIndex] = code

			romaIndex := strings.TrimSuffix(texts["roma"], "-ken")
			findMap["roma"][romaIndex] = code

		}
	}

	return findMap
}()

// List は全ての都道府県情報を返します。
func List() []Prefecture {
	var prefectures = make([]Prefecture, prefectureSize)

	for code, texts := range prefectureMap {
		prefecture := &prefecture{code, texts["kanji"], texts["kana"], texts["roma"]}
		prefectures[code-1] = prefecture
	}

	return prefectures
}

// FindByCode は与えた都道府県コードに対応する都道府県情報を返します。
// 対応する都道府県情報が見つからない場合、`ok`に`false`が設定され、都道府県情報は`nil`で返されます。
func FindByCode(code int) (Prefecture, bool) {
	texts, ok := prefectureMap[code]

	if !ok {
		return nil, false
	}

	prefecture := &prefecture{code, texts["kanji"], texts["kana"], texts["roma"]}
	return prefecture, true
}

// FindByKanji は与えた都道府県の漢字名に対応する都道府県情報を返します。
// 与える値の末尾にある「都」、「府」、「県」は省略可能です。
// 対応する都道府県情報が見つからない場合、`ok`に`false`が設定され、都道府県情報は`nil`で返されます。
func FindByKanji(kanji string) (Prefecture, bool) {
	code, ok := nameFindMap["kanji"][kanji]

	if !ok {
		return nil, false
	}

	texts := prefectureMap[code]
	prefecture := &prefecture{code, texts["kanji"], texts["kana"], texts["roma"]}
	return prefecture, true
}

// FindByKana は与えた都道府県のかな名に対応する都道府県情報を返します。
// 与える値の末尾にある「と」、「ふ」、「けん」は省略可能です。
// 対応する都道府県情報が見つからない場合、`ok`に`false`が設定され、都道府県情報は`nil`で返されます。
func FindByKana(kana string) (Prefecture, bool) {
	code, ok := nameFindMap["kana"][kana]

	if !ok {
		return nil, false
	}

	texts := prefectureMap[code]
	prefecture := &prefecture{code, texts["kanji"], texts["kana"], texts["roma"]}
	return prefecture, true
}

// FindByRoma は与えた都道府県のローマ字名に対応する都道府県情報を返します。
// 与える値の末尾にある「-to」、「-fu」、「-ken」は省略可能です。
// 与える値はUpperCamelCase、LowerCamelCaseを問いません。
// 対応する都道府県情報が見つからない場合、`ok`に`false`が設定され、都道府県情報は`nil`で返されます。
func FindByRoma(roma string) (Prefecture, bool) {
	code, ok := nameFindMap["roma"][strings.ToLower(roma)]

	if !ok {
		return nil, false
	}

	texts := prefectureMap[code]
	prefecture := &prefecture{code, texts["kanji"], texts["kana"], texts["roma"]}
	return prefecture, true
}
