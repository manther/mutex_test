package mutex

import "sync"

type Bank struct {
	WellsFargo map[string]int
}

func (b *Bank) ApplyTransaction(account string, amount int) {
	b.WellsFargo[account] = b.WellsFargo[account] + amount
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
	Bank *Bank
	Name string
}
