package storage

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/donh/identity/pkg/crypto"
	"github.com/donh/identity/pkg/jwt"
	"github.com/donh/identity/pkg/models"
	"github.com/donh/identity/pkg/nacl"
	"github.com/donh/identity/pkg/util"

	// A MySQL driver for Go's database/sql package
	_ "github.com/go-sql-driver/mysql"
)

// DatabaseInsert inserts a record into user table
func DatabaseInsert(inpytJWT string) (string, error) {
	tmp := strings.Split(inpytJWT, ".")
	if len(tmp) != 3 {
		return "JWT Format Error", models.ErrBadRequest
	}

	headerString := tmp[0]
	payloadString := tmp[1]
	signature := tmp[2]

	err := jwt.ValidateJWT(headerString, models.JWTHeader{})
	if err != nil {
		return models.ErrInternalServerError.Error(), err
	}

	payload := models.KYCDBJWTPayload{}
	err = jwt.ValidateJWT(payloadString, payload)
	if err != nil {
		return models.ErrInternalServerError.Error(), err
	}

	if payload.DID == "" {
		return "please enter your DID", models.ErrBadRequest
	}

	publicKey := crypto.PublicKeyFromSignature(signature, headerString+"."+payloadString)
	result, err := crypto.CheckPublicKey(payload.DID, publicKey)
	if !result || err != nil {
		return err.Error(), models.ErrInternalServerError
	}

	insertKey := bytes.Buffer{}
	insertValue := bytes.Buffer{}

	e := reflect.ValueOf(&payload).Elem()
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		if varName == "Timestamp" {
			continue
		}
		if varType.Name() == "int" {
			if varValue == 0 {
				insertKey.WriteString(varName + ", ")
				insertValue.WriteString("-1, ")
				continue
			}
			insertKey.WriteString(varName + ", ")
			insertValue.WriteString(strconv.Itoa(varValue.(int)) + ", ")
		} else if varType.Name() == "string" {
			if varValue == "" {
				insertKey.WriteString(varName + ", ")
				insertValue.WriteString("'NULL', ")
				continue
			}
			insertKey.WriteString(varName + ", ")
			insertValue.WriteString("'" + varValue.(string) + "', ")
		}
	}

	SQLStatement := "INSERT INTO user (" + insertKey.String()[:len(insertKey.String())-2]
	SQLStatement += ") VALUES (" + insertValue.String()[:len(insertValue.String())-2] + ");"
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	db, _ := sql.Open("mysql", dbUser+":"+dbPassword+"@/identity")
	_, err = db.Exec(
		SQLStatement,
	)
	if err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	defer db.Close()
	return "Success", nil
}

func parsePayload(payload *models.KYCDBJWTPayload) bytes.Buffer {
	buffer := bytes.Buffer{}
	e := reflect.ValueOf(&payload).Elem()
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		if varName == "DID" || varName == "Timestamp" {
			continue
		}
		if varType.Name() == "int" {
			if varValue == 0 {
				continue
			}
			buffer.WriteString(varName + "=" + strconv.Itoa(varValue.(int)) + ", ")
		} else if varType.Name() == "string" {
			if varValue == "" {
				continue
			}
			buffer.WriteString(varName + "='" + varValue.(string) + "', ")
		}
	}
	return buffer
}

// DatabaseUpdate updates a record in user table
func DatabaseUpdate(inputJWT string) (string, error) {
	tmp := strings.Split(inputJWT, ".")
	if len(tmp) != 3 {
		return "JWT Format Error", models.ErrBadRequest
	}

	headerString := tmp[0]
	payloadString := tmp[1]
	signature := tmp[2]

	err := jwt.ValidateJWT(headerString, models.JWTHeader{})
	if err != nil {
		return models.ErrInternalServerError.Error(), err
	}

	payload := models.KYCDBJWTPayload{}
	err = jwt.ValidateJWT(payloadString, payload)
	if err != nil {
		return models.ErrInternalServerError.Error(), err
	}

	if payload.DID == "" {
		return "please enter your DID", models.ErrBadRequest
	}

	publicKey := crypto.PublicKeyFromSignature(signature, headerString+"."+payloadString)
	result, err := crypto.CheckPublicKey(payload.DID, publicKey)
	if !result || err != nil {
		return "public key is not found in DID document", models.ErrInternalServerError
	}

	data := parsePayload(&payload)
	if data.String() == "" {
		return "Invalid Payload.", models.ErrBadRequest
	}

	SQLStatement := "UPDATE user SET "
	SQLStatement += data.String()[:len(data.String())-2]
	SQLStatement += " WHERE DID = '" + payload.DID + "';"
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	db, _ := sql.Open("mysql", dbUser+":"+dbPassword+"@/identity")
	_, err = db.Exec(
		SQLStatement,
	)
	if err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	defer db.Close()
	return "Success", nil
}

