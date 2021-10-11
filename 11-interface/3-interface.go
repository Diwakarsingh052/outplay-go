package main

import "fmt"

type Expense interface {
	CalculateSal() int
}

type perm struct {
	basicPay int
	pf       int
}

type contract struct {
	pay int
}

func (p perm) CalculateSal() int {
	return p.basicPay + p.pf
}

func (c contract) CalculateSal() int {
	return c.pay
}

func TotalExpense(e ...Expense) {

	sum := 0
	for _, v := range e {
		sum = sum + v.CalculateSal()
	}
	fmt.Println(sum)

}

//func TotalExpense(c ...contract) {
//
//}
//func TotalExpense1(p ...perm) {
//
//}

func main() {
	p := perm{
		basicPay: 100,
		pf:       100,
	}

	c := contract{pay: 100}

	p1 := perm{
		basicPay: 100,
		pf:       200,
	}

	TotalExpense(p, c, p1)

}
