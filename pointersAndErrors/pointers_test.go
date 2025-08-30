package pointersanderrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	gotBalance := wallet.Balance()
	want := 10
	fmt.Printf("address of balance in test is %p \n", &wallet.balance)
	if gotBalance != want {
		t.Errorf("got %d want %d", gotBalance, want)
	}
}
