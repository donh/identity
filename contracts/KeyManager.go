// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KeyManager

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// KeyManagerABI is the input ABI used to generate the binding from.
const KeyManagerABI = "[{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"payable\":false,\"inputs\":[]},{\"name\":\"KeyAdded\",\"type\":\"event\",\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"purpose\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"keyType\",\"type\":\"uint256\"}]},{\"name\":\"KeyRemoved\",\"type\":\"event\",\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"purpose\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"keyType\",\"type\":\"uint256\"}]},{\"name\":\"ExecutionRequested\",\"type\":\"event\",\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executionID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}]},{\"name\":\"Executed\",\"type\":\"event\",\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executionID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}]},{\"name\":\"Approved\",\"type\":\"event\",\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executionID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}]},{\"name\":\"KeysRequiredChanged\",\"type\":\"event\",\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"purpose\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"number\",\"type\":\"uint256\"}]},{\"name\":\"ClaimAdded\",\"type\":\"event\",\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"claimType\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"scheme\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"uri\",\"type\":\"string\"}]},{\"name\":\"ClaimRemoved\",\"type\":\"event\",\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"claimType\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"scheme\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"uri\",\"type\":\"string\"}]},{\"name\":\"ClaimChanged\",\"type\":\"event\",\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"claimType\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"scheme\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"uri\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"payable\":false,\"outputs\":[{\"name\":\"isExistent\",\"type\":\"bool\"}],\"name\":\"keyHasPurpose\",\"constant\":true,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_purpose\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"payable\":false,\"outputs\":[{\"name\":\"purpose\",\"type\":\"uint256\"},{\"name\":\"keyType\",\"type\":\"uint256\"},{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getKey\",\"constant\":true,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"payable\":false,\"outputs\":[{\"name\":\"_keys\",\"type\":\"bytes32[]\"}],\"name\":\"getKeysByPurpose\",\"constant\":true,\"inputs\":[{\"name\":\"_purpose\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"payable\":false,\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"addKey\",\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_purpose\",\"type\":\"uint256\"},{\"name\":\"_keyType\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"payable\":false,\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"removeKey\",\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_purpose\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"payable\":false,\"outputs\":[{\"name\":\"claimType\",\"type\":\"uint256\"},{\"name\":\"issuer\",\"type\":\"address\"},{\"name\":\"signatureType\",\"type\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\"},{\"name\":\"claim\",\"type\":\"bytes\"},{\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"getClaim\",\"constant\":true,\"inputs\":[{\"name\":\"_claimId\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"payable\":false,\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"name\":\"getClaimsIdByType\",\"constant\":true,\"inputs\":[{\"name\":\"_claimType\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"payable\":false,\"outputs\":[{\"name\":\"claimId\",\"type\":\"bytes32\"}],\"name\":\"addClaim\",\"constant\":false,\"inputs\":[{\"name\":\"_claimType\",\"type\":\"uint256\"},{\"name\":\"_issuer\",\"type\":\"address\"},{\"name\":\"_signatureType\",\"type\":\"uint256\"},{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_claim\",\"type\":\"bytes\"},{\"name\":\"_uri\",\"type\":\"string\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"payable\":false,\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"removeClaim\",\"constant\":false,\"inputs\":[{\"name\":\"_claimId\",\"type\":\"bytes32\"}]}]"

