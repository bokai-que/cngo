package core

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var ReservedWord map[string]string
var VariablesReplace map[string]string

var variablesIndex = 0

func removeReservedWordArray(lines []string) []string {
	for i := 0; i < len(lines); i++ {
		for index := 0; index < len(ReservedWordOrder); index++ {
			reservedWord := ReservedWordOrder[index]
			if len(reservedWord) > 0 {
				lines[i] = strings.ReplaceAll(lines[i]+" ", reservedWord, " ")
				lines[i] = strings.ReplaceAll(" "+lines[i], reservedWord, " ")
				lines[i] = strings.ReplaceAll(lines[i], reservedWord, " ")
			}
		}
	}

	return lines

}

/********************************************************************************
字符串全部替换
s string:要替换的整个字符串。
映射 map[string]string:要替换的字符串。
new string:替换成什么字符串。
返回值：返回替换后的字符串。
说明：将字符串 s 中的 old 字符串全部替换成 new 字符串，返回替换后的字符串。
********************************************************************************/
func 字符串全部替换(s string, 映射 map[string]string, new string) string {

	for old, _ := range 映射 {
		if len(old) > 0 {
			s = strings.ReplaceAll(s, old, new)
			//strings.ReplaceAll
			//s	要替换的整个字符串。
			//old	要替换的字符串。
			//new	替换成什么字符串。
			//返回值	返回替换后的字符串。
			//说明	将字符串 s 中的 old 字符串全部替换成 new 字符串，返回替换后的字符串。
		}
	}
	return s
}

func 过滤字符串(line string) string {
	stringStart := -1
	stringEnd := -1
	for strings.Contains(line, "“") {
		stringStart = strings.Index(line, "“")
		//s	原字符串。
		//substr	要检索的字符串。
		//返回值		Index() 函数返回 int 类型的值，如果包含，则返回第一次出现该字符串的索引；反之，则返回 -1。
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
		line = 字符串全部替换(line, tempReservedWord, "")
		分割字符串(line, " ")

	}
	GenVariableReplaceOrder()

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

/**************************************************************************************************************/

var ReservedWordOrder []string
var variablesOrder []string

var force bool

type codeString struct {
	isString bool
	content  string
}

//分割代码行
func splitCodeLine(line string) []codeString {
	var codeString1 []codeString

	//strings.Count 返回字符串 line 中有几个不重复的 "\""。
	quoteN := strings.Count(line, "\"")

	//strings.Contains 判断字符串 line 是否包含子串 "“" "“"
	if strings.Contains(line, "“") && strings.Contains(line, "”") {
		//strings.Split 用去掉 line 中出现的 sep 的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片
		//（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。
		//如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。
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

// ReplaceKeyWord 替换关键字
func ReplaceKeyWord(line string) string {

	//strings.Contains判断字符串line是否包含子串 substr:"导入包"
	if strings.Contains(line, "导入包") || strings.Contains(line, "import") {
		force = true
	}
	if strings.Contains(line, ")") || strings.Contains(line, "）") {
		force = false
	}

	line = replaceWithArray(line, ReservedWordOrder, ReservedWord, force)
	line = replaceWithArray(line, variablesOrder, VariablesReplace, force)
	return line
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
	ReservedWordOrder = wordOrder
	//fmt.Printf("%q\n\n", reservedWordOrder)
}

func GenVariableReplaceOrder() {
	var wordOrder stringArray

	for k, _ := range VariablesReplace {
		wordOrder = append(wordOrder, k)
	}
	sort.Stable(wordOrder)
	variablesOrder = wordOrder
}
