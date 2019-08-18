package crypto

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"strings"

	KeyManager "github.com/donh/identity/contracts"
	"github.com/donh/identity/pkg/models"
	"github.com/donh/identity/pkg/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// CheckPublicKey checks whether the public key matches the decentralized identifier (DID)
func CheckPublicKey(decentralizedIdentifier, publicKey string) (bool, error) {
	tmp := strings.Split(decentralizedIdentifier, ":")
	if len(tmp) != 4 {
		return false, models.ErrBadRequest
	}
	method := tmp[1]
	chain := tmp[2]
	contractAddressString := tmp[3]

	if method != "erc725" {
		return false, models.ErrBadRequest
	}

	client, err := ethclient.Dial("https://" + chain + ".infura.io")
	if err != nil {
		return false, err
	}

	contractAddress := common.HexToAddress(contractAddressString)
	instance, err := KeyManager.NewKeyManager(contractAddress, client)
	if err != nil {
		return false, err
	}

	keyB32 := publicKeyStringToByte32Array(publicKey)
	result, err := instance.KeyHasPurpose(&bind.CallOpts{}, keyB32, big.NewInt(1))
	if err != nil {
		return false, err
	}
	if result {
		return result, nil
	}

	return result, models.ErrNotFound
}

// KYCClaim returns a string of a claim
func KYCClaim(investorDID, expirationDate string, status int) string {
	output := make(map[string]interface{})
	output["recipientDID"] = util.Config().NaCl.Recipient.DID
	output["investorDID"] = investorDID
	output["expirationDate"] = expirationDate
	output["KYCstatus"] = status
	tmp, _ := json.Marshal(output)
	str := base64.StdEncoding.EncodeToString(tmp)
	if str[len(str)-2] == 61 {
		str = str[:len(str)-2]
	} else if str[len(str)-1] == 61 {
		str = str[:len(str)-1]
	}
	return str
}

// PublicKeyFromAddress returns a public key by its address
func PublicKeyFromAddress(address string) string {
	addr := common.HexToAddress(address).Bytes()
	hash := crypto.Keccak256Hash(addr)
	return hash.Hex()[2:]
}

// PublicKeyFromSignature returns a public key by its signature
func PublicKeyFromSignature(signature, msg string) string {
	data := []byte(msg)
	hash := crypto.Keccak256Hash(data)
	bSig, err := hexutil.Decode(signature)
	if err != nil {
		return err.Error()
	}
	sigPublicKey, err := crypto.SigToPub(hash.Bytes(), bSig)
	if err != nil {
		return err.Error()
	}
	addr := crypto.PubkeyToAddress(*sigPublicKey).Hex()
	return PublicKeyFromAddress(addr)
}

// publicKeyStringToByte32Array converts a string of a public key to a byte array
func publicKeyStringToByte32Array(publicKeyString string) [32]byte {
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
