package main

import (
	"bufio"
	"fmt"
	"gitee.com/bokai-que/cngo/core"
	"gitee.com/bokai-que/cngo/lang"
	"os"
	"strings"
)

func main() {
	filename := "./test.go"

	core.VariablesReplace = make(map[string]string)
	core.ReservedWord = make(map[string]string)
	lang.InitReservedWord()
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

func mainBak() {
	if len(os.Args) < 2 {
		fmt.Println("请把 你编写的 .cn.go 程序文件拖进来")
		return
	}

	filename := os.Args[1]

	fmt.Println("文件名：", filename)

	core.VariablesReplace = make(map[string]string)
	core.ReservedWord = make(map[string]string)
	lang.InitReservedWord()
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
