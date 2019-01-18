package jp_prefecture

type Region interface {
	Kanji() string
	Kana() string
	Roma() string
	List() []Prefecture
}
