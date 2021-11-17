/*
创建工程 初始化模块
go mod init [模块名] //不加模块名，默认就是当前目录名

		PS E:\work_go\src\test_gin> go build
		go: go.mod file not found in current directory or any parent directory; see 'go help modules'
		PS E:\work_go\src\test_gin> go mod init test_gin
		go: creating new go.mod: module test_gin
		go: to add module requirements and sums:
				go mod tidy
		PS E:\work_go\src\test_gin> go mod tidy



//-----------
Go 中命名有 internal 的 package，只有该 package 的父级 package 才可以访问该 package 的内容。
两点需要注意：
	只有直接父级package，以及父级package的子孙package可以访问，其他的都不行，再往上的祖先package也不行
	父级package也只能访问internal package使用大写暴露出的内容，小写的不行
比如, 	"test_gin/gin/internal/bytesconv" 只有在 test_gin/gin 模块范围才能使用 bytesconv
//-----------


Go 语言的基础组成有以下几个部分：

	包声明
	引入包
	函数
	变量
	语句 & 表达式
	注释
*/

/*

关键字
	下面列举了 Go 代码中会使用到的 25 个关键字或保留字：
	break	default	func	interface	select
	case	defer	go	map	struct
	chan	else	goto	package	switch
	const	fallthrough	if	range	type
	continue	for	import	return	var
除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符：
	append	bool	byte	cap	close	complex	complex64	complex128	uint16
	copy	false	float32	float64	imag	int	int8	int16	uint32
	int32	int64	iota	len	make	new	nil	panic	uint64
	print	println	real	recover	string	true	uint	uint8	uintptr
程序一般由关键字、常量、变量、运算符、类型和函数组成。

程序中可能会使用到这些分隔符：括号 ()，中括号 [] 和大括号 {}。

程序中可能会使用到这些标点符号：.、 ,、 ;、 :  和 … 。

*/

/*
//数据类型
1	布尔型 ,	布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。
2	数字类型 ,	整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。
3	字符串类型,	字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。
4	派生类型:
		包括：
		(a) 指针类型（Pointer）
		(b) 数组类型
		(c) 结构化类型(struct)
		(d) Channel 类型
		(e) 函数类型
		(f) 切片类型
		(g) 接口类型（interface）
		(h) Map 类型



数字类型 Go 也有基于架构的类型，例如：int、uint 和 uintptr。
	1	uint8 	无符号 8 位整型 (0 到 255)
	2	uint16 	无符号 16 位整型 (0 到 65535)
	3	uint32 	无符号 32 位整型 (0 到 4294967295)
	4	uint64 	无符号 64 位整型 (0 到 18446744073709551615)
	5	int8 	有符号 8 位整型 (-128 到 127)
	6	int16 	有符号 16 位整型 (-32768 到 32767)
	7	int32 	有符号 32 位整型 (-2147483648 到 2147483647)
	8	int64 	有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
浮点型
	1	float32 	IEEE-754 32位浮点型数
	2	float64 	IEEE-754 64位浮点型数
	3	complex64 	32 位实数和虚数
	4	complex128 	64 位实数和虚数
其他数字类型
	1	byte 	类似 uint8
	2	rune 	类似 int32
	3	uint 	32 或 64 位
	4	int 	与 uint 一样大小
	5	uintptr 	无符号整型，用于存放一个指针

*/

/*
type 可以给类型取别名

	type bigint int64 //给 int64 起一个别名为 bigint

	var a bigintfmt.Printf("a的类型为%T\n",a) //这里 a 的类型为 bigint


*/

/*
字面量
	是表示值的一种标记法，但是在Go语言中，字面量的含义要更广一些：

	1.用于表示基础数据类型值的各种字面量。
	2.用户构造各种自定义的复合数据类型的类型字面量

	如下面这个 字面量 表示了一个名称为 Person 的自定义结构体类型：
	type Person struct {
	     Name string
	     Age uint8
	     Address string
	}

	用于表示复合数据类型的值的复合字面量，更确切地讲，它会被用来构造类型 Struct（结构体）、Array（数组）、Slice（切片）和Map（字典）的值。

	如下面的字面量可以表示上面的那个 Person 结构体类型的值：
	Person(Name: "Eric Pan", Age: 28, Address: "Beijing China"}

*/

