package main

import "fmt"

func main() {
	//	var mp0 map[string]int
	mp1 := make(map[string]int) //初始化map
	mp1["ten"] = 10
	fmt.Println("ten: ", mp1["ten"])

	mp2 := map[string]float32{"C": 5, "GO": 4.5, "JAVA": 6} //初始化map
	valueOfGo, hasGo := mp2["GO"]                           //从map取元素，两个返回值：value, 是否存在.
	if hasGo {
		fmt.Println("Value of 'Go': ", valueOfGo)
	} else {
		fmt.Println("Go's not in the map.")
	}
	delete(mp2, "C") //删除元素
	valueOfC, hasC := mp2["C"]
	if hasC {
		fmt.Println("Value of 'C': ", valueOfC)
	} else {
		fmt.Println("'C' is not in the map.")
	}
}
