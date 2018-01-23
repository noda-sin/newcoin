package main

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
)

func getDbPath() string {
	dbpath := os.Getenv("DB_PATH")
	if len(dbpath) == 0 {
		dbpath = "coins.json"
	}
	return dbpath
}

func notifyNewCoin() error {
	logger, _ := zap.NewProduction()

	oldCoins, err := readOldCoins(getDbPath())
	if err != nil {
		oldCoins = make([]string, 0)
	}

	newCoins, err := getNewCoins()

	if err != nil {
		logger.Error("Fail to get new coin from coinmarketcap", zap.Error(err))
		return nil
	}

	for i := 0; i < len(newCoins); i++ {
		newCoin := newCoins[i]
		if contains(oldCoins, newCoin.Symbol) {
			continue
		}

		text := fmt.Sprintf("Added new coin in coinmarketcap!!\nSymbol:%s\nPrice:$%f", newCoin.Symbol, newCoin.Price)
		err = postTweet(text)

		if err != nil {
			logger.Error("Fail to post tweet", zap.String("Symbol", newCoin.Symbol), zap.Error(err))
			continue
		}

		oldCoins = append(oldCoins, newCoin.Symbol)

		time.Sleep(5 * time.Second)
	}

	writeOldCoins(getDbPath(), oldCoins)
	return nil
}