// KeyManagerBin is the compiled bytecode used for deploying new contracts.
const KeyManagerBin = `0x608060405234801561001057600080fd5b50600033604051602001808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014019150506040516020818303038152906040528051906020012090506060604051908101604052806001815260200160018152602001828152506000808381526020019081526020016000206000820151816000015560208201518160010155604082015181600201559050506001600060018152602001908152602001600020819080600181540180825580915050906001820390600052602060002001600090919290919091505550600180827f480000bb1edad8ca1470381cc334b1917fbd51c6531f3a623ea8e0ec7e38a6e960405160405180910390a450611c09806101466000396000f3fe608060405234801561001057600080fd5b50600436106100b0576000357c01000000000000000000000000000000000000000000000000000000009004806353d413c51161008357806353d413c5146103d65780636fa28249146104265780639010f726146104a9578063c9100bcb1461052c578063d202158d146106ec576100b0565b80630607f937146100b557806312aaac70146102e65780631d381240146103365780634eee424a14610390575b600080fd5b6102d0600480360360c08110156100cb57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561011c57600080fd5b82018360208201111561012e57600080fd5b8035906020019184600183028401116401000000008311171561015057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156101b357600080fd5b8201836020820111156101c557600080fd5b803590602001918460018302840111640100000000831117156101e757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561024a57600080fd5b82018360208201111561025c57600080fd5b8035906020019184600183028401116401000000008311171561027e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061073c565b6040518082815260200191505060405180910390f35b610312600480360360208110156102fc57600080fd5b8101908080359060200190929190505050610a94565b60405180848152602001838152602001828152602001935050505060405180910390f35b6103766004803603606081101561034c57600080fd5b81019080803590602001909291908035906020019092919080359060200190929190505050610ae8565b604051808215151515815260200191505060405180910390f35b6103bc600480360360208110156103a657600080fd5b8101908080359060200190929190505050610ddc565b604051808215151515815260200191505060405180910390f35b61040c600480360360408110156103ec57600080fd5b8101908080359060200190929190803590602001909291905050506111d9565b604051808215151515815260200191505060405180910390f35b6104526004803603602081101561043c57600080fd5b8101908080359060200190929190505050611552565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561049557808201518184015260208101905061047a565b505050509050019250505060405180910390f35b6104d5600480360360208110156104bf57600080fd5b81019080803590602001909291905050506115bd565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156105185780820151818401526020810190506104fd565b505050509050019250505060405180910390f35b6105586004803603602081101561054257600080fd5b8101908080359060200190929190505050611628565b604051808781526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001858152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b838110156105de5780820151818401526020810190506105c3565b50505050905090810190601f16801561060b5780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b83811015610644578082015181840152602081019050610629565b50505050905090810190601f1680156106715780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b838110156106aa57808201518184015260208101905061068f565b50505050905090810190601f1680156106d75780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b6107226004803603604081101561070257600080fd5b8101908080359060200190929190803590602001909291905050506118e8565b604051808215151515815260200191505060405180910390f35b60008587604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018281526020019250505060405160208183030381529060405280519060200120905060c0604051908101604052808881526020018773ffffffffffffffffffffffffffffffffffffffff16815260200186815260200185815260200184815260200183815250600260008381526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201556060820151816003019080519060200190610878929190611944565b506080820151816004019080519060200190610895929190611944565b5060a08201518160050190805190602001906108b29291906119c4565b50905050600360008881526020019081526020016000208190806001815401808255809150509060018203906000526020600020016000909192909190915055508573ffffffffffffffffffffffffffffffffffffffff1687827f46149b18aa084502c3f12bc75e19eda8bda8d102b82cce8474677a6d0d5f43c58888888860405180858152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561097f578082015181840152602081019050610964565b50505050905090810190601f1680156109ac5780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b838110156109e55780820151818401526020810190506109ca565b50505050905090810190601f168015610a125780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b83811015610a4b578082015181840152602081019050610a30565b50505050905090810190601f168015610a785780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a49695505050505050565b60008060008060008581526020019081526020016000206000015460008086815260200190815260200160002060010154600080878152602001908152602001600020600201549250925092509193909250565b6000836000808681526020019081526020016000206002015414151515610b77576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f4b657920616c7265616479206578697374732e0000000000000000000000000081525060200191505060405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610d16573073ffffffffffffffffffffffffffffffffffffffff1663d202158d33604051602001808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014019150506040516020818303038152906040528051906020012060016040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808381526020018281526020019250505060206040518083038186803b158015610c8357600080fd5b505afa158015610c97573d6000803e3d6000fd5b505050506040513d6020811015610cad57600080fd5b81019080805190602001909291905050501515610d15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180611b986023913960400191505060405180910390fd5b5b60606040519081016040528084815260200183815260200185815250600080868152602001908152602001600020600082015181600001556020820151816001015560408201518160020155905050600160008481526020019081526020016000208490806001815401808255809150509060018203906000526020600020016000909192909190915055508183857f480000bb1edad8ca1470381cc334b1917fbd51c6531f3a623ea8e0ec7e38a6e960405160405180910390a4600190509392505050565b6000806003600060026000868152602001908152602001600020600001548152602001908152602001600020905060008090505b8180549050811015610ec257838282815481101515610e2b57fe5b90600052602060002001541415610eb557816001838054905003815481101515610e5157fe5b90600052602060002001548282815481101515610e6a57fe5b9060005260206000200181905550816001838054905003815481101515610e8d57fe5b906000526020600020016000905581805480919060019003610eaf9190611a44565b50610ec2565b8080600101915050610e10565b50600260008481526020019081526020016000206000808201600090556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160009055600382016000610f1e9190611a70565b600482016000610f2e9190611a70565b600582016000610f3e9190611ab8565b50506002600084815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166002600085815260200190815260200160002060000154847f3cf57863a89432c61c4a27073c6ee39e8a764bff5a05aebfbcdcdc80b2e6130a600260008881526020019081526020016000206002015460026000898152602001908152602001600020600301600260008a8152602001908152602001600020600401600260008b8152602001908152602001600020600501604051808581526020018060200180602001806020018481038452878181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156110b65780601f1061108b576101008083540402835291602001916110b6565b820191906000526020600020905b81548152906001019060200180831161109957829003601f168201915b50508481038352868181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156111395780601f1061110e57610100808354040283529160200191611139565b820191906000526020600020905b81548152906001019060200180831161111c57829003601f168201915b50508481038252858181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156111bc5780601f10611191576101008083540402835291602001916111bc565b820191906000526020600020905b81548152906001019060200180831161119f57829003601f168201915b505097505050505050505060405180910390a46001915050919050565b60008260008085815260200190815260200160002060020154141515611267576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f4e6f2073756368206b65792e000000000000000000000000000000000000000081525060200191505060405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611406573073ffffffffffffffffffffffffffffffffffffffff1663d202158d33604051602001808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014019150506040516020818303038152906040528051906020012060016040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808381526020018281526020019250505060206040518083038186803b15801561137357600080fd5b505afa158015611387573d6000803e3d6000fd5b505050506040513d602081101561139d57600080fd5b81019080805190602001909291905050501515611405576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180611bbb6023913960400191505060405180910390fd5b5b600080848152602001908152602001600020600080820160009055600182016000905560028201600090555050600060016000848152602001908152602001600020905060008090505b81805490508110156115025784828281548110151561146b57fe5b906000526020600020015414156114f55781600183805490500381548110151561149157fe5b906000526020600020015482828154811015156114aa57fe5b90600052602060002001819055508160018380549050038154811015156114cd57fe5b9060005260206000200160009055818054809190600190036114ef9190611a44565b50611502565b8080600101915050611450565b506000808581526020019081526020016000206001015483857f585a4aef50f8267a92b32412b331b20f7f8b96f2245b253b9cc50dcc621d339760405160405180910390a4600191505092915050565b6060600360008381526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156115b157602002820191906000526020600020905b81548152602001906001019080831161159d575b50505050509050919050565b60606001600083815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561161c57602002820191906000526020600020905b815481526020019060010190808311611608575b50505050509050919050565b6000806000606080606061163a611b00565b6002600089815260200190815260200160002060c06040519081016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201548152602001600382018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561175c5780601f106117315761010080835404028352916020019161175c565b820191906000526020600020905b81548152906001019060200180831161173f57829003601f168201915b50505050508152602001600482018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156117fe5780601f106117d3576101008083540402835291602001916117fe565b820191906000526020600020905b8154815290600101906020018083116117e157829003601f168201915b50505050508152602001600582018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156118a05780601f10611875576101008083540402835291602001916118a0565b820191906000526020600020905b81548152906001019060200180831161188357829003601f168201915b5050505050815250509050806000015181602001518260400151836060015184608001518560a001518292508191508090509650965096509650965096505091939550919395565b600080600102600080858152602001908152602001600020600201541415611913576000905061193e565b81600080858152602001908152602001600020600001541415611939576001905061193e565b600090505b92915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061198557805160ff19168380011785556119b3565b828001600101855582156119b3579182015b828111156119b2578251825591602001919060010190611997565b5b5090506119c09190611b4d565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611a0557805160ff1916838001178555611a33565b82800160010185558215611a33579182015b82811115611a32578251825591602001919060010190611a17565b5b509050611a409190611b4d565b5090565b815481835581811115611a6b57818360005260206000209182019101611a6a9190611b72565b5b505050565b50805460018160011615610100020316600290046000825580601f10611a965750611ab5565b601f016020900490600052602060002090810190611ab49190611b4d565b5b50565b50805460018160011615610100020316600290046000825580601f10611ade5750611afd565b601f016020900490600052602060002090810190611afc9190611b4d565b5b50565b60c06040519081016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160608152602001606081525090565b611b6f91905b80821115611b6b576000816000905550600101611b53565b5090565b90565b611b9491905b80821115611b90576000816000905550600101611b78565b5090565b9056fe53656e64657220646f7365206e6f7420686176652061206d616e61676572206b65792e53656e64657220646f6573206e6f7420686176652061206d616e61676572206b65792ea165627a7a723058204c8d6feae7ca5ce9e8b5dd4d5611831e5110ff6b24e75fad57fd9ec17f2a09e80029`

