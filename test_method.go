package main
import(
	"fmt"
	"sync"
	"reflect"
)
type N int
//函数：无关联状态
func testFunc(){
	println("I am func")
}
//方法：有关联状态（与对象关联）
func (n N) testMethod()string{		//方法，对象n(N)
	println("I am method")
	n++;
	fmt.Println(n)	//26,形参拷贝，不会修改对象值，
	return fmt.Sprintf("%#x",n)	//输出对象n信息
}

func (n *N)testMethodPointer(){		//传指针，不进行拷贝，可修改对象值
	*n++	//26
	fmt.Println(*n)
}

func testAnonymous(){
	type data struct{
		sync.Mutex	//匿名字段
		buf [1024]byte
	}
	
	d := data{}
	d.Lock()	//默认处理为 sync.(*Mutex).Lock()
	defer d.Unlock()
}
		

//覆盖
type user struct{}
type manger struct{
	user
}

func (user) toString()string{
	return "user"
}
func (m manger)toString()string{	//同名方法，覆盖user方法的toString
	return m.user.toString()+";manger"
}
func testCover(){
	var m manger
	println(m.toString())		//调用对象的方法，覆盖成员的方法
	println(m.user.toString())	//指定调用成员的方法
}

//方法集
type S struct{}
type T struct{
	S
}
func (S)  sVal() {}
func (*S) sPval(){}
func (T) tVal()  {}
func (*T) pTval(){}

func methodSet(a interface{}){
	t := reflect.TypeOf(a)
	println("Method",t.NumMethod())
	fmt.Println(t)
	for i,n := 0,t.NumMethod(); i < n ; i++{
		m := t.Method(i)
		println("Method2")
		fmt.Println(m.Name,m.Type)
	}
}

func testMethodSet(){
	var t T
	methodSet(t)
	println("------------------------")
	methodSet(&t)
}

//表达式
//1、expression
type E int
func (e E)testEs(){
	fmt.Printf("testEs.e: %p,%d\n",&e,e)
}

func testExpression(){
	var e E = 100
	fmt.Printf("testExpression.e: %p,%d\n",&e,e)
	f1 := E.testEs	//func(e E)
	f1(e)

	f2 := (*E).testEs
	f2(&e)

	//Method Value方式
	f3 := &e
	e++
	f4 := e.testEs		//方法不是指针类型，会复制e，所以不会改变e值

	e++
	f5 := f3.testEs		//同上复制值

	e++
	fmt.Printf("testExpression.e: %p,%v\n",f3,e)
	f4()
	f5()
}

//空对象调用
func (E)value(){}
func (*E)pointer(){}

func testEmpty(){
	var p *E
	p.pointer()		
	(*E)(nil).pointer()		//空对象调用
	(*E).pointer(nil)
	//p.value()	//错误，指针对象不可调用非指针表达式
}


func main(){
	var a N = 25
	println(a)	//25
	println(a.testMethod())
	println(a)	//25
	a.testMethodPointer()
	println(a)	//26
	ap := &a			//原a不为指针对象，ap指针可调用非指针方法及指针方法
	ap.testMethod()		//指针对象调用非指针方法，自动匹配为非指针方法，值不可修改
	ap.testMethodPointer();
	
	//var a1 *N
	//a1.testMethod()	//原为指针对象，不可调用非指针对象方法

	//匿名字段
	testAnonymous();

	//覆盖
	testCover()

	//方法集
	testMethodSet()

	//表达式
	testExpression()

	//空对象调用
	testEmpty()
}
