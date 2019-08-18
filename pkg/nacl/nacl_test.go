package nacl

import (
	"testing"

	"github.com/donh/identity/pkg/util"
)

// func TestConvertKeyFromStringToByte(t *testing.T) {
// 	recipient := util.Config().NaCl.Recipient
// 	privateKeyByte, err := ConvertKeyFromStringToByte(recipient.PrivateKey)
// 	if err != nil {
// 		t.Error(err)
// 	} else {
// 		t.Log("privateKeyByte =", privateKeyByte)
// 	}

// 	publicKeyByte, err := ConvertKeyFromStringToByte(recipient.PublicKey)
// 	if err != nil {
// 		t.Error(err)
// 	} else {
// 		t.Log("publicKeyByte =", publicKeyByte)
// 	}
// }

func TestEncrypt(t *testing.T) {
	config := util.Config()
	recipient := config.NaCl.Recipient
	sender := config.NaCl.Sender
	senderPrivateKey, _ := ConvertKeyFromStringToByte(sender.PrivateKey)
	recipientPublicKey, _ := ConvertKeyFromStringToByte(recipient.PublicKey)
	result, err := Encrypt(senderPrivateKey, recipientPublicKey, config.User.FirstName)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(result)
	}
}

func TestDecrypt(t *testing.T) {
	config := util.Config()
	recipient := config.NaCl.Recipient
	sender := config.NaCl.Sender
	senderPrivateKey, _ := ConvertKeyFromStringToByte(sender.PrivateKey)
	recipientPublicKey, _ := ConvertKeyFromStringToByte(recipient.PublicKey)
	encryptedMessage, _ := Encrypt(senderPrivateKey, recipientPublicKey, config.User.FirstName)

	recipientPrivateKey, _ := ConvertKeyFromStringToByte(recipient.PrivateKey)
	senderPublicKey, _ := ConvertKeyFromStringToByte(sender.PublicKey)
	result, err := Decrypt(recipientPrivateKey, senderPublicKey, encryptedMessage)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(result)
	}
}

// func TestPublicKeyNaCl(t *testing.T) {
// 	result, _ := PublicKeyNaCl(util.Config().NaCl.Sender.DID)
// 	t.Log(result)
// }
