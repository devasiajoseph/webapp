package blockcypher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type UTXO struct {
	Address                      string  `json:"address"`
	TotalReceived                int64   `json:"total_received"`
	TotalSent                    int64   `json:"total_sent"`
	Balance                      int64   `json:"balance"`
	UnconfirmedBalance           int64   `json:"unconfirmed_balance"`
	TransactionsNumber           int     `json:"n_tx"`
	UnconfirmedNumberTransaction int     `json:"unconfirmed_n_tx"`
	FinalNumberTransactions      int     `json:"final_n_tx"`
	TxRefs                       []TxRef `json:"txrefs"`
}

type TxRef struct {
	TxHash        string    `json:"tx_hash"`
	BlockHeight   int       `json:"block_height"`
	TxInputN      int       `json:"tx_input_n"`
	TxOutputN     int       `json:"tx_output_n"`
	Value         int64     `json:"value"`
	RefBalance    int64     `json:"ref_balance"`
	Spent         bool      `json:"spent"`
	Confirmations int       `json:"confirmations"`
	Confirmed     time.Time `json:"confirmed"`
	DoubleSpend   bool      `json:"double_spend"`
}

func GetUTXO(address string) (UTXO, error) {
	var utxo UTXO
	response, err := http.Get(fmt.Sprintf(unspentApi, address, 1))
	if err != nil {
		fmt.Println(err)
		return utxo, err
	}

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return utxo, err
	}

	// Unmarshal the response into a slice of UTXOs
	err = json.Unmarshal(body, &utxo)
	if err != nil {
		fmt.Println(err)
		return utxo, err
	}

	return utxo, err

}
