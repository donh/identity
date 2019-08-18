package nacl

import (
	crypto_rand "crypto/rand"
	"encoding/hex"
	"io"
	"math/big"
	"strings"

	KeyManager "github.com/donh/identity/contracts"
	"github.com/donh/identity/pkg/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"golang.org/x/crypto/nacl/box"
)

// ConvertKeyFromStringToByte converts a key from string type to byte type
func ConvertKeyFromStringToByte(key string) (x [32]byte, err error) {
	keyByte := [32]byte{}
	keyHex, err := hex.DecodeString(key)
	if err != nil {
		return keyByte, err
	}

	for i := range keyHex {
		keyByte[i] = keyHex[i]
	}
	return keyByte, nil
}

// Encrypt encrypts a message by sender's private key and recipient's public key
func Encrypt(senderPrivateKey, recipientPublicKey [32]byte, message string) (string, error) {
	nonce := [24]byte{}
	if _, err := io.ReadFull(crypto_rand.Reader, nonce[:]); err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	encrypted := box.Seal(nonce[:], []byte(message), &nonce, &recipientPublicKey, &senderPrivateKey)
	return hex.EncodeToString(encrypted), nil
}

// Decrypt decrypts a message by recipient's private key and sender's public key
func Decrypt(recipientPrivateKey, senderPublicKey [32]byte, message string) (string, error) {
	nonce := [24]byte{}
	encrypted, err := hex.DecodeString(message)
	if err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	copy(nonce[:], encrypted[:24])
	decrypted, ok := box.Open(nil, encrypted[24:], &nonce, &senderPublicKey, &recipientPrivateKey)
	if !ok {
		return "Decrypt Failed.", models.ErrInternalServerError
	}
	return string(decrypted), nil
}

// PublicKeyNaCl returns a public key in NaCl form by a DID
func PublicKeyNaCl(decentralizedIdentifier string) (string, error) {
	tmp := strings.Split(decentralizedIdentifier, ":")
	if len(tmp) != 4 {
		return "invalid DID", models.ErrBadRequest
	}
	method := tmp[1]
	chain := tmp[2]
	contractAddressString := tmp[3]

	if method != "erc725" {
		return "invalid DID method", models.ErrBadRequest
	}

	client, err := ethclient.Dial("https://" + chain + ".infura.io")
	if err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	contractAddress := common.HexToAddress(contractAddressString)
	instance, err := KeyManager.NewKeyManager(contractAddress, client)
	if err != nil {
		return err.Error(), models.ErrInternalServerError
	}

	result, err := instance.GetKeysByPurpose(&bind.CallOpts{}, big.NewInt(5844))
	if err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	if len(result) == 0 {
		return "NaCl public key not found", models.ErrNotFound
	}
	return hex.EncodeToString(result[0][:]), nil
}

/*
func executioner(privateKey string) [4]string {
	output := [4]string{}
	output[0] = privateKey[0:16]
	output[1] = privateKey[16:32]
	output[2] = privateKey[32:48]
	output[3] = privateKey[48:64]
	return output
}
*/
