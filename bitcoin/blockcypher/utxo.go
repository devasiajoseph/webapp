package blockcypher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UTXO struct {
	TxHash   string  `json:"tx_hash"`
	Block    int     `json:"block_height"`
	Value    float64 `json:"value"`
	Script   string  `json:"script"`
	TxInputs []struct {
		PrevHash    string `json:"prev_hash"`
		OutputIndex int    `json:"output_index"`
	} `json:"txinputs"`
}

func GetUTXO(address string) {

	response, err := http.Get(fmt.Sprintf("https://api.blockcypher.com/v1/btc/main/addrs/%s?unspentOnly=true", address))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal the response into a slice of UTXOs
	var utxos []UTXO
	err = json.Unmarshal(body, &utxos)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the retrieved UTXOs
	fmt.Println(utxos)

}
