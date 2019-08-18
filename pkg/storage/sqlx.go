package storage

import (
	"strconv"
	"strings"
	"time"

	"github.com/donh/identity/pkg/models"
	"github.com/donh/identity/pkg/util"
	"github.com/jmoiron/sqlx"
)

// Query queries user table
func Query(fields []string, filters map[string]string) (map[string]string, error) {
	item := map[string]string{}
	condition := ""
	value := ""
	if val, ok := filters["email"]; ok {
		condition = "`email`=?"
		value = val
	} else if val, ok := filters["address"]; ok {
		condition = "`address`=?"
		value = val
	}

	if value == "" {
		return item, models.ErrBadRequest
	}
	where := "WHERE " + condition
	statement, object := setSQLStatement(fields, where)

	db, err := setDatabaseConnection()
	if err != nil {
		return item, models.ErrDatabaseError
	}

	user := models.User{}
	err = db.Get(&user, statement, value)
	if err != nil {
		return item, models.ErrDatabaseError
	}

	item = setQueryResult(object, &user)
	return item, nil
}

func setDatabaseConnection() (*sqlx.DB, error) {
	databaseConfig := util.Config().Database
	account := databaseConfig.Account
	database := databaseConfig.Database
	hostname := databaseConfig.Hostname
	password := databaseConfig.Password
	port := strconv.Itoa(databaseConfig.Port)
	connection := account + ":" + password + "@(" + hostname + ":" + port + ")/" + database
	db, err := sqlx.Connect("mysql", connection)
	return db, err
}

func setSQLStatement(fields []string, where string) (x string, y map[string]string) {
	object := map[string]string{}
	columns := map[string]string{
		"firstname": "FirstName",
		"lastname":  "LastName",
		"email":     "Email",
		"phone":     "Phone",
		"birthday":  "Birthday",
		"ssn":       "SSN",
		"country":   "Country",
		"address":   "Address",
		"did":       "DID",
		"created":   "Created",
		"updated":   "Updated",
	}
	for _, value := range fields {
		if attribute, ok := columns[value]; ok {
			object[value] = attribute
		}
	}
	slice := []string{}
	for key := range object {
		slice = append(slice, key)
	}

	s := "*"
	if len(slice) > 0 {
		s = "`" + strings.Join(slice, "`, `") + "`"
	}
	statement := "SELECT " + s + " FROM `identity`.`users` " + where + " LIMIT 1;"
	return statement, object
}

func setQueryResult(object map[string]string, user *models.User) map[string]string {
	item := map[string]string{}
	for key := range object {
		switch key {
		case "firstname":
			item[key] = user.FirstName
		case "lastname":
			item[key] = user.LastName
		case "email":
			item[key] = user.Email
		case "phone":
			item[key] = user.Phone
		case "birthday":
			item[key] = user.Birthday
		case "ssn":
			if user.SSN.Valid {
				item[key] = user.SSN.String
			}
		case "country":
			if user.Country.Valid {
				item[key] = user.Country.String
			}
		case "address":
			if user.Address.Valid {
				item[key] = user.Address.String
			}
		case "did":
			if user.DID.Valid {
				item[key] = user.DID.String
			}
		default:
		}
	}
	return item
}

func validatePayload(payload map[string]interface{}) (firstname, lastname, email, phone, birthday string, err error) {
	firstname, lastname, email, phone, birthday = "", "", "", "", ""
	if val, ok := payload["firstname"]; ok {
		firstname = val.(string)
	}
	if firstname == "" {
		return firstname, lastname, email, phone, birthday, models.ErrBadRequest
	}

	if val, ok := payload["lastname"]; ok {
		lastname = val.(string)
	}
	if lastname == "" {
		return firstname, lastname, email, phone, birthday, models.ErrBadRequest
	}

	if val, ok := payload["email"]; ok {
		email = val.(string)
	}
	if email == "" {
		return firstname, lastname, email, phone, birthday, models.ErrBadRequest
	}

	if val, ok := payload["phone"]; ok {
		phone = val.(string)
	}
	if phone == "" {
		return firstname, lastname, email, phone, birthday, models.ErrBadRequest
	}

	if val, ok := payload["birthday"]; ok {
		birthday = val.(string)
	}
	if birthday == "" {
		return firstname, lastname, email, phone, birthday, models.ErrBadRequest
	}
	return firstname, lastname, email, phone, birthday, nil
}

// Insert inserts a record into user table
func Insert(payload map[string]interface{}) error {
	firstname, lastname, email, phone, birthday, err := validatePayload(payload)
	if err != nil {
		return err
	}

	db, err := setDatabaseConnection()
	if err != nil {
		return models.ErrDatabaseError
	}

	statement := "INSERT INTO `identity`.`users` "
	statement += "(firstname, lastname, email, phone, birthday, ssn, "
	statement += "country, region, city, street, zip, address, did, created) "
	statement += "VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	result, err := db.Exec(statement, firstname, lastname, email, phone, birthday,
		payload["ssn"], payload["country"], payload["region"], payload["city"],
		payload["street"], payload["zip"], payload["address"], payload["did"], getNow())
	if err != nil {
		return models.ErrDatabaseError
	}
	RowsAffected, _ := result.RowsAffected()
	if RowsAffected > 0 {
		return nil
	}
	return models.ErrDatabaseError
}

// Update updates a record in user table
func Update(payload map[string]interface{}) error {
	firstname, lastname, email, phone, birthday, err := validatePayload(payload)
	if err != nil {
		return err
	}

	db, err := setDatabaseConnection()
	if err != nil {
		return models.ErrDatabaseError
	}

	statement := "UPDATE `identity`.`users` "
	statement += "SET `firstname`=?, `lastname`=?, `email`=?, `phone`=?, "
	statement += "`birthday`=?, `ssn`=?, `country`=?, `region`=?, `city`=?, "
	statement += "`street`=?, `zip`=?, `address`=?, `did`=?, `updated`=? "
	statement += "WHERE `email`=?;"
	result, err := db.Exec(statement, firstname, lastname, email, phone, birthday,
		payload["ssn"], payload["country"], payload["region"], payload["city"],
		payload["street"], payload["zip"], payload["address"], payload["did"], getNow(), email)
	if err != nil {
		return models.ErrDatabaseError
	}
	RowsAffected, _ := result.RowsAffected()
	if RowsAffected > 0 {
		return nil
	}
	return models.ErrDatabaseError
}

// Delete removes a record from user table
func Delete(payload map[string]interface{}) error {
	firstname, lastname, email, phone, birthday, err := validatePayload(payload)
	if err != nil {
		return err
	}

	db, err := setDatabaseConnection()
	if err != nil {
		return models.ErrDatabaseError
	}

	statement := "DELETE FROM `identity`.`users` "
	statement += "WHERE `firstname`=? AND `lastname`=? AND `email`=? AND `phone`=? AND `birthday`=?;"
	result, err := db.Exec(statement, firstname, lastname, email, phone, birthday)
	if err != nil {
		return models.ErrDatabaseError
	}
	RowsAffected, _ := result.RowsAffected()
	if RowsAffected > 0 {
		return nil
	}
	return models.ErrDatabaseError
}

func getNow() string {
	t := time.Now()
	now := t.Format("2006-01-02 15:04:05")
	return now
}
