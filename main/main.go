package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var variablesReplace map[string]string
var reservedWordOrder []string
var variablesOrder []string
var reservedWord map[string]string

func main() {
	if len(os.Args) < 2 {
		fmt.Println("请把 你编写的 .pan.go 程序文件拖进来")
		return
	}

	filename := os.Args[1]

	fmt.Println("文件名：", filename)

	variablesReplace = make(map[string]string)
	reservedWord = make(map[string]string)
	initReservedWord()
	genReservedWordOrder()

	file, err := os.Open(filename)
	if err != nil {
		return
	}
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

	findVariablesReplace(lines)

	for i := 0; i < len(lines); i++ {
		line2 := replaceKeyWord(lines[i])
		fileOutContent += line2 + "\r\n"
	}

	//path := strings.ReplaceAll(filename,".pan.go",".go")
	path := filename + ".run.go"
	WriteFile(path, fileOutContent)
}

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

func findVariablesReplace(lines []string) {
	//string_start := -1
	tempReservedWord := make(map[string]string)
	commentStat := 0

	for k, v := range reservedWord {
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
				_, isset := variablesReplace[variableName]
				if isset {
					continue
				}
				_, isset2 := reservedWord[variableName]
				if isset2 {
					continue
				}
				variablesIndexS := fmt.Sprintf("%d", variablesIndex)
				variablesReplace[variableName] = "pan_" + variablesIndexS + "_"
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

func replaceKeyWord(line string) string {

	if strings.Contains(line, "导入包") {
		force = true
	}
	if strings.Contains(line, ")") || strings.Contains(line, "）") {
		force = false
	}

	line = replaceWithArray(line, reservedWordOrder, reservedWord, force)

	line = replaceWithArray(line, variablesOrder, variablesReplace, force)
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

func setReservedWord(in string, out string) {
	reservedWord[in] = out

}

func initReservedWord() {
	setReservedWord("变量", "var")
	setReservedWord("包名", "package")
	setReservedWord("主程序", "main")
	setReservedWord("导入包", "import")
	setReservedWord("字典", "map")
	setReservedWord("数组", "[]")
	setReservedWord("生成", "make")
	setReservedWord("循环", "for")
	setReservedWord("启动循环：", "for")
	setReservedWord("若", "if")
	setReservedWord("如果", "if")
	setReservedWord("函数", "func")
	setReservedWord("功能", "func")
	setReservedWord("推迟执行", "defer")
	setReservedWord("跳出循环", "break")
	setReservedWord("默认", "default")
	setReservedWord("选择", "select")
	setReservedWord("此外", "else")
	setReservedWord("常数", "const")
	setReservedWord("往下执行", "fallthrough")
	setReservedWord("继续", "continue")
	setReservedWord("返回", "return")
	setReservedWord("选择执行", "switch")
	setReservedWord("当", "case")
	setReservedWord("当它为", "case")
	setReservedWord("等待用户输入", "input")
	setReservedWord("空", "_")
	setReservedWord("空引用", "nil")

	setReservedWord("等待队列执行完毕", "<-")
	setReservedWord("队列", "chan")
	setReservedWord("运行线程", "go")
	setReservedWord("长度", "len")
	setReservedWord("添加", "append")
	setReservedWord("关闭", "close")

	setReservedWord("生成范围", "range")
	//package
	setReservedWord("运行库", "runtime")
	setReservedWord("格式", "fmt")
	setReservedWord("打印", "Println")
	setReservedWord("系统", "os")
	setReservedWord("打开", "Open")
	setReservedWord("关闭文件", "Close")
	setReservedWord("传入参数", "Args")
	setReservedWord("工具集", "utils")
	setReservedWord("初始化函数", "Initial")
	setReservedWord("时间", "time")
	setReservedWord("转换格式", "Format")
	setReservedWord("时间格式", "utils.TIME_LAYOUT")
	setReservedWord("此刻", "Now()")
	setReservedWord("睡眠", "Sleep")
	setReservedWord("时间长度", "time.Duration")
	setReservedWord("1秒时间", "time.Second")
	setReservedWord("缓存", "bufio")
	setReservedWord("阅读器", "reader")
	setReservedWord("新建阅读器", "NewReader")
	setReservedWord("读字符串直到", "ReadString")
	setReservedWord("换行符", "'\\n'")

	setReservedWord("显示数", "%d")
	setReservedWord("显示字符串", "%s")
	setReservedWord("显示浮点", "%f")

	setReservedWord("显示", "Println")
	setReservedWord("生成", "make")

	setReservedWord("返回", "return")
	setReservedWord("定义", "type")
	setReservedWord("结构体", "struct")
	setReservedWord("定义", "type")
	setReservedWord("接口", "interface")
	setReservedWord("的类型是", "")

	setReservedWord("为", "=")
	setReservedWord("设置", "=")
	setReservedWord("等于", "=")
	setReservedWord("取引用", "&")
	setReservedWord("引用", "*")
	setReservedWord("与运算", "&")
	setReservedWord("或运算", "|")

	setReservedWord("或者", "||")
	setReservedWord("与", "|")

	setReservedWord("不相等", "!=")
	setReservedWord("相等于", "==")
	setReservedWord("相等", "==")
	setReservedWord("不", "!")
	setReservedWord("的", ".")
	//setReservedWord("。",".")
	setReservedWord("，", ",")
	setReservedWord("模块调用", ".")

	setReservedWord("注释", "//")
	setReservedWord("注释开始", "/*")
	setReservedWord("注释结束", "*/")
	setReservedWord("使用索引：", "[")
	setReservedWord("结束索引", "]")

	setReservedWord("初始化", ":=")
	setReservedWord("初始化为", ":=")
	setReservedWord("删除", "delete")
	setReservedWord("字节", "byte")
	setReservedWord("语句", "string")
	setReservedWord("字符串", "string")
	setReservedWord("整数", "int")
	setReservedWord("长整数", "int64")
	setReservedWord("浮点数", "float64")
	setReservedWord("布尔", "bool")
	setReservedWord("“", "\"")
	setReservedWord("”", "\"")
	setReservedWord("右移", ">>")
	setReservedWord("左移", "<<")
	setReservedWord("》", ">")
	setReservedWord("《", "<")
	setReservedWord("大于", ">")
	setReservedWord("小于", "<")
	setReservedWord("（", "(")
	setReservedWord("（", "(")
	setReservedWord("）", ")")
	setReservedWord("『", "{")
	setReservedWord("』", "}")
	setReservedWord("；", ";")
	setReservedWord("加", "+")
	setReservedWord("减", "-")
	setReservedWord("乘", "*")
	setReservedWord("除", "/")

	setReservedWord("一", "1")
	setReservedWord("二", "2")
	setReservedWord("三", "3")
	setReservedWord("四", "4")
	setReservedWord("五", "5")
	setReservedWord("六", "6")
	setReservedWord("七", "7")
	setReservedWord("八", "8")
	setReservedWord("九", "9")
	setReservedWord("零", "0")

}

type stringArray []string

func (s stringArray) Len() int           { return len(s) }
func (s stringArray) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s stringArray) Less(i, j int) bool { return len(s[i]) > len(s[j]) }

func genReservedWordOrder() {
	var wordOrder stringArray

	for k, _ := range reservedWord {
		wordOrder = append(wordOrder, k)
	}
	sort.Stable(wordOrder)
	reservedWordOrder = wordOrder
}

func genVariableReplaceOrder() {
	var wordOrder stringArray

	for k, _ := range variablesReplace {
		wordOrder = append(wordOrder, k)
	}
	sort.Stable(wordOrder)
	variablesOrder = wordOrder
}
