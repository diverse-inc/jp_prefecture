package jp_prefecture

import (
	"testing"
)

var regionTestData = &region{kanji: "kanji", kana: "kana", roma: "roma", prefectureCodes: []int{JISCodeHokkaido}}

func TestRegion_Kanji(t *testing.T) {
	if regionTestData.Kanji() != "kanji" {
		t.Fatal()
	}
}

func TestRegion_Kana(t *testing.T) {
	if regionTestData.Kana() != "kana" {
		t.Fatal()
	}
}

func TestRegion_Roma(t *testing.T) {
	if regionTestData.Roma() != "roma" {
		t.Fatal()
	}
}

func TestRegion_List(t *testing.T) {
	list := regionTestData.List()

	if len(list) != 1 {
		t.Fatal()
	}
	if list[0].Code() != JISCodeHokkaido {
		t.Fatal()
	}
}