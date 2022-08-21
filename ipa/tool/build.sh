#!/usr/bin/env bash

src="./mecab-ipadic-2.7.0-20070801"
data_dir=".."
dict="./${data_dir}/ipa.dict"

go run main.go -dict ${src} -out ${dict}
