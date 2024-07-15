package tools

import (
	"time"
)


type mockDB struct{}


var mockLoginDetails = map[string]LoginDetails {
	"simon": {
		AuthToken: "1234",
		Username: "simon",
	},
	"john": {
		AuthToken: "5678",
		Username: "john",
	},
	"jane": {
		AuthToken: "91011",
		Username: "jane",
	},
}

var mockCoinDetails = map[string]CoinDetails {
	"simon": {
		Username: "simon",
		Coins: 100,
	},
	"john": {
		Username: "john",
		Coins: 200,
	},
	"jane": {
		Username: "jane",
		Coins: 300,
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate a database query
	time.Sleep(1 * time.Second)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate a database query
	time.Sleep(1 * time.Second)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
