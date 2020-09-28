package main

import (
	"fmt"
	"gotask2-1/pkg/credit"
)

func main() {
	var sumCredit int64 = 1000000_00
	var creditRate int64 = 20
	var periodCredit int64 = 3
	fmt.Println(credit.Calculate(sumCredit, creditRate, periodCredit))
}
