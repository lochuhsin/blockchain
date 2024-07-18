package main

import (
	"context"
	"fmt"
	"log"

	"blockchain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func test() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	_, err = client.BalanceAt(context.Background(), account, nil)
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
	wallet := blockchain.GenerateNewWallet()
	fmt.Println(string(wallet.GetPrivateKeyBytes()))
	fmt.Println(string(wallet.GetPublicKeyBytes()))

}
func main() {
	test2()
}
