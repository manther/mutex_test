package mutex

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
	"time"
)

// Base case test. With multiple actors making transactions
// to the same bank in parellel and no syncronization
// the resulting bank balances and the expected balances
// should not match
func TestBankFail(t *testing.T) {

	// Represents a bank that multiple merchants will try to
	// execute transactions against simultaneously
	var bank = &Bank{}

	bank = &Bank{
		WellsFargo: map[string]int{"Dave": 400, "Susan": 1200, "Mike": 1000, "Steve": 300, "Jamiraqui": -32},
	}

	// The same merchant amounts can be used by all merchants.
	// its the order they are applied in that causes incorrect results.
	merchantAccounts := buildMerchAccnts()

	// Build a collection of merchants
	merchants := buildMerchants(bank, merchantAccounts)

	// This is what the ending account balances should end up as if every transaction was
	// executed in serial order
	var expectedResult = &Bank{
		WellsFargo: map[string]int{"Dave": 275, "Susan": 1700, "Mike": -500, "Steve": 525, "Jamiraqui": -5032},
	}

	// Run monthly charges accrued for customer accounts at merchant
	wg := sync.WaitGroup{}
	for _, m := range merchants {
		wg.Add(1)
		go m.merch.RunCharges(&wg)
	}
	wg.Wait()
	fmt.Println("Expected to fail. No synchronization safegaurding.")
	fmt.Println("Expected bank records at end of transaction runs: ", expectedResult)
	fmt.Println("Actual bank records at end of transaction runs:   ", bank)
	require.NotEqual(t, expectedResult, bank)
	fmt.Print("test Over.")
}

// Synchronized test.
// This test utilizes a mutex to lock down the
// Bank's customers list.
// End result should be that transactions are applied
// without any dirty reads, resulting in an expected
// end state.
func TestBankMutex(t *testing.T) {
	var mutex = &sync.Mutex{}
	// Represents a bank that multiple merchants will try to
	// execute transactions against simultaneously
	var bank = &Bank{}

	bank = &Bank{
		WellsFargo: map[string]int{"Dave": 400, "Susan": 1200, "Mike": 1000, "Steve": 300, "Jamiraqui": -32},
	}

	// The same merchant amounts can be used by all merchants.
	// its the order they are applied in that causes incorrect results.
	merchantAccounts := buildMerchAccnts()

	// Build a collection of merchants
	merchants := buildMerchants(bank, merchantAccounts)

	// This is what the ending account balances should end up as if every transaction was
	// executed in serial order
	var expectedResult = &Bank{
		WellsFargo: map[string]int{"Dave": 275, "Susan": 1700, "Mike": -500, "Steve": 525, "Jamiraqui": -5032},
	}

	wg := sync.WaitGroup{}
	for _, m := range merchants {
		wg.Add(1)
		go m.merch.RunChargesMT(mutex, &wg)
	}
	
	wg.Wait()
	fmt.Println("Expected bank records at end of transaction runs: ", expectedResult)
	fmt.Println("Actual bank records at end of transaction runs:   ", bank)
	require.Equal(t, expectedResult, bank)
}

// Test structures
// Setting up various merchants
// Each merchant fulfills the IMerchant interface
// Each merchant has a normal RunCharges, and
// a RunCharges varient that is synchronized with
// a mutex. This allow to run two similar tests
// with different end results.
type Costco struct {
	Merchant
}

func (c Costco) RunCharges(wg *sync.WaitGroup) {
	for _, account := range c.Accounts {
		time.Sleep(time.Millisecond * 100)
		c.Bank.ApplyTransaction(account.ID, account.ChargeTotal)
	}
	wg.Done()
}

func (c *Costco) RunChargesMT(mutex *sync.Mutex, wg *sync.WaitGroup) {
	for _, account := range c.Accounts {
		mutex.Lock()
		time.Sleep(time.Millisecond * 100)
		c.Bank.ApplyTransaction(account.ID, account.ChargeTotal)
		mutex.Unlock()
	}
	wg.Done()
}

type Target struct {
	Merchant
}

