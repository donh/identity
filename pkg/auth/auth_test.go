package auth

import (
	"testing"
	"time"

	"github.com/donh/identity/pkg/jwt"
	"github.com/donh/identity/pkg/util"
)

func TestDIDAuth(t *testing.T) {
	AuthJWT, _ := jwt.JWT(AuthPayload())
	result, err := DIDAuth(AuthJWT)
	if result && err == nil {
		t.Log("DID Auth Success.")
	} else {
		t.Log("DID Auth Failed.")
		t.Log(result)
	}
}

func AuthPayload() map[string]interface{} {
	payload := make(map[string]interface{})
	config := util.Config()
	payload["DID"] = config.NaCl.Sender.DID
	payload["timestamp"] = int(time.Now().Unix())
	payload["uuid"] = config.Websocket.UUID
	return payload
}
