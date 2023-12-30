package mutex

import "sync"

type Bank struct {
	WellsFargo map[string]int
}

// Doing this in two steps on purpose to create a situation
// where a thread reads can become dirty
// if I just did b.WellsFargo[account] = b.WellsFargo[account] + amount
// the ending balances would be correct, because there is no read-in-state
// to become out of date
func (b *Bank) ApplyTransaction(account string, amount int) {
	b.WellsFargo[account] = amount
}

func (b *Bank) GetBalance(account string) int {
	return b.WellsFargo[account]
}

type IMerchant interface {
	RunCharges(*sync.WaitGroup)
	RunChargesMT(*sync.Mutex, *sync.WaitGroup)
}

type MerchantAccount struct {
	ID          string
	ChargeTotal int
}

type Merchant struct {
	Accounts []MerchantAccount
	Bank     *Bank
	Name     string
}
