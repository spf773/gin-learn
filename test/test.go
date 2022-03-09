package main

import "fmt"

func main() {

	//*和&的区别 :
	//
	//    & 是取地址符号 , 即取得某个变量的地址 , 如 ; &a
	//    *是指针运算符 , 可以表示一个变量是指针类型 , 也可以表示一个指针变量所指向的存储单元 , 也就是这个地址所存储的值 .
	fmt.Println(&Rect{100, 100}) //&{100 100}

	//两种声明方式结果一致
	//var r *Rect = &Rect{100, 100}
	r := &Rect{100, 100}
	fmt.Println(r)  //&{100 100}
	fmt.Println(*r) //{100 100}
	fmt.Println(&r) //0xc000006030
}

type Rect struct {
	width  float64
	height float64
}

func (r *Rect) size() float64 {
	return r.width * r.height
}
