package jp_prefecture

type Region interface {
	Kanji() string
	Kana() string
	Roma() string
	List() []Prefecture
}

type region struct {
	kanji           string
	kana            string
	roma            string
	prefectureCodes []int
}

func (r *region) Kanji() string {
	return r.kanji
}

func (r *region) Kana() string {
	return r.kana
}

func (r *region) Roma() string {
	return r.roma
}

func (r *region) List() []Prefecture {
	var list = make([]Prefecture, len(r.prefectureCodes))

	for i, code := range r.prefectureCodes {
		list[i] = prefectureMap[code]
	}

	return list
}
