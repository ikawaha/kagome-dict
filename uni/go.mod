module github.com/ikawaha/kagome-dict/uni

go 1.24.1

require github.com/ikawaha/kagome-dict v1.1.5

//replace github.com/ikawaha/kagome-dict => ../

retract v1.1.2 // Bug in dictionary lookup.
