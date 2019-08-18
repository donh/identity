package test

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"log"
	"math/big"
	"strings"
	"time"

	KeyManager "github.com/donh/identity/contracts"
	"github.com/donh/identity/pkg/crypto"
	"github.com/donh/identity/pkg/jwt"
	"github.com/donh/identity/pkg/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Test tests KeyManager smart contract
func Test() {
	config := util.Config()
	investorDID := config.Ethereum[0].DID
	log.Println("investorDID =", investorDID)
	log.Println("config =", config)
	ledger := strings.Split(investorDID, ":")[2]
	investorAddress := strings.Split(investorDID, ":")[3]
	log.Println("investorDID =", investorDID)

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

	privateKey, err := ethCrypto.HexToECDSA(config.Ethereum[0].PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := ethCrypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(investorAddress)
	instance, err := KeyManager.NewKeyManager(address, client)
	if err != nil {
		log.Fatal(err)
	}

	_ = instance

	key, err := instance.GetKeysByPurpose(nil, big.NewInt(1))
	if err != nil {
		log.Println("err =", err)
		log.Fatal(err)
	}
	log.Println("getKeyByPurpose key =", hex.EncodeToString(key[0][:]))

	keyContent, err := instance.GetKey(nil, key[0])
	if err != nil {
		log.Println("err =", err)
		log.Fatal(err)
	}
	log.Println("getKey key =", keyContent)

	user := config.User
	log.Println("user =", user)
	tx, err := instance.AddKey(getAuth(privateKey, nonce), stringToByte32Array(user.PublicKey), big.NewInt(5), big.NewInt(1))
	nonce++
	if err != nil {
		log.Fatal(err)
	}
	log.Println("addKey transaction sent:", tx.Hash().Hex())
	log.Println("waiting...")
	time.Sleep(1000 * time.Millisecond)

	tx, err = instance.RemoveKey(getAuth(privateKey, nonce), stringToByte32Array(user.PublicKey), big.NewInt(5))
	nonce++
	if err != nil {
		log.Fatal(err)
	}
	log.Println("removeKey transaction sent:", tx.Hash().Hex())

	claimID, err := instance.GetClaimsIdByType(nil, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("getClaimsIdByType claimID =", hex.EncodeToString(claimID[0][:]))

	claim, err := instance.GetClaim(nil, claimID[0])
	if err != nil {
		log.Fatal(err)
	}
	log.Println("getClaim claim =", claim)

	KYCStatus := 0
	expirationDate := time.Now().AddDate(1, 0, 0).String()
	b64Claim := crypto.KYCClaim(investorDID, expirationDate, KYCStatus)

	_claimType := big.NewInt(1)
	_issuer := common.HexToAddress(user.Address)
	_signatureType := big.NewInt(1)
	_claim := []byte(b64Claim)
	s, _ := jwt.Signature(b64Claim)
	_signature := []byte(s)
	_uri := config.Claim.Host
	newClaim, err := instance.AddClaim(getAuth(privateKey, nonce), _claimType, _issuer, _signatureType, _signature, _claim, _uri)
	nonce++
	if err != nil {
		log.Fatal(err)
	}
	log.Println("addClaim transaction sent =", newClaim.Hash().Hex())

	_claimID := claimID[0]
	removeClaim, err := instance.RemoveClaim(getAuth(privateKey, nonce), _claimID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("removeClaim transaction sent =", removeClaim.Hash().Hex())
}

func stringToByte32Array(publicKeyString string) [32]byte {
	publicKeyByte32Array := [32]byte{}
	puk, err := hex.DecodeString(publicKeyString)
	if err != nil {
		return publicKeyByte32Array
	}
	for i := range publicKeyByte32Array {
		publicKeyByte32Array[i] = puk[i]
	}
	return publicKeyByte32Array
}

func getAuth(privateKey *ecdsa.PrivateKey, nonce uint64) *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)          // in wei
	auth.GasLimit = uint64(8000000)     // in units
	auth.GasPrice = big.NewInt(7000000) // gasPrice * big.NewInt(2)
	return auth
}
