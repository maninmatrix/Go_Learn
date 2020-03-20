package main
import (
	"fmt"
)

func main()  {
	//iota只能在常量的表达式中使用 常用来做枚举表示
	const a = iota
	fmt.Println("a value:", a)
	//每次 const 出现时，都会让 iota 初始化为0.
	const ( 
		b = iota     //b=0 
		c            //c=1   相当于c=iota
	  )
	  fmt.Printf("a value:%d, b value: %d\n", b, c)
	  //如果两个const的赋值语句的表达式是一样的，那么可以省略后一个赋值表达式
	const ( // iota被重设为0
		c0 = iota // c0 == 0
		c1 // c1 == 1
		c2 // c2 == 2
	)
	//我们可以使用下划线跳过不想要的值。
	const ( 
		OutMute  = iota // 0 
		OutMono                    // 1 
		OutStereo                  // 2 
		_ 
		_ 
		OutSurround                // 5 
	)
	//中间插队
	const ( 
    	i = iota 
    	j = 3.14 
    	k = iota 
	)
	fmt.Printf("i value:%d, k value: %d\n", i, k)
}