/*

声明变量的一般形式是使用 var 关键字：
	var identifier type
	var identifier1, identifier2 type //一次声明多个变量：



常量的定义:
	const identifier [type] = value
	你可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。
	显式类型定义： const b string = "abc"
	隐式类型定义： const b = "abc"
	多个相同类型的声明可以简写为：
	const c_name1, c_name2 = value1, value2


运算符
	&	返回变量存储地址	&a; 将给出变量的实际地址。
	*	指针变量。	*a; 是一个指针变量
		示例
		{
	   		var a int = 4
	   		var ptr *int  //指针变量
	   		ptr = &a     // 'ptr' 包含了 'a' 变量的地址
		}

	= 	赋值,浅拷贝
	:= 	声明变量并赋值, 声明变量前面不用 var

*/

/*
Go 没有三目运算符，所以不支持 ?: 形式的条件判断

switch 	语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。
switch 	语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。
switch 	默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 fallthrough
		使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true



*/

/*
Go 语言的 goto 语句可以无条件地转移到过程中指定的行。
goto 语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。
但是，在结构化程序设计中一般不主张使用 goto 语句， 以免造成程序流程的混乱，使理解和调试程序都产生困难。
	goto 语法格式如下：

	{
		goto label;
		..
		.
		label: statement;
	}

*/

/*
Go 语言函数
	Go 语言最少有个 main() 函数。
	Go 语言标准库提供了多种可动用的内置的函数。
		例如，len() 函数可以接受不同类型参数并返回该类型的长度。如果我们传入的是字符串则返回字符串的长度，如果传入的是数组，则返回数组中包含的元素个数。

定义
	func function_name( [parameter list] ) [return_types] {
	   函数体
	}
func：			关键字,函数由 func 开始声明
function_name：	函数名称，参数列表和返回值类型构成了函数签名。
parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
return_types：	返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
				函数返可以回多个值
函数体：		函数定义的代码集合。


//注意 参数没有默认值.  一般多参数的方法可定义一个类型做为参数, 后续好扩展


返回多个值示例
	func swap(x, y string) (string, string) {
	   return y, x
	}

	func main() {
	   a, b := swap("Google", "Runoob")
	   fmt.Println(a, b)
	}

参数传递类型
	值传递	值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
	引用传递	引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
	默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

		func swap(x *int, y *int) {
		   var temp int
		   temp = *x    // 保持 x 地址上的值
		   *x = *y      // 将 y 值赋给 x
		   *y = temp    // 将 temp 值赋给 y
		}

		var a,b int
		swap(&a, &b)



函数作为参数


func main() {
   getSquareRoot := func(x float64) float64 {
      return math.Sqrt(x)
   }
   getSquareRoot1 := func(x float64,ff func(float64 ) float64 ) float64 {
      return ff(x)
   }

   fmt.Println(getSquareRoot1(64,getSquareRoot))
}


函数闭包
func getSequence() func() int {
   i:=0
   return func() int {
      i+=1
     return i
   }
}

func main(){
	// nextNumber 为一个函数，函数 i 为 0
   nextNumber := getSequence()

   // 调用 nextNumber 函数，i 变量自增 1 并返回
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())

   // 创建新的函数 nextNumber1，并查看结果
   nextNumber1 := getSequence()
   fmt.Println(nextNumber1())
   fmt.Println(nextNumber1())
}

匿名函数
	f:=func() int {
      i+=1
     return i
   }

函数和方法
	一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
	所有给定类型的方法属于该类型的方法集。语法格式如下：

	func (variable_name variable_data_type) function_name() [return_type]{
	   // 函数体
	}

//示例
		// 定义结构体
		type Circle struct {
		  radius float64
		}

		func main() {
		  var c1 Circle
		  c1.radius = 10.00
		  fmt.Println("圆的面积 = ", c1.getArea())
		}

		//该 method 属于 Circle 类型对象中的方法
		func (c Circle) getArea() float64 {
		  //c.radius 即为 Circle 类型对象中的属性
		  return 3.14 * c.radius * c.radius
		}


*/

/*
不同类型的局部和全局变量默认值为：
	int	0
	float32	0
	pointer	nil




类型转换
	type_name(expression)
	type_name 为类型，expression 为表达式。

*/

