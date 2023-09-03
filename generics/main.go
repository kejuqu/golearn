package main

import "fmt"

// learn link: https://go.dev/doc/tutorial/generics

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}

	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}

	return s
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64 as types for map values.
// Go requires that map keys be comparable. comparable constraint is predeclared in Go
// func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
// 	var s V

// 	for _, v := range m {
// 		s += v
// 	}
// 	return s
// }

// Declare a type constraint 声明一个类型约束
// 约束允许任何类型实现该接口
// 如果使用声明类型约束接口，然后将其与反省函数中的类型参数一起使用，则调用函数的类型参数必须具有所有的这些方法
// 约束接口也可以推断出指定的类型
type Number interface { // 定义 Number 接口类型来作为 类型约束
	int64 | float64
}

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V

	for _, v := range m {
		s += v
	}
	return s
}

// use type constraint in Generic

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n", SumInts(ints), SumFloats(floats))

	// 	go 语言的范型是 声明在 [] 中的，这个和其他语言的 <> 有点不同
	fmt.Printf("Generic Sums: %v and %v\n", SumIntsOrFloats[string, int64](ints), SumIntsOrFloats[string, float64](floats))
	// 当然，因为 go 编译器 有类型推导，所以 [] 也不需要传递
	fmt.Printf("Generic Sums: %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
}
