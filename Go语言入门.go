// Go语言入门
//// Go的第一个程序
//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("Happy every day!")
//}

// for-range遍历数组、切片、字符串、map和通道（channel）
// 遍历数组和切片
// package main
//
// import "fmt"
//
//	func main() {
//		for key, value := range []int{0, 1, -1, -2} {
//			fmt.Printf("key:%d value:%d\n", key, value)
//		}
//	}
//
// 遍历字符串
// package main
//
// import "fmt"
//
//	func main() {
//		var str string = "hello"
//		for key, value := range str {
//			fmt.Printf("key:%d value:%d\n", key, value)
//		}
//	}
//
// 遍历map
//package main
//
//import "fmt"
//
//func main() {
//	m := map[string]int{
//		"go":  100,
//		"web": 100,
//	}
//	for key, value := range m {
//		fmt.Println(key, value)
//	}
//}

// 遍历通道
// package main
//
// import "fmt"
//
//	func main() {
//		c := make(chan int)
//		go func() {
//			c <- 7
//			c <- 8
//			c <- 9
//			close(c)
//		}()
//		for v := range c {
//			fmt.Println(v)
//		}
//	}

// 单个字节的字符串遍历
//package main
//
//import "fmt"
//
//func main() {
//	str := "happy"
//	fmt.Println(string(str[0]))     //获取字符索引为0的字符
//}

// rune切片对多字节字符单个输出
// package main
//
// import "fmt"
//
//	func main() {
//		str := "你好 世界！"
//		chars := []rune(str)
//		for _, char := range chars {
//			fmt.Println(string(char))
//		}
//	}

// for-range将字符串的一个字节一个字节迭代出来
// package main
//
// import "fmt"
//
//	func main() {
//		str := "hallo 世界！"
//		for index, char := range str {
//			fmt.Printf("%d %U %c \n", index, char, char)
//		}
//	}

// 字符串修改字节
// package main
//
// import "fmt"
//
//	func main() {
//		str := "Hi 世界"
//		by := []byte(str)
//		by[2] = ','                           //字节用单引号
//		fmt.Printf("%s\n", str)
//		fmt.Printf("%s\n", by)
//	}

// rune修改字符
//package main
//
//import "fmt"
//
//func main() {
//	str := "Hi 世界"
//	by := []rune(str)
//	by[3] = '中'
//	by[4] = '国'
//	fmt.Println(str)
//	fmt.Printf("%s", string(by))
//}

// 指针修改值
//package main
//
//import "fmt"
//
//func exchange(c, d *int) {
//	t := *c
//	*c = *d
//	*d = t
//}
//func main() {
//	a := 2
//	b := 6
//	exchange(&a, &b)
//	fmt.Println(a, b)
//}

//声明一个数组
//package main
//
//func main() {
//	var a [10]int
//}

//初始化一个数组
//package main
//
//func main() {
//	var number = [3]int{1,2,3}
//}

// 结构体类型
// 定义一个名为Book的图书结构体，并打印出结构体的字段值
//package main
//
//import "fmt"
//
//func main() {
//	type Book struct {
//		title   string
//		author  string
//		subject string
//		press   string
//	}
//	fmt.Println(Book{"Go web", "徐倩", "Go", "东软出版社"})
//	fmt.Println(Book{title: "Go web", author: "徐倩", subject: "Go", press: "东软出版社"})
//	fmt.Println(Book{title: "Go web", author: "徐倩"})
//}

// 访问结构体成员的示例
//package main
//
//import "fmt"
//
//func main() {
//	type Books struct {
//		title   string
//		author  string
//		subject string
//		press   string
//	}
//	var bookGo Books
//	var bookPhp Books
//
//	bookGo.title = "Go web"
//	bookGo.author = "徐倩"
//	bookGo.subject = "Go"
//	bookGo.press = "成都东软出版社"
//
//	bookPhp.title = "Php web"
//	bookPhp.author = "徐倩"
//	bookPhp.subject = "Php"
//	bookPhp.press = "成都东软出版社"
//
//	fmt.Printf("bookGo title : %s\n", bookGo.title)
//	fmt.Printf("bookGo author : %s\n", bookGo.author)
//	fmt.Printf("bookGo subject : %s\n", bookGo.subject)
//	fmt.Printf("bookGo press : %s\n", bookGo.press)
//
//	fmt.Printf("bookPhp title : %s\n", bookPhp.title)
//	fmt.Printf("bookPhp author : %s\n", bookPhp.author)
//	fmt.Printf("bookPhp subject : %s\n", bookPhp.subject)
//	fmt.Printf("bookPhp press : %s\n", bookPhp.press)
//}

// 将结构体作为函数参数
//package main
//
//import "fmt"
//
//type Books struct {
//	title   string
//	author  string
//	subject string
//	press   string
//}
//
//func main() {
//
//	var bookGo Books
//	var bookPhp Books
//
//	bookGo.title = "Go web"
//	bookGo.author = "徐倩"
//	bookGo.subject = "Go"
//	bookGo.press = "成都东软出版社"
//
//	bookPhp.title = "Php web"
//	bookPhp.author = "徐倩"
//	bookPhp.subject = "Php"
//	bookPhp.press = "成都东软出版社"
//
//	printBook(bookGo)
//	printBook(bookPhp)
//}
//
//func printBook(book Books) {
//	fmt.Printf("book title : %s\n", book.title)
//	fmt.Printf("book author : %s\n", book.author)
//	fmt.Printf("book subject : %s\n", book.subject)
//	fmt.Printf("book press : %s\n", book.press)
//}

