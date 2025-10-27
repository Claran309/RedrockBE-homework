package main

import (
	"fmt"
)

func main() {
	fmt.Printf("欢迎使用Go语言计算器！\n请输入两个整数和一个操作符，进行四则运算\n输入exit退出程序\n")
	var s string
	var n1, n2, ans int
	var ifF bool
	var ansf float64
	var m rune
	for s != "exit" {
		fmt.Printf("请输入第一个整数：")
		fmt.Scan(&n1)
		fmt.Printf("\n请输入操作符：")
		fmt.Scanf("%c", &m)
		fmt.Printf("\n请输入第二个整数：")
		fmt.Scan(&n2)
		switch m {
		case '+':
			ans = n1 + n2
		case '-':
			ans = n1 - n2
		case '*':
			ans = n1 * n2
		case '/':
			ifF = true
			ansf = float64(n1) / float64(n2)
		}
		if ifF == false {
			fmt.Printf("\n%d%c%d=%d", n1, m, n2, ans)
		} else {
			fmt.Printf("\n%d%c%d=%.2f", n1, m, n2, ansf)
		}
		ans = 0
		ansf = 0
		ifF = false
		fmt.Printf("\n是否继续?(exit退出):")
		fmt.Scan(&s)
	}
	fmt.Printf("\n感谢使用，再见！")
}
