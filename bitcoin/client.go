package bitcoin

import (
	"log"

	"github.com/btcsuite/btcd/rpcclient"
)

// node "localhost:8332"
// user rpcuser"
// password "rpcpassword"
func Connect(node string, user string, password string) (*rpcclient.Client, error) {

	connCfg := &rpcclient.ConnConfig{
		Host:         node,
		User:         user,
		Pass:         password,
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	client, err := rpcclient.New(connCfg, nil)

	if err != nil {
		log.Println("Error connecting node")
		log.Println(err)
	}

	return client, nil
}
