package lang

import "gitee.com/bokai-que/cngo/core"

//修改
func setReservedWord(in string, out string) {
	//reservedWord[in] = out
	core.ReservedWord[out] = in
}

func InitReservedWord() {

	//关键字
	//setReservedWord("跳出", "break")
	setReservedWord("跳出循环", "break")
	setReservedWord("当", "case")
	setReservedWord("当它为", "case")
	//setReservedWord("案例", "case")
	//setReservedWord("通道", "chan")
	setReservedWord("队列", "chan")
	//setReservedWord("常量", "const")
	setReservedWord("常数", "const")
	//setReservedWord("跳过", "continue")
	setReservedWord("继续", "continue")
	setReservedWord("默认", "default")
	//setReservedWord("延缓", "defer")
	setReservedWord("推迟执行", "defer")
	//setReservedWord("否则", "else")
	setReservedWord("此外", "else")
	//setReservedWord("继续", "fallthrough")
	setReservedWord("往下执行", "fallthrough")
	setReservedWord("循环", "for")
	setReservedWord("启动循环：", "for")
	setReservedWord("函数", "func")
	setReservedWord("功能", "func")
	setReservedWord("并发", "go")
	setReservedWord("运行线程", "go")
	setReservedWord("跳到", "goto")
	setReservedWord("若", "if")
	setReservedWord("如果", "if")
	//setReservedWord("导入", "import")
	setReservedWord("导入包", "import")
	setReservedWord("接口", "interface")
	setReservedWord("字典", "map")
	//setReservedWord("映射", "map")
	setReservedWord("包名", "package")
	setReservedWord("生成范围", "range")
	//setReservedWord("范围", "range")
	setReservedWord("返回", "return")
	setReservedWord("选择", "select")
	setReservedWord("结构体", "struct")
	//setReservedWord("结构", "struct")
	//setReservedWord("转变", "switch")
	setReservedWord("选择执行", "switch")
	setReservedWord("定义", "type")
	setReservedWord("类型", "type")
	setReservedWord("变量", "var")

	//内置类型
	//setReservedWord("整数", "int")
	//setReservedWord("整数8", "int8")
	//setReservedWord("整数16", "int16")
	//setReservedWord("整数32", "int32")
	//setReservedWord("整数64", "int64")
	setReservedWord("长整数", "int64")
	//setReservedWord("无符", "uint")
	//setReservedWord("无符8", "uint8")
	//setReservedWord("无符16", "uint16")
	//setReservedWord("无符32", "uint32")
	//setReservedWord("无符64", "uint64")
	//setReservedWord("字串", "string")
	//setReservedWord("浮点数", "float64")
	//setReservedWord("浮点32", "float32")
	//setReservedWord("浮点64", "float64")
	//setReservedWord("复数64", "complex64")
	//setReservedWord("复数128", "complex128")
	setReservedWord("布尔", "bool")
	setReservedWord("字符串", "string")
	setReservedWord("指针", "uintptr")

	//内置函数
	//setReservedWord("附加", "append")
	setReservedWord("添加", "append")
	setReservedWord("容量", "cap")
	setReservedWord("关闭", "close")
	setReservedWord("关闭文件", "Close")
	setReservedWord("复数", "complex")
	setReservedWord("复制", "copy")
	setReservedWord("删除", "delete")
	setReservedWord("虚部", "imag")
	setReservedWord("长度", "len")
	setReservedWord("构建", "make")
	setReservedWord("生成", "make")
	setReservedWord("新建", "new")
	setReservedWord("报警", "panic")
	setReservedWord("打印", "print")
	setReservedWord("换行打印", "println")
	setReservedWord("实部", "real")
	setReservedWord("复原", "recover")
	setReservedWord("取对齐", "Alignof")
	setReservedWord("取偏移", "Offsetof")
	setReservedWord("取大小", "Sizeof")
	setReservedWord("真", "true")
	setReservedWord("假", "false")
	setReservedWord("空引用", "nil")
	//setReservedWord("空", "nil")
	//setReservedWord("|", "iota")
	setReservedWord("错误", "error")
	setReservedWord("字节", "byte")
	setReservedWord("字符", "rune")

	//其他
	setReservedWord("主程序", "main")
	setReservedWord("等待用户输入", "input")
	setReservedWord("空", "_")
	setReservedWord("运行库", "runtime")
	setReservedWord("格式", "fmt")
	setReservedWord("打印", "Println")
	setReservedWord("系统", "os")
	setReservedWord("打开", "Open")
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
	setReservedWord("显示", "Println")
	setReservedWord("语句", "string")
	setReservedWord("显示数", "%d")
	setReservedWord("显示字符串", "%s")
	setReservedWord("显示浮点", "%f")
	setReservedWord("数组", "[]")
	setReservedWord("等待队列执行完毕", "<-")
	//setReservedWord("的类型是", "")
	setReservedWord("的", ".")
	setReservedWord("模块调用", ".")
	setReservedWord("。", ".")
	setReservedWord("，", ",")
	setReservedWord("换行符", "'\\n'")
	setReservedWord("“", "\"")
	setReservedWord("”", "\"")
	setReservedWord("（", "(")
	setReservedWord("）", ")")
	setReservedWord("『", "{")
	setReservedWord("』", "}")
	setReservedWord("；", ";")
	setReservedWord("加", "+")
	setReservedWord("减", "-")
	setReservedWord("乘", "*")
	setReservedWord("除", "/")
	setReservedWord("》", ">")
	setReservedWord("大于", ">")
	setReservedWord("《", "<")
	setReservedWord("小于", "<")
	setReservedWord("右移", ">>")
	setReservedWord("左移", "<<")
	setReservedWord("注释", "//")
	setReservedWord("注释开始", "/*")
	setReservedWord("注释结束", "*/")
	setReservedWord("使用索引：", "[")
	setReservedWord("结束索引", "]")
	setReservedWord("初始化", ":=")
	setReservedWord("初始化为", ":=")
	setReservedWord("为", "=")
	setReservedWord("设置", "=")
	setReservedWord("等于", "=")
	setReservedWord("引用", "*")
	setReservedWord("取引用", "&")
	setReservedWord("与运算", "&")
	setReservedWord("或运算", "|")
	setReservedWord("或者", "||")
	setReservedWord("与", "|")
	setReservedWord("不相等", "!=")
	setReservedWord("相等于", "==")
	setReservedWord("相等", "==")
	setReservedWord("不", "!")
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