// DeployKeyManager deploys a new Ethereum contract, binding an instance of KeyManager to it.
func DeployKeyManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KeyManager, error) {
	parsed, err := abi.JSON(strings.NewReader(KeyManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KeyManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KeyManager{KeyManagerCaller: KeyManagerCaller{contract: contract}, KeyManagerTransactor: KeyManagerTransactor{contract: contract}, KeyManagerFilterer: KeyManagerFilterer{contract: contract}}, nil
}

// KeyManager is an auto generated Go binding around an Ethereum contract.
type KeyManager struct {
	KeyManagerCaller     // Read-only binding to the contract
	KeyManagerTransactor // Write-only binding to the contract
	KeyManagerFilterer   // Log filterer for contract events
}

// KeyManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeyManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeyManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeyManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeyManagerSession struct {
	Contract     *KeyManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeyManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeyManagerCallerSession struct {
	Contract *KeyManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// KeyManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeyManagerTransactorSession struct {
	Contract     *KeyManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// KeyManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeyManagerRaw struct {
	Contract *KeyManager // Generic contract binding to access the raw methods on
}

// KeyManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeyManagerCallerRaw struct {
	Contract *KeyManagerCaller // Generic read-only contract binding to access the raw methods on
}

// KeyManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeyManagerTransactorRaw struct {
	Contract *KeyManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeyManager creates a new instance of KeyManager, bound to a specific deployed contract.
func NewKeyManager(address common.Address, backend bind.ContractBackend) (*KeyManager, error) {
	contract, err := bindKeyManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeyManager{KeyManagerCaller: KeyManagerCaller{contract: contract}, KeyManagerTransactor: KeyManagerTransactor{contract: contract}, KeyManagerFilterer: KeyManagerFilterer{contract: contract}}, nil
}

// NewKeyManagerCaller creates a new read-only instance of KeyManager, bound to a specific deployed contract.
func NewKeyManagerCaller(address common.Address, caller bind.ContractCaller) (*KeyManagerCaller, error) {
	contract, err := bindKeyManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeyManagerCaller{contract: contract}, nil
}

// NewKeyManagerTransactor creates a new write-only instance of KeyManager, bound to a specific deployed contract.
func NewKeyManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*KeyManagerTransactor, error) {
	contract, err := bindKeyManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeyManagerTransactor{contract: contract}, nil
}

// NewKeyManagerFilterer creates a new log filterer instance of KeyManager, bound to a specific deployed contract.
func NewKeyManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*KeyManagerFilterer, error) {
	contract, err := bindKeyManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeyManagerFilterer{contract: contract}, nil
}

// bindKeyManager binds a generic wrapper to an already deployed contract.
func bindKeyManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KeyManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyManager *KeyManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KeyManager.Contract.KeyManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyManager *KeyManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyManager.Contract.KeyManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyManager *KeyManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyManager.Contract.KeyManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyManager *KeyManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KeyManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyManager *KeyManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyManager *KeyManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyManager.Contract.contract.Transact(opts, method, params...)
}

// GetClaim is a free data retrieval call binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(bytes32 _claimId) constant returns(uint256 claimType, address issuer, uint256 signatureType, bytes signature, bytes claim, string uri)
func (_KeyManager *KeyManagerCaller) GetClaim(opts *bind.CallOpts, _claimId [32]byte) (struct {
	ClaimType     *big.Int
	Issuer        common.Address
	SignatureType *big.Int
	Signature     []byte
	Claim         []byte
	Uri           string
}, error) {
	ret := new(struct {
		ClaimType     *big.Int
		Issuer        common.Address
		SignatureType *big.Int
		Signature     []byte
		Claim         []byte
		Uri           string
	})
	out := ret
	err := _KeyManager.contract.Call(opts, out, "getClaim", _claimId)
	return *ret, err
}

// GetClaim is a free data retrieval call binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(bytes32 _claimId) constant returns(uint256 claimType, address issuer, uint256 signatureType, bytes signature, bytes claim, string uri)
func (_KeyManager *KeyManagerSession) GetClaim(_claimId [32]byte) (struct {
	ClaimType     *big.Int
	Issuer        common.Address
	SignatureType *big.Int
	Signature     []byte
	Claim         []byte
	Uri           string
}, error) {
	return _KeyManager.Contract.GetClaim(&_KeyManager.CallOpts, _claimId)
}

// GetClaim is a free data retrieval call binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(bytes32 _claimId) constant returns(uint256 claimType, address issuer, uint256 signatureType, bytes signature, bytes claim, string uri)
func (_KeyManager *KeyManagerCallerSession) GetClaim(_claimId [32]byte) (struct {
	ClaimType     *big.Int
	Issuer        common.Address
	SignatureType *big.Int
	Signature     []byte
	Claim         []byte
	Uri           string
}, error) {
	return _KeyManager.Contract.GetClaim(&_KeyManager.CallOpts, _claimId)
}

// GetClaimsIdByType is a free data retrieval call binding the contract method 0x6fa28249.
//
// Solidity: function getClaimsIdByType(uint256 _claimType) constant returns(bytes32[])
func (_KeyManager *KeyManagerCaller) GetClaimsIdByType(opts *bind.CallOpts, _claimType *big.Int) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _KeyManager.contract.Call(opts, out, "getClaimsIdByType", _claimType)
	return *ret0, err
}

// GetClaimsIdByType is a free data retrieval call binding the contract method 0x6fa28249.
//
// Solidity: function getClaimsIdByType(uint256 _claimType) constant returns(bytes32[])
func (_KeyManager *KeyManagerSession) GetClaimsIdByType(_claimType *big.Int) ([][32]byte, error) {
	return _KeyManager.Contract.GetClaimsIdByType(&_KeyManager.CallOpts, _claimType)
}

