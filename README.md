# identity [![CircleCI Build Status](https://circleci.com/gh/donh/identity.svg?style=shield)](https://circleci.com/gh/donh/identity) [![Go Report Card](https://goreportcard.com/badge/github.com/donh/identity)](https://goreportcard.com/report/github.com/donh/identity) [![GolangCI](https://golangci.com/badges/github.com/donh/identity.svg)](https://golangci.com/r/github.com/donh/identity) [![MIT Licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/CircleCI-Public/circleci-demo-go/master/LICENSE.md)

## Installation
- Go 1.12
- go get
  - go get github.com/ethereum/go-ethereum
  - go get github.com/go-sql-driver/mysql
  - go get github.com/gofrs/uuid
  - go get github.com/gorilla/mux
  - go get github.com/gorilla/websocket
  - go get github.com/jmoiron/sqlx
  - go get github.com/rs/cors
  - go get golang.org/x/crypto/nacl/box
  - go get gopkg.in/gomail.v2
  - go get gopkg.in/yaml.v2
- npm install
  - ethereumjs-testrpc
  - openzeppelin-solidity
  - truffle
  - truffle-hdwallet-provider
- pip3 install
  - web3

## Run
- Set configuration
  ```bash
  $ cp config/config.yml.example config/config.yml
  ```
- Start
  ```bash
  $ go run ./identity.go
  ```

## Lint
- Go
  ```bash
  $ go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
  ```
  ```bash
  $ golangci-lint run --fix
  ```
