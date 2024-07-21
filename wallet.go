package blockchain

import (
	"context"
	"crypto/ecdsa"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/sirupsen/logrus"
)

type SimpleWallet struct {
	privateKey *ecdsa.PrivateKey
	address    common.Address
}

func (w *SimpleWallet) GetPrivateKeyBytes() []byte {
	return crypto.FromECDSA(w.privateKey)
}

func (w *SimpleWallet) GetPublicKeyBytes() []byte {
	pubKey := crypto.FromECDSAPub(&w.privateKey.PublicKey)
	return pubKey[1:] // 64 to 64 bytes
}

func (w *SimpleWallet) GetPrivateKeyHex() string {
	return encodeHex(crypto.FromECDSA(w.privateKey))
}

func (w *SimpleWallet) GetPublicKeyHex() string {
	pubKey := crypto.FromECDSAPub(&w.privateKey.PublicKey)
	return encodeHex(pubKey[1:]) // 64 to 64 bytes
}

func (w *SimpleWallet) GetAddressHex() string {
	return w.address.Hex()
}

func (w *SimpleWallet) GetAddressString() []byte {
	return w.address.Bytes()
}

func (w *SimpleWallet) GetBalanceWei() (*big.Int, error) {
	client = GetClient()
	balance, err := client.BalanceAt(context.Background(), w.address, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (w *SimpleWallet) GetBalanceEth() (*big.Float, error) {
	client = GetClient()
	balance, err := client.BalanceAt(context.Background(), w.address, nil)
	if err != nil {
		return nil, err
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethValue, nil
}

func NewWallet() SimpleWallet {
	pk, err := crypto.GenerateKey()
	if err != nil {
		logrus.Fatal("Unable to generate private key", err)
	}
	return SimpleWallet{
		privateKey: pk,
		address:    crypto.PubkeyToAddress(pk.PublicKey),
	}
}

func InitWallet(privateKey string) (SimpleWallet, error) {

	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return *new(SimpleWallet), err
	}
	/**
	 * when converting public key to Address, there are actually three steps
	 * 1. original public key is 65 bytes long, as we need to abandon the first 0x4a (in hex encoding) to 64 bytes
	 * 	  key = key[1:]
	 *
	 * 2. the actual address is a Keccak256 hashed value with 20 bytes long ,
	 *    so get the value via hashVal := keccak256(key)
	 *
	 * 3. the hashed value is 32 bytes long, so get the value
	 * 	  via address := hashVal[12:]
	 *
	 * The above step is what actually crypto.PubkeyToAddress does.
	 */

	return SimpleWallet{
		privateKey: key,
		address:    crypto.PubkeyToAddress(key.PublicKey),
	}, nil
}

func encodeHex(key []byte) string {
	return hexutil.Encode(key)
}

func decodeHex(key string) ([]byte, error) {
	return hexutil.Decode(key)
}
