package jp_prefecture

import (
	"testing"
)

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
