package bitcoin

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

type BtcAccount struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

func NewAccount() BtcAccount {
	var btca BtcAccount
	// Create a new private key.
	// Create a new private key.
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		fmt.Println(err)
		return btca
	}

	publicKey := privateKey.PubKey()
	hash := btcutil.Hash160(publicKey.SerializeCompressed())

	segwitAddress, err := btcutil.NewAddressWitnessPubKeyHash(hash, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println("here")
		fmt.Println(err)
		return btca
	}

	fmt.Printf("Private Key (hex): %x\n", privateKey.Serialize())
	fmt.Printf("Public Key (hex): %x\n", publicKey.SerializeCompressed())
	fmt.Printf("SegWit Address: %s\n", segwitAddress.EncodeAddress())

	btca.PrivateKey = hex.EncodeToString(privateKey.Serialize())
	btca.PublicKey = hex.EncodeToString(publicKey.SerializeCompressed())
	btca.Address = segwitAddress.EncodeAddress()
	return btca
}
