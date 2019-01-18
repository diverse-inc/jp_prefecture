package jp_prefecture

import (
	"strings"
)

const (
	JISCodeHokkaido  = 1
	JISCodeAomoti    = 2
	JISCodeIwate     = 3
	JISCodeMiyagi    = 4
	JISCodeAkita     = 5
	JISCodeYamagata  = 6
	JISCodeFukushima = 7
	JISCodeIbaraki   = 8
	JISCodeTochigi   = 9
	JISCodeGunma     = 10
	JISCodeSaitama   = 11
	JISCodeChiba     = 12
	JISCodeTokyo     = 13
	JISCodeKanagawa  = 14
	JISCodeNiigata   = 15
	JISCodeToyama    = 16
	JISCodeIshikawa  = 17
	JISCodeFukui     = 18
	JISCodeYamanashi = 19
	JISCodeNagano    = 20
	JISCodeGifu      = 21
	JISCodeShizuoka  = 22
	JISCodeAichi     = 23
	JISCodeMie       = 24
	JISCodeShiga     = 25
	JISCodeKyoto     = 26
	JISCodeOsaka     = 27
	JISCodeHyogo     = 28
	JISCodeNara      = 29
	JISCodeWakayama  = 30
	JISCodeTottori   = 31
	JISCodeShimane   = 32
	JISCodeOkayama   = 33
	JISCodeHitoshima = 34
	JISCodeYamaguchi = 35
	JISCodeTokushima = 36
	JISCodeKagawa    = 37
	JISCodeEhime     = 38
	JISCodeKochi     = 39
	JISCodeFukuoka   = 40
	JISCodeSaga      = 41
	JISCodeNagasaki  = 42
	JISCodeKumamoto  = 43
	JISCodeOita      = 44
	JISCodeMiyazaki  = 45
	JISCodeKagoshima = 46
	JISCodeOlinawa   = 47
)

