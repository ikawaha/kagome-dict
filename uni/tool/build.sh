#!/usr/bin/env bash

src="./unidic-mecab-2.1.2_src"
intermidiate_dir="./dict"
intermidiate_prefix="unidict."
testdata_dir="../testdata"
dict="${testdata_dir}/uni.dict"
dst_pkg=data
dst_dir="../internal/${dst_pkg}"

rm -f ${dst_dir}/*.go ${dict}
mkdir -p ${testdata_dir}
go run main.go -dict ${src} -out ${dict}
rm -rf ${intermidiate_dir}
mkdir -p ${intermidiate_dir}
split -a 2 -b 512k ${dict} ${intermidiate_dir}/${intermidiate_prefix}
rm -rf ${dst_dir}
mkdir -p ${dst_dir}
go-bindata -o ${dst_dir}/bindata.go -nocompress -separate -pkg=${dst_pkg} ${intermidiate_dir}
rm -rf ${intermidiate_dir}