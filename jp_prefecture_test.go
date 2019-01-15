package jp_prefecture

import (
	"testing"
)

func TestFindByCode(t *testing.T) {
	var (
		prefecture Prefecture
		ok bool
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
