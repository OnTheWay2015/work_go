package test__123

import (
	"encoding/json"
	"fmt" // 导入内置 fmt 包
	"reflect"
	"sync"
	"time"
)

type Circle struct {
	radius  float64
	radius1 float64
	ptr     *Circle
	f       float64
	s       string
	i       int
}
type Circle1 struct {
	Radius  float64
	Radius1 float64
	f       float64
	s       string
	i       int
	Ptr     *Circle1
	SS      string
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

func test_struct2(c *Circle) {
	c.radius = 1234
}
func test_struct1(c Circle) {
	c.radius = 1234
}

func test_struct() {
	a := Circle{}
	test_struct1(a)
	fmt.Println("1:", a)
	test_struct2(&a)
	fmt.Println("2:", a)
}

func test_map6(m *map[string]int) {
	m = &map[string]int{"aa": 200} //只是修改了局部变量 m 的值
}
func test_map5(m *map[string]int) {
	*m = map[string]int{"aa": 100} //*m 读取了传入参数的实际地址
}

func test_map4(m *map[string]int) {
	(*m)["aa"] = 1111
}

func test_map3(m interface{}) {
	switch mm := m.(type) {
	case map[string]int:
		mm["aa"] = 1111
		break

	}

}
func test_map2(m map[string]int) {
	m["aa"] = 1234
}

//https://blog.csdn.net/cyk2396/article/details/78890185
func test_map1() {
	a := map[string]int{"aa": 1}
	test_map2(a) //a.aa 被修改, 这看似是传引用，但其实map也是传值的，它的原理和数组切片类似。map内部维护着一个指针，该指针指向真正的map存储空间
	fmt.Println("2 a.aa :", a["aa"])
	test_map3(a) //a.aa 被修改,这看似是传引用，但其实map也是传值的，它的原理和数组切片类似。map内部维护着一个指针，该指针指向真正的map存储空间
	fmt.Println("3 a.aa:", a["aa"])

	//传的引用
	test_map4(&a)
	fmt.Println("4 a.aa:", a["aa"])

	//传的引用
	test_map5(&a)
	fmt.Println("5 a.aa:", a["aa"])

	//传的引用
	test_map5(&a)
	fmt.Println("6 a.aa:", a["aa"])

	b := map[string]map[string]int{"kk": {"aa": 1}}

	c := b["kk"]
	c["aa"] = 123

	fmt.Println("b.kk.aa:", b["kk"]["aa"])

	b["kk"] = nil
	b["kk"] = map[string]int{"kk": 123}
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

func test_reflect3(vv interface{}) {
	rrv := reflect.ValueOf(vv)
	rrvk := rrv.Kind()
	rrve := rrv.Elem()

	//rrrr := rrve.Pointer()
	println(rrvk.String(), rrve.String())

	rrt := reflect.TypeOf(vv)
	rv := reflect.ValueOf(rrv)
	//println("valueof rv:", rv)
	name1 := rrt.Field(0).Name
	name2 := rrt.Field(1).Name
	println(name1, name2)
	rt := reflect.TypeOf(rrv)
	for i := 0; i < rv.NumField(); i++ {
		name := rt.Field(i).Name
		tp := rv.Field(i).Type()
		if 65 > name[0] || name[0] > 95 {
			println(" cant export key:", name)
			continue
		}
		vif := rv.Field(i).Interface() //必须是可以导出的字段
		fmt.Printf("name:%s,type:%s, value:%v \n", name, tp, vif)
	}

	s := rv.FieldByName("s")
	fmt.Printf("%T %v\n", s.String(), s.String())

	i := rv.FieldByName("i")
	fmt.Printf("%T %v\n", i.Int(), i.Int())

	f := rv.FieldByName("f")
	fmt.Printf("%T %v\n", f.Float(), f.Float())
	//rv.FieldByName("s").SetString("aaaa")

}
func test_reflect2(m interface{}) {
	vvv, ok := m.(*Circle1)
	if ok {
		vvv.Radius = 34567 //修改的是局部变量
		//return
	}

	vv := *vvv
	vv.f = 123
	vv.i = 321
	rv := reflect.ValueOf(vv)
	//println("valueof rv:", rv)

	rt := reflect.TypeOf(vv)
	////https://blog.csdn.net/raoxiaoya/article/details/111181104
	for i := 0; i < rv.NumField(); i++ {
		name := rt.Field(i).Name
		tp := rv.Field(i).Type()
		//kvt := rv.Field(i)
		//if kvt.IsNil() {
		//	continue
		//}

		if 65 > name[0] || name[0] > 95 {
			println(" cant export key:", name)
			continue
		}
		//println(kvt.String())
		vif := rv.Field(i).Interface() //必须是可以导出的字段
		fmt.Printf("name:%s,type:%s, value:%v \n", name, tp, vif)
		//fmt.Printf("name:%s,type:%s \n", name, tp)
	}

	s := rv.FieldByName("s")
	fmt.Printf("%T %v\n", s.String(), s.String())

	i := rv.FieldByName("i")
	fmt.Printf("%T %v\n", i.Int(), i.Int())

	f := rv.FieldByName("f")
	fmt.Printf("%T %v\n", f.Float(), f.Float())

}

// 测试unit
//func TestReflect(t *testing.T)  {
// reflectNew((*A)(nil))
//}

//反射创建新对象
func test_reflect4(target interface{}) {

	a := Circle{}
	p := &a
	pp := &p
	ppp := &pp
	fmt.Println("ppp:", ppp)
	if target == nil {
		fmt.Println("参数不能未空")
		return
	}

	t := reflect.TypeOf(target)
	if t.Kind() == reflect.Ptr { //指针类型获取真正type需要调用Elem
		t = t.Elem()
	}
	fmt.Println("kind:", t.Kind().String())

	newStruc := reflect.New(t) // 调用反射创建对象

	v := newStruc.Elem()
	vv := v.FieldByName("SS") //必须要能导出的字段
	vv.SetString("xxxx")
	//newStruc.Elem().FieldByName("s").SetString("Lily") //设置值

	newVal1 := newStruc.Elem().FieldByName("s") //获取值
	fmt.Println(newVal1.String())
	newVal2 := newStruc.Elem().FieldByName("SS") //获取值
	fmt.Println(newVal2.String())
	test_reflect5(v.Interface().(Circle1)) //使用
	test_reflect6(v.Interface())           //使用
}

func test_reflect5(a1 Circle1) {
	fmt.Println("Circle1.SS:", a1.SS)
}
func test_reflect6(a1 interface{}) {
	a := a1.(Circle1)
	fmt.Println("Circle1.SS:", a.SS)

}

func test_reflect() {
	t := Circle1{}

	//test_reflect1()

	test_reflect4(t)
}
func test_reflect1() {

	//t:= reflect.TypeOf((*cellnet.SessionConnected)(nil)).Elem()
	t := reflect.TypeOf((*Foo1)(nil))

	// Elem() returns a type's element type.
	// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
	e := t.Elem()

	k := t.Kind()

	fmt.Println("e:", e, "k:", k)

	tt := reflect.TypeOf((*int32)(nil))
	ee := tt.Elem()
	kk := tt.Kind()
	fmt.Println("ee:", ee, "kk:", kk)

	a := Circle1{}
	aa := reflect.TypeOf(&a)

	aaa := aa.Elem()

	fmt.Println(aaa)

	a.Radius = 12345

	//aaaa := reflect.ValueOf(a)
	//aaaai := aaaa.Interface()
	//test_reflect2(&a)

	fmt.Println(a)
	//aaa
}

func test_json() {
	a := map[string]interface{}{
		"key": 11,
		"GO":  22,
	}

	jstr, err := json.Marshal(a)
	if err != nil {
		fmt.Println("obj => JSON str ERR:", err.Error())
		return
	}
	fmt.Println("obj => JSON str :", string(jstr))

	//b := []byte{}
	b := []byte(`{"key":456,"GO":567}`)
	m := map[string]interface{}{}
	err1 := json.Unmarshal(b, &m)
	//err1 := json.Unmarshal(b, m) //err, 内部有判断，需要传入指针
	if err1 != nil {
		fmt.Println("JSON str =>obj ERR:", err1.Error())
		return
	}
	fmt.Println("JSON str =>obj:", m)
}

type tttt struct {
	a int32
}

type ttttSt struct {
	tmap *map[int32]*[]*tttt
	tary *[]*tttt
	tm   int64
}

func test_remember() {

}
func TestSingle() {
	//test_range()
	//test_struct()
	//test_map()
	//test_map1()
	//test_interface()
	//aaaa(123)

	//testgogo()
	//test_interface1()
	test_reflect()
	//test_json()

	fmt.Println("main end")
}
