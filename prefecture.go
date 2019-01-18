package jp_prefecture

import (
	"strings"
)

// prefecture は都道府県情報を定義する構造体です。
type prefecture struct {
	code  int
	kanji string
	kana  string
	roma  string
}

// Prefecture は外部に公開するための都道府県インターフェースです。
type Prefecture interface {
	Code() int
	Kanji() string
	Kana() string
	Roma() string
}

// Code は都道府県コードを返します。
func (p *prefecture) Code() int {
	return p.code
}

// Kanji は都道府県の漢字名を返します、末尾には「都」、「府」、「県」が付与された状態で値が返ります。
func (p *prefecture) Kanji() string {
	return p.kanji
}

// KanjiShort は都道府県名の末尾から「都」、「府」、「県」を除いた漢字名を返します。
func (p *prefecture) KanjiShort() string {
	switch p.code {
	case JISCodeHokkaido:
		return p.kanji
	case JISCodeTokyo:
		return strings.TrimSuffix(p.kanji, "都")
	case JISCodeKyoto, JISCodeOsaka:
		return strings.TrimSuffix(p.kanji, "府")
	default:
		return strings.TrimSuffix(p.kanji, "県")
	}
}

// Kana は都道府県のかな名を返します、末尾には「と」、「ふ」、「けん」が付与された状態で値が返ります。
func (p *prefecture) Kana() string {
	return p.kana
}

// Roma は都道府県のローマ字名を返します、末尾には「-to」、「-fu」、「-ken」が付与された状態で値が返ります。
func (p *prefecture) Roma() string {
	return p.roma
}