// GetClaimsIdByType is a free data retrieval call binding the contract method 0x6fa28249.
//
// Solidity: function getClaimsIdByType(uint256 _claimType) constant returns(bytes32[])
func (_KeyManager *KeyManagerCallerSession) GetClaimsIdByType(_claimType *big.Int) ([][32]byte, error) {
	return _KeyManager.Contract.GetClaimsIdByType(&_KeyManager.CallOpts, _claimType)
}

// GetKey is a free data retrieval call binding the contract method 0x12aaac70.
//
// Solidity: function getKey(bytes32 _key) constant returns(uint256 purpose, uint256 keyType, bytes32 key)
func (_KeyManager *KeyManagerCaller) GetKey(opts *bind.CallOpts, _key [32]byte) (struct {
	Purpose *big.Int
	KeyType *big.Int
	Key     [32]byte
}, error) {
	ret := new(struct {
		Purpose *big.Int
		KeyType *big.Int
		Key     [32]byte
	})
	out := ret
	err := _KeyManager.contract.Call(opts, out, "getKey", _key)
	return *ret, err
}

// GetKey is a free data retrieval call binding the contract method 0x12aaac70.
//
// Solidity: function getKey(bytes32 _key) constant returns(uint256 purpose, uint256 keyType, bytes32 key)
func (_KeyManager *KeyManagerSession) GetKey(_key [32]byte) (struct {
	Purpose *big.Int
	KeyType *big.Int
	Key     [32]byte
}, error) {
	return _KeyManager.Contract.GetKey(&_KeyManager.CallOpts, _key)
}

// GetKey is a free data retrieval call binding the contract method 0x12aaac70.
//
// Solidity: function getKey(bytes32 _key) constant returns(uint256 purpose, uint256 keyType, bytes32 key)
func (_KeyManager *KeyManagerCallerSession) GetKey(_key [32]byte) (struct {
	Purpose *big.Int
	KeyType *big.Int
	Key     [32]byte
}, error) {
	return _KeyManager.Contract.GetKey(&_KeyManager.CallOpts, _key)
}

// GetKeysByPurpose is a free data retrieval call binding the contract method 0x9010f726.
//
// Solidity: function getKeysByPurpose(uint256 _purpose) constant returns(bytes32[] _keys)
func (_KeyManager *KeyManagerCaller) GetKeysByPurpose(opts *bind.CallOpts, _purpose *big.Int) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _KeyManager.contract.Call(opts, out, "getKeysByPurpose", _purpose)
	return *ret0, err
}

// GetKeysByPurpose is a free data retrieval call binding the contract method 0x9010f726.
//
// Solidity: function getKeysByPurpose(uint256 _purpose) constant returns(bytes32[] _keys)
func (_KeyManager *KeyManagerSession) GetKeysByPurpose(_purpose *big.Int) ([][32]byte, error) {
	return _KeyManager.Contract.GetKeysByPurpose(&_KeyManager.CallOpts, _purpose)
}

// GetKeysByPurpose is a free data retrieval call binding the contract method 0x9010f726.
//
// Solidity: function getKeysByPurpose(uint256 _purpose) constant returns(bytes32[] _keys)
func (_KeyManager *KeyManagerCallerSession) GetKeysByPurpose(_purpose *big.Int) ([][32]byte, error) {
	return _KeyManager.Contract.GetKeysByPurpose(&_KeyManager.CallOpts, _purpose)
}

// KeyHasPurpose is a free data retrieval call binding the contract method 0xd202158d.
//
// Solidity: function keyHasPurpose(bytes32 _key, uint256 _purpose) constant returns(bool isExistent)
func (_KeyManager *KeyManagerCaller) KeyHasPurpose(opts *bind.CallOpts, _key [32]byte, _purpose *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KeyManager.contract.Call(opts, out, "keyHasPurpose", _key, _purpose)
	return *ret0, err
}

// KeyHasPurpose is a free data retrieval call binding the contract method 0xd202158d.
//
// Solidity: function keyHasPurpose(bytes32 _key, uint256 _purpose) constant returns(bool isExistent)
func (_KeyManager *KeyManagerSession) KeyHasPurpose(_key [32]byte, _purpose *big.Int) (bool, error) {
	return _KeyManager.Contract.KeyHasPurpose(&_KeyManager.CallOpts, _key, _purpose)
}

// KeyHasPurpose is a free data retrieval call binding the contract method 0xd202158d.
//
// Solidity: function keyHasPurpose(bytes32 _key, uint256 _purpose) constant returns(bool isExistent)
func (_KeyManager *KeyManagerCallerSession) KeyHasPurpose(_key [32]byte, _purpose *big.Int) (bool, error) {
	return _KeyManager.Contract.KeyHasPurpose(&_KeyManager.CallOpts, _key, _purpose)
}

// AddClaim is a paid mutator transaction binding the contract method 0x0607f937.
//
// Solidity: function addClaim(uint256 _claimType, address _issuer, uint256 _signatureType, bytes _signature, bytes _claim, string _uri) returns(bytes32 claimId)
func (_KeyManager *KeyManagerTransactor) AddClaim(opts *bind.TransactOpts, _claimType *big.Int, _issuer common.Address, _signatureType *big.Int, _signature []byte, _claim []byte, _uri string) (*types.Transaction, error) {
	return _KeyManager.contract.Transact(opts, "addClaim", _claimType, _issuer, _signatureType, _signature, _claim, _uri)
}

// AddClaim is a paid mutator transaction binding the contract method 0x0607f937.
//
// Solidity: function addClaim(uint256 _claimType, address _issuer, uint256 _signatureType, bytes _signature, bytes _claim, string _uri) returns(bytes32 claimId)
func (_KeyManager *KeyManagerSession) AddClaim(_claimType *big.Int, _issuer common.Address, _signatureType *big.Int, _signature []byte, _claim []byte, _uri string) (*types.Transaction, error) {
	return _KeyManager.Contract.AddClaim(&_KeyManager.TransactOpts, _claimType, _issuer, _signatureType, _signature, _claim, _uri)
}

