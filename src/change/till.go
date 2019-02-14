package change

import (
	"math"
)

// Result ...
type Result struct {
	status string // is till open, closed or no money left
	change []Coin // list of coins with value
}

// Coins ...
type Coins struct {
	value   float64 // coin base value
	amount  float64 // amount of type coin in till
	returns float64 // how much of this coin to return
}

// Coin ...
type Coin struct {
	name     string  // name of the coin
	sumTotal float64 // amount of them
}

// CashRegister ... freecodecamp.org Cash Register
func CashRegister(price float64, cash float64, cashInDraw []Coin) *Result {
	change := cash - price
	totalCash := 0.0

	result := &Result{status: "OPEN", change: make([]Coin, 0)}

	coinNames := []string{"PENNY", "NICKEL", "DIME", "QUARTER", "ONE", "FIVE", "TEN", "TWENTY", "ONE HUNDRED"}

	coins := map[string]*Coins{
		"PENNY":       &Coins{value: 0.01, amount: 0, returns: 0},
		"NICKEL":      &Coins{value: 0.05, amount: 0, returns: 0},
		"DIME":        &Coins{value: 0.10, amount: 0, returns: 0},
		"QUARTER":     &Coins{value: 0.25, amount: 0, returns: 0},
		"ONE":         &Coins{value: 1.0, amount: 0, returns: 0},
		"FIVE":        &Coins{value: 5.0, amount: 0, returns: 0},
		"TEN":         &Coins{value: 10.0, amount: 0, returns: 0},
		"TWENTY":      &Coins{value: 20.0, amount: 0, returns: 0},
		"ONE HUNDRED": &Coins{value: 100.0, amount: 0, returns: 0},
	}

	// SETUP --
	for idx := range cashInDraw {
		name := cashInDraw[idx].name
		sumTotal := cashInDraw[idx].sumTotal

		if _, ok := coins[name]; ok {
			coins[name].amount = math.Round(sumTotal / coins[name].value)
			totalCash = math.Round((totalCash+sumTotal)*100) / 100
		}
	}

	//  NOT ENOUGH
	if change > totalCash {
		result.status = "INSUFFICIENT_FUNDS"
		return result
	}

	// SAME
	if change == totalCash {
		result.status = "CLOSED"

		for _, c := range coinNames {

			if coins[c].amount >= 0.01 {
				v := math.Round((coins[c].value*coins[c].amount)*100) / 100
				result.change = append(result.change, Coin{name: c, sumTotal: v})
			} else {
				result.change = append(result.change, Coin{name: c})
			}

		}
		return result
	}

	// GET CHANGE
	idx := len(coinNames) - 1

	for change > 0.0 {
		currentCoin := coinNames[idx]

		if coins[currentCoin].amount > 0.0 && idx >= 0 {

			calculate := math.Round((change-coins[currentCoin].value)*100) / 100

			// target
			if calculate == 0 {
				coins[currentCoin].returns += coins[currentCoin].value
				coins[currentCoin].amount--
				// hasChange = true
				break
			}

			// go to next coin or update
			if calculate < 0 && idx >= 0 {
				idx--
			} else {
				coins[currentCoin].returns += coins[currentCoin].value
				coins[currentCoin].amount--
				change = calculate
			}

		} else {
			idx--
		}
	}

	for _, c := range coinNames {
		if coins[c].returns > 0 {
			result.change = append(result.change, Coin{name: c, sumTotal: coins[c].returns})
		}
	}

	return result
}
