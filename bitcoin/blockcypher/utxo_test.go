package blockcypher

import (
	"fmt"
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