// prefectureMap は都道府県の情報を”コード：名称のMap（漢字、かな、ローマ字）”形式で定義した値です。
var prefectureMap = map[int]*prefecture{
	1:  {code: 1, kanji: "北海道", kana: "ほっかいどう", roma: "hokkaido"},
	2:  {code: 2, kanji: "青森県", kana: "あおもりけん", roma: "aomori-ken"},
	3:  {code: 3, kanji: "岩手県", kana: "いわてけん", roma: "iwate-ken"},
	4:  {code: 4, kanji: "宮城県", kana: "みやぎけん", roma: "miyagi-ken"},
	5:  {code: 5, kanji: "秋田県", kana: "あきたけん", roma: "akita-ken"},
	6:  {code: 6, kanji: "山形県", kana: "やまがたけん", roma: "yamagata-ken"},
	7:  {code: 7, kanji: "福島県", kana: "ふくしまけん", roma: "fukushima-ken"},
	8:  {code: 8, kanji: "茨城県", kana: "いばらきけん", roma: "ibaraki-ken"},
	9:  {code: 9, kanji: "栃木県", kana: "とちぎけん", roma: "tochigi-ken"},
	10: {code: 10, kanji: "群馬県", kana: "ぐんまけん", roma: "gunma-ken"},
	11: {code: 11, kanji: "埼玉県", kana: "さいたまけん", roma: "saitama-ken"},
	12: {code: 12, kanji: "千葉県", kana: "ちばけん", roma: "chiba-ken"},
	13: {code: 13, kanji: "東京都", kana: "とうきょうと", roma: "tokyo-to"},
	14: {code: 14, kanji: "神奈川県", kana: "かながわけん", roma: "kanagawa-ken"},
	15: {code: 15, kanji: "新潟県", kana: "にいがたけん", roma: "niigata-ken"},
	16: {code: 16, kanji: "富山県", kana: "とやまけん", roma: "toyama-ken"},
	17: {code: 17, kanji: "石川県", kana: "いしかわけん", roma: "ishikawa-ken"},
	18: {code: 18, kanji: "福井県", kana: "ふくいけん", roma: "fukui-ken"},
	19: {code: 19, kanji: "山梨県", kana: "やまなしけん", roma: "yamanashi-ken"},
	20: {code: 20, kanji: "長野県", kana: "ながのけん", roma: "nagano-ken"},
	21: {code: 21, kanji: "岐阜県", kana: "ぎふけん", roma: "gifu-ken"},
	22: {code: 22, kanji: "静岡県", kana: "しずおかけん", roma: "shizuoka-ken"},
	23: {code: 23, kanji: "愛知県", kana: "あいちけん", roma: "aichi-ken"},
	24: {code: 24, kanji: "三重県", kana: "みえけん", roma: "mie-ken"},
	25: {code: 25, kanji: "滋賀県", kana: "しがけん", roma: "shiga-ken"},
	26: {code: 26, kanji: "京都府", kana: "きょうとふ", roma: "kyoto-fu"},
	27: {code: 27, kanji: "大阪府", kana: "おおさかふ", roma: "osaka-fu"},
	28: {code: 28, kanji: "兵庫県", kana: "ひょうごけん", roma: "hyogo-ken"},
	29: {code: 29, kanji: "奈良県", kana: "ならけん", roma: "nara-ken"},
	30: {code: 30, kanji: "和歌山県", kana: "わかやまけん", roma: "wakayama-ken"},
	31: {code: 31, kanji: "鳥取県", kana: "とっとりけん", roma: "tottori-ken"},
	32: {code: 32, kanji: "島根県", kana: "しまねけん", roma: "shimane-ken"},
	33: {code: 33, kanji: "岡山県", kana: "おかやまけん", roma: "okayama-ken"},
	34: {code: 34, kanji: "広島県", kana: "ひろしまけん", roma: "hiroshima-ken"},
	35: {code: 35, kanji: "山口県", kana: "やまぐちけん", roma: "yamaguchi-ken"},
	36: {code: 36, kanji: "徳島県", kana: "とくしまけん", roma: "tokushima-ken"},
	37: {code: 37, kanji: "香川県", kana: "かがわけん", roma: "kagawa-ken"},
	38: {code: 38, kanji: "愛媛県", kana: "えひめけん", roma: "ehime-ken"},
	39: {code: 39, kanji: "高知県", kana: "こうちけん", roma: "kochi-ken"},
	40: {code: 40, kanji: "福岡県", kana: "ふくおかけん", roma: "fukuoka-ken"},
	41: {code: 41, kanji: "佐賀県", kana: "さがけん", roma: "saga-ken"},
	42: {code: 42, kanji: "長崎県", kana: "ながさきけん", roma: "nagasaki-ken"},
	43: {code: 43, kanji: "熊本県", kana: "くまもとけん", roma: "kumamoto-ken"},
	44: {code: 44, kanji: "大分県", kana: "おおいたけん", roma: "oita-ken"},
	45: {code: 45, kanji: "宮崎県", kana: "みやざきけん", roma: "miyazaki-ken"},
	46: {code: 46, kanji: "鹿児島県", kana: "かごしまけん", roma: "kagoshima-ken"},
	47: {code: 47, kanji: "沖縄県", kana: "おきなわけん", roma: "okinawa-ken"},
}

const (
	prefectureSize = 47

	prefectureCodeOfHokkaido = 1
	prefectureCodeOfTokyo    = 13
	prefectureCodeOfKyoto    = 26
	prefectureCodeOfOsaka    = 27

	findMapKeyKanji = 0
	findMayKeyKana  = 1
	findMapKeyRoma  = 2
)

