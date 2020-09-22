package card

import (
	"bank/pkg/bank/types"
	"fmt"
	//"bank/pkg/bank/card"
)

func ExampleWithdraw_positive() {

	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)

	//Output:
	//1000000
}

func ExampleWithdraw_noMoney() {

	result := Withdraw(types.Card{Balance: 0, Active: true}, 10_000_00)
	fmt.Println(result.Balance)

	//Output:
	//0
}

func ExampleWithdraw_inactive() {

	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)
	fmt.Println(result.Balance)

	//Output:
	//2000000
}

func ExampleWithdraw_limit() {

	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 20_000_00)
	fmt.Println(result.Balance)

	//Output:
	//2000000
}

func ExampleDeposit_positive() {
	result := types.Card{Balance: 30_000_00, Active: true}
	Deposit(&result, 30_000_00)

	fmt.Println(result.Balance)

	//Output:
	//6000000
}

func ExampleDeposit_inactive() {
	result := types.Card{Balance: 30_000_00, Active: false}
	Deposit(&result, 30_000_00)

	fmt.Println(result.Balance)

	//Output:
	//3000000
}

func ExampleDeposit_limit() {
	result := types.Card{Balance: 30_000_00, Active: false}
	Deposit(&result, 60_000_00)

	fmt.Println(result.Balance)

	//Output:
	//3000000
}

func ExampleAddBonus_positive() {
	card := types.Card{Active: true, Balance: 30_000_00, MinBalance: 10_000_00}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)

	//Output:
	//3002465
}

func ExampleAddBonus_inActive() {
	card := types.Card{Active: false, Balance: 30_000_00, MinBalance: 10_000_00}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)

	//Output:
	//3000000
}

func ExampleAddBonus_negativeBalance() {
	card := types.Card{Active: true, Balance: -30_000_00, MinBalance: 10_000_00}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)

	//Output:
	//-3000000
}

func ExampleAddBonus_limit() {
	card := types.Card{Active: true, Balance: 30_000_00, MinBalance: 3_000_000_00}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)

	//Output:
	//3000000
}

func ExampleTotal_positive() {
	cards := []types.Card{
		{
			Balance: 30_000_00,
			Active:  true,
		},
		{
			Balance: 30_000_00,
			Active:  true,
		},
		{
			Balance: 30_000_00,
			Active:  true,
		},
	}

	result := Total(cards)

	fmt.Println(result)

	//Output:
	//9000000
}

func ExampleTotal_negative() {
	cards := []types.Card{
		{
			Balance: -30_000_00,
			Active:  true,
		},
		{
			Balance: 30_000_00,
			Active:  true,
		},
	}

	result := Total(cards)

	fmt.Println(result)

	//Output:
	//3000000
}

func ExampleTotal_noMoney() {
	cards := []types.Card{
		{
			Balance: 30_000_00,
			Active:  true,
		},
		{
			Balance: 0,
			Active:  true,
		},
	}

	result := Total(cards)

	fmt.Println(result)

	//Output:
	//3000000
}

func ExampleTotal_inActive() {
	cards := []types.Card{
		{
			Balance: 30_000_00,
			Active:  false,
		},
		{
			Balance: 0,
			Active:  true,
		},
	}

	result := Total(cards)

	fmt.Println(result)

	//Output:
	//0
}

func ExamplePaymentSources(){
	cards := []types.Card{
		{
			Balance: 30_000_00,
			Active:  true,
			Currency: "TJS",
			PAN: "1111 1111 1111 1111",
		},
		{
			Balance: 0,
			Active:  true,
			Currency: "TJS",
			PAN: "1111 1111 1111 1112",
		},
		{
			Balance: 10_000_00,
			Active:  true,
			Currency: "TJS",
			PAN: "1111 1111 1111 1113",
		},
		{
			Balance: 10_000_00,
			Active:  false,
			Currency: "TJS",
			PAN: "1111 1111 1111 1114",
		},
	}
	
	result := PaymentSources(cards)

	for _, paymentSource := range result{
		fmt.Println(paymentSource.Number)
	}

	//Output:
	//1111 1111 1111 1111
	//1111 1111 1111 1113
}


