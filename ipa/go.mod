module github.com/ikawaha/kagome-dict/ipa

go 1.24.0

require github.com/ikawaha/kagome-dict v1.1.7

//replace github.com/ikawaha/kagome-dict => ../

retract v1.2.3 // Bug in dictionary lookup.
