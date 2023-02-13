package bitcoin

import (
	"crypto/sha256"
	"log"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/base58"
)

type BtcAccount struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

// doubleSHA256 calculates sha256(sha256(b)).
func DoubleSHA256(b []byte) []byte {
	first := sha256.Sum256(b)
	second := sha256.Sum256(first[:])
	return second[:]
}

func NewAccount() BtcAccount {
	var btca BtcAccount
	// Create a new private key.
	// Create a new private key.
	privateKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		log.Fatalf("failed to generate private key: %s", err)
	}

	// Encode the private key in WIF format.
	wif, err := btcutil.NewWIF(privateKey, &chaincfg.MainNetParams, true)
	if err != nil {
		log.Fatalf("failed to encode private key as WIF: %s", err)
	}

	// Get the public key from the private key.
	publicKey := privateKey.PubKey().SerializeCompressed()

	publicKeyHash := sha256.Sum256(publicKey)
	// Create the witness program for the pay-to-witness-public-key-hash (p2wpkh) address.
	witnessProgram := []byte{0x00, 0x14}
	witnessProgram = append(witnessProgram, publicKeyHash[:]...)
	// Create the address for the pay-to-witness-public-key-hash (p2wpkh) address.
	address, err := btcutil.NewAddressWitnessPubKeyHash(witnessProgram, &chaincfg.MainNetParams)
	if err != nil {
		log.Fatalf("failed to generate address: %s", err)
	}

	// Encode the address in base58check format.
	addressBase58 := base58.Encode(address.ScriptAddress())

	// Output the address.
	log.Printf("Address: %s\n", addressBase58)
	log.Printf("Private Key (WIF): %s\n", wif.String())
	return btca
}