// DIDByAddress returns an Ethereum address by its DID
func DIDByAddress(address string) (string, error) {
	result := bytes.Buffer{}
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	db, _ := sql.Open("mysql", dbUser+":"+dbPassword+"@/identity")
	rows, err := db.Query("SELECT DID FROM user WHERE address = ?", address)
	if err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	defer db.Close()
	hasValue := false
	for rows.Next() {
		hasValue = true
		DID := ""
		if err := rows.Scan(&DID); err != nil {
			return err.Error(), models.ErrInternalServerError
		}
		result.WriteString(DID + "#")
	}
	if !hasValue {
		return "Ethereum address not found.", models.ErrNotFound
	}
	if err := rows.Err(); err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	resultSlice := strings.Split(result.String()[:len(result.String())-1], "#")
	resultJSON, _ := json.Marshal(resultSlice)
	return string(resultJSON), nil
}

// DIDByEmail returns an email address by its DID
func DIDByEmail(email string) (string, error) {
	result := bytes.Buffer{}
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	db, _ := sql.Open("mysql", dbUser+":"+dbPassword+"@/identity")
	rows, err := db.Query("SELECT DID FROM user WHERE Email = ?", email)
	if err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	defer db.Close()
	hasValue := false
	for rows.Next() {
		hasValue = true
		DID := ""
		if err := rows.Scan(&DID); err != nil {
			return err.Error(), models.ErrInternalServerError
		}
		result.WriteString(DID + "#")
	}
	if !hasValue {
		return "Email address not found.", models.ErrNotFound
	}
	if err := rows.Err(); err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	return result.String()[:len(result.String())-1], nil
}

// KYCClaimInsert inserts a claim into KYC table
func KYCClaimInsert(claim, signature string) (string, error) {
	recipientDID := util.Config().NaCl.Recipient.DID
	claimSlice := make(map[string]interface{})
	_ = jwt.ValidateJWT(claim, claimSlice)
	publicKey := crypto.PublicKeyFromSignature(signature, claim)
	result, err := crypto.CheckPublicKey(recipientDID, publicKey)
	if !result || err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	investorDID, _ := claimSlice["investorDID"].(string)
	KYCStatus, _ := claimSlice["KYCStatus"].(float64)
	SQLStatement := "INSERT INTO KYC VALUES ('" + investorDID + "', '" + claim
	SQLStatement += "." + signature + "', " + strconv.FormatFloat(KYCStatus, 'f', 0, 64) + ")"
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	db, _ := sql.Open("mysql", dbUser+":"+dbPassword+"@/identity")
	_, err = db.Exec(
		SQLStatement,
	)
	if err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	defer db.Close()
	return "Success", nil
}

// KYCClaimQuery returns claims of a DID
// queryType == 0, get content of claim
// queryType == 1, get claim information for ERC-725
func KYCClaimQuery(decentralizedIdentifier string, queryType int) (string, error) {
	result := bytes.Buffer{}
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	db, _ := sql.Open("mysql", dbUser+":"+dbPassword+"@/identity")
	rows, err := db.Query("SELECT claim FROM KYC WHERE DID = ?", decentralizedIdentifier)
	if err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	defer db.Close()
	hasValue := false
	for rows.Next() {
		hasValue = true
		claim := ""
		if err := rows.Scan(&claim); err != nil {
			return err.Error(), models.ErrInternalServerError
		}
		result.WriteString(claim + "#")
	}
	if !hasValue {
		return "DID not found.", models.ErrNotFound
	}
	if err := rows.Err(); err != nil {
		return err.Error(), models.ErrInternalServerError
	}

	if queryType == 0 {
		claimB64 := result.String()[:len(result.String())-1]
		claimB64 = strings.Split(claimB64, ".")[0]
		claimSlice := make(map[string]interface{})
		_ = jwt.ValidateJWT(claimB64, claimSlice)
		output, _ := json.Marshal(claimSlice)
		return string(output), nil
	} else if queryType == 1 {
		recipientDID := util.Config().NaCl.Recipient.DID
		recipientDIDSlice := strings.Split(recipientDID, ":")
		claimB64 := result.String()[:len(result.String())-1]

		claimSlice := make(map[string]interface{})
		claimSlice["claim"] = strings.Split(claimB64, ".")[0]
		claimSlice["claimType"] = 1
		claimSlice["issuer"] = recipientDIDSlice[len(recipientDIDSlice)-1]
		claimSlice["signature"] = strings.Split(claimB64, ".")[1]
		claimSlice["signatureType"] = 1
		claimSlice["url"] = util.Config().Claim.Host

		output, _ := json.Marshal(claimSlice)
		return string(output), nil
	}
	return "Failed", models.ErrInternalServerError
}

