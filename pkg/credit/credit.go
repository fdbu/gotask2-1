package credit

import (
	"math"
)

func Calculate(sumCredit, creditRate, periodCredit int64) (monthlyPayment, overPayment, totalPayment int64) {
	monthlyCreditRate := float64(creditRate) / 12 / 100
	numberPeriods := float64(periodCredit) * 12
	annuityRate := (monthlyCreditRate * math.Pow(1+monthlyCreditRate, numberPeriods)) / (math.Pow(1+monthlyCreditRate, numberPeriods) - 1)
	monthlyPayment = int64(float64(sumCredit) * annuityRate)
	totalPayment = monthlyPayment * int64(numberPeriods)
	overPayment = totalPayment - sumCredit
	return
}
