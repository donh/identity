package build

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	KeyManager "github.com/donh/identity/contracts"
	"github.com/donh/identity/pkg/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Deploy deploys KeyManager smart contract to blockchain
func Deploy(ledger string) {
	ledgers := map[string]string{
		"localhost": "http://127.0.0.1:7545",
		"rinkeby":   "https://rinkeby.infura.io",
		"ropsten":   "https://ropsten.infura.io",
	}

	URL := ""
	if val, ok := ledgers[ledger]; ok {
		URL = val
	} else {
		URL = ledgers["ropsten"]
	}

	client, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(util.Config().Ethereum[0].PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(600000) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := KeyManager.DeployKeyManager(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := address.Hex()
	log.Println("transaction hash =", tx.Hash().Hex())
	contractDID := "did:erc725:" + ledger + ":" + contractAddress
	log.Println("contract DID =", contractDID)

	_ = instance
}
