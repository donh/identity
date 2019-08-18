package email

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"database/sql"
	"encoding/hex"
	"os"

	"github.com/donh/identity/pkg/models"
	"github.com/donh/identity/pkg/util"

	// A MySQL driver for Go's database/sql package
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gomail.v2"
)

const emailNotFound = "Email address not found."
const success = "Success"

func code(email string) (string, error) {
	result := bytes.Buffer{}
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	db, _ := sql.Open("mysql", dbUser+":"+dbPassword+"@/identity")
	rows, err := db.Query("SELECT VerificationCode FROM emailVerification WHERE Email = ?", email)
	if err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	defer db.Close()
	hasValue := false
	for rows.Next() {
		hasValue = true
		verificationCode := ""
		if err := rows.Scan(&verificationCode); err != nil {
			return err.Error(), models.ErrInternalServerError
		}
		result.WriteString(verificationCode + "#")
	}
	if !hasValue {
		return emailNotFound, models.ErrNotFound
	}
	if err := rows.Err(); err != nil {
		return err.Error(), models.ErrDatabaseError
	}
	return result.String()[:len(result.String())-1], nil
}

// Send sends a verification email to user
func Send(recipientEmail string) (string, error) {
	email := util.Config().Email
	verificationCode, err := code(recipientEmail)
	if err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	m := gomail.NewMessage()
	m.SetHeader("From", email.Sender)
	m.SetHeader("To", recipientEmail)
	m.SetHeader("Subject", email.Subject)
	body := email.Content + "Confirmation Code: " + verificationCode + email.Regard
	m.SetBody("text/html", body)
	d := gomail.NewDialer(email.Domain, email.Port, email.Account, email.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	return success, nil
}

// Status reutrns the verification status of an email address
func Status(email string) (string, error) {
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	db, _ := sql.Open("mysql", dbUser+":"+dbPassword+"@/identity")
	rows, err := db.Query("SELECT status FROM emailVerification WHERE Email = ?", email)
	if err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	defer db.Close()
	for rows.Next() {
		status := 0
		if err := rows.Scan(&status); err != nil {
			return err.Error(), models.ErrInternalServerError
		}
		if status == 1 {
			return "True", nil
		}
	}
	return emailNotFound, models.ErrNotFound
}

// VerificationCode generates email verification code
func VerificationCode(email string) (string, error) {
	rendomString, _ := util.UUID()
	str := rendomString + "IDHUB" + email
	hasher := md5.New()
	_, _ = hasher.Write([]byte(str))
	verificationCode := hex.EncodeToString(hasher.Sum(nil))[0:6]

	SQLStatement := "INSERT INTO emailVerification VALUES ('" + email + "', '" + verificationCode + "', false);"
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	db, _ := sql.Open("mysql", dbUser+":"+dbPassword+"@/identity")
	_, err := db.Exec(
		SQLStatement,
	)
	if err != nil {
		return err.Error(), models.ErrDatabaseError
	}
	defer db.Close()
	return success, nil
}

// Verify verifies an email address
func Verify(email, verificationCode string) (string, error) {
	realCode, err := code(email)
	if err != nil {
		return err.Error(), models.ErrInternalServerError
	}
	if realCode != verificationCode {
		return "Invalid verification code.", models.ErrBadRequest
	}
	SQLStatement := "UPDATE emailVerification SET status = true WHERE email = '" + email + "';"
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
	return success, nil
}
