package jp_prefecture

import (
	"testing"
)

var prefectureTestData = &prefecture{1, "kanji", "kana", "roma"}

func TestPrefecture_Code(t *testing.T) {
	if prefectureTestData.Code() != 1 {
		t.Fatal()
	}
}

func TestPrefecture_Kanji(t *testing.T) {
	if prefectureTestData.Kanji() != "kanji" {
		t.Fatal()
	}
}

func TestPrefecture_KanjiShort(t *testing.T) {
	testData := map[Prefecture]string{
		prefectureMap[JISCodeHokkaido]: "北海道",
		prefectureMap[JISCodeTokyo]:    "東京",
		prefectureMap[JISCodeKyoto]:    "京都",
		prefectureMap[JISCodeOsaka]:    "大阪",
		prefectureMap[JISCodeOkinawa]:  "沖縄",
	}

	for pref, expect := range testData {
		kanji := pref.KanjiShort()
		if kanji != expect {
			t.Fatal()
		}
	}
}

func TestPrefecture_Kana(t *testing.T) {
	if prefectureTestData.Kana() != "kana" {
		t.Fatal()
	}
}

func TestPrefecture_KanaShort(t *testing.T) {
	testData := map[Prefecture]string{
		prefectureMap[JISCodeHokkaido]: "ほっかいどう",
		prefectureMap[JISCodeTokyo]:    "とうきょう",
		prefectureMap[JISCodeKyoto]:    "きょうと",
		prefectureMap[JISCodeOsaka]:    "おおさか",
		prefectureMap[JISCodeOkinawa]:  "おきなわ",
	}

	for pref, expect := range testData {
		kana := pref.KanaShort()
		if kana != expect {
			t.Fatal()
		}
	}
}

func TestPrefecture_Roma(t *testing.T) {
	if prefectureTestData.Roma() != "roma" {
		t.Fatal()
	}
}
