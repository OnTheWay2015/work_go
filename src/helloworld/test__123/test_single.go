package test__123

import (
	"fmt" // 导入内置 fmt 包
	"reflect"
	"sync"
	"time"
)

type Circle struct {
	radius  float64
	radius1 float64
	ptr     *Circle
}

func (c Circle) aaa() {
	fmt.Println(c.radius)
}

func main1() {
	fmt.Println("Hello World !")
	var c1 Circle
	c1.radius = 111
	var ptr *Circle = &c1 // &后面接变量名，表示取出该变量的内存地址

	ptr.radius = 100
	balance := [5]Circle{2: c1, 3: *ptr}
	balance1 := [...]Circle{2: c1, 3: *ptr}

	fmt.Println(balance)
	fmt.Println(balance1)

}

func test_range() {
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	var num int
	num = 1
	fmt.Println("num:", num)

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)
	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for _, v := range kvs {
		fmt.Printf("-> %s\n", v)
	}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func test_map() {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
	for k, _ := range countryCapitalMap {
		delete(countryCapitalMap, k)
	}
}

type ff interface {
	f
}

type f interface {
	act() string
}

type FFoo struct {
	name  string
	value int
}

type Foo1 struct {
	name1 string
	ff    FFoo
}
type Foo2 struct {
	name2 string
	ff    FFoo
}
type Foo3 struct {
	name3 string
	ff    FFoo
}

func (o Foo1) act() string {
	o.ff.value = 123
	o.name1 = "123"

	return o.name1
}
func (o Foo2) act() string {
	o.ff.value = 123
	o.name2 = "123"
	return o.name2
}

func (o *Foo3) act() string {
	o.ff.value = 123
	o.name3 = "123"
	return o.name3
}
func test_interface1() {
	x1 := Foo1{}
	x3 := &Foo3{}

	fmt.Println(x1)
	fmt.Println(x3)

	x1.act()
	x3.act()

	fmt.Println(x1)
	fmt.Println(x3)

}

func test_interface() {
	//ppp := tt{f: Foo1{""}}
	//ppp.f.act()

	x := Foo1{}
	x2 := Foo2{}
	xxx := []f{x, x2}
	for i, v := range xxx {
		fmt.Println(i)
		fmt.Println(v.act())

	}

}

func aaaa(a interface{}) {
	//fmt.Println(a)
	i, ok := a.(f)
	if ok {
		i.act()
	}
}

func go_sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	time.Sleep(1000 * time.Millisecond)
	c <- sum // 把 sum 发送到通道 c
	c <- sum // 把 sum 发送到通道 c
	fmt.Println("test_sum end")
	//close(c)
}

var wg sync.WaitGroup
var c chan int = make(chan int, 3)

func test_goclose() {
	time.Sleep(5000 * time.Millisecond)
	close(c)
}

func test_go() {

	s := []int{7, 2, 8, -9, 4, 0}

	go go_sum(s[:len(s)/2], c)
	//x, y := <-c, <-c // 从通道 c 中接收
	for true {
		v, ok := <-c
		if ok {
			fmt.Printf("v:%d\n", v)
		} else {
			break
		}

	}
	fmt.Println("test_go end")
	wg.Done()
}

func testgogo() {
	wg.Add(1) // 因为有一个动作，所以增加1个计数
	go test_goclose()
	go test_go()
	wg.Wait()
}

func test_reflect() {
	//t:= reflect.TypeOf((*cellnet.SessionConnected)(nil)).Elem()
	t := reflect.TypeOf((*Foo1)(nil))
	e := t.Elem()
	k := t.Kind()

	fmt.Println("e:", e, "k:", k)

	tt := reflect.TypeOf((*int32)(nil))
	ee := tt.Elem()
	kk := tt.Kind()
	fmt.Println("ee:", ee, "kk:", kk)

}

func TestSingle() {
	//test_range()
	//test_map()
	//test_interface()
	//aaaa(123)

	//testgogo()
	//test_interface1()
	test_reflect()
	fmt.Println("main end")
}
