#!bin/bash

if [ $# -ne 1 ]; then
	echo "作成するディレクトリ名を引数に入れてください"
	exit 1
fi

dic=$1
mkdir $dic
cp ./template.go $dic/main.go
