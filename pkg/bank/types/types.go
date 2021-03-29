package types

//Money the minimal money unit
type Money int64

//Currency the cod of currency
type Currency string

//code currency
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

//PAN the card number
type PAN string

//Crad info abaout paycard
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

//Paymet представляет информацю о платеже.
type Payment struct {
	ID     int
	Amount Money
}

//PaymentSource источник оплаты
type PaymentSource struct {
	Type    string // card
	Number  string // card Number
	Balance Money  // diram
}
