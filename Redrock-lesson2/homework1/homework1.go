package main

import "fmt"

type Product struct {
	Name  string
	Price float64
	Stock int
}

func (p Product) TotalValue() float64 {
	return p.Price * float64(p.Stock)
}

func (p Product) IsInStock() bool {
	if p.Stock > 0 {
		return true
	}
	return false
}

func (p Product) Info() string {
	return fmt.Sprintf("商品: %s, 单价: %.1f￥, 库存: %d", p.Name, p.Price, p.Stock)
}

func (p *Product) Restock(amount int) {
	p.Stock += amount
}

func (p *Product) Sell(amount int) (success bool, massage string) {
	if p.Stock >= amount {
		p.Stock -= amount
		return true, "售卖成功"
	}
	return false, "库存不足"
}

func main() {
	var product Product = Product{"Go编程书", 89.5, 10}
	if suc, _ := product.Sell(5); suc == true {
		fmt.Printf("成功！剩余库存：%d\n", product.Stock)
	} else {
		fmt.Printf("失败，库存不足！\n")
	}
	product.Restock(20)
	fmt.Printf("进货20本，当前库存：%d\n", product.Stock)
	if suc, _ := product.Sell(30); suc == true {
		fmt.Printf("成功！剩余库存：%d\n", product.Stock)
	} else {
		fmt.Printf("失败，库存不足！\n")
	}
	fmt.Printf("\n商品信息：\n")
	fmt.Println(product.Info())
	fmt.Printf("库存总价值：%.2f￥", product.TotalValue())
}
