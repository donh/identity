package auth

import (
	"encoding/json"
	"math/big"
	"strings"
	"time"

	KeyManager "github.com/donh/identity/contracts"
	"github.com/donh/identity/pkg/crypto"
	"github.com/donh/identity/pkg/jwt"
	"github.com/donh/identity/pkg/models"
	"github.com/donh/identity/pkg/storage"
	"github.com/donh/identity/pkg/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/websocket"
)

// Clients are  web socket connections
var Clients = make(map[string]*websocket.Conn)

// DIDAuth decides whether DID authentication is successful
func DIDAuth(inputJWT string) (bool, error) {
	tmp := strings.Split(inputJWT, ".")
	if len(tmp) != 3 {
		return false, models.ErrBadRequest
	}

	headerString := tmp[0]
	payloadString := tmp[1]
	signature := tmp[2]

	err := jwt.ValidateJWT(headerString, models.JWTHeader{})
	if err != nil {
		return false, err
	}

	payload := models.DIDAuthJWTPayload{}
	err = jwt.ValidateJWT(payloadString, payload)
	if err != nil {
		return false, err
	}

	timeout := util.Config().Timeout
	if int(time.Now().Unix())-payload.Timestamp > int(timeout) {
		webSocket := models.WebSocketWrapper{
			Data:       "JWT timeout Error.",
			Event:      "login",
			ResultType: 2,
		}
		wsResult, _ := json.Marshal(webSocket)
		setWebsocket(payload.UUID, string(wsResult))
		return false, models.ErrTimeout
	}
	publicKey := crypto.PublicKeyFromSignature(signature, headerString+"."+payloadString)
	result, err := crypto.CheckPublicKey(payload.DID, publicKey)
	if result {
		data := ""
		data, err = storage.UserInfo(payload.DID)
		if err == nil {
			webSocket := models.WebSocketWrapper{
				Data:       data,
				Event:      "login",
				ResultType: 0,
			}
			value, _ := json.Marshal(webSocket)
			setWebsocket(payload.UUID, string(value))
			return true, nil
		}
		webSocket := models.WebSocketWrapper{
			Data:       data,
			Event:      "login",
			ResultType: 2,
		}
		wsResult, _ := json.Marshal(webSocket)
		setWebsocket(payload.UUID, string(wsResult))
		return false, err
	}
	webSocket := models.WebSocketWrapper{
		Data:       err.Error(),
		Event:      "login",
		ResultType: 2,
	}
	wsResult, _ := json.Marshal(webSocket)
	setWebsocket(payload.UUID, string(wsResult))
	return false, err
}

// GetClaimSlice gets claims of a DID
func GetClaimSlice(decentralizedIdentifier string) (string, error) {
	tmp := strings.Split(decentralizedIdentifier, ":")
	if len(tmp) != 4 {
		return "Invalid DID.", models.ErrBadRequest
	}
	method := tmp[1]
	chain := tmp[2]
	contractAddressString := tmp[3]

	if method != "erc725" {
		return "Invalid DID method.", models.ErrBadRequest
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
	result, err := instance.GetClaimsIdByType(&bind.CallOpts{}, big.NewInt(1))
	if err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	for i := range result {
		claim, _ := instance.GetClaim(&bind.CallOpts{}, result[i])
		claimSlice := make(map[string]interface{})
		claimString := string(claim.Claim)
		_ = jwt.ValidateJWT(claimString, claimSlice)
		publicKey := crypto.PublicKeyFromSignature(string(claim.Signature), claimString)
		recipientDID := util.Config().NaCl.Recipient.DID
		hasKey, _ := crypto.CheckPublicKey(recipientDID, publicKey)
		if hasKey {
			output, _ := json.Marshal(claimSlice)
			return string(output), nil
		}
	}
	return "Claim not found.", models.ErrNotFound
}

func setWebsocket(universallyUniqueIdentifier, message string) {
	if _, ok := Clients[universallyUniqueIdentifier]; ok {
		_ = Clients[universallyUniqueIdentifier].WriteMessage(1, []byte(message))
		delete(Clients, universallyUniqueIdentifier)
	}
}
