package uni

import (
	"testing"

	"github.com/ikawaha/kagome-dict/dict"
)

const (
	UniDictEntrySize = 756466
	testDictPath     = "./uni.dict"
)

func Test_DictShrink(t *testing.T) {
	d := DictShrink()
	if want, got := UniDictEntrySize, len(d.Morphs); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
	if want, got := 0, len(d.Contents); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func Test_LoadShrink(t *testing.T) {
	d, err := dict.LoadShrink(testDictPath)
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if want, got := UniDictEntrySize, len(d.Morphs); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
	if want, got := 0, len(d.Contents); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func Test_New(t *testing.T) {
	d := Dict()
	if want, got := UniDictEntrySize, len(d.Morphs); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
	if want, got := UniDictEntrySize, len(d.Contents); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func Test_LoadDictFile(t *testing.T) {
	d, err := dict.LoadDictFile(testDictPath)
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if want, got := UniDictEntrySize, len(d.Morphs); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
	if want, got := UniDictEntrySize, len(d.Contents); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func Test_Singleton(t *testing.T) {
	a := Dict()
	b := Dict()
	if a != b {
		t.Errorf("got %p and %p, expected singleton", a, b)
	}
}

func Test_ContentsMeta(t *testing.T) {
	d := Dict()
	if want, got := 0, d.ContentsMeta[dict.POSStartIndex]; want != int(got) {
		t.Errorf("pos start index: want %d, got %d", want, got)
	}
	if want, got := 4, d.ContentsMeta[dict.POSHierarchy]; want != int(got) {
		t.Errorf("pos end index: want %d, got %d", want, got)
	}
	if want, got := 4, d.ContentsMeta[dict.InflectionalType]; want != int(got) {
		t.Errorf("base form index: want %d, got %d", want, got)
	}
	if want, got := 5, d.ContentsMeta[dict.InflectionalForm]; want != int(got) {
		t.Errorf("base form index: want %d, got %d", want, got)
	}
	if want, got := 0, d.ContentsMeta[dict.ReadingIndex]; want != int(got) { // undefined
		t.Errorf("reading index: want %d, got %d", want, got)
	}
	if want, got := 10, d.ContentsMeta[dict.BaseFormIndex]; want != int(got) {
		t.Errorf("base form index: want %d, got %d", want, got)
	}
	if want, got := 9, d.ContentsMeta[dict.PronunciationIndex]; want != int(got) {
		t.Errorf("pronunciation index: want %d, got %d", want, got)
	}
}

/*
func Test_InflectionalType(t *testing.T) {
	tnz, err := tokenizer.New(Dict())
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}

	tokens := tnz.Tokenize("寿司を食べたい")
	if want, got := 6, len(tokens); want != got {
		t.Fatalf("token length: want %d, got %d", want, got)
	}
	//BOS
	if got, ok := tokens[0].InflectionalType(); ok {
		t.Errorf("want !ok, got %q", got)
	}
	//食べ 動詞,一般,*,*,下一段-バ行,連用形-一般,タベル,食べる,食べ,タベ,食べる,タベル,和,*,*,*,*
	if got, ok := tokens[3].InflectionalType(); !ok {
		t.Error("want ok, but !ok")
	} else if want := "下一段-バ行"; want != got {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func Test_InflectionalForm(t *testing.T) {
	tnz, err := tokenizer.New(Dict())
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}

	tokens := tnz.Tokenize("寿司を食べたい")
	if want, got := 6, len(tokens); want != got {
		t.Fatalf("token length: want %d, got %d", want, got)
	}
	//BOS
	if got, ok := tokens[0].InflectionalForm(); ok {
		t.Errorf("want !ok, got %q", got)
	}
	//食べ 動詞,一般,*,*,下一段-バ行,連用形-一般,タベル,食べる,食べ,タベ,食べる,タベル,和,*,*,*,*
	if got, ok := tokens[3].InflectionalForm(); !ok {
		t.Error("want ok, but !ok")
	} else if want := "連用形-一般"; want != got {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func Test_BaseForm(t *testing.T) {
	tnz, err := tokenizer.New(Dict())
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}

	tokens := tnz.Tokenize("寿司を食べたい")
	if want, got := 6, len(tokens); want != got {
		t.Fatalf("token length: want %d, got %d", want, got)
	}
	//BOS
	if got, ok := tokens[0].BaseForm(); ok {
		t.Errorf("want !ok, got %q", got)
	}
	//食べ->食べる
	if got, ok := tokens[3].BaseForm(); !ok {
		t.Error("want ok, but !ok")
	} else if want := "食べる"; want != got {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func Test_Reading(t *testing.T) {
	tnz, err := tokenizer.New(Dict())
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}

	tokens := tnz.Tokenize("公園に行く")
	if want, got := 5, len(tokens); want != got {
		t.Fatalf("token length: want %d, got %d", want, got)
	}
	//BOS
	if got, ok := tokens[0].Reading(); ok {
		t.Errorf("want !ok, got %q", got)
	}
	//公園->コウエン // undefined
	if got, ok := tokens[1].Reading(); ok {
		t.Errorf("want !ok, but ok, got %s", got)
	}
}

func Test_Pronunciation(t *testing.T) {
	tnz, err := tokenizer.New(Dict())
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}

	tokens := tnz.Tokenize("公園に行く")
	if want, got := 5, len(tokens); want != got {
		t.Fatalf("token length: want %d, got %d", want, got)
	}
	//BOS
	if got, ok := tokens[0].Pronunciation(); ok {
		t.Errorf("want !ok, got %q", got)
	}
	//公園->コウエン
	if got, ok := tokens[1].Pronunciation(); !ok {
		t.Error("want ok, but !ok")
	} else if want := "コーエン"; want != got {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func Test_POS(t *testing.T) {
	tnz, err := tokenizer.New(Dict())
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}

	tokens := tnz.Tokenize("公園に行く")
	if want, got := 5, len(tokens); want != got {
		t.Fatalf("token length: want %d, got %d", want, got)
	}
	//BOS
	if got := tokens[0].POS(); len(got) > 0 {
		t.Errorf("want empty, got %+v", got)
	}
	//行く
	if want, got := []string{"動詞", "非自立可能", "*", "*"}, tokens[3].POS(); !reflect.DeepEqual(want, got) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}

func Test_FeatureIndex(t *testing.T) {
	testdata := []struct {
		index   FeatureIndex
		feature string
	}{
		{index: CType, feature: "下一段-バ行"},
		{index: CForm, feature: "連用形-一般"},
		{index: LForm, feature: "タベル"},
		{index: Lemma, feature: "食べる"},
		{index: Orth, feature: "食べ"},
		{index: Pron, feature: "タベ"},
		{index: OrthBase, feature: "食べる"},
		{index: PronBase, feature: "タベル"},
		{index: Goshu, feature: "和"},
		{index: IType, feature: "*"},
		{index: IForm, feature: "*"},
		{index: FType, feature: "*"},
		{index: FForm, feature: "*"},
		// aliases
		{index: InflectionalType, feature: "下一段-バ行"},
		{index: InflectionalForm, feature: "連用形-一般"},
		{index: BaseForm, feature: "食べる"},
		{index: Pronunciation, feature: "タベ"},
	}
	tnz, err := tokenizer.New(Dict())
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}

	tokens := tnz.Tokenize("寿司を食べたい")
	if want, got := 6, len(tokens); want != got {
		t.Fatalf("token length: want %d, got %d", want, got)
	}
	//食べ 動詞,一般,*,*,下一段-バ行,連用形-一般,タベル,食べる,食べ,タベ,食べる,タベル,和,*,*,*,*
	token := tokens[3]
	for _, v := range testdata {
		if got, ok := token.FeatureAt(v.index); !ok {
			t.Errorf("index %v: want ok, but !ok", v.index)
		} else if want := v.feature; want != got {
			t.Fatalf("index %v: want %s, got %s", v.index, want, got)
		}
	}
}
*/