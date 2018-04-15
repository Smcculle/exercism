package account

import (
	"sync"
)

type Account struct {
	balance int64
	lock    sync.RWMutex
	closed  bool
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{balance: initialDeposit}
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {

	a.lock.Lock()
	defer a.lock.Unlock()

	if a.closed {
		return
	}

	if a.balance+amount < 0 {
		newBalance, ok = a.balance, false
		return
	}

	a.balance += amount
	newBalance, ok = a.balance, true

	return
}

func (a *Account) Balance() (balance int64, ok bool) {

	a.lock.RLock()
	defer a.lock.RUnlock()

	if a.closed {
		return
	}

	balance, ok = a.balance, true
	return
}

func (a *Account) Close() (payout int64, ok bool) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.closed {
		return
	}

	a.closed = true
	payout, ok = a.balance, true
	a.balance = 0
	return
}