// nameFindMap は漢字、かな、ローマ字で都道府県情報を検索しやすくするためのインデックスマップです。
var nameFindMap = func() map[uint8]map[string]int {
	findMap := map[uint8]map[string]int{
		findMapKeyKanji: make(map[string]int, prefectureSize),
		findMayKeyKana:  make(map[string]int, prefectureSize),
		findMapKeyRoma:  make(map[string]int, prefectureSize),
	}

	for code, prefecture := range prefectureMap {
		findMap[findMapKeyKanji][prefecture.kanji] = code
		findMap[findMayKeyKana][prefecture.kana] = code
		findMap[findMapKeyRoma][prefecture.roma] = code

		switch code {
		case prefectureCodeOfHokkaido:
			findMap[findMapKeyKanji][prefecture.kanji] = code
			findMap[findMayKeyKana][prefecture.kana] = code
			findMap[findMapKeyRoma][prefecture.roma] = code

		case prefectureCodeOfTokyo:
			kanjiIndex := strings.TrimSuffix(prefecture.kanji, "都")
			findMap[findMapKeyKanji][kanjiIndex] = code

			kanaIndex := strings.TrimSuffix(prefecture.kana, "と")
			findMap[findMayKeyKana][kanaIndex] = code

			romaIndex := strings.TrimSuffix(prefecture.roma, "-to")
			findMap[findMapKeyRoma][romaIndex] = code

		case prefectureCodeOfKyoto, prefectureCodeOfOsaka:
			kanjiIndex := strings.TrimSuffix(prefecture.kanji, "府")
			findMap[findMapKeyKanji][kanjiIndex] = code

			kanaIndex := strings.TrimSuffix(prefecture.kana, "ふ")
			findMap[findMayKeyKana][kanaIndex] = code

			romaIndex := strings.TrimSuffix(prefecture.roma, "-fu")
			findMap[findMapKeyRoma][romaIndex] = code

		default:
			kanjiIndex := strings.TrimSuffix(prefecture.kanji, "県")
			findMap[findMapKeyKanji][kanjiIndex] = code

			kanaIndex := strings.TrimSuffix(prefecture.kana, "けん")
			findMap[findMayKeyKana][kanaIndex] = code

			romaIndex := strings.TrimSuffix(prefecture.roma, "-ken")
			findMap[findMapKeyRoma][romaIndex] = code

		}
	}

	return findMap
}()

// List は全ての都道府県情報を返します。
func List() []Prefecture {
	var prefectures = make([]Prefecture, prefectureSize)

	for code, prefecture := range prefectureMap {
		prefectures[code-1] = prefecture
	}

	return prefectures
}

// FindByCode は与えた都道府県コードに対応する都道府県情報を返します。
// 対応する都道府県情報が見つからない場合、`ok`に`false`が設定され、都道府県情報は`nil`で返されます。
func FindByCode(code int) (Prefecture, bool) {
	prefecture, ok := prefectureMap[code]

	if !ok {
		return nil, false
	}

	return prefecture, true
}

// FindByKanji は与えた都道府県の漢字名に対応する都道府県情報を返します。
// 与える値の末尾にある「都」、「府」、「県」は省略可能です。
// 対応する都道府県情報が見つからない場合、`ok`に`false`が設定され、都道府県情報は`nil`で返されます。
func FindByKanji(kanji string) (Prefecture, bool) {
	code, ok := nameFindMap[findMapKeyKanji][kanji]

	if !ok {
		return nil, false
	}

	return prefectureMap[code], true
}

// FindByKana は与えた都道府県のかな名に対応する都道府県情報を返します。
// 与える値の末尾にある「と」、「ふ」、「けん」は省略可能です。
// 対応する都道府県情報が見つからない場合、`ok`に`false`が設定され、都道府県情報は`nil`で返されます。
func FindByKana(kana string) (Prefecture, bool) {
	code, ok := nameFindMap[findMayKeyKana][kana]

	if !ok {
		return nil, false
	}

	return prefectureMap[code], true
}

// FindByRoma は与えた都道府県のローマ字名に対応する都道府県情報を返します。
// 与える値の末尾にある「-to」、「-fu」、「-ken」は省略可能です。
// 与える値はUpperCamelCase、LowerCamelCaseを問いません。
// 対応する都道府県情報が見つからない場合、`ok`に`false`が設定され、都道府県情報は`nil`で返されます。
func FindByRoma(roma string) (Prefecture, bool) {
	code, ok := nameFindMap[findMapKeyRoma][strings.ToLower(roma)]

	if !ok {
		return nil, false
	}

	return prefectureMap[code], true
}
