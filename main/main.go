package main

import (
	"bufio"
	"fmt"
	"go/scanner"
	"go/token"
	"io"
	"io/ioutil"
	"os"
)

func test() {
	//var src = []byte(`println("你好，世界")`)

	src, _ := ioutil.ReadFile("./test.go") //读取文件

	//fmt.Printf("%s\n", src)
	var fset = token.NewFileSet() // positions are relative to fset
	//fmt.Printf("%s\n", fset)
	var file = fset.AddFile("hello.go", fset.Base(), len(src)) // register input "file"
	//fmt.Printf("%s\n", file)
	var s scanner.Scanner
	s.Init(file, src, nil, scanner.ScanComments)
	pos, tok, lit := s.Scan()
	fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	if tok == tokcn {
		lit = ""
	}
	fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
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

func 读文件行() {
	//打开文件
	file, err := os.Open("./test.go") //只是用来读的时候，用os.Open。相对路径，针对于同目录下。
	if err != nil {
		fmt.Printf("打开文件失败,err:%v\n", err)
		return
	}
	defer file.Close() //关闭文件,为了避免文件泄露和忘记写关闭文件

	//使用buffio读取文件内容
	reader := bufio.NewReader(file) //创建新的读的对象
	for {
		line, err := reader.ReadString('\n') //注意是字符，换行符。
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil { //错误处理
			fmt.Printf("读取文件失败,错误为:%v", err)
			return
		}
		fmt.Println(line)
	}
}

func 读文件() {
	iot, err := ioutil.ReadFile("./test.go") //读取文件
	if err != nil {                          //做错误判断
		fmt.Printf("读取文件错误,错误为:%v\n", err)
		return
	}
	fmt.Println(string(iot)) //打印文件内容
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
	test()
}
