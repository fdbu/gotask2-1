package credit_test // взяли пакет credit, добавили _test

import (
	"fmt"
	"gotask2-1/pkg/credit"
)

func ExampleCalculate() { // имя функции - Example + имя проверяемой функции
	fmt.Println(credit.Calculate(1_000_000_00, 20, 3))
	fmt.Println(credit.Calculate(500_000_00, 20, 2))
	// Output:
	// 3716358 33788888 133788888
	// 2544790 11074960 61074960
}
