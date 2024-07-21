package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"blockchain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

func test() {
	client := blockchain.GetClient()
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	_, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(balance)
	// blockNumber := big.NewInt(0)
	// balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(balanceAt) // 25729324269165216042
	// fbalance := new(big.Float)
	// fbalance.SetString(balance.String())
	// ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	// fmt.Println(ethValue) // 25.729324269165216041
}

func test2() {
	wallet := blockchain.NewWallet()
	fmt.Println(len(wallet.GetPrivateKeyBytes()))
	fmt.Println(len(wallet.GetPublicKeyBytes()))
	fmt.Println(len(wallet.GetAddressString()))
	fmt.Println(wallet.GetAddressHex())
	fmt.Println(wallet.GetPublicKeyHex())
	fmt.Println(wallet.GetPrivateKeyHex())
	bigint, err := wallet.GetBalanceWei()
	if err != nil {
		fmt.Println("Unable to get wallet balance")
	} else {
		fmt.Println(bigint.Int64())
	}

}
func main() {
	var (
		network string
	)

	flag.StringVar(&network, "network", blockchain.LOCAL_NETWORK, "the network that the wallet is trying to connect to")
	flag.Parse()

	fmt.Printf("Network: %s\n", network)
	if len(network) == 0 {
		logrus.Info("No network specified, using default network", blockchain.LOCAL_NETWORK)
		blockchain.InitDefaultClient()
	} else {
		blockchain.InitClient(network)
	}
	test2()
}