func (c Target) RunCharges(wg *sync.WaitGroup) {
	for _, account := range c.Accounts {
		time.Sleep(time.Millisecond * 100)
		c.Bank.ApplyTransaction(account.ID, account.ChargeTotal)
	}
	wg.Done()
}

func (c *Target) RunChargesMT(mutex *sync.Mutex, wg *sync.WaitGroup) {
	for _, account := range c.Accounts {
		mutex.Lock()
		time.Sleep(time.Millisecond * 100)
		c.Bank.ApplyTransaction(account.ID, account.ChargeTotal)
		mutex.Unlock()
	}
	wg.Done()
}

type CVS struct {
	Merchant
}

func (c CVS) RunCharges(wg *sync.WaitGroup) {
	for _, account := range c.Accounts {
		time.Sleep(time.Millisecond * 100)
		c.Bank.ApplyTransaction(account.ID, account.ChargeTotal)
	}
	wg.Done()
}

func (c *CVS) RunChargesMT(mutex *sync.Mutex, wg *sync.WaitGroup) {
	for _, account := range c.Accounts {
		mutex.Lock()
		time.Sleep(time.Millisecond * 100)
		c.Bank.ApplyTransaction(account.ID, account.ChargeTotal)
		mutex.Unlock()
	}
	wg.Done()
}

type GuitarCenter struct {
	Merchant
}

func (c GuitarCenter) RunCharges(wg *sync.WaitGroup) {
	for _, account := range c.Accounts {
		time.Sleep(time.Millisecond * 100)
		c.Bank.ApplyTransaction(account.ID, account.ChargeTotal)
	}
	wg.Done()
}

func (c *GuitarCenter) RunChargesMT(mutex *sync.Mutex, wg *sync.WaitGroup) {
	for _, account := range c.Accounts {
		mutex.Lock()
		time.Sleep(time.Millisecond * 100)
		c.Bank.ApplyTransaction(account.ID, account.ChargeTotal)
		mutex.Unlock()
	}
	wg.Done()
}

type Starbucks struct {
	Merchant
}

func (c Starbucks) RunCharges(wg *sync.WaitGroup) {
	for _, account := range c.Accounts {
		time.Sleep(time.Millisecond * 100)
		c.Bank.ApplyTransaction(account.ID, account.ChargeTotal)
	}
	wg.Done()
}

func (c *Starbucks) RunChargesMT(mutex *sync.Mutex, wg *sync.WaitGroup) {
	for _, account := range c.Accounts {
		mutex.Lock()
		time.Sleep(time.Millisecond * 100)
		c.Bank.ApplyTransaction(account.ID, account.ChargeTotal)
		mutex.Unlock()
	}
	wg.Done()
}

type merchants struct {
	merch IMerchant
}

// Test fixtures
func buildMerchAccnts() []MerchantAccount {
	return []MerchantAccount{
		{
			ID:          "Dave",
			ChargeTotal: -25,
		},
		{
			ID:          "Susan",
			ChargeTotal: 100,
		},
		{
			ID:          "Mike",
			ChargeTotal: -300,
		},
		{
			ID:          "Steve",
			ChargeTotal: 45,
		},
		{
			ID:          "Jamiraqui",
			ChargeTotal: -1000,
		},
	}
}

// A collection of merchants to iterract over and
// execute RunCharges against a single bank
// All share a common list of accounts held at the bank
func buildMerchants(bank *Bank, merchantAccounts []MerchantAccount) []merchants {
	return []merchants{
		{merch: &Costco{
			Merchant: Merchant{merchantAccounts, bank, "Costco"},
		}}, {merch: &Target{
			Merchant: Merchant{merchantAccounts, bank, "Target"},
		}}, {merch: &CVS{
			Merchant: Merchant{merchantAccounts, bank, "CVS"},
		}}, {merch: &GuitarCenter{
			Merchant: Merchant{merchantAccounts, bank, "GuitarCenter"},
		}}, {merch: &Starbucks{
			Merchant: Merchant{merchantAccounts, bank, "Starbucks"},
		}},
	}
}
