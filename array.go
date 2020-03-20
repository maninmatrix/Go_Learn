package main
import (
	"fmt"
)

func main()  {
	//数组声明方法1 var arr = [n]type{}
	var a = [5]byte{} //长度为5的数组，每个元素为一个字节
	fmt.Println("a=", a)
	//数组声明方法2 var arr = [n]type{}
	 b := [5]byte{} //长度为5的数组，每个元素为一个字节
	fmt.Println("b=", b)
	fmt.Println("------------------切片------------------------")
	//切片是一个很小的对象，它对底层的数组(内部是通过数组保存数据的)进行了抽象，并提供相关的操作方法。
	//切片是一个有三个字段的数据结构，这些数据结构包含 Golang 需要操作底层数组的元数据：
	//这 3 个字段分别是
	//1.指向底层数组的指针、
	//2.切片访问的元素的个数(即长度)
	//3.切片允许增长到的元素个数(即容量)。
	//切片声明方法1  arr := []type{} 
	slice1 := make([]int, 5)
	fmt.Println("slice1=", slice1)
	slice2 := []int{1,2,3,4}
	for i,v :=range slice2{
		fmt.Printf("slice2 index: %d,value: %d\n",i, v)
	}
	fmt.Println("------------------------------------------")
	slice2[1] = 3
	for i,v :=range slice2{
		fmt.Printf("slice2 index: %d,value: %d\n",i, v)
	}
	// 使用切片字面量创建空的整型切片
	slice3 := []int{}
	fmt.Println("slice3=", slice3)
	//当使用字面量来声明切片时，其语法与使用字面量声明数组非常相似。
	//二者的区别是：如果在 [] 运算符里指定了一个值，那么创建的就是数组而不是切片。只有在 [] 中不指定值的时候，创建的才是切片。
	// 创建有 3 个元素的整型数组
	myArray := [3]int{10, 20, 30}
	// 创建长度和容量都是 3 的整型切片
	mySlice := []int{10, 20, 30}
	fmt.Println("myArray=", myArray)
	fmt.Println("mySlice=", mySlice)
	// 创建一个整型切片
	// 其长度和容量都是 5 个元素
	myNum := []int{10, 20, 30, 40, 50}
	// 创建一个新切片 通过切片创建新的切片
	// slice[i:j:k] 其中 i 表示从 slice 的第几个元素开始切，j 控制切片的长度(j-i)，k 控制切片的容量(k-i)，
	//如果没有给定 k，则表示切到底层数组的最尾部
	//如果设置的容量比可用的容量还大，就会得到一个运行时错误：
	newNum := myNum[1:5:5]
	for i,v :=range newNum{
		fmt.Printf("**** index: %d,value: %d\n",i, v)
	}
	//如果设置的容量比可用的容量还大，就会得到一个运行时错误
	newNum2 := myNum[1:4]
	for i,v :=range newNum2{
		//fmt.Println("index: ", i, "value: ",v)
		fmt.Printf("newNum index: %d,value: %d\n",i, v)
	}
	//共享底层数组的切片
//需要注意的是：现在两个切片 myNum 和 newNum 共享同一个底层数组。如果一个切片修改了该底层数组的共享
//部分，另一个切片数值也会改变
	myNum[1] = 100
	for i,v :=range newNum{
		fmt.Printf("modify index: %d,value: %d\n",i, v)
	}

	//切片扩容
	//相对于数组而言，使用切片的一个好处是：可以按需增加切片的容量。Golang 内置的 append() 函数会处理增加长度时的所有操作细节。
	//要使用 append() 函数，需要一个被操作的切片和一个要追加的值，当 append() 函数返回时，会返回一个包含修改结果的新切片。
	//函数 append() 总是会增加新切片的长度，而容量有可能会改变，也可能不会改变，这取决于被操作的切片的可用容量。
	for i,v :=range myNum{
		fmt.Printf("now myNum index: %d,value: %d\n",i, v)
	}

	newNum = append(newNum, 60)
	for i,v :=range newNum{
		fmt.Printf("append newNum index: %d,value: %d\n",i, v)
	}
	for i,v :=range myNum{
		fmt.Printf("after append myNum index: %d,value: %d\n",i, v)
	}
	//内置函数 append() 在操作切片时会首先使用可用容量。一旦没有可用容量，就会分配一个新的底层数组。
	//切片间正在共享同一个底层数组。一旦发生这种情况，对切片进行修改，很可能会导致随机且奇怪的问题，
	//如果在创建切片时设置切片的容量和长度一样，就可以强制让新切片的第一个 append 操作创建新的底层数组，
	//与原有的底层数组分离。这样就可以安全地进行后续的修改操作了：
	fruit := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	myFruit := fruit[2:3:3]
	// 向 myFruit 追加新字符串
	myFruit = append(myFruit, "Kiwi")
	fmt.Println("myFruit:",myFruit)
	fmt.Print("fruit:",fruit)
}