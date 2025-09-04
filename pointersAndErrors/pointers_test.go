package pointersanderrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	gotBalance := wallet.Balance()
	// var want Bitcoin = 10
	want := Bitcoin(10)
	fmt.Printf("address of balance in test is %p \n", &wallet.balance)
	if gotBalance != want {
		t.Errorf("got %s want %d", gotBalance, want)
	}
}