// PictureUpdate updates the image of user's application documents
func PictureUpdate(decentralizedIdentifier string, picture []byte) (string, error) {
	SQLStatement := "UPDATE user SET ApplicationDocuments=? WHERE DID = '" + decentralizedIdentifier + "'"
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	db, _ := sql.Open("mysql", dbUser+":"+dbPassword+"@/identity")
	_, err := db.Exec(
		SQLStatement,
		picture,
	)
	if err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	defer db.Close()
	return "Success", nil
}

func renderOutput(fields, decentralizedIdentifier string, dest []interface{}, output map[string]interface{}) map[string]interface{} {
	userPublicKey, _ := nacl.PublicKeyNaCl(decentralizedIdentifier)
	sliceFields := strings.Split(fields, ",")
	for i, key := range sliceFields {
		switch dest[i].(type) {
		// For now, only string type is encrypted.
		// Future works: implement type conversion for int, float64, and other types.
		case *int:
			output[key] = sliceFields[i]
			continue
		case *float64:
			output[key] = sliceFields[i]
			continue
		case *string:
			recipientPrivateKeyByte, _ := nacl.ConvertKeyFromStringToByte(util.Config().NaCl.Recipient.PrivateKey)
			userPublicKeyByte, _ := nacl.ConvertKeyFromStringToByte(userPublicKey)
			decrypted, err := nacl.Decrypt(recipientPrivateKeyByte, userPublicKeyByte, sliceFields[i])
			if err != nil {
				output[sliceFields[i]] = sliceFields[i]
			} else {
				output[sliceFields[i]] = decrypted
			}
			continue
		default:
			output[key] = sliceFields[i]
		}
	}
	output["DID"] = decentralizedIdentifier
	output["Timestamp"] = int(time.Now().Unix())
	return output
}

// UserInfo returns user data by DID
func UserInfo(decentralizedIdentifier string) (string, error) {
	payload := models.KYCDBJWTPayload{}
	fields := ""
	e := reflect.ValueOf(&payload).Elem()
	for i := 0; i < e.NumField(); i++ {
		if e.Type().Field(i).Name == "DID" || e.Type().Field(i).Name == "Timestamp" {
			continue
		}
		fields = fields + e.Type().Field(i).Name + ","
	}
	fields = fields[:len(fields)-1]

	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	db, _ := sql.Open("mysql", dbUser+":"+dbPassword+"@/identity")
	rows, err := db.Query("SELECT " + fields + " FROM user WHERE did = '" + decentralizedIdentifier + "';")
	if err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	defer db.Close()
	hasValue := false
	sliceFields := strings.Split(fields, ",")
	dest := make([]interface{}, len(sliceFields))
	for i := range sliceFields {
		dest[i] = &sliceFields[i] // Put pointers to each string in the interface slice
	}
	output := make(map[string]interface{})
	for rows.Next() {
		hasValue = true
		if err := rows.Scan(dest...); err != nil {
			return err.Error(), models.ErrInternalServerError
		}
	}
	if !hasValue {
		output["DID"] = decentralizedIdentifier
		output["Timestamp"] = int(time.Now().Unix())
		result, _ := json.Marshal(output)
		return string(result), nil
	}
	if err := rows.Err(); err != nil {
		return err.Error(), models.ErrInternalServerError
	}

	output = renderOutput(fields, decentralizedIdentifier, dest, output)
	result, _ := json.Marshal(output)
	return string(result), nil
}
