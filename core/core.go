package core

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var reservedWordOrder []string
var variablesOrder []string
var ReservedWord map[string]string
var VariablesReplace map[string]string

var variablesIndex = 0

func removeReservedWordArray(lines []string) []string {
	for i := 0; i < len(lines); i++ {
		for index := 0; index < len(reservedWordOrder); index++ {
			reservedWord := reservedWordOrder[index]
			if len(reservedWord) > 0 {
				lines[i] = strings.ReplaceAll(lines[i]+" ", reservedWord, " ")
				lines[i] = strings.ReplaceAll(" "+lines[i], reservedWord, " ")
				lines[i] = strings.ReplaceAll(lines[i], reservedWord, " ")
			}
		}
	}

	return lines

}

func removeReservedWord(line string, rep map[string]string) string {

	for reservedWord, _ := range rep {
		if len(reservedWord) > 0 {
			line = strings.ReplaceAll(line, reservedWord, "")
		}
	}
	return line
}

func removeStrings(line string) string {
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

func FindVariablesReplace(lines []string) {
	//string_start := -1
	tempReservedWord := make(map[string]string)
	commentStat := 0

	for k, v := range ReservedWord {
		tempReservedWord[k] = ""
		tempReservedWord[v] = ""
	}

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		fmt.Println("第 ", i, "行：", line)

		line = removeStrings(line)

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

		line = removeReservedWord(line, tempReservedWord)
		linesLice := strings.Split(line, " ")

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
	genVariableReplaceOrder()

}

type codeString struct {
	isString bool
	content  string
}

func splitCodeLine(line string) []codeString {
	var codeString1 []codeString

	quoteN := strings.Count(line, "\"")
	//fmt.Println("quoteN",quoteN)

	if strings.Contains(line, "“") && strings.Contains(line, "”") {
		line1 := strings.Split(line, "“")
		var codeStringT codeString
		codeStringT.isString = false
		codeStringT.content = line1[0]
		codeString1 = append(codeString1, codeStringT)
		for i := 1; i < len(line1); i++ {
			line1[i] = line1[i]
			line2 := strings.Split(line1[i], "”")
			if len(line2) > 1 {
				var codeStringT codeString
				codeStringT.isString = true
				codeStringT.content = "\"" + line2[0] + "\""
				codeString1 = append(codeString1, codeStringT)
				var codeStringT1 codeString
				codeStringT1.isString = false
				codeStringT1.content = line2[1]
				codeString1 = append(codeString1, codeStringT1)

				if len(line2) > 2 {
					fmt.Println("seems wrong")
				}
			} else {
				fmt.Println("seems wrong")
			}
		}
	}

	if quoteN%2 == 0 && quoteN > 1 {
		line1 := strings.Split(line, "\"")
		var codeStringT codeString
		codeStringT.isString = false
		codeStringT.content = line1[0]
		codeString1 = append(codeString1, codeStringT)
		for i := 1; i < len(line1); i++ {
			line1[i] = line1[i]
			var codeStringT codeString
			if i%2 == 1 {
				codeStringT.isString = true
				codeStringT.content = "\"" + line1[i]
				codeString1 = append(codeString1, codeStringT)
			} else {
				codeStringT.isString = false
				codeStringT.content = "\"" + line1[i]
				if i != len(line1)-1 {
					// codeStringT.content += "\""
				}
				codeString1 = append(codeString1, codeStringT)
			}
		}
	}

	return codeString1
}

func replaceWithArray(line string, repOrder []string, rep map[string]string, force bool) string {
	splitO := splitCodeLine(line)
	if len(splitO) > 0 {
		//fmt.Println("------\nline",line)
		//fmt.Println("split out",split_o)
	}

	if !force && len(splitO) > 0 {
		line1 := ""
		for _, v := range splitO {
			if v.isString == false {
				//fmt.Println(k,v.content,"force",force)
				for _, key := range repOrder {
					v.content = strings.ReplaceAll(v.content, key, rep[key])
				}
			}
			line1 += v.content
		}
		line = line1
	} else {
		for _, key := range repOrder {
			line = strings.ReplaceAll(line, key, rep[key])
		}
	}

	//fmt.Println("line out",line)

	return line
}

var force bool

func ReplaceKeyWord(line string) string {

	if strings.Contains(line, "导入包") {
		force = true
	}
	if strings.Contains(line, ")") || strings.Contains(line, "）") {
		force = false
	}

	line = replaceWithArray(line, reservedWordOrder, ReservedWord, force)

	line = replaceWithArray(line, variablesOrder, VariablesReplace, force)
	return line
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

type stringArray []string

func (s stringArray) Len() int           { return len(s) }
func (s stringArray) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s stringArray) Less(i, j int) bool { return len(s[i]) > len(s[j]) }

func GenReservedWordOrder() {
	var wordOrder stringArray

	for k, _ := range ReservedWord {
		wordOrder = append(wordOrder, k)
	}
	sort.Stable(wordOrder)
	reservedWordOrder = wordOrder
}

func genVariableReplaceOrder() {
	var wordOrder stringArray

	for k, _ := range VariablesReplace {
		wordOrder = append(wordOrder, k)
	}
	sort.Stable(wordOrder)
	variablesOrder = wordOrder
}
