package blockchain

import (
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

var client *ethclient.Client
var once sync.Once
var cMu sync.Mutex

const LOCAL_NETWORK = "http://localhost:8545"

func InitDefaultClient() {
	once.Do(func() {
		if client == nil {
			cli, err := ethclient.Dial(LOCAL_NETWORK)
			if err != nil {
				logrus.Fatal("Unable to connect to default network")
				return
			}
			client = cli
		}
	})
}

func GetClient() *ethclient.Client {
	return client
}

func InitClient(network string) {
	cMu.Lock()
	defer cMu.Unlock()
	cli, err := ethclient.Dial(network)
	if err != nil {
		logrus.Fatal("Unable to connect to default network")
		client = nil
	}
	client = cli
}
