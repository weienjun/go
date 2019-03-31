package main
import (
	"fmt"	
	"reflect"
)

//空结构体作为通道，用于事件通知
func emptyStruct(){
	exit := make(chan struct{})

	go func(){
		println("empty struct func")
		exit <- struct{}{}
	}()
	
	<- exit
	println("end.")
}

func anonymousData(){
	type attr struct{
		perm int
		name string		//与外部字段名相同，必须显示访问
	}
	type file struct{
		name string
		attr	//匿名字段
	}
	
	f := file{
		name : "aaa",
		attr : attr{	//初始化匿名字段
			perm : 11,
			name : "ab",
		},
	}

	f.perm = 100	//调用匿名字段
	f.attr.name = "aabb"	//与外部字段名相同，必须显示访问
	fmt.Println(f)
	
}

//标签
func testTag(){
	type user struct{
		name string		`姓名`	//打标签
		age int		`年龄`
	}

	u := user{"aa",111}
	v := reflect.ValueOf(u)
	t := v.Type()

	for i , n := 0,t.NumField(); i < n; i++{
		fmt.Printf("%s: %v\n",t.Field(i).Tag,v.Field(i))
	}
}

func main(){
	type user struct{
		name string
		age int
	}
	people1 := user{"Tom",11}	//按顺序初始化
	//people2 := user{"Str"}	//不可初始化部分参数
	fmt.Println(people1)
	people2 := user{			//推荐按命名初始化，防止出错
		name : "Str",
		age : 12,
	}
	fmt.Println(people2)

	type file struct {
		name string
		attr struct {	//匿名结构体
			owner int
			perm int
		}
	}
	f := file{
		name : "book",
		//	attr :{		//注:匿名结构体不可按此方法初始化
		//		owner : 1,
		//		perm : 2,
		//	}
	}
	f.attr.owner = 1	//初始化匿名结构体
	f.attr.perm = 2
	fmt.Println(f)

	//指针修改结构体成员
	type teacher struct{
		name string
		age int
	}
	t1 := teacher{
		name : "aa",
		age : 22,
	}
	p1 := &t1		//一级指针，可修改成员
	p1.age = 33
	fmt.Println(t1,*p1)
	t2 := &teacher{		//一级指针
		name : "bb",
		age : 19,
	}
	p2 := &t2			//二级指针
	//*p2.age = 44		//多及指针不可修改成员
	fmt.Println(t2,*p2)

	//空结构体
	var ks [10]struct{}
	fk := ks[:]		//空结构可操作元素
	ks[1] = struct{}{}	//空结构体可操作元素，ks[1]存空结构体类型
	fk[2] = struct{}{}
	fmt.Println(fk[2],len(ks),cap(ks))

	//空结构体作为通道
	emptyStruct()

	//匿名字段
	anonymousData()
	
	//标签
	testTag();
}
