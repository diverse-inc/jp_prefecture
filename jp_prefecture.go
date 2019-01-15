package jp_prefecture

import (
	"strings"
)

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

var nameFindMap = func() map[string]map[string]int {
	findMap := map[string]map[string]int{
		"kanji": make(map[string]int, 47),
		"kana":  make(map[string]int, 47),
		"roma":  make(map[string]int, 47),
	}

	for code, texts := range prefectureMap {
		findMap["kanji"][texts["kanji"]] = code
		findMap["kana"][texts["kana"]] = code
		findMap["roma"][texts["roma"]] = code

		switch code {
		case 1:
			findMap["kanji"][texts["kanji"]] = code
			findMap["kana"][texts["kana"]] = code
			findMap["roma"][texts["roma"]] = code

		case 13:
			kanjiIndex := strings.TrimSuffix(texts["kanji"], "都")
			findMap["kanji"][kanjiIndex] = code

			kanaIndex := strings.TrimSuffix(texts["kana"], "と")
			findMap["kana"][kanaIndex] = code

			romaIndex := strings.TrimSuffix(texts["roma"], "-to")
			findMap["roma"][romaIndex] = code

		case 26, 27:
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

type prefecture struct {
	code  int
	kanji string
	kana  string
	roma  string
}

type Prefecture interface {
	Code() int
	Kanji() string
	Kana() string
	Roma() string
}

func (p *prefecture) Code() int {
	return p.code
}

func (p *prefecture) Kanji() string {
	return p.kanji
}

func (p *prefecture) Kana() string {
	return p.kana
}

func (p *prefecture) Roma() string {
	return p.roma
}

func List() []Prefecture {
	var prefectures = make([]Prefecture, 47)

	for code, texts := range prefectureMap {
		prefecture := &prefecture{code, texts["kanji"], texts["kana"], texts["roma"]}
		prefectures[code-1] = prefecture
	}

	return prefectures
}

func FindByCode(code int) (Prefecture, bool) {
	texts, ok := prefectureMap[code]

	if !ok {
		return nil, false
	}

	prefecture := &prefecture{code, texts["kanji"], texts["kana"], texts["roma"]}
	return prefecture, true
}

func FindByKanji(kanji string) (Prefecture, bool) {
	code, ok := nameFindMap["kanji"][kanji]

	if !ok {
		return nil, false
	}

	texts := prefectureMap[code]
	prefecture := &prefecture{code, texts["kanji"], texts["kana"], texts["roma"]}
	return prefecture, true
}

func FindByKana(kana string) (Prefecture, bool) {
	code, ok := nameFindMap["kana"][kana]

	if !ok {
		return nil, false
	}

	texts := prefectureMap[code]
	prefecture := &prefecture{code, texts["kanji"], texts["kana"], texts["roma"]}
	return prefecture, true
}

func FindByRoma(roma string) (Prefecture, bool) {
	code, ok := nameFindMap["roma"][strings.ToLower(roma)]

	if !ok {
		return nil, false
	}

	texts := prefectureMap[code]
	prefecture := &prefecture{code, texts["kanji"], texts["kana"], texts["roma"]}
	return prefecture, true
}