/*
Go 数组类型
	数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整型、字符串或者自定义类型。

声明数组
	声明需要指定元素类型及元素个数，语法格式如下：
		var variable_name [SIZE] variable_type   //一维数组

		//balance 长度为 10 类型为 float32：
		var balance [10] float32
初始化数组
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	也可以通过 字面量 在声明数组的同时快速初始化数组：
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度：
	var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
		或
	balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
如果设置了数组的长度，我们还可以通过指定下标来初始化元素：

//  将索引为 1 和 3 的元素初始化,其他会用默认值填充
balance := [5]float32{1:2.0,3:7.0}
初始化数组中 {} 中的元素个数不能大于 [] 中的数字。

如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：



*/

/*
Slice 是对数组的抽象,动态数组
定义
	1.声明一个未指定大小的数组来创建：
	var identifier []type

	2.使用 make() 函数来创建
		var slice1 []type = make([]type, len)
		简写
		slice1 := make([]type, len)
		make([]T, length, capacity) //可以指定容量，其中 capacity 为可选参数。


初始化, 一个切片在未初始化之前默认为 nil，长度为 0
	s :=[] int {1,2,3 }
	直接初始化切片，[] 表示是切片类型，{1,2,3} 初始化值依次是 1,2,3，其 cap=len=3。

	s := arr[:]
	初始化切片 s，是数组 arr 的引用。

	s := arr[startIndex:endIndex]
	将 arr 中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片。

	s := arr[startIndex:]
	默认 endIndex 时将表示一直到arr的最后一个元素。

	s := arr[:endIndex]
	默认 startIndex 时将表示从 arr 的第一个元素开始。

	s1 := s[startIndex:endIndex]
	通过切片 s 初始化切片 s1。

	s :=make([]int,len,cap)
	通过内置函数 make() 初始化切片s，[]int 标识为其元素类型为 int 的切片。


len() //长度
cap() //容量
append(tar, xxx) //向 tar 追加元素
copy(tar,src) //拷贝 src 到 tar

清空 todo

Slice 截取
	可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]
	numbers[4:] // 从 4 到最后.
	numbers[:2]  // 0(包含) 到 索引 2(不包含)
	numbers[2:5] // 2(包含) 到 索引 5(不包含)

*/

/*
range 关键字
	用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
	在数组和 slice 中它返回元素的 索引 和 索引对应的值，在集合中返回 key-value 对。

	for i, num := range nums {
		fmt.Printf("idx %d ->val %d\n", i, num)
	}
	for k, v := range kvs {
		fmt.Printf(" %s -> %s\n", k, v)
	}

//传统遍历 注意 for 后面没有括号
for i := 0; i < len(a); i++ {
	fmt.Println(&a[i])
}





*/

/*

Map

//声明变量，默认 map 是 nil, 需要初始化或 make创建才能使用
var map_variable map[key_data_type]value_data_type

//使用 make 函数
map_variable := make(map[key_data_type]value_data_type)

清空, 新创建一个新 map 赋值,就是空的了

//判断是否为空
if len(map) == 0 {
    ....
}

//判断是否存在某个key

_,ok:= map1["slogan"]
if ok {
	fmt.Println("存在")
}else{
	fmt.Println("不存在")
}

*/

/*
interface
	可把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
interface{} 类型是没有方法的接口 (这是一个类型,相当于 c# 的 object )
	所有类型都至少实现了 0 个方法，所以 所有类型都实现了空接口。如果函数以 interface{} 值作为参数，那么可以为该函数提供任何值.

	//判断是否实现了某类型接口
		//安全的断言方法： <目标接口类型值>, ok := <空接口值>.(目标接口类型)
		//非安全的断言方法：<目标接口类型值> := <空接口值>.(目标接口类型)

	//x.(type) 类型断言,返回 x 的类型. 当 x 是 interface{} 时非常有用

	Go 中的所有东西都是按值传递的。每次调用函数时，传入的数据都会被复制。
	对于具有值接收者的方法，在调用该方法时将复制该值。例如下面的方法：
		func (t *T)MyMethod(s string) {
		    // ...
		}
	所以,一般用指针参数，减少复制
		func (t *T)MyMethod(s *string) {
		    // ...
		}





// 定义接口
type interface_name interface {
   method_name () [return_type]
}

// 定义结构体
type struct_name1 struct {
   // variables
}
type struct_name2 struct {
   // variables
}

type struct_name3 struct {
	struct_name1    //继承
	stn2 struct_name2 //组合
   // variables
}



// 实现接口方法
func (struct_name_variable struct_name1) method_name() [return_type] {
}
...
func (struct_name_variable struct_name2) method_name() [return_type] {
}

//使用
objs := []interface_name {struct_name1{}, struct_name1{}}
for _, obj := range objs {
	obj.method_name()
}


//嵌套写法
aaa := map[string]interface{}{
	"tp":          1,
	"obj":         map[string]interface{}{"a": 1},
	"arr_obj":     []map[string]interface{}{{"b": 1}, {}},
	"arr_arr":     [][]map[string]interface{}{},
	"arr_arr_arr": [][][]map[string]interface{}{},
}


*/

