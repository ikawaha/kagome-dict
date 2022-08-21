#!/usr/bin/env bash

src="./unidic-mecab-2.1.2_src"
data_dir=".."
dict="./${data_dir}/uni.dict"

go run main.go -dict ${src} -out ${dict}
