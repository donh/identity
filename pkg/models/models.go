package models

import (
	"time"
)

// DIDAuthJWTPayload is the format of a JWT payload
type DIDAuthJWTPayload struct {
	DID       string
	Timestamp int
	UUID      string
}

// JWTHeader is the format of a JWT header
type JWTHeader struct {
	Alg string
	Typ string
}

// KYCDBJWTPayload is the format of a JWT payload for KYC
type KYCDBJWTPayload struct {
	DID                    string
	Address                string
	Timestamp              int
	Resident               string
	Email                  string
	FirstName              string
	MiddleName             string
	LastName               string
	SocialSecurityNumber   string
	DateOfBirth            string
	ResidentialAddress     string
	InvestorType           int
	AccreditationCriterion int
	EvidenceType           int
	ApplicationDocuments   string
	Description            string
	RegisterDate           string
}

// ResponseWrapper is the format of a wrapper for API Response
type ResponseWrapper struct {
	Result interface{}
	Status int
	Error  string
	Time   time.Time
}

// WebSocketWrapper is the format of a wrapper for web socket
type WebSocketWrapper struct {
	Data       interface{}
	Event      string
	ResultType int
}
