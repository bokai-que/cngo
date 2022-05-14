package main

import (
	"bufio"
	"fmt"
	"gitee.com/bokai-que/cngo/core"
	"gitee.com/bokai-que/cngo/lang"
	"os"
	"strings"
)

func openFile() *os.File {
	//if len(os.Args) < 2 {
	//	fmt.Println("请把 你编写的 .cn.go 程序文件拖进来")
	//	return nil
	//}
	//filename := os.Args[1]
	filename := "./test.go"
	fmt.Println("文件名：", filename)

	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	return file
}

func main() {
	file := openFile()
	if file == nil {
		return
	}
	defer file.Close()
	core.VariablesReplace = make(map[string]string)
	core.ReservedWord = make(map[string]string)
	lang.ZhCN()
	//fmt.Printf("%q\n", core.ReservedWord)
	core.GenReservedWordOrder()
	reader := bufio.NewReader(file)
	fileOutContent := ""
	var lines []string
	for {
		line, err := reader.ReadString('\n')
		line1 := string(line)
		line1 = strings.ReplaceAll(line1, "\r\n", "")
		line1 = strings.ReplaceAll(line1, "\n", "")
		lines = append(lines, line1)
		//fmt.Println(string(line))
		if err != nil {
			break
		}
	}
	//fmt.Printf("%q\n", lines)
	core.FindVariablesReplace(lines)

	for i := 0; i < len(lines); i++ {
		//fmt.Printf("%q\n", lines[i])
		line2 := core.ReplaceKeyWord(lines[i])
		//fmt.Printf("%q\n", line2)
		fileOutContent += line2 + "\r\n"
	}
	//fmt.Printf("%q", fileOutContent)
}

func mainBak() {
	if len(os.Args) < 2 {
		fmt.Println("请把 你编写的 .cn.go 程序文件拖进来")
		return
	}

	filename := os.Args[1]

	fmt.Println("文件名：", filename)

	core.VariablesReplace = make(map[string]string)
	core.ReservedWord = make(map[string]string)
	lang.ZhCN()
	core.GenReservedWordOrder()

	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	//scanner := bufio.NewScanner(file)
	fileOutContent := ""
	var lines []string
	//for scanner.Scan() {
	for {
		//line, prefix, err := reader.ReadLine();
		//line := scanner.Text()
		line, err := reader.ReadString('\n')
		line1 := string(line)
		line1 = strings.ReplaceAll(line1, "\r\n", "")
		line1 = strings.ReplaceAll(line1, "\n", "")
		lines = append(lines, line1)
		//fmt.Println(string(line))
		if err != nil {
			break
		}
	}

	core.FindVariablesReplace(lines)

	for i := 0; i < len(lines); i++ {
		line2 := core.ReplaceKeyWord(lines[i])
		fileOutContent += line2 + "\r\n"
	}

	//path := strings.ReplaceAll(filename,".cn.go",".go")
	path := filename + ".run.go"
	core.WriteFile(path, fileOutContent)
}
