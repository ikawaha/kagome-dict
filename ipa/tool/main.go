package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ikawaha/kagome-dict/dict"
	"github.com/ikawaha/kagome-dict/dict/builder"
	"golang.org/x/text/encoding/japanese"
)

const (
	CommandName     = "builder"
	ArchiveFileName = "ipa.dict"

	usageMessage = "%s -dict <dict_path> [-other <other_dict_path>] -out <output_file_name>\n"
)

var (
	RecordInfo = builder.MorphRecordInfo{
		ColSize:                 13,
		SurfaceIndex:            0,
		LeftIDIndex:             1,
		RightIDIndex:            2,
		WeightIndex:             3,
		POSStartIndex:           4,
		OtherContentsStartIndex: 8,
		// extra
		Meta: map[string]int8{
			dict.POSStartIndex:      0,
			dict.POSHierarchy:       4,
			dict.InflectionalType:   4,
			dict.InflectionalForm:   5,
			dict.BaseFormIndex:      6,
			dict.ReadingIndex:       7,
			dict.PronunciationIndex: 8,
		},
	}
	UnkRecordInfo = builder.UnkRecordInfo{
		ColSize:                 11,
		CategoryIndex:           0,
		LeftIDIndex:             1,
		RightIndex:              2,
		WeigthIndex:             3,
		POSStartIndex:           4,
		OtherContentsStartIndex: 8,
	}
	FileEncoding = japanese.EUCJP // set nil if utf8 (default)
)

type Paths []string

func (p Paths) String() string {
	return strings.Join(p, ",")
}

func (p *Paths) Set(v string) error {
	*p = append(*p, v)
	return nil
}

type Option struct {
	flagSet        *flag.FlagSet
	DictPath       string
	OtherPath      Paths
	OutputFileName string
}

// NewOption create an option.
// ContinueOnError ErrorHandling // Return a descriptive error.
// ExitOnError                   // Call os.Exit(2).
// PanicOnError                  // Call panic with a descriptive error.flag.ContinueOnError
func NewOption(w io.Writer, eh flag.ErrorHandling) *Option {
	o := &Option{
		flagSet: flag.NewFlagSet(CommandName, eh),
	}
	// option settings
	o.flagSet.SetOutput(w)
	o.flagSet.StringVar(&o.DictPath, "dict", "", "target dict path, ex. mecab-ipadic-2.7.0-20070801")
	o.flagSet.Var(&o.OtherPath, "other", "other dict path, neologd etc...")
	o.flagSet.StringVar(&o.OutputFileName, "out", ArchiveFileName, "output file name")
	return o
}

func (o *Option) Parse(args []string) error {
	if err := o.flagSet.Parse(args); err != nil {
		return err
	}
	// validations
	if nonFlag := o.flagSet.Args(); len(nonFlag) != 0 {
		return fmt.Errorf("invalid argument: %v", nonFlag)
	}
	if o.DictPath == "" {
		return fmt.Errorf("invalid argument: dict path is empty")
	}
	return nil
}

func Usage() {
	fmt.Fprintf(os.Stderr, usageMessage, CommandName)
}

func Run(args []string) error {
	opt := NewOption(os.Stderr, flag.ExitOnError)
	if err := opt.Parse(args); err != nil {
		Usage()
		opt.flagSet.PrintDefaults()
		return fmt.Errorf("%s, %v", CommandName, err)
	}
	config := builder.NewConfig(opt.DictPath, opt.OtherPath, FileEncoding, &RecordInfo, &UnkRecordInfo)
	dict, err := builder.Build(config)
	if err != nil {
		return fmt.Errorf("build failed: %v", err)
	}
	f, err := os.Create(opt.OutputFileName)
	if err != nil {
		return fmt.Errorf("create dict failed: %v", err)
	}
	defer f.Close()
	zw := zip.NewWriter(f)
	defer zw.Close()
	if err := dict.Save(zw); err != nil {
		return fmt.Errorf("save dict failed: %v", err)

	}
	return nil
}

func main() {
	if err := Run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
