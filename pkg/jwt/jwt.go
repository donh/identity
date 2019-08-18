package jwt

import (
	"encoding/base64"
	"encoding/json"

	"github.com/donh/identity/pkg/models"
	"github.com/donh/identity/pkg/util"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func base64Decode(message string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// ValidateJWT validates the format of a JWT
func ValidateJWT(input string, format interface{}) error {
	var err error
	prefix := [3]string{"", "=", "=="}
	for key := range prefix {
		message := prefix[key] + input
		messageByte, errDecode := base64Decode(message)
		if errDecode != nil {
			err = errDecode
			continue
		}
		errUnmarshal := json.Unmarshal(messageByte, &format)
		if errUnmarshal == nil {
			return nil
		}
		err = errUnmarshal
	}
	return err
}

func setHeader() string {
	header := make(map[string]interface{})
	header["alg"] = "Secp256k1"
	header["typ"] = "JWT"
	tmp, _ := json.Marshal(header)
	str := base64.StdEncoding.EncodeToString(tmp)
	if str[len(str)-2] == 61 {
		str = str[:len(str)-2]
	} else if str[len(str)-1] == 61 {
		str = str[:len(str)-1]
	}
	return str
}

// JWT generates a JWT by combining header, payload, and signature
func JWT(payload map[string]interface{}) (string, error) {
	header := setHeader()
	tmp, _ := json.Marshal(payload)
	payloadString := base64.StdEncoding.EncodeToString(tmp)
	if payloadString[len(payloadString)-2] == 61 {
		payloadString = payloadString[:len(payloadString)-2]
	} else if payloadString[len(payloadString)-1] == 61 {
		payloadString = payloadString[:len(payloadString)-1]
	}

	signature, err := Signature(header + "." + payloadString)
	if err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	return (header + "." + payloadString + "." + signature), nil
}

// Signature signs a message with a private key
func Signature(str string) (string, error) {
	privateKey, err := crypto.HexToECDSA(util.Config().Ethereum[0].PrivateKey)
	if err != nil {
		return err.Error(), models.ErrInternalServerError
	}

	data := []byte(str)
	hash := crypto.Keccak256Hash(data)

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	return hexutil.Encode(signature), nil
}
