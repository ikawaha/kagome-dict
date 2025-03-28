module tool

go 1.23.0

toolchain go1.24.1

replace (
	github.com/ikawaha/kagome-dict => ./../..
	github.com/ikawaha/kagome-dict/ipa => ./../../ipa
)

require (
	github.com/ikawaha/kagome-dict v1.0.10
	github.com/ikawaha/kagome-dict/ipa v1.0.11
	golang.org/x/text v0.23.0
)
