package bitcoin

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

func PrivateKeyToAddress(privateKeyStr string) (string, error) {
	addr := ""

	// Decode the private key string to bytes
	privateKeyBytes, err := hex.DecodeString(privateKeyStr)
	if err != nil {
		return addr, err
	}

	// Create a private key from the bytes
	privateKey, _ := btcec.PrivKeyFromBytes(privateKeyBytes)

	// Get the public key from the private key
	publicKey := privateKey.PubKey()

	// Create a new address from the public key
	address, err := btcutil.NewAddressPubKey(publicKey.SerializeCompressed(), &chaincfg.MainNetParams)
	if err != nil {
		return addr, err
	}

	// Print the address
	return address.EncodeAddress(), err
}