/*

继承
一个结构体嵌到另一个结构体，称作组合
匿名和组合的区别
如果一个struct嵌套了另一个匿名结构体，那么这个结构可以直接访问匿名结构体的方法，从而实现继承
如果一个struct嵌套了另一个【有名】的结构体，那么这个模式叫做组合
如果一个struct嵌套了多个匿名结构体，那么这个结构可以直接访问多个匿名结构体的方法，从而实现多重继承








*/

/*
new 和 make

make	只能用于 slice，map，channel 三种类型
	 	make(T, args) 返回的是初始化之后的 T 类型的值，这个新值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。

new
	new(T) 为一个 T 类型新值分配空间并将此空间初始化为 T 的零值，返回的是新值的地址，也就是 T 类型的指针 *T，该指针指向 T 的新分配的零值。
	不要使用 new 来构造 map

	可以不用 new, 直接用 var ,声明时会自动初始化 结构体下每个类型的零值，也可通过 字面量 完成初始化。
	var st1 Foo 或  st2 := Foo{}

	//声明初始化
	var foo1 Foo
	fmt.Printf("foo1 --> %#v\n ", foo1)
	foo1.age = 1
	fmt.Println(foo1)

	//struct literal 初始化
	foo2 := Foo{}
	fmt.Printf("foo2 --> %#v\n ", foo2)
	foo2.age = 2
	fmt.Println(foo2)

	//指针初始化
	foo3 := &Foo{}
	fmt.Printf("foo3 --> %#v\n ", foo3)
	foo3.age = 3
	fmt.Println(foo3)

	//new 初始化
	foo4 := new(Foo)
	fmt.Printf("foo4 --> %#v\n ", foo4)
	foo4.age = 4
	fmt.Println(foo4)

	//声明指针并用 new 初始化
	var foo5 *Foo = new(Foo)
	fmt.Printf("foo5 --> %#v\n ", foo5)
	foo5.age = 5
	fmt.Println(foo5)

*/

/*
goroutine
	协程, 调度是由 Golang 运行时进行管理的。
语法格式：
	go 函数名( 参数列表 )

*/

/*
channel
	用来传递数据的一个数据结构。
	可用于两个 goroutine 之间,通过传递一个指定类型的值来同步运行和通讯。
	操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
	默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。

		ch := make(chan int)
		ch <- v    // 把 v 发送到通道 ch
		v := <-ch  // 从 ch 接收数据, 并把值赋给 v



	通过 make 的第二个参数指定缓冲区大小：
		ch := make(chan int, 100)

	带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。
	不过由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，否则缓冲区一满，数据发送端就无法再发送数据了。

	注意:	如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。
			如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内, 如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。 接收方在有值可以接收之前会一直阻塞。

创建无缓冲channel
	chreadandwrite :=make(chan int) //读写
	chonlyread := make(<-chan int) //创建只读channel
	chonlywrite := make(chan<- int) //创建只写channel




遍历通道与关闭通道
	通过 range 关键字来实现遍历读取到的数据，类似于与数组或切片。格式如下：
		for true {
			v, ok := <-ch   //v 为接收值
			if (ok) {

			}else{
					//通道使用 close() 后，通道再接收数据 ok 就为 false , 这时跳出循环
				break
			}
		}
		//其他地方 close(ch)


主线程等待协程退出
		var wg sync.WaitGroup
		func test(){
			wg.done() //计数减1
		}
		func main(){
			wg.Add(1) // 增加1个计数
			go test()
			wg.Wait()
		}
*/

