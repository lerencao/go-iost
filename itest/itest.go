package itest

import (
	"math/rand"
)

type ITest struct {
	bank    *Account
	keys    []*Key
	clients []*Client
}

func New(bank *Account, keys []*Key, clients []*Client) *ITest {
	return &ITest{
		bank:    bank,
		keys:    keys,
		clients: clients,
	}
}

func (t *ITest) CreateAccount(name string) (*Account, error) {
	kIndex := rand.Intn(len(t.keys))
	key := t.keys[kIndex]
	cIndex := rand.Intn(len(t.clients))
	client := t.clients[cIndex]

	account, err := client.CreateAccount(name, key, t.bank)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (t *ITest) Transfer(token, sender, recipient, amount string) error {
	return nil
}

func (t *ITest) SetContract(abi, code string) error {
	return nil
}

func (t *ITest) GetTransaction(hash string) (*Transaction, error) {
	cIndex := rand.Intn(len(t.clients))
	client := t.clients[cIndex]

	transaction, err := client.GetTransaction(hash)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *ITest) GetAccount(name string) (*Account, error) {
	cIndex := rand.Intn(len(t.clients))
	client := t.clients[cIndex]

	account, err := client.GetAccount(name)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (t *ITest) SendTransaction(transaction *Transaction) (string, error) {
	cIndex := rand.Intn(len(t.clients))
	client := t.clients[cIndex]

	hash, err := client.SendTransaction(transaction)
	if err != nil {
		return "", err
	}

	return hash, nil
}