// AddClaim is a paid mutator transaction binding the contract method 0x0607f937.
//
// Solidity: function addClaim(uint256 _claimType, address _issuer, uint256 _signatureType, bytes _signature, bytes _claim, string _uri) returns(bytes32 claimId)
func (_KeyManager *KeyManagerTransactorSession) AddClaim(_claimType *big.Int, _issuer common.Address, _signatureType *big.Int, _signature []byte, _claim []byte, _uri string) (*types.Transaction, error) {
	return _KeyManager.Contract.AddClaim(&_KeyManager.TransactOpts, _claimType, _issuer, _signatureType, _signature, _claim, _uri)
}

// AddKey is a paid mutator transaction binding the contract method 0x1d381240.
//
// Solidity: function addKey(bytes32 _key, uint256 _purpose, uint256 _keyType) returns(bool success)
func (_KeyManager *KeyManagerTransactor) AddKey(opts *bind.TransactOpts, _key [32]byte, _purpose *big.Int, _keyType *big.Int) (*types.Transaction, error) {
	return _KeyManager.contract.Transact(opts, "addKey", _key, _purpose, _keyType)
}

// AddKey is a paid mutator transaction binding the contract method 0x1d381240.
//
// Solidity: function addKey(bytes32 _key, uint256 _purpose, uint256 _keyType) returns(bool success)
func (_KeyManager *KeyManagerSession) AddKey(_key [32]byte, _purpose *big.Int, _keyType *big.Int) (*types.Transaction, error) {
	return _KeyManager.Contract.AddKey(&_KeyManager.TransactOpts, _key, _purpose, _keyType)
}

