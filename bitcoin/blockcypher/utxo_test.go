package blockcypher

import (
	"fmt"
	"log"
	"testing"
)

func TestUtxoApi(t *testing.T) {

	utxo, err := GetUTXO("1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2")
	if err != nil {
		t.Errorf("Error getting utxo from blockcypher api")
	}

	fmt.Println(utxo.Address)
	fmt.Println(utxo.Balance)
}

func TestBalance(t *testing.T) {
	b, err := GetBalance("1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2")
	if err != nil {
		log.Println(err)
		t.Errorf("Error getting balance")
	}

	log.Println(b)
}
