package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func test() {
	var src = []byte(`println("你好，世界")`)
	//fmt.Printf("%s\n", src)
	var fset = token.NewFileSet() // positions are relative to fset
	//fmt.Printf("%s\n", fset)
	var file = fset.AddFile("hello.go", fset.Base(), len(src)) // register input "file"
	//fmt.Printf("%s\n", file)
	var s scanner.Scanner
	s.Init(file, src, nil, scanner.ScanComments)
	fmt.Printf("%s\n", s)
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}

	//a := token.Position{Filename: "hello.go", Line: 1, Column: 2}
	//b := token.Position{Filename: "hello.go", Line: 1}
	//c := token.Position{Filename: "hello.go"}
	//
	//d := token.Position{Line: 1, Column: 2}
	//e := token.Position{Line: 1}
	//f := token.Position{Column: 2}
	//
	//fmt.Println(a.String())
	//fmt.Println(b.String())
	//fmt.Println(c.String())
	//fmt.Println(d.String())
	//fmt.Println(e.String())
	//fmt.Println(f.String())
}

func commandline() {
	// 中文版go语言：cngo
	// 用法：cngo file.cngo

	//判断输入是否有误

	//判断文件是否存在

	//打开文件交给语法和词法器匹配原GO代码

	//保存文件

}

func main() {
	commandline()
}
