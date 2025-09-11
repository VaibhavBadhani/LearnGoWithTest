package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		gotBalance := wallet.Balance()
		if gotBalance != want {
			t.Errorf("got %s want %d", gotBalance, want)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Errorf("wanted a error but didnt get one")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{20}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("insufficient balance", func(t *testing.T) {
		wallet := Wallet{20}
		err := wallet.Withdraw((Bitcoin(100)))
		assertError(t, err)
		assertBalance(t, wallet, Bitcoin(20))
	})
}
