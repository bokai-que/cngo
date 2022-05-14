package kernel

import (
	"fmt"
	"gitee.com/bokai-que/cngo/core"
	"sort"
	"strings"
)

var ReservedWordOrder []string
var variablesOrder []string

var force bool

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

func ReplaceKeyWord(line string) string {

	if strings.Contains(line, "导入包") {
		force = true
	}
	if strings.Contains(line, ")") || strings.Contains(line, "）") {
		force = false
	}

	line = replaceWithArray(line, ReservedWordOrder, core.ReservedWord, force)

	line = replaceWithArray(line, variablesOrder, core.VariablesReplace, force)
	return line
}

type stringArray []string

func (s stringArray) Len() int           { return len(s) }
func (s stringArray) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s stringArray) Less(i, j int) bool { return len(s[i]) > len(s[j]) }

func GenReservedWordOrder() {
	var wordOrder stringArray

	for k, _ := range core.ReservedWord {
		wordOrder = append(wordOrder, k)

	}
	sort.Stable(wordOrder)
	ReservedWordOrder = wordOrder
	//fmt.Printf("%q\n\n", reservedWordOrder)
}

func GenVariableReplaceOrder() {
	var wordOrder stringArray

	for k, _ := range core.VariablesReplace {
		wordOrder = append(wordOrder, k)
	}
	sort.Stable(wordOrder)
	variablesOrder = wordOrder
}
