package blockchain

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	address    common.Address
}

func (w *Wallet) GetPrivateKeyBytes() []byte {
	return crypto.FromECDSA(w.privateKey)
}

func (w *Wallet) GetPublicKeyBytes() []byte {
	return crypto.FromECDSAPub(&w.privateKey.PublicKey)
}

func (w *Wallet) CheckBalance() {

}

func GenerateNewWallet() Wallet {
	pk, err := crypto.GenerateKey()
	if err != nil {
		logrus.Fatal("Unable to generate private key", err)
	}

	publicKey := crypto.FromECDSAPub(&pk.PublicKey)
	return Wallet{
		privateKey: pk,
		address:    common.HexToAddress(string(publicKey)),
	}
}

func InitWallet(priavtekey []byte) Wallet {
	// do something
	return Wallet{}
}
