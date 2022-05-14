package core

import (
	"bufio"
	"fmt"
	"gitee.com/bokai-que/cngo/kernel"
	"os"
	"strings"
)

var ReservedWord map[string]string
var VariablesReplace map[string]string

var variablesIndex = 0

func removeReservedWordArray(lines []string) []string {
	for i := 0; i < len(lines); i++ {
		for index := 0; index < len(kernel.ReservedWordOrder); index++ {
			reservedWord := kernel.ReservedWordOrder[index]
			if len(reservedWord) > 0 {
				lines[i] = strings.ReplaceAll(lines[i]+" ", reservedWord, " ")
				lines[i] = strings.ReplaceAll(" "+lines[i], reservedWord, " ")
				lines[i] = strings.ReplaceAll(lines[i], reservedWord, " ")
			}
		}
	}

	return lines

}

func 过滤关键字(line string, rep map[string]string) string {

	for reservedWord, _ := range rep {
		if len(reservedWord) > 0 {
			line = strings.ReplaceAll(line, reservedWord, "")
		}
	}
	return line
}

func 过滤字符串(line string) string {
	stringStart := -1
	stringEnd := -1
	for strings.Contains(line, "“") {
		stringStart = strings.Index(line, "“")
		if stringStart > 0 {
			stringEnd = strings.Index(line, "”")
			//lineslice := strings.Split(line,"“")
			//lineslice2 := strings.Split(line,"“")
			//remove string
			if stringEnd > 0 {
				line = line[0:stringStart] + line[stringEnd+1:len(line)]
			} else {
				break
			}
		}
	}

	for strings.Contains(line, "\"") {
		stringStart = strings.Index(line, "\"")
		if stringStart > 0 {

			stringEnd = strings.Index(line[stringStart+1:], "\"") + stringStart + 1
			if stringEnd > 0 {
				line = line[0:stringStart] + line[stringEnd+1:len(line)]
			} else {
				break
			}
		}
	}
	return line
}

func 分割字符串(line, sep string) {

	linesLice := strings.Split(line, sep)

	for j := 0; j < len(linesLice); j++ {
		variableName := linesLice[j]
		variableName = strings.ReplaceAll(variableName, "�", "")
		variableName = strings.ReplaceAll(variableName, " ", "")
		variableName = strings.ReplaceAll(variableName, "\t", "")
		if len(variableName) > 0 {
			_, isset := VariablesReplace[variableName]
			if isset {
				continue
			}
			_, isset2 := ReservedWord[variableName]
			if isset2 {
				continue
			}
			variablesIndexS := fmt.Sprintf("%d", variablesIndex)
			VariablesReplace[variableName] = "pan_" + variablesIndexS + "_"
			variablesIndex++
			fmt.Println(" 变量名:[", variableName, "] ,配置为 :", "pan_"+variablesIndexS)
			//fmt.Println(" find variable name:",variableName," ,config to :","variable_"+variables_index_s)
		}
	}
}

func FindVariablesReplace(lines []string) {
	//string_start := -1
	tempReservedWord := make(map[string]string)
	commentStat := 0

	for k, v := range ReservedWord {
		tempReservedWord[k] = ""
		tempReservedWord[v] = ""
	}
	//fmt.Printf("%q\n", tempReservedWord)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		fmt.Println("第 ", i, "行：", line)

		line = 过滤字符串(line)

		if strings.Contains(line, "注释开始") {
			fmt.Println("注释开始")
			commentStat = 1
			continue
		}

		if strings.Contains(line, "注释结束") && commentStat == 1 {
			fmt.Println("注释结束")
			commentStat = 0
		} else if strings.Contains(line, "注释") {
			continue
		}

		if commentStat == 1 {
			continue
		}
		line = 过滤关键字(line, tempReservedWord)
		分割字符串(line, " ")

	}
	kernel.GenVariableReplaceOrder()

}

func WriteFile(filePath string, content string) {

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err", err)
		return
	}

	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)

	write.WriteString(content)

	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}
