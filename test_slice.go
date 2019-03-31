package main

import(
	"fmt"
)
func main(){
	x := [...]int {1,2,3,4,5}
	c := x[2:4]		//第一个数从0算起，第二个从1算起
	fmt.Printf("%v\n",x)
	fmt.Printf("%v\n",c)	//2,4
	//println(c)	//错误，不支持输出
	for i := 0; i < len(c); i++ {
		println(c[i])
	}

	c1 := make([]int,3,5)	//len=3,cap=5
	c2 := make([]int,3)		//len=cap=3
	c3 := []int{10,20,5:30}	//按初始化元素分配空间
	fmt.Println(c1,len(c1),cap(c1))
	fmt.Println(c2,len(c2),cap(c2))
	fmt.Println(c3,len(c3),cap(c3))
	
	var c4 []int	//定义切片，并未进行初始化 为nil
	c5 := []int{}	//定义切片，同时初始化不 为nil
	println(c4 == nil, c5 == nil)

	c6 :=[][]int{	//定义切片,元素也为切片
		{1,2},
		{10,20,30},
		{100},
	}
	fmt.Println(c6[1])
	c6[2] = append(c6[2],200,300)	//取下标为2的末尾追加
	fmt.Println(c6[2])

	//append追加数据返回新切片
	c7 := make([]int,0,5)
	xc := append(c7,10)		//c7末尾追加10，返回新切片
	xc2 := append(xc,20,30,40,50,60,70)		//超出原长度，重新分配内存,一般是原来cap的二倍
	fmt.Println(c7,len(c7),cap(c7))
	fmt.Println(xc,len(xc),cap(xc))
	fmt.Println(xc2,len(xc2),cap(xc2))
	fmt.Printf("c7: %p,%v\n",&c7,c7)
	fmt.Printf("xc: %p,%v\n",&xc[0],xc)
	fmt.Printf("xc2: %p,%v\n",&xc2[0],xc2)

	c8 := []int{1,2,3,4,5,6,7,8}
	fmt.Println(c8)
	sc := c8[5:8]			//取c8中的5——8地址数据，同一地址
	cn := copy(c8[4:],sc)	//将sc中的数据拷贝到c8中,同一底层地址数据拷贝
	fmt.Println(cn,c8)
	sc2 := make([]int,6)	//新切片，不同地址拷贝
	cn = copy(sc2,c8)
	fmt.Println(sc2,c8)

	b := make([]byte,3)		//byte类型切片
	bn := copy(b,"abcde")	//将字符串拷贝到byte中
	fmt.Println(bn,b)		//97(a),98 99
}
