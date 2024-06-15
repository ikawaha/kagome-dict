module tool

go 1.19

replace (
    github.com/ikawaha/kagome-dict => ./../..
    github.com/ikawaha/kagome-dict/uni => ./..
)

require (
	github.com/ikawaha/kagome-dict v1.0.10
	github.com/ikawaha/kagome-dict/uni v1.0.10
	golang.org/x/text v0.15.0
)
