package uni

import (
	"archive/zip"
	"bytes"
	"embed"
	"io/fs"
	"sync"

	"github.com/ikawaha/kagome-dict/dict"
)

type FeatureIndex = int

const (
	// Features are information given to a word, such as follows:
	// 公園	名詞,普通名詞,一般,*,*,*,コウエン,公園,公園,コーエン,公園,コーエン,漢,*,*,*,*
	// に	助詞,格助詞,*,*,*,*,ニ,に,に,ニ,に,ニ,和,*,*,*,*
	// 行っ	動詞,非自立可能,*,*,五段-カ行,連用形-促音便,イク,行く,行っ,イッ,行く,イク,和,*,*,*,*
	// た	助動詞,*,*,*,助動詞-タ,終止形-一般,タ,た,た,タ,た,タ,和,*,*,*,*
	// EOS

	// POSHierarchy represents part-of-speech hierarchy
	// e.g. Columns 動詞,非自立可能,*,* are POSs which hierarchy depth is 4.
	POSHierarchy = 4
	// CType represents  活用型 (e.g. 五段-カ行).
	CType = 4
	// CForm represents 活用形 (e.g. 連用形-促音便).
	CForm = 5
	// LForm represents 読み (e.g. コウエン).
	LForm = 6
	// Lemma represents 語彙素 (e.g. 公園, 行く).
	Lemma = 7
	// Orth represents 書字形出現形.
	Orth = 8
	// Pron represents 発音形出現形.
	Pron = 9
	// OrthBase represents 書字形基本型.
	OrthBase = 10
	// PronBase represents 発音形基本型.
	PronBase = 11
	// Goshu represents 語種.
	Goshu = 12
	// IType represents 語頭変化型.
	IType = 13
	// IForm represents 語頭変化形.
	IForm = 14
	// FType represents 語末変化型.
	FType = 15
	// FForm represents 語末変化形.
	FForm = 16

	// Aliases

	// InflectionalType represents 活用型 (e.g. 五段-カ行), an alias for CType.
	InflectionalType FeatureIndex = 4
	// InflectionalForm represents 活用形 (e.g. 連用形-促音便), an alias for CForm.
	InflectionalForm = 5
	// BaseForm represents 基本形 (e.g. 行く), an alias for Lemma.
	BaseForm = 7
	// Pronunciation represents 発音 (e.g. コーエン), an alias for Pron.
	Pronunciation = 9
)

type systemDict struct {
	once sync.Once
	dict *dict.Dict
}

var (
	full   systemDict
	shrink systemDict
)

// Dict returns a dictionary.
func Dict() *dict.Dict {
	full.once.Do(func() {
		full.dict = loadDict(true)
		shrink.once.Do(func() {
			shrink.dict = full.dict
		})
	})
	return full.dict
}

// DictShrink returns a dictionary without content part.
// note. If an unshrinked dictionary already exists, this function returns it.
func DictShrink() *dict.Dict {
	shrink.once.Do(func() {
		shrink.dict = loadDict(false)
	})
	return shrink.dict
}

//go:embed uni.dict
var dictsrc embed.FS

func loadDict(full bool) *dict.Dict {
	b, err := fs.ReadFile(dictsrc, "uni.dict")
	if err != nil {
		panic(err)
	}
	r := bytes.NewReader(b)
	zr,err := zip.NewReader(r, r.Size())
	if err != nil {
		panic(err)
	}
	d, err := dict.Load(zr, full)
	if err != nil {
		panic(err)
	}
	return d
}