/*
import  引用的是"目录"(相对于 go mod init 当前项目模块名  的路径, 比如当前模块名是 helloworld, 包 test 在 ./tttt 下,那引用的路径就是  helloworld/tttt  ) ,  "当前项目模块名" + 子目录
		引入的同一个目录下，只能有一个包名
package 声明一个包名,注意如果要包方法在其他包中可以调用，包方法需要首字母大写


	import "fmt"最常用的一种形式
	import "./test"导入同一目录下test包中的内容
	import f "fmt"导入fmt，并给他启别名ｆ
	import . "fmt"，将fmt启用别名"."，这样就可以直接使用其内容，而不用再添加ｆｍｔ，如fmt.Println可以直接写成Println
	import  _ "fmt" 表示不使用该包，而是只是使用该包的init函数，并不显示的使用该包的其他内容。
			注意：这种形式的 import，当 import时 就执行了fmt 包中的 init 函数，而不能够使用该包的其他函数。


注: 包名不要和已有的重名

import (
	"fmt"
	test "helloworld/pbtest"
	"helloworld/test__123
)

*/

/*

defer  延时处理,在方法结束时再处理. 如果一个方法里定义了多个defer, 处理顺序和栈一样，后定义的先处理.
panic
recover

	Go不支持传统的 try…catch…finally 这种异常，在Go中，可使用多值返回来返回错误。不要用异常代替错误，更不要用来控制流程。
	在极个别的情况下，也就是说，遇到真正的异常的情况下（比如除数为 0了）。才使用Go中引入的Exception处理：defer, panic, 。
	这几个异常的使用场景可以这么简单描述：Go中可以抛出一个 panic 的异常，然后在 defer 中通过 recover 捕获这个异常，然后正常处理。


*/

/*
select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。

select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。

语法
Go 编程语言中 select 语句的语法如下：

select {
    case communication clause  :
       statement(s);
    case communication clause  :
       statement(s);
     //你可以定义任意数量的 case
    default : // 可选
       statement(s);
}
以下描述了 select 语句的语法：

每个 case 都必须是一个通信
所有 channel 表达式都会被求值
所有被发送的表达式都会被求值
如果任意某个通信可以进行，它就执行，其他被忽略。
如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
否则：
如果有 default 子句，则执行该语句。
如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值

*/

/*
//反射
//"reflect"


*/

/*
//context
https://www.cnblogs.com/zhangboyu/p/7456606.html

每个 Goroutine 在执行之前，都要先知道程序当前的执行状态，通常将这些执行状态封装在一个Context变量中，传递给要执行的Goroutine中。
上下文则几乎已经成为传递与请求同生存周期变量的标准方法。
在网络编程下，当接收到一个网络请求Request，处理Request时，我们可能需要开启不同的 Goroutine 来获取数据与逻辑处理，即一个请求 Request，会在多个 Goroutine 中处理。
而这些Goroutine可能需要共享Request的一些信息；同时当Request被取消或者超时的时候，所有从这个Request创建的所有Goroutine也应该被结束。

context 包的核心就是 Context 接口，其定义如下：
	type Context interface {
	    Deadline() (deadline time.Time, ok bool)
	    Done() <-chan struct{}
	    Err() error
	    Value(key interface{}) interface{}
	}
	//Deadline(), 返回 context 的截止时间，通过此时间，函数就可以决定是否进行接下来的操作，如果时间太短，就可以不往下做了，否则浪费系统资源。
					当然，也可以用这个 deadline 来设置一个 I/O 操作的超时时间。

	//Done(), 返回一个 channel，可以表示 context 被取消的信号, 当这个 channel 被关闭时，说明 context 被取消了。
				注意，这是一个只读的channel。 我们又知道，读一个关闭的 channel 会读出相应类型的零值。
				并且源码里没有地方会向这个 channel 里面塞入值。 换句话说，这是一个 receive-only 的 channel。
				因此在子协程里读这个 channel，除非被关闭，否则读不出来任何东西。也正是利用了这一点，子协程从 channel 里读出了值（零值）后，就可以做一些收尾工作，尽快退出。
	//Err(), 返回一个错误，表示 channel 被关闭的原因。例如是被取消，还是超时。
	//Value(), 获取之前设置的 key 对应的 value。
			type vars struct {
			    lock    sync.Mutex
			    db      *sql.DB
			}
			//Then you can add this struct in context:
				ctx := context.WithValue(context.Background(), "values", vars{lock: mylock, db: mydb})
			//And you can retrieve it:
				ctxVars, ok := r.Context().Value("values").(vars)






//api --------
		context.Background() 返回一个空的 Context
			我们可以用这个 空的 Context 作为 goroutine 的 root 节点（如果把整个 goroutine 的关系看作 树状）
			一般由接收请求的第一个 Goroutine 创建，是与进入请求对应的 Context 根节点。
			它不能被取消、没有值、也没有过期时间，常常作为处理 Request 的顶层 context 存在。
			通过WithCancel、WithTimeout函数来创建子对象，其可以获得 cancel、timeout 的能力。

		context.TODO() 也返回一个 emptyCtx 类型的对象。
			在目前还不清楚要使用的上下文时，或上下文尚不可用时，使用 context.TODO() 生成的 Context 接口类型的对象


		context.WithCancel(parent) 函数，创建一个可取消的子 Context
			函数返回值有两个：子Context Cancel 取消函数
		func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)  创建一个有生命周期的 context

//--------


*/

