package bitcoin

import (
	"fmt"
	"log"
	"testing"
)

func TestUTXO(t *testing.T) {
	cl, err := Connect("144.217.71.189:8333", "", "")

	if err != nil {
		fmt.Println(err)
		t.Errorf("Unable to connect to node")
	}

	utxo, err := GetUTXO(cl, "1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2")

	if err != nil {
		log.Println(err)
		t.Errorf("Unable to connect to node")
	}
	fmt.Println(utxo)
}

func TestUtxoApi(t *testing.T) {
	_, err := GetUTXOAPI("1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2", 20, 0)
	if err != nil {
		fmt.Println(err)
		t.Errorf("Unable to get utxos")
	}
	GetUTXOBc("1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2")
}
