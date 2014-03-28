package main

import "fmt"

func main() {
	/*
		slice:
		array[i:j]表示array[i] - array[j-1]
		属于引用类型，随原数组改变而改变
	*/
	var array = []byte{'a', 'b', 'c', 'd', 'e'}
	var as1, as2, as3, as4, as5 []byte
	as1 = array[:3]                                 //从首个元素开始
	as2 = array[1:]                                 //直到最后一个元素
	as3 = array[:]                                  //所有元素
	as4 = append(array, 'f')                        //追加元素
	fmt.Printf("length of as1: %d\n", len(as1))     //读取长度
	fmt.Printf("capacity of as2: %d\n", cap(array)) //最大容量
	fmt.Printf("length of as3: %d\n", len(as3))
	fmt.Printf("length of as4: %d\n", len(as4))
	fmt.Printf("length of copied: %d\n", copy(as5, array[1:3])) //复制元素，返回复制的个数
	fmt.Printf("length of as5: %d\n", len(as5))
}
