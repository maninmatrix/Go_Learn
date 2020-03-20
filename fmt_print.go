package main
import (
	"fmt"
)
type Power struct{
	age int
	high int
	name string
}

func main()  {
	var name = "zhang"
	//Println 输出到控制台并换行
	fmt.Println(name)
	//Print 输出到控制台,不接受任何格式化操作
	fmt.Print(name)
	fmt.Println("没有换行")
	//Printf : 只可以打印出格式化的字符串。只可以直接输出字符串类型的变量（不可以输出别的类型）
	var i Power = Power{age: 10, high: 178, name: "NewMan"}
	fmt.Println(i)
	fmt.Println("----------Printf-------------")
	//%T    相应值的类型的Go语法表示 
	 fmt.Printf("相应值的类型的Go语法表示:%T\n", i)
	//%v     值的默认格式。
	fmt.Printf("值的默认格式:%v\n", i)
	//%+v   类似%v，但输出结构体时会添加字段名
	fmt.Printf("类似%%v，但输出结构体时会添加字段名:%+v\n", i)
	//%#v　 相应值的Go语法表示 
	fmt.Printf("相应值的Go语法表示 :%#v\n", i)
	//%%    百分号,字面上的%,非占位符含义
	fmt.Printf("百分号,字面上的:%%")
	fmt.Println()
	fmt.Println("----------整数类型 %b %c %d %o %q %x %X %U-------------")
	//%b     二进制表示 
	fmt.Printf("二进制表示:%b\n",123)
	//%c     相应Unicode码点所表示的字符 
	fmt.Printf("相应Unicode码点所表示的字符:%c\n",123)
	//%d     十进制表示 
	fmt.Printf("十进制表示:%d\n",123)
	//%o     八进制表示 
	fmt.Printf("八进制表示:%o\n",123)
	//%q     单引号围绕的字符字面值，由Go语法安全地转义 
	fmt.Printf("八进制表示:%q\n",123)
	//%x     十六进制表示，字母形式为小写 a-f 
	fmt.Printf("十六进制表示，字母形式为小写 a-f:%x\n",123)
	//%X     十六进制表示，字母形式为大写 A-F 
	fmt.Printf("十六进制表示，字母形式为大写 A-F :%X\n",123)
	//%U     Unicode格式：123，等同于 "U+007B"
	fmt.Printf("Unicode格式 :%U\n",123)
	fmt.Println("----------浮点数 %e %E %f %g %G-------------")
	//%e     科学计数法，例如 -1234.456e+78 
	fmt.Printf("科学计数法:%e\n",12675757563.5345432567)
	//%E     科学计数法，例如 -1234.456E+78 
	fmt.Printf("科学计数法:%E\n",12675757563.5345432567)
	//%f     有小数点而无指数，例如 123.456 
	fmt.Printf("有小数点而无指数:%f\n",12675757563.5345432567)
	//%g     根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出 
	fmt.Printf("根据情况选择 %%e 或 %%f:%g\n",12675757563.5345432567)
	//%G     根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出
	fmt.Printf("根据情况选择 %%e 或 %%f:%G\n",12675757563.5345432567)
	fmt.Println("----------字符串-------------")
	//%s     字符串或切片的无解译字节 
	fmt.Printf("字符串或切片的无解译字节:%s\n","I'm a girl")
	//%q     双引号围绕的字符串，由Go语法安全地转义 
	fmt.Printf("双引号围绕的字符串:%q\n","I'm a girl")
	//%x     十六进制，小写字母，每字节两个字符 
	fmt.Printf("十六进制，小写字母:%x\n","I'm a girl")
	//%X     十六进制，大写字母，每字节两个字符
	fmt.Printf("十六进制，大写字母:%X\n","I'm a girl")
	fmt.Println("----------指针-------------")
	//%p     十六进制表示，前缀 0x
	a := 1
	b := &a
	fmt.Printf("a的地址：%p",b)

}