package jp_prefecture

import (
	"strings"
)

const (
	JISCodeHokkaido  = 1
	JISCodeAomori    = 2
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
	JISCodeHiroshima = 34
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
	JISCodeOkinawa   = 47
)

// prefectureMap は都道府県の情報を”コード：名称のMap（漢字、かな、ローマ字）”形式で定義した値です。
var prefectureMap = map[int]*prefecture{
	JISCodeHokkaido:  {code: JISCodeHokkaido, kanji: "北海道", kana: "ほっかいどう", roma: "hokkaido"},
	JISCodeAomori:    {code: JISCodeAomori, kanji: "青森県", kana: "あおもりけん", roma: "aomori-ken"},
	JISCodeIwate:     {code: JISCodeIwate, kanji: "岩手県", kana: "いわてけん", roma: "iwate-ken"},
	JISCodeMiyagi:    {code: JISCodeMiyagi, kanji: "宮城県", kana: "みやぎけん", roma: "miyagi-ken"},
	JISCodeAkita:     {code: JISCodeAkita, kanji: "秋田県", kana: "あきたけん", roma: "akita-ken"},
	JISCodeYamagata:  {code: JISCodeYamagata, kanji: "山形県", kana: "やまがたけん", roma: "yamagata-ken"},
	JISCodeFukushima: {code: JISCodeFukushima, kanji: "福島県", kana: "ふくしまけん", roma: "fukushima-ken"},
	JISCodeIbaraki:   {code: JISCodeIbaraki, kanji: "茨城県", kana: "いばらきけん", roma: "ibaraki-ken"},
	JISCodeTochigi:   {code: JISCodeTochigi, kanji: "栃木県", kana: "とちぎけん", roma: "tochigi-ken"},
	JISCodeGunma:     {code: JISCodeGunma, kanji: "群馬県", kana: "ぐんまけん", roma: "gunma-ken"},
	JISCodeSaitama:   {code: JISCodeSaitama, kanji: "埼玉県", kana: "さいたまけん", roma: "saitama-ken"},
	JISCodeChiba:     {code: JISCodeChiba, kanji: "千葉県", kana: "ちばけん", roma: "chiba-ken"},
	JISCodeTokyo:     {code: JISCodeTokyo, kanji: "東京都", kana: "とうきょうと", roma: "tokyo-to"},
	JISCodeKanagawa:  {code: JISCodeKanagawa, kanji: "神奈川県", kana: "かながわけん", roma: "kanagawa-ken"},
	JISCodeNiigata:   {code: JISCodeNiigata, kanji: "新潟県", kana: "にいがたけん", roma: "niigata-ken"},
	JISCodeToyama:    {code: JISCodeToyama, kanji: "富山県", kana: "とやまけん", roma: "toyama-ken"},
	JISCodeIshikawa:  {code: JISCodeIshikawa, kanji: "石川県", kana: "いしかわけん", roma: "ishikawa-ken"},
	JISCodeFukui:     {code: JISCodeFukui, kanji: "福井県", kana: "ふくいけん", roma: "fukui-ken"},
	JISCodeYamanashi: {code: JISCodeYamanashi, kanji: "山梨県", kana: "やまなしけん", roma: "yamanashi-ken"},
	JISCodeNagano:    {code: JISCodeNagano, kanji: "長野県", kana: "ながのけん", roma: "nagano-ken"},
	JISCodeGifu:      {code: JISCodeGifu, kanji: "岐阜県", kana: "ぎふけん", roma: "gifu-ken"},
	JISCodeShizuoka:  {code: JISCodeShizuoka, kanji: "静岡県", kana: "しずおかけん", roma: "shizuoka-ken"},
	JISCodeAichi:     {code: JISCodeAichi, kanji: "愛知県", kana: "あいちけん", roma: "aichi-ken"},
	JISCodeMie:       {code: JISCodeMie, kanji: "三重県", kana: "みえけん", roma: "mie-ken"},
	JISCodeShiga:     {code: JISCodeShiga, kanji: "滋賀県", kana: "しがけん", roma: "shiga-ken"},
	JISCodeKyoto:     {code: JISCodeKyoto, kanji: "京都府", kana: "きょうとふ", roma: "kyoto-fu"},
	JISCodeOsaka:     {code: JISCodeOsaka, kanji: "大阪府", kana: "おおさかふ", roma: "osaka-fu"},
	JISCodeHyogo:     {code: JISCodeHyogo, kanji: "兵庫県", kana: "ひょうごけん", roma: "hyogo-ken"},
	JISCodeNara:      {code: JISCodeNara, kanji: "奈良県", kana: "ならけん", roma: "nara-ken"},
	JISCodeWakayama:  {code: JISCodeWakayama, kanji: "和歌山県", kana: "わかやまけん", roma: "wakayama-ken"},
	JISCodeTottori:   {code: JISCodeTottori, kanji: "鳥取県", kana: "とっとりけん", roma: "tottori-ken"},
	JISCodeShimane:   {code: JISCodeShimane, kanji: "島根県", kana: "しまねけん", roma: "shimane-ken"},
	JISCodeOkayama:   {code: JISCodeOkayama, kanji: "岡山県", kana: "おかやまけん", roma: "okayama-ken"},
	JISCodeHiroshima: {code: JISCodeHiroshima, kanji: "広島県", kana: "ひろしまけん", roma: "hiroshima-ken"},
	JISCodeYamaguchi: {code: JISCodeYamaguchi, kanji: "山口県", kana: "やまぐちけん", roma: "yamaguchi-ken"},
	JISCodeTokushima: {code: JISCodeTokushima, kanji: "徳島県", kana: "とくしまけん", roma: "tokushima-ken"},
	JISCodeKagawa:    {code: JISCodeKagawa, kanji: "香川県", kana: "かがわけん", roma: "kagawa-ken"},
	JISCodeEhime:     {code: JISCodeEhime, kanji: "愛媛県", kana: "えひめけん", roma: "ehime-ken"},
	JISCodeKochi:     {code: JISCodeKochi, kanji: "高知県", kana: "こうちけん", roma: "kochi-ken"},
	JISCodeFukuoka:   {code: JISCodeFukuoka, kanji: "福岡県", kana: "ふくおかけん", roma: "fukuoka-ken"},
	JISCodeSaga:      {code: JISCodeSaga, kanji: "佐賀県", kana: "さがけん", roma: "saga-ken"},
	JISCodeNagasaki:  {code: JISCodeNagasaki, kanji: "長崎県", kana: "ながさきけん", roma: "nagasaki-ken"},
	JISCodeKumamoto:  {code: JISCodeKumamoto, kanji: "熊本県", kana: "くまもとけん", roma: "kumamoto-ken"},
	JISCodeOita:      {code: JISCodeOita, kanji: "大分県", kana: "おおいたけん", roma: "oita-ken"},
	JISCodeMiyazaki:  {code: JISCodeMiyazaki, kanji: "宮崎県", kana: "みやざきけん", roma: "miyazaki-ken"},
	JISCodeKagoshima: {code: JISCodeKagoshima, kanji: "鹿児島県", kana: "かごしまけん", roma: "kagoshima-ken"},
	JISCodeOkinawa:   {code: JISCodeOkinawa, kanji: "沖縄県", kana: "おきなわけん", roma: "okinawa-ken"},
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
