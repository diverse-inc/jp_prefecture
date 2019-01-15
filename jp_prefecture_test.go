package jp_prefecture

import (
	"testing"
)

func TestList(t *testing.T) {
	list := List()

	if len(list) != 47 {
		t.Fatal()
	}

	for i, prefecture := range list {
		if prefecture == nil {
			t.Fatal()
		}
		if prefecture.Code() != i+1 {
			t.Fatal()
		}
	}
}

func TestFindByCode(t *testing.T) {
	var (
		prefecture Prefecture
		ok         bool
	)

	prefecture, ok = FindByCode(1)
	if !ok {
		t.Fatal()
	}
	if prefecture == nil {
		t.Fatal()
	}

	prefecture, ok = FindByCode(0)
	if ok {
		t.Fatal()
	}
	if prefecture != nil {
		t.Fatal()
	}
}

func TestFindByKanji(t *testing.T) {
	findValues := []string{
		"北海道", "東京", "東京都", "大阪", "大阪府", "京都", "京都府", "沖縄", "沖縄県",
	}

	for _, value := range findValues {
		prefecture, ok := FindByKanji(value)
		if !ok {
			t.Fatal()
		}
		if prefecture == nil {
			t.Fatal()
		}
	}

	prefecture, ok := FindByKanji("北海")
	if ok {
		t.Fatal()
	}
	if prefecture != nil {
		t.Fatal()
	}
}

func TestFindByKana(t *testing.T) {
	findValues := []string{
		"ほっかいどう", "とうきょう", "とうきょうと", "おおさか", "おおさかふ",
		"きょうと", "きょうとふ", "おきなわ", "おきなわけん",
	}

	for _, value := range findValues {
		prefecture, ok := FindByKana(value)
		if !ok {
			t.Fatal()
		}
		if prefecture == nil {
			t.Fatal()
		}
	}

	prefecture, ok := FindByKana("ほっかい")
	if ok {
		t.Fatal()
	}
	if prefecture != nil {
		t.Fatal()
	}
}

func TestFindByRoma(t *testing.T) {
	findValues := []string{
		"hokkaido", "tokyo", "tokyo-to", "osaka", "osaka-fu",
		"kyoto", "kyoto-fu", "okinawa", "okinawa-ken",

		"Hokkaido", "TOKYO", "Tokyo-To", "oSAKA",
	}

	for _, value := range findValues {
		prefecture, ok := FindByRoma(value)
		if !ok {
			t.Fatal()
		}
		if prefecture == nil {
			t.Fatal()
		}
	}

	prefecture, ok := FindByRoma("tokyoto")
	if ok {
		t.Fatal()
	}
	if prefecture != nil {
		t.Fatal()
	}
}

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

func TestPrefecture_Kana(t *testing.T) {
	if prefectureTestData.Kana() != "kana" {
		t.Fatal()
	}
}

func TestPrefecture_Roma(t *testing.T) {
	if prefectureTestData.Roma() != "roma" {
		t.Fatal()
	}
}
