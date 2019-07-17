package main

import (
	"errors"
	"fmt"
)

// product
type CreditCard interface {
	GetCardType() string
	GetCreditLimit() int
	GetAnnualCharge() int
}

// concrete product
type MoneyBackCreditCard struct{}

func (m *MoneyBackCreditCard) GetCardType() string {
	return "moneyback"
}

func (m *MoneyBackCreditCard) GetCreditLimit() int {
	return 15000
}

func (m *MoneyBackCreditCard) GetAnnualCharge() int {
	return 500
}

type TitaniumCreditCard struct{}

func (t *TitaniumCreditCard) GetCardType() string {
	return "titanium"
}

func (t *TitaniumCreditCard) GetCreditLimit() int {
	return 30000
}

func (t *TitaniumCreditCard) GetAnnualCharge() int {
	return 1500
}

// creator
type CreditCardAbstractFactory interface {
	CreateCreditCard() CreditCard
}

// concrete creator
type MoneyBackCreditCardFactory struct{}

func (m *MoneyBackCreditCardFactory) CreateCreditCard() CreditCard {
	return &MoneyBackCreditCard{}
}

type TitaniumCreditCardFactory struct{}

func (t *TitaniumCreditCardFactory) CreateCreditCard() CreditCard {
	return &TitaniumCreditCard{}
}

type CreditCardFactory struct{}

func (c *CreditCardFactory) CreateObject(num int) (CreditCard, error) {
	if num == 1 {
		return new(MoneyBackCreditCardFactory).CreateCreditCard(), nil
	} else if num == 2 {
		return new(TitaniumCreditCardFactory).CreateCreditCard(), nil
	} else {
		return nil, errors.New("num is not valid")
	}
}

func main() {
	// moneyBankFactory := &MoneyBackCreditCardFactory{}
	// moneyBank := moneyBankFactory.CreateCreditCard()
	// fmt.Println("money bank card type : ", moneyBank.GetCardType())
	// fmt.Println("money bank credit limit : ", moneyBank.GetCreditLimit())
	// fmt.Println("money bank annual charge : ", moneyBank.GetAnnualCharge())

	// fmt.Println("++++++++++++++")

	// titaniumFactory := &TitaniumCreditCardFactory{}
	// titanium := titaniumFactory.CreateCreditCard()
	// fmt.Println("titanium card type : ", titanium.GetCardType())
	// fmt.Println("titanium credit limit : ", titanium.GetCreditLimit())
	// fmt.Println("titanium annual charge : ", titanium.GetAnnualCharge())

	ccFactory := &CreditCardFactory{}
	cc, err := ccFactory.CreateObject(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(cc.GetCardType())

}
