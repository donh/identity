package models

type claimStruct struct {
	Host string `yaml:"host"`
}

type addressAndPrivateKey struct {
	Address    string `yaml:"address"`
	DID        string `yaml:"DID"`
	PrivateKey string `yaml:"privateKey"`
}

type databaseStruct struct {
	Account  string `yaml:"account"`
	Database string `yaml:"database"`
	Hostname string `yaml:"hostname"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
}

type emailStruct struct {
	Account   string `yaml:"account"`
	Content   string `yaml:"content"`
	Domain    string `yaml:"domain"`
	Password  string `yaml:"password"`
	Port      int    `yaml:"port"`
	Recipient string `yaml:"recipient"`
	Regard    string `yaml:"regard"`
	Sender    string `yaml:"sender"`
	Subject   string `yaml:"subject"`
}

type naclStruct struct {
	Recipient recipientAndSender `yaml:"recipient"`
	Sender    recipientAndSender `yaml:"sender"`
}

type recipientAndSender struct {
	DID        string `yaml:"DID"`
	PrivateKey string `yaml:"privateKey"`
	PublicKey  string `yaml:"publicKey"`
}

type userStruct struct {
	Address   string `yaml:"address"`
	Birthday  string `yaml:"birthday"`
	City      string `yaml:"city"`
	Country   string `yaml:"country"`
	Email     string `yaml:"email"`
	FirstName string `yaml:"firstName"`
	LastName  string `yaml:"lastName"`
	Passport  string `yaml:"passport"`
	Phone     string `yaml:"phone"`
	PublicKey string `yaml:"publicKey"`
}

type websocketStruct struct {
	UUID string `yaml:"UUID"`
}

// ConfigStruct is the format of config/config.yml
type ConfigStruct struct {
	Claim     claimStruct            `yaml:"claim"`
	Database  databaseStruct         `yaml:"database"`
	Email     emailStruct            `yaml:"email"`
	Ethereum  []addressAndPrivateKey `yaml:"Ethereum"`
	JWT       string                 `yaml:"JWT"`
	NaCl      naclStruct             `yaml:"NaCl"`
	Port      int                    `yaml:"port"`
	Timeout   float64                `yaml:"timeout"`
	User      userStruct             `yaml:"user"`
	Websocket websocketStruct        `yaml:"websocket"`
}
