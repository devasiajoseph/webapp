package bitcoin

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
)

type BlockcypherAddress struct {
	Address string `json:"address"`
	Balance int64  `json:"balance"`
	Final   int    `json:"final_balance"`
}

type UTXOs struct {
	Notice         string `json:"notice"`
	UnspentOutputs []UTXO `json:"unspent_outputs"`
}

type UTXO struct {
	TxHash        string `json:"tx_hash"`
	TxOutputN     uint32 `json:"tx_output_n"`
	Value         int64  `json:"value"`
	Confirmations int64  `json:"confirmations"`
}

type UTXOBc struct {
	TxHash   string  `json:"tx_hash"`
	Block    int     `json:"block_height"`
	Value    float64 `json:"value"`
	Script   string  `json:"script"`
	TxInputs []struct {
		PrevHash    string `json:"prev_hash"`
		OutputIndex int    `json:"output_index"`
	} `json:"txinputs"`
}

type UnspentTransaction struct {
	TxID          string  `json:"txid"`
	Vout          uint32  `json:"vout"`
	ScriptPubKey  string  `json:"scriptPubKey"`
	Amount        float64 `json:"amount"`
	Confirmations int64   `json:"confirmations"`
}

func GetUTXO(client *rpcclient.Client, address string) ([]UnspentTransaction, error) {
	var addrs []btcutil.Address
	var unspentTransactions []UnspentTransaction
	pubKeyHashAddress, err := btcutil.DecodeAddress(address, &chaincfg.MainNetParams)
	if err != nil {
		log.Println(err)
		return unspentTransactions, err
	}
	addrs = append(addrs, pubKeyHashAddress)
	unspentOutputs, err := client.ListUnspentMinMaxAddresses(0, 9999999, addrs)
	if err != nil {
		return nil, err
	}

	for _, output := range unspentOutputs {
		unspentTransaction := UnspentTransaction{
			TxID:          output.TxID,
			Vout:          output.Vout,
			ScriptPubKey:  output.ScriptPubKey,
			Amount:        output.Amount,
			Confirmations: output.Confirmations,
		}
		unspentTransactions = append(unspentTransactions, unspentTransaction)
	}

	return unspentTransactions, nil
}

func GetBalance(address string) (int64, error) {
	var balance int64
	url := fmt.Sprintf("https://api.blockcypher.com/v1/btc/main/addrs/%s", address)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while making API request:", err)
		return balance, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error while reading API response:", err)
		return balance, err
	}

	var data BlockcypherAddress
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error while unmarshaling API response:", err)
		return balance, err
	}

	fmt.Println("Address:", data.Address)
	fmt.Println("Balance:", data.Balance, "satoshis")
	return balance, err
}

func GetUTXOAPI(address string, limit int, after int) (UTXOs, error) {

	var utxos UTXOs

	// URL for the blockchain explorer API
	apiUrl := fmt.Sprintf("https://blockchain.info/unspent?active=%s", address)

	// Make a GET request to the API
	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error fetching UTXOs:", err)
		return utxos, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return utxos, err
	}

	// Unmarshal the JSON response into a map

	err = json.Unmarshal(body, &utxos)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return utxos, err
	}
	var balance int64
	// Print the UTXOs

	fmt.Println("UTXOs:")
	for _, utxo := range utxos.UnspentOutputs {
		balance += utxo.Value
		fmt.Printf("  txid: %s, vout: %d, amount: %d, confirmations: %d\n", utxo.TxHash, utxo.TxOutputN, utxo.Value, utxo.Confirmations)
	}
	fmt.Println("Balance")
	fmt.Println(balance)
	return utxos, err
}

func GetUTXOBc(address string) {

	response, err := http.Get(fmt.Sprintf("https://api.blockcypher.com/v1/btc/main/addrs/%s?unspentOnly=true", address))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal the response into a slice of UTXOs
	var utxos []UTXOBc
	err = json.Unmarshal(body, &utxos)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the retrieved UTXOs
	fmt.Println(utxos)

}
