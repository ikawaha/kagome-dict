module unidicttool

go 1.24.1

replace (
	github.com/ikawaha/kagome-dict => ./../..
	github.com/ikawaha/kagome-dict/uni => ./..
)

require (
	github.com/ikawaha/kagome-dict v1.1.2
	github.com/ikawaha/kagome-dict/uni v1.0.10
	golang.org/x/text v0.23.0
)
