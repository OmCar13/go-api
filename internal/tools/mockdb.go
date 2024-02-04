package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"Omkar": {
		AuthToken: "123ABC",
		Username:  "omkar",
	},
	"Kushal": {
		AuthToken: "456DEF",
		Username:  "kushal",
	},
	"Jack": {
		AuthToken: "789GHI",
		Username:  "jack",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"Omkar": {
		Coins:    100,
		Username: "omkar",
	},
	"Kushal": {
		Coins:    200,
		Username: "kushal",
	},
	"Jack": {
		Coins:    300,
		Username: "jack",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {

	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {

	time.Sleep(time.Second * 1)

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
