package main

import(
	"fmt"
	"reflect"
	"unsafe"
)

func state(format string,ptr interface{}){
	p := reflect.ValueOf(ptr).Pointer()
	h := (*uintptr)(unsafe.Pointer(p))
	fmt.Printf(format,*h)
}

func toString(bs []byte)string {
	return *(*string)(unsafe.Pointer(&bs))
}

func main(){
	str := "雨痕\x61\142\u0041"
	lens := len(str)
	println(str,"--len:",lens)

	var s string
	s = `11111
			22222
				3333`
	fmt.Printf("%s,len=%d\n",s,len(s))
	
	var s1 string = "123"
	 var s2 = "234"
	 s3 := "345"
	 s4 := s1+s2+s3
	 println(s4)
	
	s5 := "12345"
	q1 := s5[:3]	
	q2 := s5[2:4]
	q3 := s5[3:]
	println(s5,q1,q2,q3)
	
	s5 = "12雨痕"
	for i := 0; i < len(s5); i++{
		fmt.Printf("%d  [%c]\n",i,s5[i])
	}
	println();
	for i , c := range(s5){
		fmt.Printf("%d  [%c]\n",i,c)
	}
	//类型转换
	sptr := "hello word"
	state("sptr:%x\n",&sptr)
	bs := []byte(sptr)
	ps := string(bs)
	state("string to []byte, bs:%x\n",&bs)
	state("[]byte to string, ps:%x\n",&ps)
	rs := []rune(sptr)
	ps2 := string(rs)
	state("string to []rune, rs:%x\n",&rs)
	state("[]rune to string, ps2:%x\n",&ps2)
	//非安全方法
	fbs := []byte("hello word")
	fs := toString(fbs)
	fmt.Printf("fbs:%x\n",&fbs)
	fmt.Printf("fs:%x\n",&fs)
	//使用append追加到[]byte中
	var ads []byte
	abs := append(ads,"abc"...)
	fmt.Println(abs)
}