- Solidity
  - [Solium](https://github.com/duaraghav8/Ethlint)
  ```bash
  $ sudo npm install -g ethlint
  ```
  ```bash
  $ solium -d contracts/ --fix
  ```
- Python
  - [pylint](https://www.pylint.org/)
  ```bash
  $ pip install pylint
  ```
  ```bash
  $ pylint **/*.py
  ```
  ```bash
  $ find . -iname "*.py" | xargs pylint
  ```
  - [yapf](https://github.com/google/yapf)
  ```bash
  $ pip install yapf
  ```
  ```bash
  $ yapf -ir .
  ```

## Test
- Go
```bash
go test -v fmt
go test -v ./...
go test -cover -v ./...
go test $(go list ./... | grep -v /vendor/) -v -coverprofile .testCoverage.txt
```
- Smart contract
```bash
$ cd contracts
$ truffle build
$ abigen --bin=KeyManager.bin --abi=KeyManager.abi --pkg=KeyManager --out=KeyManager.go

$ cd ..
$ go run ./identity.go build
$ go run ./identity.go deploy ropsten
$ go run ./identity.go test
```

## Docker
- Development environment
  - Truffle
  - python-web3
- Build docker image
```bash
$ sudo docker build . -t "truffle:test"
```
- Start Docker Compose
```bash
$ sudo docker-compose up --force-recreate -d
```
- Enter docker
```bash
$ sudo docker exec -it $DockerID bash
```
- Start API server
```bash
$ go run ./identity.go
```
- Shut down Docker Compose
```bash
$ sudo docker-compose down -v
```

## Introduction
- Three models of digital identity
  - Siloed identity
  - Identity provider (IDP)
  - [Self-sovereign identity](https://wiki.p2pfoundation.net/Self-Sovereign_Identity) (SSI)
- [Decentralized Identifier](https://w3c-ccg.github.io/did-spec/) (DID)
  - Decentralized Identifier (DID): A globally unique identifier that does not require a centralized registration authority because it is registered with distributed ledger technology or other form of decentralized network.
- DID methods

| Driver Name | Driver Version | DID Spec Version | DID Method Spec Version | Docker Image |
| ----------- | -------------- | ---------------- | ----------------------- | ------------ |
| [did-btcr](https://github.com/decentralized-identity/universal-resolver/tree/master/drivers/btcr/) | 0.1-SNAPSHOT | [0.11](https://w3c-ccg.github.io/did-spec/) | [0.1](https://w3c-ccg.github.io/didm-btcr) | [universalresolver/driver-did-btcr](https://hub.docker.com/r/universalresolver/driver-did-btcr/)
| [did-sov](https://github.com/decentralized-identity/universal-resolver/tree/master/drivers/sov/) | 0.1-SNAPSHOT | [0.11](https://w3c-ccg.github.io/did-spec/) | [0.1](https://github.com/mikelodder7/sovrin/blob/master/spec/did-method-spec-template.html) | [universalresolver/driver-did-sov](https://hub.docker.com/r/universalresolver/driver-did-sov/)
| [did-erc725](https://github.com/decentralized-identity/universal-resolver/tree/master/drivers/erc725/) | 0.1-SNAPSHOT | [0.11](https://w3c-ccg.github.io/did-spec/) | [0.1](https://github.com/WebOfTrustInfo/rebooting-the-web-of-trust-spring2018/blob/master/topics-and-advance-readings/DID-Method-erc725.md) | [universalresolver/driver-did-erc725](https://hub.docker.com/r/universalresolver/driver-did-erc725/)
| [did-stack](https://github.com/decentralized-identity/universal-resolver/tree/master/drivers/stack/) | 0.1 | [0.11](https://w3c-ccg.github.io/did-spec/) | (missing) | [universalresolver/driver-did-stack](https://hub.docker.com/r/universalresolver/driver-did-stack/)
| [did-dom](https://github.com/decentralized-identity/universal-resolver/tree/master/drivers/dom/) | 0.1-SNAPSHOT | [0.11](https://w3c-ccg.github.io/did-spec/) | (missing) | [universalresolver/driver-did-dom](https://hub.docker.com/r/universalresolver/driver-did-dom/)
| [did-uport](https://github.com/uport-project/uport-did-driver) | 1.1.0 | [0.11](https://w3c-ccg.github.io/did-spec/) | [1.0](https://docs.google.com/document/d/1vS6UBUDwxYR8tLTNo4HUhGe2qb9Q95QLiJTt9NkwZ8M/) | [uport/uni-resolver-driver-did-uport](https://hub.docker.com/r/uport/uni-resolver-driver-did-uport/)
| did-v1 |  | [0.11](https://w3c-ccg.github.io/did-spec/) | [1.0](https://w3c-ccg.github.io/didm-veres-one/) |
| did-ipid |  | [0.11](https://w3c-ccg.github.io/did-spec/) | [0.1](https://github.com/jonnycrunch/ipid) |
| [did-jolo](https://github.com/jolocom/jolocom-did-driver) | 0.1 | [0.11](https://w3c-ccg.github.io/did-spec/) | [0.1](https://github.com/jolocom/jolocom-did-driver/blob/master/jolocom-did-method-specification.md) | [jolocomgmbh/jolocom-did-driver](https://hub.docker.com/r/jolocomgmbh/jolocom-did-driver) |
| [did-hacera](https://github.com/hacera/hacera-did-driver) | 0.1 | [0.11](https://w3c-ccg.github.io/did-spec/) | (missing) | [hacera/hacera-did-driver](https://hub.docker.com/r/hacera/hacera-did-driver) |

- [DID document](https://w3c-ccg.github.io/did-spec/#dfn-did-document)
  - A set of data that describes the subject of a DID, including mechanisms, such as public keys and pseudonymous biometrics, that the DID subject can use to authenticate itself and prove their association with the DID. A DID Document may also contain other attributes or claims describing the subject. These documents are graph-based data structures that are typically expressed using [JSON-LD], but may be expressed using other compatible graph-based data formats.
- [Contract function interface](https://github.com/ethereum/EIPs/issues/734)
  - The following describes standard functions for a key manager to be used in conjunction with ERC725. This contract can hold keys to sign actions (transactions, documents, logins, access, etc), as well as execute instructions through an ERC 725 proxy account.
```bash
contract ERC734 {
  uint256 constant MANAGEMENT_KEY = 1;
  uint256 constant EXECUTION_KEY = 2;

  event KeyAdded(bytes32 indexed key, uint256 indexed purpose, uint256 indexed keyType);
  event KeyRemoved(bytes32 indexed key, uint256 indexed purpose, uint256 indexed keyType);
  event ExecutionRequested(uint256 indexed executionId, address indexed to, uint256 indexed value, bytes data);
  event Executed(uint256 indexed executionId, address indexed to, uint256 indexed value, bytes data);
  event Approved(uint256 indexed executionId, bool approved);
  event KeysRequiredChanged(uint256 purpose, uint256 number);

  struct Key {
    uint256 purpose;
    uint256 keyType;
    bytes32 key;
  }

  function getKey(bytes32 _key) public constant returns(uint256[] purposes, uint256 keyType, bytes32 key);
  function keyHasPurpose(bytes32 _key, uint256 _purpose) public constant returns (bool exists);
  function getKeysByPurpose(uint256 _purpose) public constant returns (bytes32[] keys);
  function addKey(bytes32 _key, uint256 _purpose, uint256 _keyType) public returns (bool success);
  function removeKey(bytes32 _key, uint256 _purpose) public returns (bool success);
  function changeKeysRequired(uint256 purpose, uint256 number) external;
  function getKeysRequired(uint256 purpose) external view returns(uint256);
  function execute(address _to, uint256 _value, bytes _data) public returns (uint256 executionId);
  function approve(uint256 _id, bool _approve) public returns (bool success);
}
```

## Components
- Wallet
  - private key management
- DID Auth
  - private key ownership
- KYC
  - validate user's identity

## API
- API list
  - /api/v1/users/login
  - /api/v1/common/now
  - /api/v1/common/uuid
  - /api/v1/dids
  - /api/v1/emails/send
  - /api/v1/emails/activation
  - /api/v1/emails/{email}/did
  - /api/v1/emails/{email}/status
  - /api/v1/emails/{email}/resend
  - /api/v1/kyc
  - /api/v1/kyc/{did}/claim
  - /api/v1/kyc/{did}/query
  - /api/v1/kyc/{did}/erc725/claim
  - /api/v1/publickeys
  - /api/v1/users
  - /api/v1/users/authentication
- Start API
```
$ go run api/api.go
```

### API Example
- **/api/v1/common/now**
  - Get current timestamp
    - 用於 DID Auth challenge
  - method
    - GET
  - Call API
  ```
  $ curl -X GET \
  http://127.0.0.1:8080/api/v1/common/now \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json' \
  -H 'Host: 127.0.0.1:8080' \
  -H 'Postman-Token: ea418c2a-1d5e-480b-a4fa-2170c66dfcf7,d067b3bb-8b37-450c-b997-c797f42314b0' \
  -H 'User-Agent: PostmanRuntime/7.11.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache'
  ```
  - API response
  ```
  {
    "Result": 1557801487,
    "Status": 200,
    "Errors": "",
    "Time": "2019-05-14T02:38:07.251338041Z"
  }
  ```
- **/api/v1/common/uuid**
  - Get a random UUID
    - used as a mapping key for websocket connection
  - method
    - GET
  - Call API
  ```
  $ curl -X GET \
  http://127.0.0.1:8080/api/v1/common/uuid \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json' \
  -H 'Host: 127.0.0.1:8080' \
  -H 'Postman-Token: 26ef169c-3435-4e07-a694-cd3ae52c825b,deed5a83-aa35-4068-931d-7d2532fbe0b9' \
  -H 'User-Agent: PostmanRuntime/7.11.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache'
  ```
  - API response
  ```
  {
    "Result": "7a781f95-73af-4e59-90b9-586600ec5e88",
    "Status": 200,
    "Errors": "",
    "Time": "2019-05-14T02:44:10.539052653Z"
  }
  ```
- **/api/v1/dids**
  - Get address by DID
    - used for MetaMask login
  - method
    - GET
  - parameters
    - address
  - Call API
  ```
  $ curl -X GET \
  'http://127.0.0.1:8080/api/v1/dids?address=0x270830cE9169CB067d62bC232864E5bDe0Ba8721' \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json' \
  -H 'Host: 127.0.0.1:8080' \
  -H 'Postman-Token: 8168ecaf-3a4b-406c-a082-060d3d79710a,974a5af4-dcda-4e64-9bdf-c0759497f952' \
  -H 'User-Agent: PostmanRuntime/7.11.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache'
  ```
  - API response
    - 200
    ```
    {
      "Result": "did:erc725:ropsten:0x70316C92D7C7E10bfa258F4408F38a97f10dee7a",
      "Status": 200,
      "Errors": "",
      "Time": "2019-05-15T02:17:17.039140303Z"
    }
    ```
    - 400
    ```
    {
      "Result": "Failed",
      "Status": 400,
      "Errors": "Ethereum address not found.",
      "Time": "2019-05-15T02:23:32.088574854Z"
    }
    ```
- **/api/v1/emails/send**
  - Send verification email
  - method
    - POST
  - parameters
    - email
  - Call API
  ```
  $ curl -X POST \
  http://127.0.0.1:8080/api/v1/emails/send \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json' \
  -H 'Host: 127.0.0.1:8080' \
  -H 'Postman-Token: cbb7c25c-82c3-4803-bdce-17bcb1ee7b81,09c5416c-b770-4fc6-8cbd-e6938970c6ee' \
  -H 'User-Agent: PostmanRuntime/7.11.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache' \
  -H 'content-length: 36' \
  -d '{
  "email": "user@example.com"
  }'
  ```
  - API response
    - 200
    ```
    {
      "Result": "Success",
      "Status": 200,
      "Errors": "",
      "Time": "2019-05-15T03:08:22.116407848Z"
    }
    ```
    - 400
      - Email exists
    ```
    {
      "Result": "Failed",
      "Status": 400,
      "Errors": "Error 1062: Duplicate entry 'user@example.com' for key 'PRIMARY'; Verification code error.",
      "Time": "2019-05-15T02:26:15.596045708Z"
    }
    ```
- **/api/v1/emails/activation**
  - Activate user's account
  - method
    - PUT
  - parameters
    - email
    - code
  - Call API
  ```
  $ curl -X PUT \
  http://127.0.0.1:8080/api/v1/emails/activation \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json' \
  -H 'Host: 127.0.0.1:8080' \
  -H 'Postman-Token: 002cd192-ecbc-439a-a53c-da70758e8ddb,7ee8f196-3058-4e85-a869-52f9daffa787' \
  -H 'User-Agent: PostmanRuntime/7.11.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache' \
  -H 'content-length: 55' \
  -d '{
	"email": "user@example.com",
	"code": "130553"
  }'
  ```
  - API response
    - 200
    ```
    {
      "Result": "Success",
      "Status": 200,
      "Errors": "",
      "Time": "2019-05-15T08:18:35.596235577Z"
    }
    ```
    - 400
    ```
    {
      "Result": "Failed",
      "Status": 400,
      "Errors": "Email not found.",
      "Time": "2019-05-15T08:23:59.730348523Z"
    }
    ```
- **/api/v1/emails/{email}/did**
  - Get DID by email
  - method
    - GET
  - parameters
    - email
  - Call API
  ```
  $ curl -X GET \
  http://127.0.0.1:8080/api/v1/emails/user@example.com/did \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json' \
  -H 'Host: 127.0.0.1:8080' \
  -H 'Postman-Token: f286994b-987c-4049-a2ea-bc8b7ce4e60f,2c91d91d-72de-4719-864b-aa871c2a24e6' \
  -H 'User-Agent: PostmanRuntime/7.11.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache'
  ```
  - API response
    - 200
    ```
    {
      "Result": "did:erc725:ropsten:0x70316C92D7C7E10bfa258F4408F38a97f10dee7a",
      "Status": 200,
      "Errors": "",
      "Time": "2019-05-15T03:54:16.449318658Z"
    }
    ```
    - 400
    ```
    {
      "Result": "Failed",
      "Status": 400,
      "Errors": "Email not found.",
      "Time": "2019-05-15T04:00:33.002841961Z"
    }
    ```
- **/api/v1/emails/{email}/resend**
  - Resend verification email
  - method
    - GET
  - parameters
    - email
  - Call API
  ```
  $ curl -X GET \
  http://127.0.0.1:8080/api/v1/emails/user@example.com/resend \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json' \
  -H 'Host: 127.0.0.1:8080' \
  -H 'Postman-Token: edbe1b9d-7c02-4419-850f-2bfc177c047f,15d92364-ea5a-47a0-b74a-b8818f26a470' \
  -H 'User-Agent: PostmanRuntime/7.11.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache'
  ```
  - API response
    - 200
    ```
    {
      "Result": "Success",
      "Status": 200,
      "Errors": "",
      "Time": "2019-05-15T04:05:10.325914442Z"
    }
    ```
    - 400
    ```
    {
      "Result": "Failed",
      "Status": 400,
      "Errors": "Email not found.; Send email error.",
      "Time": "2019-05-15T04:06:36.973290483Z"
    }
    ```
- **/api/v1/emails/{email}/status**
  - Get email verification status
  - method
    - GET
  - parameters
    - email
  - Call API
  ```
  $ curl -X GET \
  http://127.0.0.1:8080/api/v1/emails/user@example.com/status \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json' \
  -H 'Host: 127.0.0.1:8080' \
  -H 'Postman-Token: 07ab78b8-6c5f-4361-bc0d-377dbf9d9a06,9addab34-97dd-4cc3-862d-7d0a1307df35' \
  -H 'User-Agent: PostmanRuntime/7.11.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache'
  ```
  - API response
    - 200
    ```
    {
      "Result": "False",
      "Status": 200,
      "Errors": "",
      "Time": "2019-05-15T04:12:10.634070066Z"
    }
    ```
    - 400
    ```
    {
      "Result": "Failed",
      "Status": 400,
      "Errors": "Email not found.",
      "Time": "2019-05-15T04:13:10.73916517Z"
    }
    ```
- **/api/v1/kyc**
  - Issuer signs on an investor's KYC claim
  - method
    - POST
  - parameters
    - claim
    - signature
  - Call API
  ```
  $ curl -X POST \
  http://127.0.0.1:8080/api/v1/kyc \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json' \
  -H 'Host: 127.0.0.1:8080' \
  -H 'Postman-Token: 61d41061-e187-42e3-9593-d7da9ac9a254,1463b565-fa57-4697-b428-915959784762' \
  -H 'User-Agent: PostmanRuntime/7.13.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache' \
  -H 'content-length: 420' \
  -d '{
	"claim": "eyJJREh1YkRJRCI6ImRpZDplcmM3MjU6cm9wc3RlbjoweEIxRkY3N0MxNEY5ZWQxM2NiQzc0MWQ3MDc3ODA2MTMzN2E2YTc2QjAiLCJpbnZlc3RvckRJRCI6ImRpZDplcmM3MjU6cm9wc3RlbjoweENhN0RFNzFjQTBDMUVDQ2JlOUQzYTE4RkI1OUZCZDREMjRhN2Y1NTUiLCJzdGF0dXMiOjAsInRpbWVMaW1pdCI6IjIwMTkuMDkuMDkifQ",
	"signature": "0x2e858beba6682d465886b00ef2c2107b3044cfed33146f23699e89995e8d684002328a10873f9f8c3fe49159c36eb5ce08beb8c10dead1a3547b5d22d89df3f901"
  }'
  ```
  - API response
  ```
  {
    "Result": "Success",
    "Status": 200,
    "Errors": "",
    "Time": "2019-06-13T01:59:02.276354528Z"
  }
  ```
- **/api/v1/kyc/{did}/claim**
  - Issuer generates a claim for an investor's KYC
  - method
    - GET
  - parameters
    - did
    - expiration
    - status
      - approved
      - rejected
      - pending
  - returns
    - KYC claim JSON includes four keys:
      - recipientDID
      - investorDID
      - KYCStatus
      - expirationDate
  - Call API
  ```
  $ curl -X GET \
  'http://127.0.0.1:8080/api/v1/kyc/did:erc725:ropsten:0xCa7DE71cA0C1ECCbe9D3a18FB59FBd4D24a7f555/claim?expiration=20190909&status=0' \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Host: 127.0.0.1:8080' \
  -H 'Postman-Token: d6a688ea-27d9-4fed-a0bc-f55f74836461,a015352a-4298-4b0d-9f21-c57ce7982d9d' \
  -H 'User-Agent: PostmanRuntime/7.13.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache'
  ```
  - API response
  ```
  {
    "Result": "eyJJREh1YkRJRCI6ImRpZDplcmM3MjU6cm9wc3RlbjoweEIxRkY3N0MxNEY5ZWQxM2NiQzc0MWQ3MDc3ODA2MTMzN2E2YTc2QjAiLCJLWUNzdGF0dXMiOjAsImV4cGlyYXRpb25EYXRlIjoiMjAxOTA5MDkiLCJpbnZlc3RvckRJRCI6ImRpZDplcmM3MjU6cm9wc3RlbjoweENhN0RFNzFjQTBDMUVDQ2JlOUQzYTE4RkI1OUZCZDREMjRhN2Y1NTUifQ",
    "Status": 200,
    "Errors": "",
    "Time": "2019-06-13T01:55:59.774548642Z"
  }
  ```
- **/api/v1/kyc/{did}/query**
  - Query investor's claim
  - method
    - GET
  - parameters
    - did
    - type
  - Call API
  ```
  $ curl -X GET \
  'http://127.0.0.1:8080/api/v1/kyc/did:erc725:ropsten:0xCa7DE71cA0C1ECCbe9D3a18FB59FBd4D24a7f555/query?type=0' \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Host: 127.0.0.1:8080' \
  -H 'Postman-Token: 21c0ae89-5e12-4467-8470-41ecc8cb8f15,b51e931d-d6b7-4b10-a600-b65bf353df51' \
  -H 'User-Agent: PostmanRuntime/7.13.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache'
  ```
  - API response
  ```
  {
    "Result": "{\"recipientDID\":\"did:erc725:ropsten:0xB1FF77C14F9ed13cbC741d70778061337a6a76B0\",\"investorDID\":\"did:erc725:ropsten:0xCa7DE71cA0C1ECCbe9D3a18FB59FBd4D24a7f555\",\"status\":0,\"expirationDate\":\"2019.09.09\"}",
    "Status": 200,
    "Errors": "",
    "Time": "2019-06-13T02:00:34.994785315Z"
  }
  ```
- **/api/v1/kyc/{did}/erc725/claim**
  - Query investor's ERC-275 claim by DID on blockchain
  - method
    - GET
  - parameters
    - did
  - Call API
  ```
  $ curl -X GET \
  http://127.0.0.1:8080/api/v1/kyc/did:erc725:ropsten:0xCa7DE71cA0C1ECCbe9D3a18FB59FBd4D24a7f555/erc725/claim \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Host: 127.0.0.1:8080' \
  -H 'Postman-Token: 5df588dc-741e-463e-a35b-77a84b16e3c8,35f0cfac-6467-44e2-bb74-021f63ce7740' \
  -H 'User-Agent: PostmanRuntime/7.13.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache'
  ```
  - API response
  ```
  {
    "Result": "{\"recipientDID\":\"did:erc725:ropsten:0xB1FF77C14F9ed13cbC741d70778061337a6a76B0\",\"investorDID\":\"did:erc725:ropsten:0xCa7DE71cA0C1ECCbe9D3a18FB59FBd4D24a7f555\",\"status\":0,\"expirationDate\":\"2019.09.09\"}",
    "Status": 200,
    "Errors": "",
    "Time": "2019-06-13T02:02:01.697746696Z"
  }
  ```
- **/api/v1/publickeys**
  - Mapping an address to a public key
  - method
    - GET
  - parameters
    - address
  - Call API
  ```
  $ curl -X GET \
  'http://127.0.0.1:8080/api/v1/publickeys?address=0x270830cE9169CB067d62bC232864E5bDe0Ba8721' \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json' \
  -H 'Host: 127.0.0.1:8080' \
  -H 'Postman-Token: 60024857-ff03-4f5f-b9b6-1e3d6672ee34,1b1c7efe-664a-4e43-b2a8-36952a8148ae' \
  -H 'User-Agent: PostmanRuntime/7.11.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache'
  ```
  - API response
  ```
  {
    "Result": "4a5be29364d66788174ef43d9caacfa95c05cefe0133387520e16f61b5dcad5d",
    "Status": 200,
    "Errors": "",
    "Time": "2019-05-15T04:20:07.70286819Z"
  }
  ```
- **/api/v1/users**
  - Get
  ```
  $ curl -X GET \
  'http://127.0.0.1:8080/api/v1/users?email=%27user@example.com%27&fields=did,address' \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json' \
  -H 'Host: 127.0.0.1:8080' \
  -H 'Postman-Token: c8d7fb13-4ac5-46ee-9bd3-7f22d89be020,da208664-3997-4998-affd-74fbb2478781' \
  -H 'User-Agent: PostmanRuntime/7.11.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache'
  ```
    - API response
      - 200
      ```
      {
        "Result": "[[\"did:erc725:ropsten:0x70316C92D7C7E10bfa258F4408F38a97f10dee7a\",\"0x270830cE9169CB067d62bC232864E5bDe0Ba8721\"]]",
        "Status": 200,
        "Errors": "",
        "Time": "2019-05-15T04:36:12.50597636Z"
      }
      ```
      - 400
      ```
      {
        "Result": "Failed",
        "Status": 400,
        "Errors": "Error 1054: Unknown column 'address' in 'field list'",
        "Time": "2019-05-15T04:38:01.238351542Z"
      }
      ```
  - Add
    - method
      - POST
    - parameters
      - JWT
        - content
        ```
        header['alg'] = 'Secp256k1'
        header['typ'] = 'JWT'
        payload['did'] = 'did:erc725:ropsten:0x70316C92D7C7E10bfa258F4408F38a97f10dee7a'
        payload['Address'] = '0x270830cE9169CB067d62bC232864E5bDe0Ba8721'
        payload['Email'] = 'user@example.com'
        payload['Nationality'] = 'Taiwan'
        payload['timestamp'] = int(time.time())
        ```
    - Call API
    ```
    $ curl -X POST \
    http://127.0.0.1:8080/api/v1/users \
    -H 'Accept: */*' \
    -H 'Cache-Control: no-cache' \
    -H 'Connection: keep-alive' \
    -H 'Content-Type: application/json' \
    -H 'Host: 127.0.0.1:8080' \
    -H 'Postman-Token: 8f1ba4dd-4e05-4c01-9bb2-6b84a373ac8c,0c2e9cc8-f871-4365-a870-b1a02d1c7878' \
    -H 'User-Agent: PostmanRuntime/7.11.0' \
    -H 'accept-encoding: gzip, deflate' \
    -H 'cache-control: no-cache' \
    -H 'content-length: 402' \
    -d '{
	"jwt": "eyJhbGciOiAiU2VjcDI1NmsxIiwgInR5cCI6ICJKV1QifQ.eyJkaWQiOiAiZGlkOmVyYzcyNTpyb3BzdGVuOjB4NzAzMTZDOTJEN0M3RTEwYmZhMjU4RjQ0MDhGMzhhOTdmMTBkZWU3YSIsICJOYXRpb25hbGl0eSI6ICJBbWVyaWNhIiwgInRpbWVzdGFtcCI6IDE1NTYwODc2OTMsICJFbWFpbCI6ICJ5ZW5rdWFubGVlQGdtYWlsLmNvbSJ9.0xe1c7267ff5b82eb1af3e80cab7730726ed788fe4b63a3aaf3c3976c5b4d6e7b7561be510b5f81589fd8708157f6222c3d1d42bf978d6de6d5be1c27a39f2bf2401"
    }'
    ```
    - API response
      - 200
      ```
      {
        "Result": "Success",
        "Status": 200,
        "Errors": "",
        "Time": "2019-05-15T04:52:09.933413999Z"
      }
      ```
      - 400
      ```
      {
        "Result": "Failed",
        "Status": 400,
        "Errors": "Error 1062: Duplicate entry 'did:erc725:ropsten:0x70316C92D7C7E10bfa258F4408F38a97f10dee7a' for key 'PRIMARY'",
        "Time": "2019-05-15T04:50:04.666030975Z"
      }
      ```
  - Update
    - method
      - PUT
    - parameters
      - JWT
    - Call API
    ```
    $ curl -X PUT \
    http://127.0.0.1:8080/api/v1/users \
    -H 'Accept: */*' \
    -H 'Cache-Control: no-cache' \
    -H 'Connection: keep-alive' \
    -H 'Content-Type: application/json' \
    -H 'Host: 127.0.0.1:8080' \
    -H 'Postman-Token: 6b44c8f2-1cc7-43c6-9490-0e3e5befdc37,9185ae58-f815-4e12-847a-16c6b13c3a07' \
    -H 'User-Agent: PostmanRuntime/7.11.0' \
    -H 'accept-encoding: gzip, deflate' \
    -H 'cache-control: no-cache' \
    -H 'content-length: 402' \
    -d '{
	"jwt": "eyJhbGciOiAiU2VjcDI1NmsxIiwgInR5cCI6ICJKV1QifQ.eyJkaWQiOiAiZGlkOmVyYzcyNTpyb3BzdGVuOjB4NzAzMTZDOTJEN0M3RTEwYmZhMjU4RjQ0MDhGMzhhOTdmMTBkZWU3YSIsICJOYXRpb25hbGl0eSI6ICJBbWVyaWNhIiwgInRpbWVzdGFtcCI6IDE1NTYwODc2OTMsICJFbWFpbCI6ICJ5ZW5rdWFubGVlQGdtYWlsLmNvbSJ9.0xe1c7267ff5b82eb1af3e80cab7730726ed788fe4b63a3aaf3c3976c5b4d6e7b7561be510b5f81589fd8708157f6222c3d1d42bf978d6de6d5be1c27a39f2bf2401"
    }'
    ```
    - API response
    ```
    {
      "Result": "Success",
      "Status": 200,
      "Errors": "",
      "Time": "2019-05-15T04:53:51.895958901Z"
    }
    ```
- **/api/v1/users/authentication**
  - method
    - GET
  - parameters
    - JWT
    ```
    $ go run pkg/jwt/jwt.go
    output:
        eyJ0eXAiOiAiSldUIiwgImFsZyI6ICJTZWNwMjU2azEifQ.eyJ0aW1lc3RhbXAiOiAxNTU3ODk2NzQ5LCAiZGlkIjogImRpZDplcmM3MjU6cm9wc3RlbjoweDcwMzE2QzkyRDdDN0UxMGJmYTI1OEY0NDA4RjM4YTk3ZjEwZGVlN2EiLCAidXVpZCI6ICI3YTc4MWY5NS03M2FmLTRlNTktOTBiOS01ODY2MDBlYzVlODgifQ.0x3164a16843d14e26b1d5a7fcdb61c3d38885c1241bb128e99169ed108aae66715b961f95d6c2e3556904789ca9ac3e78c7e0e3a19214fb14c59fa7ca192d03f600
    ```
  - Call API
  ```
  $ curl -X GET \
  'http://127.0.0.1:8080/api/v1/users/authentication?jwt=eyJ0eXAiOiAiSldUIiwgImFsZyI6ICJTZWNwMjU2azEifQ.eyJ0aW1lc3RhbXAiOiAxNTU3ODk2NzQ5LCAiZGlkIjogImRpZDplcmM3MjU6cm9wc3RlbjoweDcwMzE2QzkyRDdDN0UxMGJmYTI1OEY0NDA4RjM4YTk3ZjEwZGVlN2EiLCAidXVpZCI6ICI3YTc4MWY5NS03M2FmLTRlNTktOTBiOS01ODY2MDBlYzVlODgifQ.0x3164a16843d14e26b1d5a7fcdb61c3d38885c1241bb128e99169ed108aae66715b961f95d6c2e3556904789ca9ac3e78c7e0e3a19214fb14c59fa7ca192d03f600' \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json' \
  -H 'Host: 127.0.0.1:8080' \
  -H 'Postman-Token: cd31b0e1-4b6e-479c-994e-047a7d148092,063d7bbf-abad-446a-93a6-02483dc86bdd' \
  -H 'User-Agent: PostmanRuntime/7.11.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache'
  ```
  - API response
    - 200
    ```
    {
      "Result": "Success",
      "Status": 200,
      "Errors": "",
      "Time": "2019-05-15T05:06:12.62067406Z"
    }
    ```
    - 400
    ```
    {
      "Result": "Failed",
      "Status": 400,
      "Errors": "JWT timeout Error.",
      "Time": "2019-05-15T05:07:23.382054211Z"
    }
    ```
- **/api/v1/users/login**
  - websocket
  ```
  c, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8080/api/v1/users/login", nil)
  err = c.WriteMessage(websocket.TextMessage, []byte(UUID))
  ```
  - pkg/jwt/jwt.go
  ```
  header['alg'] = 'Secp256k1'
  header['typ'] = 'JWT'
  payload['did'] = 'did:erc725:ropsten:0x70316C92D7C7E10bfa258F4408F38a97f10dee7a'
  payload['timestamp'] = int(time.time())
  payload['uuid'] = '7a781f95-73af-4e59-90b9-586600ec5e88'
  ```
