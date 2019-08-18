pragma solidity >=0.4.21 <0.6.0;

interface ERC734 {
    event KeyAdded(
        bytes32 indexed key,
        uint256 indexed purpose,
        uint256 indexed keyType
    );
    event KeyRemoved(
        bytes32 indexed key,
        uint256 indexed purpose,
        uint256 indexed keyType
    );
    event ExecutionRequested(
        uint256 indexed executionID,
        address indexed to,
        uint256 indexed value,
        bytes data
    );
    event Executed(
        uint256 indexed executionID,
        address indexed to,
        uint256 indexed value,
        bytes data
    );
    event Approved(uint256 indexed executionID, bool approved);
    event KeysRequiredChanged(uint256 purpose, uint256 number);
    event ClaimAdded(bytes32 indexed claimId, uint256 indexed claimType,
        uint256 scheme, address indexed issuer, bytes signature, bytes data,
        string uri
    );
    event ClaimRemoved(bytes32 indexed claimId, uint256 indexed claimType,
        uint256 scheme, address indexed issuer, bytes signature, bytes data,
        string uri
    );
    event ClaimChanged(bytes32 indexed claimId, uint256 indexed claimType,
        uint256 scheme, address indexed issuer, bytes signature, bytes data,
        string uri
    );

    struct Key {
        uint256 purpose; // MANAGEMENT_KEY = 1, EXECUTION_KEY = 2
        uint256 keyType; // ECDSA = 1, RSA = 2
        bytes32 key;
    }

    struct Claim {
        uint256 claimType;
        address issuer; // msg.sender
        uint256 signatureType; // The type of signature
        bytes signature; // this.address + claimType + claim
        bytes claim;
        string uri;
    }

    function keyHasPurpose(bytes32 _key, uint256 _purpose)
        external
        view
        returns (bool exists);

    function getKey(bytes32 _key)
        external
        view
        returns (uint256 purpose, uint256 keyType, bytes32 key);

    function getKeysByPurpose(uint256 _purpose)
        external
        view
        returns (bytes32[] memory keys);

    function addKey(bytes32 _key, uint256 _purpose, uint256 _keyType)
        external
        returns (bool success);

    function removeKey(bytes32 _key, uint256 _purpose)
        external
        returns (bool success);
    // function changeKeysRequired(uint256 purpose, uint256 number) external;
    // function getKeysRequired(uint256 purpose) external view returns (uint256);
    // function execute(address _to, uint256 _value, bytes calldata _data) external returns (uint256 executionID);
    // function approve(uint256 _id, bool _approve) external returns (bool success);
}

contract KeyManager is ERC734 {
    uint256 constant MANAGEMENT_KEY = 1;
    uint256 constant EXECUTION_KEY = 2;
    uint256 constant CLAIM_SIGNER_KEY = 3;
    uint256 constant ECDSA_TYPE = 1;
    uint256 constant RSA_TYPE = 2;
    mapping(bytes32 => Key) keys;
    mapping(uint256 => bytes32[]) keysByPurpose;
    mapping (bytes32 => Claim) claims;
    mapping (uint256 => bytes32[]) claimsByType;
    constructor() public {
        bytes32 _key = keccak256(abi.encodePacked(msg.sender));
        keys[_key] = Key({
            purpose: MANAGEMENT_KEY,
            keyType: ECDSA_TYPE,
            key: _key
        });
        keysByPurpose[MANAGEMENT_KEY].push(_key);
        emit KeyAdded(_key, MANAGEMENT_KEY, ECDSA_TYPE);
    }

    function keyHasPurpose(bytes32 _key, uint256 _purpose)
        external
        view
        returns (bool isExistent)
    {
        if (keys[_key].key == 0) return false;
        if (keys[_key].purpose == _purpose) {
            return true;
        } else {
            return false;
        }
    }

    function getKey(bytes32 _key)
        external
        view
        returns (uint256 purpose, uint256 keyType, bytes32 key)
    {
        return (keys[_key].purpose, keys[_key].keyType, keys[_key].key);
    }

    function getKeysByPurpose(uint256 _purpose)
        external
        view
        returns (bytes32[] memory _keys)
    {
        return keysByPurpose[_purpose];
    }

    function addKey(bytes32 _key, uint256 _purpose, uint256 _keyType)
        external
        returns (bool success)
    {
        require(keys[_key].key != _key, "Key already exists.");
        if (msg.sender != address(this)) {
            require(
                this.keyHasPurpose(keccak256(abi.encodePacked(msg.sender)), 1),
                "Sender dose not have a manager key."
            );
        }
        keys[_key] = Key({key: _key, purpose: _purpose, keyType: _keyType});
        keysByPurpose[_purpose].push(_key);
        emit KeyAdded(_key, _purpose, _keyType);
        return true;
    }

    function removeKey(bytes32 _key, uint256 _purpose)
        external
        returns (bool success)
    {
        require(keys[_key].key == _key, "No such key.");
        if (msg.sender != address(this)) {
            require(
                this.keyHasPurpose(keccak256(abi.encodePacked(msg.sender)), 1),
                "Sender does not have a manager key."
            );
        }
        delete keys[_key];
        bytes32[] storage keysCache = keysByPurpose[_purpose];
        for (uint i = 0; i < keysCache.length; i++) {
            if (keysCache[i] == _key) {
                keysCache[i] = keysCache[keysCache.length - 1];
                delete keysCache[keysCache.length - 1];
                keysCache.length--;
                break;
            }
        }
        emit KeyRemoved(_key, _purpose, keys[_key].keyType);
        return true;
    }

    function getClaim(bytes32 _claimId)
        external
        view
        returns(uint256 claimType, address issuer, uint256 signatureType,
            bytes memory signature, bytes memory claim, string memory uri)
    {
        Claim memory _claim = claims[_claimId];
        return (_claim.claimType, _claim.issuer, _claim.signatureType, _claim.signature, _claim.claim, _claim.uri);
    }

    function getClaimsIdByType(uint256 _claimType) external view returns(bytes32[] memory) {
        return claimsByType[_claimType];
    }

    function addClaim(
        uint256 _claimType,
        address _issuer,
        uint256 _signatureType,
        bytes memory _signature,
        bytes memory _claim,
        string memory _uri
    )
        public
        returns (bytes32 claimId)
    {
        claimId = keccak256(abi.encodePacked(_issuer, _claimType));
        claims[claimId] = Claim({
            claimType: _claimType,
            issuer: _issuer,
            signatureType: _signatureType,
            signature: _signature,
            claim: _claim,
            uri: _uri
        });
        claimsByType[_claimType].push(claimId);
        emit ClaimAdded(claimId, _claimType, _signatureType, _issuer, _signature, _claim, _uri);
    }

    function removeClaim(bytes32 _claimId) external returns (bool success) {
        bytes32[] storage keysCache = claimsByType[claims[_claimId].claimType];
        for (uint i = 0; i < keysCache.length; i++) {
            if (keysCache[i] == _claimId) {
                keysCache[i] = keysCache[keysCache.length - 1];
                delete keysCache[keysCache.length - 1];
                keysCache.length--;
                break;
            }
        }
        delete claims[_claimId];
        emit ClaimRemoved(
            _claimId, claims[_claimId].claimType, claims[_claimId].signatureType,
            claims[_claimId].issuer, claims[_claimId].signature,
            claims[_claimId].claim, claims[_claimId].uri
        );
        return true;
    }
}