package main

import (
	"fmt"
	test "helloworld/pbtest"
	"helloworld/test__123"

	//"helloworld/test__123"
	"log"
	"reflect"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	"github.com/davyxu/cellnet/peer"
	_ "github.com/davyxu/cellnet/peer/tcp" // 注册TCP Peer
	"github.com/davyxu/cellnet/proc"
	_ "github.com/davyxu/cellnet/proc/tcp" // 注册TCP Processor
	"github.com/davyxu/cellnet/util"
)

func main_test() {
	//test__123.TestSingle()
	//test__123.TestNet()
	//test__123.Say()
	//test__123.Test_mongodb()
	//test__123.Test_gin()
	test__123.Test_beego()
	//test__123.Test_nacos()
	//test__123.Test_sample_nacos()
}

func main_gin() {

	//r := gin.Default()
	//r.GET("/", func(c *gin.Context) {
	//	c.String(200, "Hello")
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080
}

const peerAddress = "127.0.0.1:17701"

type TestEchoACK struct {
	Msg   string
	Value int32
}

type TestI interface {
	act()
}

func (self *TestEchoACK) act() {
	fmt.Println("TestEchoACK act!")
	self.Value = 123
}

func testFunc(a TestI) {
	a.act()
}

func ttttttt() {
	a := &TestEchoACK{"aa", 100}

	testFunc(a)

	fmt.Printf("TestEchoACK %+v\n", a)
}

// 服务器逻辑
func server_cellnet() {

	// 创建服务器的事件队列，所有的消息，事件都会被投入这个队列处理
	queue := cellnet.NewEventQueue()

	// 创建一个服务器的接受器(Acceptor)，接受客户端的连接
	peerIns := peer.NewGenericPeer("tcp.Acceptor", "server", peerAddress, queue)

	// 将接受器Peer与tcp.ltv的处理器绑定，并设置事件处理回调
	// tcp.ltv处理器负责处理消息收发，使用私有的封包格式以及日志，RPC等处理
	proc.BindProcessorHandler(peerIns, "tcp.ltv", func(ev cellnet.Event) {

		// 处理Peer收到的各种事件
		switch msg := ev.Message().(type) {
		case *cellnet.SessionAccepted: // 接受一个连接
			fmt.Println("server accepted")
		case *test.ContentACK: // 收到连接发送的消息

			fmt.Printf("server recv pb: %+v\n", msg)

		case *TestEchoACK: // 收到连接发送的消息

			fmt.Printf("server recv %+v\n", msg)

			// 发送回应消息
			ev.Session().Send(&TestEchoACK{
				Msg:   msg.Msg,
				Value: msg.Value,
			})

		case *cellnet.SessionClosed: // 会话连接断开
			fmt.Println("session closed: ", ev.Session().ID())
		}

	})

	// 启动Peer，服务器开始侦听
	peerIns.Start()

	// 开启事件队列，开始处理事件，此函数不阻塞
	queue.StartLoop()
}
func main_cellnet() {
	server_cellnet()
	for true {
		time.Sleep(time.Microsecond * 30)
	}
}

func init() {
	cc := codec.MustGetCodec("binary")
	tt := reflect.TypeOf((*TestEchoACK)(nil)).Elem()
	ii := int(util.StringHash("TestEchoACK"))
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{Codec: cc, Type: tt, ID: ii})
}

func actdebug() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

}

func test_protobuf() {
	var pbmsg test.ContentACK
	pbmsg.Msg = "hello"
	pbmsg.Value = 123

}

func main() {
	main_test()
	//actdebug()
	//main_cellnet()
	//ttttttt()
}
