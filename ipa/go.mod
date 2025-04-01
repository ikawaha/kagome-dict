module github.com/ikawaha/kagome-dict/ipa

go 1.23.0

require github.com/ikawaha/kagome-dict v1.1.6

//replace github.com/ikawaha/kagome-dict => ../

retract v1.2.3 // Bug in dictionary lookup.
