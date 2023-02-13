package bitcoin

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

func KeyToWIF(privateKeyStr string) {

	// Decode the private key from string
	privateKey, _ := btcec.PrivKeyFromBytes([]byte(privateKeyStr))

	// Create a WIF from the private key bytes
	wif, err := btcutil.NewWIF(privateKey, &chaincfg.MainNetParams, true)
	if err != nil {
		fmt.Println("Error creating WIF:", err)
		return
	}

	// Print the WIF
	fmt.Println("Private Key WIF:", wif.String())
}
