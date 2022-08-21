package ipa

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
	// 公園	名詞,一般,*,*,*,*,公園,コウエン,コーエン
	// に	助詞,格助詞,一般,*,*,*,に,ニ,ニ
	// 行っ	動詞,自立,*,*,五段・カ行促音便,連用タ接続,行く,イッ,イッ
	// た	助動詞,*,*,*,特殊・タ,基本形,た,タ,タ
	// EOS

	// POSHierarchy represents part-of-speech hierarchy
	// e.g. Columns 動詞,自立,*,* are POSs which hierarchy depth is 4.
	POSHierarchy = 4
	// InflectionalType represents 活用型 (e.g. 五段・カ行促音便)
	InflectionalType FeatureIndex = 4
	// InflectionalForm represents 活用形 (e.g. 連用タ接続)
	InflectionalForm = 5
	// BaseForm represents 基本形 (e.g. 行く)
	BaseForm = 6
	// Reading represents 読み (e.g. コウエン)
	Reading = 7
	// Pronunciation represents 発音 (e.g. コーエン)
	Pronunciation = 8
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

//go:embed ipa.dict
var dictsrc embed.FS

func loadDict(full bool) *dict.Dict {
	b, err := fs.ReadFile(dictsrc, "ipa.dict")
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
