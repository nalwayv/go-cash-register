package change

import (
	"testing"
)

func check(resA, resB *Result) bool {

	if resA.status != resB.status {
		return false
	}

	for _, v1 := range resA.change {
		name := v1.name
		total := v1.sumTotal

		for _, v2 := range resB.change {

			// same name
			if name == v2.name {
				if total != v2.sumTotal {
					return false
				}
				break
			}
		}
	}

	return true
}

func TestChange1(t *testing.T) {

	price := 19.5
	cash := 20.0

	cashindraw := []Coin{
		{"PENNY", 1.01},
		{"NICKEL", 2.05},
		{"DIME", 3.10},
		{"QUARTER", 4.25},
		{"ONE", 90.0},
		{"FIVE", 55.0},
		{"TEN", 20.0},
		{"TWENTY", 60.0},
		{"ONE HUNDRED", 100},
	}

	result := CashRegister(price, cash, cashindraw)

	expect := &Result{
		status: "OPEN",
		change: []Coin{
			{"PENNY", 0.0},
			{"NICKEL", 0},
			{"DIME", 0},
			{"QUARTER", 0.50},
			{"ONE", 0},
			{"FIVE", 0},
			{"TEN", 0},
			{"TWENTY", 0},
			{"ONE HUNDRED", 0}}}

	if !check(result, expect) {
		t.Error("fail")
	}

}

func TestChange2(t *testing.T) {

	price := 3.26
	cash := 100.0

	cashindraw := []Coin{
		{"PENNY", 1.01},
		{"NICKEL", 2.05},
		{"DIME", 3.10},
		{"QUARTER", 4.25},
		{"ONE", 90.0},
		{"FIVE", 55.0},
		{"TEN", 20.0},
		{"TWENTY", 60.0},
		{"ONE HUNDRED", 100},
	}

	result := CashRegister(price, cash, cashindraw)

	expect := &Result{
		status: "OPEN",
		change: []Coin{
			{"PENNY", 0.04},
			{"DIME", 0.20},
			{"QUARTER", 0.50},
			{"ONE", 1},
			{"FIVE", 15},
			{"TEN", 20},
			{"TWENTY", 60}}}

	if !check(result, expect) {
		t.Error("fail")
	}

}

func TestChange3(t *testing.T) {

	price := 19.50
	cash := 20.0

	cashindraw := []Coin{
		{"PENNY", 0.01},
		{"NICKEL", 0.0},
		{"DIME", 0.0},
		{"QUARTER", 0.0},
		{"ONE", 0.0},
		{"FIVE", 0.0},
		{"TEN", 0.0},
		{"TWENTY", 0.0},
		{"ONE HUNDR ", 0.0}}

	result := CashRegister(price, cash, cashindraw)

	expect := &Result{status: "INSUFFICIENT_FUNDS"}

	if !check(result, expect) {
		t.Error("fail")
	}

}

func TestChange4(t *testing.T) {

	price := 19.5
	cash := 20.0
	cashindraw := []Coin{
		{"PENNY", 0.50},
		{"NICKEL", 0.0},
		{"DIME", 0.0},
		{"QUARTER", 0.0},
		{"ONE", 0.0},
		{"FIVE", 0.0},
		{"TEN", 0.0},
		{"TWENTY", 0.0},
		{"ONE HUNDRED", 0.0},
	}

	result := CashRegister(price, cash, cashindraw)

	expect := &Result{
		status: "CLOSED",
		change: []Coin{
			{"PENNY", 0.5},
			{"NICKEL", 0.0},
			{"DIME", 0.0},
			{"QUARTER", 0.0},
			{"ONE", 0.0},
			{"FIVE", 0.0},
			{"TEN", 0.0},
			{"TWENTY", 0.0},
			{"ONE HUNDRED", 0.0}}}

	if !check(result, expect) {
		t.Error("fail")
	}

}
