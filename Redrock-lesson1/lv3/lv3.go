package main

import (
	"fmt"
)

var m = map[string]int{
	"I":  1,
	"V":  5,
	"X":  10,
	"L":  50,
	"C":  100,
	"D":  500,
	"M":  1000,
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func romanToInt(s string) int {
	var ans int
	for i := 0; i < len(s); {
		if i != len(s)-1 { // 防止越界
			check := s[i : i+2]
			if value, ok := m[check]; ok { // 存在特殊情况
				//fmt.Printf("TwoTap = %s, value = %d\n", check, value)
				ans += value
				i += 2
				continue
			}
		}
		//fmt.Print("OneTap\n")
		check := s[i : i+1]
		ans += m[check]
		i++
	}
	return ans
}

func main() {
	var roman string
	for {
		fmt.Printf("请输入罗马数字(输入exit停止):")
		fmt.Scan(&roman)
		if roman == "exit" {
			break
		}
		fmt.Printf("\n该罗马数字对应的阿拉伯数字为:%d\n", romanToInt(roman))
	}
}
