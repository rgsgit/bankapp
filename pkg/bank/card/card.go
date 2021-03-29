package card

import (
	"bank/pkg/bank/types"
)

const depositLimit = 50_000_00
const bonusLimit = 5_000_00

//IssueCard issue card
func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}

	return card

}

//Withdraw withdraw balance from cardpay
func Withdraw(card types.Card, amount types.Money) types.Card {
	if amount <= 0 {
		return card
	}
	if !card.Active {
		return card
	}
	if amount > card.Balance {
		return card
	}
	if amount >= 2000000 {
		return card
	}

	card.Balance -= amount
	return card

}

//Deposit зачисление средств на карту
func Deposit(card *types.Card, amount types.Money) {
	if amount <= 0 {
		return
	}
	if !card.Active {
		return
	}
	if amount > depositLimit {
		return
	}

	card.Balance += amount

}

//AddBonus функция добавления бонуса
func AddBonus(card *types.Card, precent int, daysInMonth int, daysInYear int) {
	if !card.Active {
		return
	}

	if card.Balance <= 0 {
		return
	}

	bonus := (int(card.MinBalance) * precent / 100) * daysInMonth / daysInYear

	if bonus > bonusLimit {
		return
	}

	//card.Balance = card.Balance + (types.Money)(bonus)
	Deposit(card, types.Money(bonus))
}

//Total вычисляет общую сумму на всех картах.
//Отрицательные суммы и суммы на всех заблокированных картах игнорируются.
func Total(cards []types.Card) types.Money {
	sum := types.Money(0)
	for _, card := range cards {
		if card.Balance > 0 && card.Active {
			sum += card.Balance
		}

	}
	return sum

}

func PaymentSources(cards []types.Card) []types.PaymentSource {

	result := []types.PaymentSource{}
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			newItem := types.PaymentSource{Type: string(card.Currency), Balance: card.Balance, Number: string(card.PAN)}
			result = append(result, newItem)
		}
	}

	return result
}