// AddKey is a paid mutator transaction binding the contract method 0x1d381240.
//
// Solidity: function addKey(bytes32 _key, uint256 _purpose, uint256 _keyType) returns(bool success)
func (_KeyManager *KeyManagerTransactorSession) AddKey(_key [32]byte, _purpose *big.Int, _keyType *big.Int) (*types.Transaction, error) {
	return _KeyManager.Contract.AddKey(&_KeyManager.TransactOpts, _key, _purpose, _keyType)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x4eee424a.
//
// Solidity: function removeClaim(bytes32 _claimId) returns(bool success)
func (_KeyManager *KeyManagerTransactor) RemoveClaim(opts *bind.TransactOpts, _claimId [32]byte) (*types.Transaction, error) {
	return _KeyManager.contract.Transact(opts, "removeClaim", _claimId)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x4eee424a.
//
// Solidity: function removeClaim(bytes32 _claimId) returns(bool success)
func (_KeyManager *KeyManagerSession) RemoveClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _KeyManager.Contract.RemoveClaim(&_KeyManager.TransactOpts, _claimId)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x4eee424a.
//
// Solidity: function removeClaim(bytes32 _claimId) returns(bool success)
func (_KeyManager *KeyManagerTransactorSession) RemoveClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _KeyManager.Contract.RemoveClaim(&_KeyManager.TransactOpts, _claimId)
}

// RemoveKey is a paid mutator transaction binding the contract method 0x53d413c5.
//
// Solidity: function removeKey(bytes32 _key, uint256 _purpose) returns(bool success)
func (_KeyManager *KeyManagerTransactor) RemoveKey(opts *bind.TransactOpts, _key [32]byte, _purpose *big.Int) (*types.Transaction, error) {
	return _KeyManager.contract.Transact(opts, "removeKey", _key, _purpose)
}

// RemoveKey is a paid mutator transaction binding the contract method 0x53d413c5.
//
// Solidity: function removeKey(bytes32 _key, uint256 _purpose) returns(bool success)
func (_KeyManager *KeyManagerSession) RemoveKey(_key [32]byte, _purpose *big.Int) (*types.Transaction, error) {
	return _KeyManager.Contract.RemoveKey(&_KeyManager.TransactOpts, _key, _purpose)
}

// RemoveKey is a paid mutator transaction binding the contract method 0x53d413c5.
//
// Solidity: function removeKey(bytes32 _key, uint256 _purpose) returns(bool success)
func (_KeyManager *KeyManagerTransactorSession) RemoveKey(_key [32]byte, _purpose *big.Int) (*types.Transaction, error) {
	return _KeyManager.Contract.RemoveKey(&_KeyManager.TransactOpts, _key, _purpose)
}

// KeyManagerApprovedIterator is returned from FilterApproved and is used to iterate over the raw logs and unpacked data for Approved events raised by the KeyManager contract.
type KeyManagerApprovedIterator struct {
	Event *KeyManagerApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KeyManagerApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyManagerApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KeyManagerApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KeyManagerApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyManagerApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyManagerApproved represents a Approved event raised by the KeyManager contract.
type KeyManagerApproved struct {
	ExecutionID *big.Int
	Approved    bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterApproved is a free log retrieval operation binding the contract event 0xb3932da477fe5d6c8ff2eafef050c0f3a1af18fc07121001482600f36f3715d8.
//
// Solidity: event Approved(uint256 indexed executionID, bool approved)
func (_KeyManager *KeyManagerFilterer) FilterApproved(opts *bind.FilterOpts, executionID []*big.Int) (*KeyManagerApprovedIterator, error) {

	var executionIDRule []interface{}
	for _, executionIDItem := range executionID {
		executionIDRule = append(executionIDRule, executionIDItem)
	}

	logs, sub, err := _KeyManager.contract.FilterLogs(opts, "Approved", executionIDRule)
	if err != nil {
		return nil, err
	}
	return &KeyManagerApprovedIterator{contract: _KeyManager.contract, event: "Approved", logs: logs, sub: sub}, nil
}

// WatchApproved is a free log subscription operation binding the contract event 0xb3932da477fe5d6c8ff2eafef050c0f3a1af18fc07121001482600f36f3715d8.
//
// Solidity: event Approved(uint256 indexed executionID, bool approved)
func (_KeyManager *KeyManagerFilterer) WatchApproved(opts *bind.WatchOpts, sink chan<- *KeyManagerApproved, executionID []*big.Int) (event.Subscription, error) {

	var executionIDRule []interface{}
	for _, executionIDItem := range executionID {
		executionIDRule = append(executionIDRule, executionIDItem)
	}

	logs, sub, err := _KeyManager.contract.WatchLogs(opts, "Approved", executionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyManagerApproved)
				if err := _KeyManager.contract.UnpackLog(event, "Approved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// KeyManagerClaimAddedIterator is returned from FilterClaimAdded and is used to iterate over the raw logs and unpacked data for ClaimAdded events raised by the KeyManager contract.
type KeyManagerClaimAddedIterator struct {
	Event *KeyManagerClaimAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KeyManagerClaimAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyManagerClaimAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KeyManagerClaimAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KeyManagerClaimAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyManagerClaimAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyManagerClaimAdded represents a ClaimAdded event raised by the KeyManager contract.
type KeyManagerClaimAdded struct {
	ClaimId   [32]byte
	ClaimType *big.Int
	Scheme    *big.Int
	Issuer    common.Address
	Signature []byte
	Data      []byte
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimAdded is a free log retrieval operation binding the contract event 0x46149b18aa084502c3f12bc75e19eda8bda8d102b82cce8474677a6d0d5f43c5.
//
// Solidity: event ClaimAdded(bytes32 indexed claimId, uint256 indexed claimType, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_KeyManager *KeyManagerFilterer) FilterClaimAdded(opts *bind.FilterOpts, claimId [][32]byte, claimType []*big.Int, issuer []common.Address) (*KeyManagerClaimAddedIterator, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var claimTypeRule []interface{}
	for _, claimTypeItem := range claimType {
		claimTypeRule = append(claimTypeRule, claimTypeItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _KeyManager.contract.FilterLogs(opts, "ClaimAdded", claimIdRule, claimTypeRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &KeyManagerClaimAddedIterator{contract: _KeyManager.contract, event: "ClaimAdded", logs: logs, sub: sub}, nil
}

// WatchClaimAdded is a free log subscription operation binding the contract event 0x46149b18aa084502c3f12bc75e19eda8bda8d102b82cce8474677a6d0d5f43c5.
//
// Solidity: event ClaimAdded(bytes32 indexed claimId, uint256 indexed claimType, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_KeyManager *KeyManagerFilterer) WatchClaimAdded(opts *bind.WatchOpts, sink chan<- *KeyManagerClaimAdded, claimId [][32]byte, claimType []*big.Int, issuer []common.Address) (event.Subscription, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var claimTypeRule []interface{}
	for _, claimTypeItem := range claimType {
		claimTypeRule = append(claimTypeRule, claimTypeItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _KeyManager.contract.WatchLogs(opts, "ClaimAdded", claimIdRule, claimTypeRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyManagerClaimAdded)
				if err := _KeyManager.contract.UnpackLog(event, "ClaimAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// KeyManagerClaimChangedIterator is returned from FilterClaimChanged and is used to iterate over the raw logs and unpacked data for ClaimChanged events raised by the KeyManager contract.
type KeyManagerClaimChangedIterator struct {
	Event *KeyManagerClaimChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KeyManagerClaimChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyManagerClaimChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KeyManagerClaimChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KeyManagerClaimChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyManagerClaimChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyManagerClaimChanged represents a ClaimChanged event raised by the KeyManager contract.
type KeyManagerClaimChanged struct {
	ClaimId   [32]byte
	ClaimType *big.Int
	Scheme    *big.Int
	Issuer    common.Address
	Signature []byte
	Data      []byte
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimChanged is a free log retrieval operation binding the contract event 0x3bab293fc00db832d7619a9299914251b8747c036867ec056cbd506f60135b13.
//
// Solidity: event ClaimChanged(bytes32 indexed claimId, uint256 indexed claimType, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_KeyManager *KeyManagerFilterer) FilterClaimChanged(opts *bind.FilterOpts, claimId [][32]byte, claimType []*big.Int, issuer []common.Address) (*KeyManagerClaimChangedIterator, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var claimTypeRule []interface{}
	for _, claimTypeItem := range claimType {
		claimTypeRule = append(claimTypeRule, claimTypeItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _KeyManager.contract.FilterLogs(opts, "ClaimChanged", claimIdRule, claimTypeRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &KeyManagerClaimChangedIterator{contract: _KeyManager.contract, event: "ClaimChanged", logs: logs, sub: sub}, nil
}

// WatchClaimChanged is a free log subscription operation binding the contract event 0x3bab293fc00db832d7619a9299914251b8747c036867ec056cbd506f60135b13.
//
// Solidity: event ClaimChanged(bytes32 indexed claimId, uint256 indexed claimType, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_KeyManager *KeyManagerFilterer) WatchClaimChanged(opts *bind.WatchOpts, sink chan<- *KeyManagerClaimChanged, claimId [][32]byte, claimType []*big.Int, issuer []common.Address) (event.Subscription, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var claimTypeRule []interface{}
	for _, claimTypeItem := range claimType {
		claimTypeRule = append(claimTypeRule, claimTypeItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _KeyManager.contract.WatchLogs(opts, "ClaimChanged", claimIdRule, claimTypeRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyManagerClaimChanged)
				if err := _KeyManager.contract.UnpackLog(event, "ClaimChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// KeyManagerClaimRemovedIterator is returned from FilterClaimRemoved and is used to iterate over the raw logs and unpacked data for ClaimRemoved events raised by the KeyManager contract.
type KeyManagerClaimRemovedIterator struct {
	Event *KeyManagerClaimRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KeyManagerClaimRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyManagerClaimRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KeyManagerClaimRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KeyManagerClaimRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyManagerClaimRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyManagerClaimRemoved represents a ClaimRemoved event raised by the KeyManager contract.
type KeyManagerClaimRemoved struct {
	ClaimId   [32]byte
	ClaimType *big.Int
	Scheme    *big.Int
	Issuer    common.Address
	Signature []byte
	Data      []byte
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimRemoved is a free log retrieval operation binding the contract event 0x3cf57863a89432c61c4a27073c6ee39e8a764bff5a05aebfbcdcdc80b2e6130a.
//
// Solidity: event ClaimRemoved(bytes32 indexed claimId, uint256 indexed claimType, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_KeyManager *KeyManagerFilterer) FilterClaimRemoved(opts *bind.FilterOpts, claimId [][32]byte, claimType []*big.Int, issuer []common.Address) (*KeyManagerClaimRemovedIterator, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var claimTypeRule []interface{}
	for _, claimTypeItem := range claimType {
		claimTypeRule = append(claimTypeRule, claimTypeItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _KeyManager.contract.FilterLogs(opts, "ClaimRemoved", claimIdRule, claimTypeRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &KeyManagerClaimRemovedIterator{contract: _KeyManager.contract, event: "ClaimRemoved", logs: logs, sub: sub}, nil
}

// WatchClaimRemoved is a free log subscription operation binding the contract event 0x3cf57863a89432c61c4a27073c6ee39e8a764bff5a05aebfbcdcdc80b2e6130a.
//
// Solidity: event ClaimRemoved(bytes32 indexed claimId, uint256 indexed claimType, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_KeyManager *KeyManagerFilterer) WatchClaimRemoved(opts *bind.WatchOpts, sink chan<- *KeyManagerClaimRemoved, claimId [][32]byte, claimType []*big.Int, issuer []common.Address) (event.Subscription, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var claimTypeRule []interface{}
	for _, claimTypeItem := range claimType {
		claimTypeRule = append(claimTypeRule, claimTypeItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _KeyManager.contract.WatchLogs(opts, "ClaimRemoved", claimIdRule, claimTypeRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyManagerClaimRemoved)
				if err := _KeyManager.contract.UnpackLog(event, "ClaimRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// KeyManagerExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the KeyManager contract.
type KeyManagerExecutedIterator struct {
	Event *KeyManagerExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KeyManagerExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyManagerExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KeyManagerExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KeyManagerExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyManagerExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyManagerExecuted represents a Executed event raised by the KeyManager contract.
type KeyManagerExecuted struct {
	ExecutionID *big.Int
	To          common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x1f920dbda597d7bf95035464170fa58d0a4b57f13a1c315ace6793b9f63688b8.
//
// Solidity: event Executed(uint256 indexed executionID, address indexed to, uint256 indexed value, bytes data)
func (_KeyManager *KeyManagerFilterer) FilterExecuted(opts *bind.FilterOpts, executionID []*big.Int, to []common.Address, value []*big.Int) (*KeyManagerExecutedIterator, error) {

	var executionIDRule []interface{}
	for _, executionIDItem := range executionID {
		executionIDRule = append(executionIDRule, executionIDItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _KeyManager.contract.FilterLogs(opts, "Executed", executionIDRule, toRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &KeyManagerExecutedIterator{contract: _KeyManager.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x1f920dbda597d7bf95035464170fa58d0a4b57f13a1c315ace6793b9f63688b8.
//
// Solidity: event Executed(uint256 indexed executionID, address indexed to, uint256 indexed value, bytes data)
func (_KeyManager *KeyManagerFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *KeyManagerExecuted, executionID []*big.Int, to []common.Address, value []*big.Int) (event.Subscription, error) {

	var executionIDRule []interface{}
	for _, executionIDItem := range executionID {
		executionIDRule = append(executionIDRule, executionIDItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _KeyManager.contract.WatchLogs(opts, "Executed", executionIDRule, toRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyManagerExecuted)
				if err := _KeyManager.contract.UnpackLog(event, "Executed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// KeyManagerExecutionRequestedIterator is returned from FilterExecutionRequested and is used to iterate over the raw logs and unpacked data for ExecutionRequested events raised by the KeyManager contract.
type KeyManagerExecutionRequestedIterator struct {
	Event *KeyManagerExecutionRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KeyManagerExecutionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyManagerExecutionRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KeyManagerExecutionRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KeyManagerExecutionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyManagerExecutionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyManagerExecutionRequested represents a ExecutionRequested event raised by the KeyManager contract.
type KeyManagerExecutionRequested struct {
	ExecutionID *big.Int
	To          common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutionRequested is a free log retrieval operation binding the contract event 0x8afcfabcb00e47a53a8fc3e9f23ff47ee1926194bb1350dd007c50b412a6cee8.
//
// Solidity: event ExecutionRequested(uint256 indexed executionID, address indexed to, uint256 indexed value, bytes data)
func (_KeyManager *KeyManagerFilterer) FilterExecutionRequested(opts *bind.FilterOpts, executionID []*big.Int, to []common.Address, value []*big.Int) (*KeyManagerExecutionRequestedIterator, error) {

	var executionIDRule []interface{}
	for _, executionIDItem := range executionID {
		executionIDRule = append(executionIDRule, executionIDItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _KeyManager.contract.FilterLogs(opts, "ExecutionRequested", executionIDRule, toRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &KeyManagerExecutionRequestedIterator{contract: _KeyManager.contract, event: "ExecutionRequested", logs: logs, sub: sub}, nil
}

// WatchExecutionRequested is a free log subscription operation binding the contract event 0x8afcfabcb00e47a53a8fc3e9f23ff47ee1926194bb1350dd007c50b412a6cee8.
//
// Solidity: event ExecutionRequested(uint256 indexed executionID, address indexed to, uint256 indexed value, bytes data)
func (_KeyManager *KeyManagerFilterer) WatchExecutionRequested(opts *bind.WatchOpts, sink chan<- *KeyManagerExecutionRequested, executionID []*big.Int, to []common.Address, value []*big.Int) (event.Subscription, error) {

	var executionIDRule []interface{}
	for _, executionIDItem := range executionID {
		executionIDRule = append(executionIDRule, executionIDItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _KeyManager.contract.WatchLogs(opts, "ExecutionRequested", executionIDRule, toRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyManagerExecutionRequested)
				if err := _KeyManager.contract.UnpackLog(event, "ExecutionRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// KeyManagerKeyAddedIterator is returned from FilterKeyAdded and is used to iterate over the raw logs and unpacked data for KeyAdded events raised by the KeyManager contract.
type KeyManagerKeyAddedIterator struct {
	Event *KeyManagerKeyAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KeyManagerKeyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyManagerKeyAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KeyManagerKeyAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KeyManagerKeyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyManagerKeyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyManagerKeyAdded represents a KeyAdded event raised by the KeyManager contract.
type KeyManagerKeyAdded struct {
	Key     [32]byte
	Purpose *big.Int
	KeyType *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKeyAdded is a free log retrieval operation binding the contract event 0x480000bb1edad8ca1470381cc334b1917fbd51c6531f3a623ea8e0ec7e38a6e9.
//
// Solidity: event KeyAdded(bytes32 indexed key, uint256 indexed purpose, uint256 indexed keyType)
func (_KeyManager *KeyManagerFilterer) FilterKeyAdded(opts *bind.FilterOpts, key [][32]byte, purpose []*big.Int, keyType []*big.Int) (*KeyManagerKeyAddedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var purposeRule []interface{}
	for _, purposeItem := range purpose {
		purposeRule = append(purposeRule, purposeItem)
	}
	var keyTypeRule []interface{}
	for _, keyTypeItem := range keyType {
		keyTypeRule = append(keyTypeRule, keyTypeItem)
	}

	logs, sub, err := _KeyManager.contract.FilterLogs(opts, "KeyAdded", keyRule, purposeRule, keyTypeRule)
	if err != nil {
		return nil, err
	}
	return &KeyManagerKeyAddedIterator{contract: _KeyManager.contract, event: "KeyAdded", logs: logs, sub: sub}, nil
}

// WatchKeyAdded is a free log subscription operation binding the contract event 0x480000bb1edad8ca1470381cc334b1917fbd51c6531f3a623ea8e0ec7e38a6e9.
//
// Solidity: event KeyAdded(bytes32 indexed key, uint256 indexed purpose, uint256 indexed keyType)
func (_KeyManager *KeyManagerFilterer) WatchKeyAdded(opts *bind.WatchOpts, sink chan<- *KeyManagerKeyAdded, key [][32]byte, purpose []*big.Int, keyType []*big.Int) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var purposeRule []interface{}
	for _, purposeItem := range purpose {
		purposeRule = append(purposeRule, purposeItem)
	}
	var keyTypeRule []interface{}
	for _, keyTypeItem := range keyType {
		keyTypeRule = append(keyTypeRule, keyTypeItem)
	}

	logs, sub, err := _KeyManager.contract.WatchLogs(opts, "KeyAdded", keyRule, purposeRule, keyTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyManagerKeyAdded)
				if err := _KeyManager.contract.UnpackLog(event, "KeyAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// KeyManagerKeyRemovedIterator is returned from FilterKeyRemoved and is used to iterate over the raw logs and unpacked data for KeyRemoved events raised by the KeyManager contract.
type KeyManagerKeyRemovedIterator struct {
	Event *KeyManagerKeyRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KeyManagerKeyRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyManagerKeyRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KeyManagerKeyRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KeyManagerKeyRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyManagerKeyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyManagerKeyRemoved represents a KeyRemoved event raised by the KeyManager contract.
type KeyManagerKeyRemoved struct {
	Key     [32]byte
	Purpose *big.Int
	KeyType *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKeyRemoved is a free log retrieval operation binding the contract event 0x585a4aef50f8267a92b32412b331b20f7f8b96f2245b253b9cc50dcc621d3397.
//
// Solidity: event KeyRemoved(bytes32 indexed key, uint256 indexed purpose, uint256 indexed keyType)
func (_KeyManager *KeyManagerFilterer) FilterKeyRemoved(opts *bind.FilterOpts, key [][32]byte, purpose []*big.Int, keyType []*big.Int) (*KeyManagerKeyRemovedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var purposeRule []interface{}
	for _, purposeItem := range purpose {
		purposeRule = append(purposeRule, purposeItem)
	}
	var keyTypeRule []interface{}
	for _, keyTypeItem := range keyType {
		keyTypeRule = append(keyTypeRule, keyTypeItem)
	}

	logs, sub, err := _KeyManager.contract.FilterLogs(opts, "KeyRemoved", keyRule, purposeRule, keyTypeRule)
	if err != nil {
		return nil, err
	}
	return &KeyManagerKeyRemovedIterator{contract: _KeyManager.contract, event: "KeyRemoved", logs: logs, sub: sub}, nil
}

// WatchKeyRemoved is a free log subscription operation binding the contract event 0x585a4aef50f8267a92b32412b331b20f7f8b96f2245b253b9cc50dcc621d3397.
//
// Solidity: event KeyRemoved(bytes32 indexed key, uint256 indexed purpose, uint256 indexed keyType)
func (_KeyManager *KeyManagerFilterer) WatchKeyRemoved(opts *bind.WatchOpts, sink chan<- *KeyManagerKeyRemoved, key [][32]byte, purpose []*big.Int, keyType []*big.Int) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var purposeRule []interface{}
	for _, purposeItem := range purpose {
		purposeRule = append(purposeRule, purposeItem)
	}
	var keyTypeRule []interface{}
	for _, keyTypeItem := range keyType {
		keyTypeRule = append(keyTypeRule, keyTypeItem)
	}

	logs, sub, err := _KeyManager.contract.WatchLogs(opts, "KeyRemoved", keyRule, purposeRule, keyTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyManagerKeyRemoved)
				if err := _KeyManager.contract.UnpackLog(event, "KeyRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// KeyManagerKeysRequiredChangedIterator is returned from FilterKeysRequiredChanged and is used to iterate over the raw logs and unpacked data for KeysRequiredChanged events raised by the KeyManager contract.
type KeyManagerKeysRequiredChangedIterator struct {
	Event *KeyManagerKeysRequiredChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KeyManagerKeysRequiredChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyManagerKeysRequiredChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KeyManagerKeysRequiredChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KeyManagerKeysRequiredChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyManagerKeysRequiredChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyManagerKeysRequiredChanged represents a KeysRequiredChanged event raised by the KeyManager contract.
type KeyManagerKeysRequiredChanged struct {
	Purpose *big.Int
	Number  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKeysRequiredChanged is a free log retrieval operation binding the contract event 0x9f6d363a5a7fef1e6f9d2ac96e0554ca3494d7b59adfe0bf665020ddf9350d1a.
//
// Solidity: event KeysRequiredChanged(uint256 purpose, uint256 number)
func (_KeyManager *KeyManagerFilterer) FilterKeysRequiredChanged(opts *bind.FilterOpts) (*KeyManagerKeysRequiredChangedIterator, error) {

	logs, sub, err := _KeyManager.contract.FilterLogs(opts, "KeysRequiredChanged")
	if err != nil {
		return nil, err
	}
	return &KeyManagerKeysRequiredChangedIterator{contract: _KeyManager.contract, event: "KeysRequiredChanged", logs: logs, sub: sub}, nil
}

// WatchKeysRequiredChanged is a free log subscription operation binding the contract event 0x9f6d363a5a7fef1e6f9d2ac96e0554ca3494d7b59adfe0bf665020ddf9350d1a.
//
// Solidity: event KeysRequiredChanged(uint256 purpose, uint256 number)
func (_KeyManager *KeyManagerFilterer) WatchKeysRequiredChanged(opts *bind.WatchOpts, sink chan<- *KeyManagerKeysRequiredChanged) (event.Subscription, error) {

	logs, sub, err := _KeyManager.contract.WatchLogs(opts, "KeysRequiredChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyManagerKeysRequiredChanged)
				if err := _KeyManager.contract.UnpackLog(event, "KeysRequiredChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
