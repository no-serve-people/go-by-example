// _switch_ ，方便的条件分支语句。
package main

import (
	"fmt"
	"time"
)

// 一个基本的 `switch`。
func main() {
	i := 2
	fmt.Println("write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	// 在一个 `case` 语句中，你可以使用逗号来分隔多个表
	// 达式。在这个例子中，我们很好的使用了可选的
	// `default` 分支。
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it is the weekend")
	default:
		fmt.Println("it is weekday")
	}
	// 不带表达式的 `switch` 是实现 if/else 逻辑的另一种
	// 方式。这里展示了 `case` 表达式是如何使用非常量的。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it is before noon")
	default:
		fmt.Println("it is after noon")
	}
}
