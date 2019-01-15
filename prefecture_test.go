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
