package card

import (
	"bank/pkg/bank/types"
)

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

const withdrawLimit = 20_000_00

//WithDraw withdraw amount from the card balance
func Withdraw(card types.Card, amount types.Money) types.Card {

	if !card.Active {
		return card
	}
	if amount > card.Balance {
		return card
	}
	if amount > withdrawLimit {
		return card
	}
	if amount <= 0 {
		return card
	}

	card.Balance -= amount
	return card
}

const depositLimit = 50_000_00

//Deposit money transfer to card balance
func Deposit(card *types.Card, amount types.Money) {

	if !card.Active {
		return
	}

	if amount <= 0 {
		return
	}
	if amount > depositLimit {
		return
	}

	card.Balance += amount

}

const bonusLimit = 5_000_00

//AddBonus bonus for minimum balance in card
func AddBonus(card *types.Card, precent int, daysInMonth int, daysInYear int) {
	if !card.Active {
		return
	}

	if card.Balance <= 0 {
		return
	}

	bonus := int(card.MinBalance) * precent / 100 * daysInMonth / daysInYear

	if bonus <= 0 || bonus > bonusLimit {
		return
	}

	/*if bonus > bonusLimit {
		bonus = bonusLimit
	}*/

	Deposit(card, types.Money(bonus))
}

//Total sum of all balance of the acive cards
func Total(cards []types.Card) types.Money {

	result := types.Money(0)
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			result += card.Balance
		}
	}

	return result
}

func PaymentSources(cards []types.Card) []types.PaymentSource{

	result := []types.PaymentSource{}
	for _, card := range cards {
		if card.Active && card.Balance >0 {
			newItem := types.PaymentSource{Type:string(card.Currency), Balance: card.Balance,  Number:string(card.PAN),}
			result =append(result, newItem)
		}
	}

	return result
}
