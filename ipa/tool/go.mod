module ipadicttool

go 1.24.1

replace (
	github.com/ikawaha/kagome-dict => ./../..
	github.com/ikawaha/kagome-dict/ipa => ./../../ipa
)

require (
	github.com/ikawaha/kagome-dict v1.1.2
	golang.org/x/text v0.23.0
)
