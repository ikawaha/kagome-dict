module tool

go 1.19

replace (
	github.com/ikawaha/kagome-dict => ./../..
	github.com/ikawaha/kagome-dict/ipa => ./../../ipa
)

require (
	github.com/ikawaha/kagome-dict v1.0.10
	github.com/ikawaha/kagome-dict/ipa v1.0.11
	golang.org/x/text v0.15.0
)
