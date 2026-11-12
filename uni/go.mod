module github.com/ikawaha/kagome-dict/uni

go 1.26.0

require github.com/ikawaha/kagome-dict v1.1.8

//replace github.com/ikawaha/kagome-dict => ../

retract v1.2.3 // Bug in dictionary lookup.
