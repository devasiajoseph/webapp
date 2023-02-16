package blockcypher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var unspentApi = "https://api.blockcypher.com/v1/btc/main/addrs/%s?unspentOnly=true&page=%d"
var transactionHistoryApi = "https://api.blockcypher.com/v1/btc/main/addrs/%s?page=%d"
var transactionDetailsApi = "https://api.blockcypher.com/v1/btc/main/txs/%s"
var balanceApi = "https://api.blockcypher.com/v1/btc/main/addrs/%s"

type BlockcypherAddress struct {
	Address                         string `json:"address"`
	TotalReceived                   int64  `json:"total_received"`
	TotalSent                       int64  `json:"total_sent"`
	Balance                         int64  `json:"balance"`
	UnconfirmedBalance              int64  `json:"unconfirmed_balance"`
	FinalBalance                    int    `json:"final_balance"`
	NumberOfTransactions            int    `json:"n_tx"`
	UnconfirmedNumberOfTransactions int    `json:"unconfirmed_n_tx"`
	FinalNumberOfTransactions       int    `json:"final_n_tx"`
	Error                           bool   `json:"error"`
}

func GetBalance(address string) (BlockcypherAddress, error) {
	var data BlockcypherAddress
	url := fmt.Sprintf(balanceApi, address)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while making API request:", err)
		return data, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error while reading API response:", err)
		return data, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error while unmarshaling API response:", err)
		return data, err
	}

	fmt.Println("Address:", data.Address)
	fmt.Println("Balance:", data.Balance, "satoshis")
	return data, err
}
