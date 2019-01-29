package jp_prefecture

// Region は外部に公開するための地域インターフェースです。
type Region interface {
	Kanji() string
	Kana() string
	Roma() string
	List() []Prefecture
}

// region は地域情報を定義する構造体です。
type region struct {
	kanji           string
	kana            string
	roma            string
	prefectureCodes []int
}

// Kanji は地域の漢字名を返します。
func (r *region) Kanji() string {
	return r.kanji
}

// Kana は地域のかな名を返します。
func (r *region) Kana() string {
	return r.kana
}

// Roma は地域のローマ字名を返します。
func (r *region) Roma() string {
	return r.roma
}

// List は地域の都道府県情報を返します。
func (r *region) List() []Prefecture {
	var list = make([]Prefecture, len(r.prefectureCodes))

	for i, code := range r.prefectureCodes {
		list[i] = prefectureMap[code]
	}

	return list
}
