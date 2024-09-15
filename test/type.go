package main

// 出力用の関数があるパッケージをインポート
import (
	"fmt"
)

func main() {
	// 数値の宣言
	var i int = 1
	/*
		最大値　最小値
		int8	127		-128
		int16	32767	-32768
		int32	2147483647	-2147483648
		int64	9223372036854775807	-9223372036854775808
	*/
	fmt.Println(i)

	// 浮動小数点数の宣言
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	// 文字列の宣言
	var s string = "Hello, World!"
	fmt.Println(s)

	//　論理値の宣言
	var t, f bool = true, false
	fmt.Println(t, f)

	// 配列の宣言
	var arr1 [3]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3[0])
}
