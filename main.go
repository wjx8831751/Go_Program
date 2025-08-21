package main

import (
	"fmt"
	"strings"
)

// 定义一个Person结构体
type Person struct {
	Name string
	Age  int
}

// 为Person结构体定义一个方法
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I'm %d years old", p.Name, p.Age)
}

// 一个简单的加法函数
func add(a, b int) int {
	return a + b
}

func main() {
	// 基本输出
	fmt.Println("Hello, Go!")

	// 变量声明和使用
	var message string = "Welcome to this simple demo"
	fmt.Println(message)

	// 简短变量声明
	number := 42
	fmt.Printf("The answer is: %d\n", number)

	// 结构体实例化和使用
	person := Person{Name: "Alice", Age: 30}
	fmt.Println(person.Greet())

	// 函数调用
	sum := add(5, 7)
	fmt.Printf("5 + 7 = %d\n", sum)

	// 字符串操作
	text := "hello, world"
	fmt.Println("Uppercase:", strings.ToUpper(text))

	// 简单的循环
	fmt.Print("Counting: ")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