// 从数组生成切片
//package main
//
//import "fmt"
//
//func main() {
//	var a = [3]int{1, 2, 3}
//	fmt.Println(a, a[1:3])
//}

// 声明切片
//package main
//
//import "fmt"
//
//func main() {
//	var sliceStr []string
//	var sliceInt []int
//	var emptySliceNum []int
//	//输出3个切片
//	fmt.Println(sliceStr)
//	fmt.Println(sliceInt)
//	fmt.Println(emptySliceNum)
//	//输出3个切片大小
//	fmt.Println(len(sliceStr), len(sliceInt), len(emptySliceNum))
//	//切片判定为空的结果
//	fmt.Println(sliceStr == nil)
//	fmt.Println(sliceInt == nil)
//	fmt.Println(emptySliceNum == nil)
//}

// make函数向新的切片添加内容
//package main
//
//import "fmt"
//
//func main() {
//	var slice1 []int
//	var slice2 []string
//	slice1 = make([]int, 6)
//	slice2 = make([]string, 6, 10)
//	fmt.Println(slice1)
//	fmt.Println(slice2)
//}

// map类型的使用
// package main
//
// import "fmt"
//
//	func main() {
//		var literalMap map[string]string
//		var assignedMap map[string]string
//		literalMap = map[string]string{"first": "go", "second": "web"}
//		assignedMap = literalMap
//		createdMap := make(map[string]float32)
//		createdMap = map[string]float32{"k1": 99, "k2": 199}
//		assignedMap["second"] = "project"
//		fmt.Printf("Map literalMap at first is %s\n", literalMap["first"])
//		fmt.Printf("Map assignedMap at second is %s\n", assignedMap["second"])
//		fmt.Printf("Map literalMap at third is %s\n", literalMap["third"])
//		fmt.Printf("Map createdMap at first is %f\n", createdMap["k1"])
//	}

//函数返回最小值
//package main
//
//import "fmt"
//
//func main() {
//	var arr = []int{6, 7, 8}
//	var a int
//	a = min(arr)
//	fmt.Printf("最小值是：%d", a)
//}
//func min(arr []int) (min int) {
//	min = arr[0]
//	for _, v := range arr {
//		if v < min {
//			min = v
//		}
//	}
//	return
//}

// 函数返回多个值
//package main
//
//import "fmt"
//
//func compute(x, y int) (int, int) {
//	return x + y, x * y
//}
//func main() {
//	var a, b = 6, 8
//	a, b = compute(a, b)
//	fmt.Println(a, b)
//}

// 可变参数
//package main
//
//import "fmt"
//
//func myFunc(arg ...string) {
//	//arg告诉Go可接受不定量参数，均为string类型
//	//可以通过for-range语句遍历
//}
//func main() {
//	for _, v := range arg {
//		fmt.Printf("And the string is: %s\n", v)
//	}
//}

//可通过for-range遍历

//参数传递

// 匿名函数使用案例
//package main
//
//import "fmt"
//
//func main() {
//	x, y := 6, 8
//	defer func(a int) {
//		fmt.Println("defer x, y = ", a, y) //y为闭包引用
//	}(x) //a就是x,而y是闭包引用
//	x += 10
//	y += 100
//	fmt.Println(x, y)
//}

// 匿名函数的调用
//package main
//
//import "fmt"
//
//func main() {
//	//定义匿名函数并赋值给f变量
//	f := func(data int) {
//		fmt.Println("hi,this is a closure", data)
//	}
//	//调用
//	f(6)
//	//直接声明并调用
//	func(data int) {
//		fmt.Println("hi,this is a closure directly", data)
//	}(8)
//}

// 遍历切片中每个元素，通过给定的函数访问元素
//package main
//
//import "fmt"
//
//func visitPrint(list []int, f func(int)) {
//	for _, value := range list {
//		f(value)
//	}
//}
//func main() {
//	sli := []int{1, 6, 8}
//	visitPrint(sli, func(value int) {
//		fmt.Println(value)
//	})
//}

// defer反序的示例
//package main
//
//import "fmt"
//
//func main() {
//	deferCall()
//}
//func deferCall() {
//	defer func1()
//	defer func2()
//	defer func3()
//}
//func func1() {
//	fmt.Println("A")
//}
//func func2() {
//	fmt.Println("B")
//}
//func func3() {
//	fmt.Println("C")
//}

// 面向对象编程，封装计算三角形面积
//package main
//
//import "fmt"
//
//type Triangle struct {
//	Bottom float32
//	Height float32
//}
//
//func (t *Triangle) Area() float32 { //Area方法
//	return (t.Bottom * t.Height) / 2
//}
//func main() {
//	r := Triangle{6, 8}
//	fmt.Println(r.Area())
//}


// goroutine
//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func HelloWorld() {
//	fmt.Println("this is a goroutine msg")
//}
//func main() {
//	go HelloWorld()
//	time.Sleep(1 * time.Second)
//	fmt.Println("end")
//}


