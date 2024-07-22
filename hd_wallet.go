package blockchain

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

type HDWallet struct {
	masterKey *hdkeychain.ExtendedKey
	index     int
}

func newWallet(mnemonic string) (*HDWallet, error) {
	fixedSeed := bip39.NewSeed(mnemonic, "")
	extendKey, err := hdkeychain.NewMaster(fixedSeed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}
	return &HDWallet{
		masterKey: extendKey,
		index:     0,
	}, nil
}

func GetNewMnemonic() string {
	panic("not implemented error")
}

func NewHDWalletWithMnemonic(mnemonic string) (*HDWallet, error) {
	return newWallet(mnemonic)
}

func NewHDWallet() (*HDWallet, error) {
	mnemonic := GetNewMnemonic()
	return newWallet(mnemonic)
}
