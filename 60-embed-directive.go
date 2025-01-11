package main

import (
    "embed"
)
//go:embed 是一个编译指令，它允许程序在构建时将任意文件和文件夹包含到 Go 二进制文件中
//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

    print(fileString)
    print(string(fileByte))

    content1, _ := folder.ReadFile("folder/file1.hash")
    print(string(content1))

    content2, _ := folder.ReadFile("folder/file2.hash")
    print(string(content2))
}