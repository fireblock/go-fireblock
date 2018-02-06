// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"keyuid\",\"type\":\"bytes32\"}],\"name\":\"getKey\",\"outputs\":[{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"revoked\",\"type\":\"bool\"},{\"name\":\"begin\",\"type\":\"int64\"},{\"name\":\"end\",\"type\":\"int64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cardUID\",\"type\":\"bytes32\"}],\"name\":\"getCardB32\",\"outputs\":[{\"name\":\"useruid\",\"type\":\"bytes32\"},{\"name\":\"upgrade\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"useruid\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getUserCardAt\",\"outputs\":[{\"name\":\"cardId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getProofX\",\"outputs\":[{\"name\":\"x0\",\"type\":\"bytes32\"},{\"name\":\"x1\",\"type\":\"bytes32\"},{\"name\":\"x2\",\"type\":\"bytes32\"},{\"name\":\"x3\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"useruid\",\"type\":\"bytes32\"}],\"name\":\"getUserCardLen\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"keyuid\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"setKeyX\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"revokeProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"keyuid\",\"type\":\"bytes32\"}],\"name\":\"getKeyX\",\"outputs\":[{\"name\":\"x0\",\"type\":\"bytes32\"},{\"name\":\"x1\",\"type\":\"bytes32\"},{\"name\":\"x2\",\"type\":\"bytes32\"},{\"name\":\"x3\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"cardId\",\"type\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes32\"},{\"name\":\"metadata\",\"type\":\"bytes32\"}],\"name\":\"setProof\",\"outputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"v\",\"type\":\"string\"}],\"name\":\"keccak256StringToB32\",\"outputs\":[{\"name\":\"key\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cardId\",\"type\":\"bytes32\"},{\"name\":\"newCardId\",\"type\":\"bytes32\"}],\"name\":\"upgradeCard\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"keyuid\",\"type\":\"bytes32\"}],\"name\":\"isKeyRevoked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"keyuid\",\"type\":\"bytes32\"}],\"name\":\"revokeKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"keyuid\",\"type\":\"bytes32\"},{\"name\":\"date\",\"type\":\"int64\"}],\"name\":\"isKeyValidAt\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"setProofX\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cardId\",\"type\":\"bytes32\"},{\"name\":\"useruid\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setCard\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cardUID\",\"type\":\"bytes32\"}],\"name\":\"getCard\",\"outputs\":[{\"name\":\"card\",\"type\":\"string\"},{\"name\":\"useruid\",\"type\":\"bytes32\"},{\"name\":\"upgrade\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"keyuid\",\"type\":\"bytes32\"}],\"name\":\"hasKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cardId\",\"type\":\"bytes32\"}],\"name\":\"resetUpgradeCard\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getCurrentCard\",\"outputs\":[{\"name\":\"cardId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"contract_\",\"type\":\"address\"}],\"name\":\"setContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cardId\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"setCardX\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"keyuid\",\"type\":\"bytes32\"},{\"name\":\"keySID\",\"type\":\"bytes32\"},{\"name\":\"begin\",\"type\":\"int64\"},{\"name\":\"end\",\"type\":\"int64\"}],\"name\":\"setKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"cardId\",\"type\":\"bytes32\"}],\"name\":\"getProof\",\"outputs\":[{\"name\":\"puid\",\"type\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"string\"},{\"name\":\"metadata\",\"type\":\"string\"},{\"name\":\"revoked\",\"type\":\"bool\"},{\"name\":\"blockWrite\",\"type\":\"int64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"hasCurrentCard\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"keyuid\",\"type\":\"bytes32\"}],\"name\":\"closeKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"},{\"name\":\"k2\",\"type\":\"bytes32\"},{\"name\":\"k3\",\"type\":\"bytes32\"}],\"name\":\"getString3\",\"outputs\":[{\"name\":\"s\",\"type\":\"string\"},{\"name\":\"s2\",\"type\":\"string\"},{\"name\":\"s3\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"cardId\",\"type\":\"bytes32\"}],\"name\":\"getProofB32\",\"outputs\":[{\"name\":\"puid\",\"type\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes32\"},{\"name\":\"metadata\",\"type\":\"bytes32\"},{\"name\":\"revoked\",\"type\":\"bool\"},{\"name\":\"blockWrite\",\"type\":\"int64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"addressToB32\",\"outputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cardId\",\"type\":\"bytes32\"}],\"name\":\"hasCard\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"isProofRevoked\",\"outputs\":[{\"name\":\"revoked\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"str\",\"type\":\"string\"}],\"name\":\"setString1\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getString1\",\"outputs\":[{\"name\":\"s\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"hasProof\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"keyuid\",\"type\":\"bytes32\"}],\"name\":\"isKeyValid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cardUID\",\"type\":\"bytes32\"}],\"name\":\"getCardX\",\"outputs\":[{\"name\":\"x0\",\"type\":\"bytes32\"},{\"name\":\"x1\",\"type\":\"bytes32\"},{\"name\":\"x2\",\"type\":\"bytes32\"},{\"name\":\"x3\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"k\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"cardId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"signature\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"metadata\",\"type\":\"bytes32\"}],\"name\":\"ProofRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"cardId\",\"type\":\"bytes32\"}],\"name\":\"ProofRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cardId\",\"type\":\"bytes32\"}],\"name\":\"CardRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"cardId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"newCardId\",\"type\":\"bytes32\"}],\"name\":\"CardUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"keyuid\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"keySID\",\"type\":\"bytes32\"}],\"name\":\"KeyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"keyuid\",\"type\":\"bytes32\"}],\"name\":\"KeyRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"keyuid\",\"type\":\"bytes32\"}],\"name\":\"KeyClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// AddressToB32 is a free data retrieval call binding the contract method 0xa8ad14b0.
//
// Solidity: function addressToB32(a address) constant returns(k bytes32)
func (_Store *StoreCaller) AddressToB32(opts *bind.CallOpts, a common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "addressToB32", a)
	return *ret0, err
}

// AddressToB32 is a free data retrieval call binding the contract method 0xa8ad14b0.
//
// Solidity: function addressToB32(a address) constant returns(k bytes32)
func (_Store *StoreSession) AddressToB32(a common.Address) ([32]byte, error) {
	return _Store.Contract.AddressToB32(&_Store.CallOpts, a)
}

// AddressToB32 is a free data retrieval call binding the contract method 0xa8ad14b0.
//
// Solidity: function addressToB32(a address) constant returns(k bytes32)
func (_Store *StoreCallerSession) AddressToB32(a common.Address) ([32]byte, error) {
	return _Store.Contract.AddressToB32(&_Store.CallOpts, a)
}

// GetCard is a free data retrieval call binding the contract method 0x62bd42db.
//
// Solidity: function getCard(cardUID bytes32) constant returns(card string, useruid bytes32, upgrade bytes32, key bytes32)
func (_Store *StoreCaller) GetCard(opts *bind.CallOpts, cardUID [32]byte) (struct {
	Card    string
	Useruid [32]byte
	Upgrade [32]byte
	Key     [32]byte
}, error) {
	ret := new(struct {
		Card    string
		Useruid [32]byte
		Upgrade [32]byte
		Key     [32]byte
	})
	out := ret
	err := _Store.contract.Call(opts, out, "getCard", cardUID)
	return *ret, err
}

// GetCard is a free data retrieval call binding the contract method 0x62bd42db.
//
// Solidity: function getCard(cardUID bytes32) constant returns(card string, useruid bytes32, upgrade bytes32, key bytes32)
func (_Store *StoreSession) GetCard(cardUID [32]byte) (struct {
	Card    string
	Useruid [32]byte
	Upgrade [32]byte
	Key     [32]byte
}, error) {
	return _Store.Contract.GetCard(&_Store.CallOpts, cardUID)
}

// GetCard is a free data retrieval call binding the contract method 0x62bd42db.
//
// Solidity: function getCard(cardUID bytes32) constant returns(card string, useruid bytes32, upgrade bytes32, key bytes32)
func (_Store *StoreCallerSession) GetCard(cardUID [32]byte) (struct {
	Card    string
	Useruid [32]byte
	Upgrade [32]byte
	Key     [32]byte
}, error) {
	return _Store.Contract.GetCard(&_Store.CallOpts, cardUID)
}

// GetCardB32 is a free data retrieval call binding the contract method 0x16a7fad8.
//
// Solidity: function getCardB32(cardUID bytes32) constant returns(useruid bytes32, upgrade bytes32, key bytes32)
func (_Store *StoreCaller) GetCardB32(opts *bind.CallOpts, cardUID [32]byte) (struct {
	Useruid [32]byte
	Upgrade [32]byte
	Key     [32]byte
}, error) {
	ret := new(struct {
		Useruid [32]byte
		Upgrade [32]byte
		Key     [32]byte
	})
	out := ret
	err := _Store.contract.Call(opts, out, "getCardB32", cardUID)
	return *ret, err
}

// GetCardB32 is a free data retrieval call binding the contract method 0x16a7fad8.
//
// Solidity: function getCardB32(cardUID bytes32) constant returns(useruid bytes32, upgrade bytes32, key bytes32)
func (_Store *StoreSession) GetCardB32(cardUID [32]byte) (struct {
	Useruid [32]byte
	Upgrade [32]byte
	Key     [32]byte
}, error) {
	return _Store.Contract.GetCardB32(&_Store.CallOpts, cardUID)
}

// GetCardB32 is a free data retrieval call binding the contract method 0x16a7fad8.
//
// Solidity: function getCardB32(cardUID bytes32) constant returns(useruid bytes32, upgrade bytes32, key bytes32)
func (_Store *StoreCallerSession) GetCardB32(cardUID [32]byte) (struct {
	Useruid [32]byte
	Upgrade [32]byte
	Key     [32]byte
}, error) {
	return _Store.Contract.GetCardB32(&_Store.CallOpts, cardUID)
}

// GetCardX is a free data retrieval call binding the contract method 0xf372cd08.
//
// Solidity: function getCardX(cardUID bytes32) constant returns(x0 bytes32, x1 bytes32, x2 bytes32, x3 bytes32)
func (_Store *StoreCaller) GetCardX(opts *bind.CallOpts, cardUID [32]byte) (struct {
	X0 [32]byte
	X1 [32]byte
	X2 [32]byte
	X3 [32]byte
}, error) {
	ret := new(struct {
		X0 [32]byte
		X1 [32]byte
		X2 [32]byte
		X3 [32]byte
	})
	out := ret
	err := _Store.contract.Call(opts, out, "getCardX", cardUID)
	return *ret, err
}

// GetCardX is a free data retrieval call binding the contract method 0xf372cd08.
//
// Solidity: function getCardX(cardUID bytes32) constant returns(x0 bytes32, x1 bytes32, x2 bytes32, x3 bytes32)
func (_Store *StoreSession) GetCardX(cardUID [32]byte) (struct {
	X0 [32]byte
	X1 [32]byte
	X2 [32]byte
	X3 [32]byte
}, error) {
	return _Store.Contract.GetCardX(&_Store.CallOpts, cardUID)
}

// GetCardX is a free data retrieval call binding the contract method 0xf372cd08.
//
// Solidity: function getCardX(cardUID bytes32) constant returns(x0 bytes32, x1 bytes32, x2 bytes32, x3 bytes32)
func (_Store *StoreCallerSession) GetCardX(cardUID [32]byte) (struct {
	X0 [32]byte
	X1 [32]byte
	X2 [32]byte
	X3 [32]byte
}, error) {
	return _Store.Contract.GetCardX(&_Store.CallOpts, cardUID)
}

// GetCurrentCard is a free data retrieval call binding the contract method 0x704ad47c.
//
// Solidity: function getCurrentCard(user address) constant returns(cardId bytes32)
func (_Store *StoreCaller) GetCurrentCard(opts *bind.CallOpts, user common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "getCurrentCard", user)
	return *ret0, err
}

// GetCurrentCard is a free data retrieval call binding the contract method 0x704ad47c.
//
// Solidity: function getCurrentCard(user address) constant returns(cardId bytes32)
func (_Store *StoreSession) GetCurrentCard(user common.Address) ([32]byte, error) {
	return _Store.Contract.GetCurrentCard(&_Store.CallOpts, user)
}

// GetCurrentCard is a free data retrieval call binding the contract method 0x704ad47c.
//
// Solidity: function getCurrentCard(user address) constant returns(cardId bytes32)
func (_Store *StoreCallerSession) GetCurrentCard(user common.Address) ([32]byte, error) {
	return _Store.Contract.GetCurrentCard(&_Store.CallOpts, user)
}

// GetKey is a free data retrieval call binding the contract method 0x12aaac70.
//
// Solidity: function getKey(keyuid bytes32) constant returns(key string, revoked bool, begin int64, end int64)
func (_Store *StoreCaller) GetKey(opts *bind.CallOpts, keyuid [32]byte) (struct {
	Key     string
	Revoked bool
	Begin   int64
	End     int64
}, error) {
	ret := new(struct {
		Key     string
		Revoked bool
		Begin   int64
		End     int64
	})
	out := ret
	err := _Store.contract.Call(opts, out, "getKey", keyuid)
	return *ret, err
}

// GetKey is a free data retrieval call binding the contract method 0x12aaac70.
//
// Solidity: function getKey(keyuid bytes32) constant returns(key string, revoked bool, begin int64, end int64)
func (_Store *StoreSession) GetKey(keyuid [32]byte) (struct {
	Key     string
	Revoked bool
	Begin   int64
	End     int64
}, error) {
	return _Store.Contract.GetKey(&_Store.CallOpts, keyuid)
}

// GetKey is a free data retrieval call binding the contract method 0x12aaac70.
//
// Solidity: function getKey(keyuid bytes32) constant returns(key string, revoked bool, begin int64, end int64)
func (_Store *StoreCallerSession) GetKey(keyuid [32]byte) (struct {
	Key     string
	Revoked bool
	Begin   int64
	End     int64
}, error) {
	return _Store.Contract.GetKey(&_Store.CallOpts, keyuid)
}

// GetKeyX is a free data retrieval call binding the contract method 0x46a9e287.
//
// Solidity: function getKeyX(keyuid bytes32) constant returns(x0 bytes32, x1 bytes32, x2 bytes32, x3 bytes32)
func (_Store *StoreCaller) GetKeyX(opts *bind.CallOpts, keyuid [32]byte) (struct {
	X0 [32]byte
	X1 [32]byte
	X2 [32]byte
	X3 [32]byte
}, error) {
	ret := new(struct {
		X0 [32]byte
		X1 [32]byte
		X2 [32]byte
		X3 [32]byte
	})
	out := ret
	err := _Store.contract.Call(opts, out, "getKeyX", keyuid)
	return *ret, err
}

// GetKeyX is a free data retrieval call binding the contract method 0x46a9e287.
//
// Solidity: function getKeyX(keyuid bytes32) constant returns(x0 bytes32, x1 bytes32, x2 bytes32, x3 bytes32)
func (_Store *StoreSession) GetKeyX(keyuid [32]byte) (struct {
	X0 [32]byte
	X1 [32]byte
	X2 [32]byte
	X3 [32]byte
}, error) {
	return _Store.Contract.GetKeyX(&_Store.CallOpts, keyuid)
}

// GetKeyX is a free data retrieval call binding the contract method 0x46a9e287.
//
// Solidity: function getKeyX(keyuid bytes32) constant returns(x0 bytes32, x1 bytes32, x2 bytes32, x3 bytes32)
func (_Store *StoreCallerSession) GetKeyX(keyuid [32]byte) (struct {
	X0 [32]byte
	X1 [32]byte
	X2 [32]byte
	X3 [32]byte
}, error) {
	return _Store.Contract.GetKeyX(&_Store.CallOpts, keyuid)
}

// GetProof is a free data retrieval call binding the contract method 0x8a16a35f.
//
// Solidity: function getProof(hash bytes32, cardId bytes32) constant returns(puid bytes32, sig string, metadata string, revoked bool, blockWrite int64)
func (_Store *StoreCaller) GetProof(opts *bind.CallOpts, hash [32]byte, cardId [32]byte) (struct {
	Puid       [32]byte
	Sig        string
	Metadata   string
	Revoked    bool
	BlockWrite int64
}, error) {
	ret := new(struct {
		Puid       [32]byte
		Sig        string
		Metadata   string
		Revoked    bool
		BlockWrite int64
	})
	out := ret
	err := _Store.contract.Call(opts, out, "getProof", hash, cardId)
	return *ret, err
}

// GetProof is a free data retrieval call binding the contract method 0x8a16a35f.
//
// Solidity: function getProof(hash bytes32, cardId bytes32) constant returns(puid bytes32, sig string, metadata string, revoked bool, blockWrite int64)
func (_Store *StoreSession) GetProof(hash [32]byte, cardId [32]byte) (struct {
	Puid       [32]byte
	Sig        string
	Metadata   string
	Revoked    bool
	BlockWrite int64
}, error) {
	return _Store.Contract.GetProof(&_Store.CallOpts, hash, cardId)
}

// GetProof is a free data retrieval call binding the contract method 0x8a16a35f.
//
// Solidity: function getProof(hash bytes32, cardId bytes32) constant returns(puid bytes32, sig string, metadata string, revoked bool, blockWrite int64)
func (_Store *StoreCallerSession) GetProof(hash [32]byte, cardId [32]byte) (struct {
	Puid       [32]byte
	Sig        string
	Metadata   string
	Revoked    bool
	BlockWrite int64
}, error) {
	return _Store.Contract.GetProof(&_Store.CallOpts, hash, cardId)
}

// GetProofB32 is a free data retrieval call binding the contract method 0xa8a67ede.
//
// Solidity: function getProofB32(hash bytes32, cardId bytes32) constant returns(puid bytes32, sig bytes32, metadata bytes32, revoked bool, blockWrite int64)
func (_Store *StoreCaller) GetProofB32(opts *bind.CallOpts, hash [32]byte, cardId [32]byte) (struct {
	Puid       [32]byte
	Sig        [32]byte
	Metadata   [32]byte
	Revoked    bool
	BlockWrite int64
}, error) {
	ret := new(struct {
		Puid       [32]byte
		Sig        [32]byte
		Metadata   [32]byte
		Revoked    bool
		BlockWrite int64
	})
	out := ret
	err := _Store.contract.Call(opts, out, "getProofB32", hash, cardId)
	return *ret, err
}

// GetProofB32 is a free data retrieval call binding the contract method 0xa8a67ede.
//
// Solidity: function getProofB32(hash bytes32, cardId bytes32) constant returns(puid bytes32, sig bytes32, metadata bytes32, revoked bool, blockWrite int64)
func (_Store *StoreSession) GetProofB32(hash [32]byte, cardId [32]byte) (struct {
	Puid       [32]byte
	Sig        [32]byte
	Metadata   [32]byte
	Revoked    bool
	BlockWrite int64
}, error) {
	return _Store.Contract.GetProofB32(&_Store.CallOpts, hash, cardId)
}

// GetProofB32 is a free data retrieval call binding the contract method 0xa8a67ede.
//
// Solidity: function getProofB32(hash bytes32, cardId bytes32) constant returns(puid bytes32, sig bytes32, metadata bytes32, revoked bool, blockWrite int64)
func (_Store *StoreCallerSession) GetProofB32(hash [32]byte, cardId [32]byte) (struct {
	Puid       [32]byte
	Sig        [32]byte
	Metadata   [32]byte
	Revoked    bool
	BlockWrite int64
}, error) {
	return _Store.Contract.GetProofB32(&_Store.CallOpts, hash, cardId)
}

// GetProofX is a free data retrieval call binding the contract method 0x31744e64.
//
// Solidity: function getProofX(k bytes32) constant returns(x0 bytes32, x1 bytes32, x2 bytes32, x3 bytes32)
func (_Store *StoreCaller) GetProofX(opts *bind.CallOpts, k [32]byte) (struct {
	X0 [32]byte
	X1 [32]byte
	X2 [32]byte
	X3 [32]byte
}, error) {
	ret := new(struct {
		X0 [32]byte
		X1 [32]byte
		X2 [32]byte
		X3 [32]byte
	})
	out := ret
	err := _Store.contract.Call(opts, out, "getProofX", k)
	return *ret, err
}

// GetProofX is a free data retrieval call binding the contract method 0x31744e64.
//
// Solidity: function getProofX(k bytes32) constant returns(x0 bytes32, x1 bytes32, x2 bytes32, x3 bytes32)
func (_Store *StoreSession) GetProofX(k [32]byte) (struct {
	X0 [32]byte
	X1 [32]byte
	X2 [32]byte
	X3 [32]byte
}, error) {
	return _Store.Contract.GetProofX(&_Store.CallOpts, k)
}

// GetProofX is a free data retrieval call binding the contract method 0x31744e64.
//
// Solidity: function getProofX(k bytes32) constant returns(x0 bytes32, x1 bytes32, x2 bytes32, x3 bytes32)
func (_Store *StoreCallerSession) GetProofX(k [32]byte) (struct {
	X0 [32]byte
	X1 [32]byte
	X2 [32]byte
	X3 [32]byte
}, error) {
	return _Store.Contract.GetProofX(&_Store.CallOpts, k)
}

// GetString1 is a free data retrieval call binding the contract method 0xd3d9a829.
//
// Solidity: function getString1(k bytes32) constant returns(s string)
func (_Store *StoreCaller) GetString1(opts *bind.CallOpts, k [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "getString1", k)
	return *ret0, err
}

// GetString1 is a free data retrieval call binding the contract method 0xd3d9a829.
//
// Solidity: function getString1(k bytes32) constant returns(s string)
func (_Store *StoreSession) GetString1(k [32]byte) (string, error) {
	return _Store.Contract.GetString1(&_Store.CallOpts, k)
}

// GetString1 is a free data retrieval call binding the contract method 0xd3d9a829.
//
// Solidity: function getString1(k bytes32) constant returns(s string)
func (_Store *StoreCallerSession) GetString1(k [32]byte) (string, error) {
	return _Store.Contract.GetString1(&_Store.CallOpts, k)
}

// GetString3 is a free data retrieval call binding the contract method 0xa7dd645e.
//
// Solidity: function getString3(k bytes32, k2 bytes32, k3 bytes32) constant returns(s string, s2 string, s3 string)
func (_Store *StoreCaller) GetString3(opts *bind.CallOpts, k [32]byte, k2 [32]byte, k3 [32]byte) (struct {
	S  string
	S2 string
	S3 string
}, error) {
	ret := new(struct {
		S  string
		S2 string
		S3 string
	})
	out := ret
	err := _Store.contract.Call(opts, out, "getString3", k, k2, k3)
	return *ret, err
}

// GetString3 is a free data retrieval call binding the contract method 0xa7dd645e.
//
// Solidity: function getString3(k bytes32, k2 bytes32, k3 bytes32) constant returns(s string, s2 string, s3 string)
func (_Store *StoreSession) GetString3(k [32]byte, k2 [32]byte, k3 [32]byte) (struct {
	S  string
	S2 string
	S3 string
}, error) {
	return _Store.Contract.GetString3(&_Store.CallOpts, k, k2, k3)
}

// GetString3 is a free data retrieval call binding the contract method 0xa7dd645e.
//
// Solidity: function getString3(k bytes32, k2 bytes32, k3 bytes32) constant returns(s string, s2 string, s3 string)
func (_Store *StoreCallerSession) GetString3(k [32]byte, k2 [32]byte, k3 [32]byte) (struct {
	S  string
	S2 string
	S3 string
}, error) {
	return _Store.Contract.GetString3(&_Store.CallOpts, k, k2, k3)
}

// GetUserCardAt is a free data retrieval call binding the contract method 0x224af50b.
//
// Solidity: function getUserCardAt(useruid bytes32, index uint256) constant returns(cardId bytes32)
func (_Store *StoreCaller) GetUserCardAt(opts *bind.CallOpts, useruid [32]byte, index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "getUserCardAt", useruid, index)
	return *ret0, err
}

// GetUserCardAt is a free data retrieval call binding the contract method 0x224af50b.
//
// Solidity: function getUserCardAt(useruid bytes32, index uint256) constant returns(cardId bytes32)
func (_Store *StoreSession) GetUserCardAt(useruid [32]byte, index *big.Int) ([32]byte, error) {
	return _Store.Contract.GetUserCardAt(&_Store.CallOpts, useruid, index)
}

// GetUserCardAt is a free data retrieval call binding the contract method 0x224af50b.
//
// Solidity: function getUserCardAt(useruid bytes32, index uint256) constant returns(cardId bytes32)
func (_Store *StoreCallerSession) GetUserCardAt(useruid [32]byte, index *big.Int) ([32]byte, error) {
	return _Store.Contract.GetUserCardAt(&_Store.CallOpts, useruid, index)
}

// GetUserCardLen is a free data retrieval call binding the contract method 0x34650b9f.
//
// Solidity: function getUserCardLen(useruid bytes32) constant returns(uint256)
func (_Store *StoreCaller) GetUserCardLen(opts *bind.CallOpts, useruid [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "getUserCardLen", useruid)
	return *ret0, err
}

// GetUserCardLen is a free data retrieval call binding the contract method 0x34650b9f.
//
// Solidity: function getUserCardLen(useruid bytes32) constant returns(uint256)
func (_Store *StoreSession) GetUserCardLen(useruid [32]byte) (*big.Int, error) {
	return _Store.Contract.GetUserCardLen(&_Store.CallOpts, useruid)
}

// GetUserCardLen is a free data retrieval call binding the contract method 0x34650b9f.
//
// Solidity: function getUserCardLen(useruid bytes32) constant returns(uint256)
func (_Store *StoreCallerSession) GetUserCardLen(useruid [32]byte) (*big.Int, error) {
	return _Store.Contract.GetUserCardLen(&_Store.CallOpts, useruid)
}

// HasCard is a free data retrieval call binding the contract method 0xb18dd740.
//
// Solidity: function hasCard(cardId bytes32) constant returns(bool)
func (_Store *StoreCaller) HasCard(opts *bind.CallOpts, cardId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "hasCard", cardId)
	return *ret0, err
}

// HasCard is a free data retrieval call binding the contract method 0xb18dd740.
//
// Solidity: function hasCard(cardId bytes32) constant returns(bool)
func (_Store *StoreSession) HasCard(cardId [32]byte) (bool, error) {
	return _Store.Contract.HasCard(&_Store.CallOpts, cardId)
}

// HasCard is a free data retrieval call binding the contract method 0xb18dd740.
//
// Solidity: function hasCard(cardId bytes32) constant returns(bool)
func (_Store *StoreCallerSession) HasCard(cardId [32]byte) (bool, error) {
	return _Store.Contract.HasCard(&_Store.CallOpts, cardId)
}

// HasCurrentCard is a free data retrieval call binding the contract method 0x94d62c29.
//
// Solidity: function hasCurrentCard(user address) constant returns(bool)
func (_Store *StoreCaller) HasCurrentCard(opts *bind.CallOpts, user common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "hasCurrentCard", user)
	return *ret0, err
}

// HasCurrentCard is a free data retrieval call binding the contract method 0x94d62c29.
//
// Solidity: function hasCurrentCard(user address) constant returns(bool)
func (_Store *StoreSession) HasCurrentCard(user common.Address) (bool, error) {
	return _Store.Contract.HasCurrentCard(&_Store.CallOpts, user)
}

// HasCurrentCard is a free data retrieval call binding the contract method 0x94d62c29.
//
// Solidity: function hasCurrentCard(user address) constant returns(bool)
func (_Store *StoreCallerSession) HasCurrentCard(user common.Address) (bool, error) {
	return _Store.Contract.HasCurrentCard(&_Store.CallOpts, user)
}

// HasKey is a free data retrieval call binding the contract method 0x696eb375.
//
// Solidity: function hasKey(keyuid bytes32) constant returns(bool)
func (_Store *StoreCaller) HasKey(opts *bind.CallOpts, keyuid [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "hasKey", keyuid)
	return *ret0, err
}

// HasKey is a free data retrieval call binding the contract method 0x696eb375.
//
// Solidity: function hasKey(keyuid bytes32) constant returns(bool)
func (_Store *StoreSession) HasKey(keyuid [32]byte) (bool, error) {
	return _Store.Contract.HasKey(&_Store.CallOpts, keyuid)
}

// HasKey is a free data retrieval call binding the contract method 0x696eb375.
//
// Solidity: function hasKey(keyuid bytes32) constant returns(bool)
func (_Store *StoreCallerSession) HasKey(keyuid [32]byte) (bool, error) {
	return _Store.Contract.HasKey(&_Store.CallOpts, keyuid)
}

// HasProof is a free data retrieval call binding the contract method 0xe3d1e6d6.
//
// Solidity: function hasProof(k bytes32) constant returns(bool)
func (_Store *StoreCaller) HasProof(opts *bind.CallOpts, k [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "hasProof", k)
	return *ret0, err
}

// HasProof is a free data retrieval call binding the contract method 0xe3d1e6d6.
//
// Solidity: function hasProof(k bytes32) constant returns(bool)
func (_Store *StoreSession) HasProof(k [32]byte) (bool, error) {
	return _Store.Contract.HasProof(&_Store.CallOpts, k)
}

// HasProof is a free data retrieval call binding the contract method 0xe3d1e6d6.
//
// Solidity: function hasProof(k bytes32) constant returns(bool)
func (_Store *StoreCallerSession) HasProof(k [32]byte) (bool, error) {
	return _Store.Contract.HasProof(&_Store.CallOpts, k)
}

// IsKeyRevoked is a free data retrieval call binding the contract method 0x57254421.
//
// Solidity: function isKeyRevoked(keyuid bytes32) constant returns(bool)
func (_Store *StoreCaller) IsKeyRevoked(opts *bind.CallOpts, keyuid [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "isKeyRevoked", keyuid)
	return *ret0, err
}

// IsKeyRevoked is a free data retrieval call binding the contract method 0x57254421.
//
// Solidity: function isKeyRevoked(keyuid bytes32) constant returns(bool)
func (_Store *StoreSession) IsKeyRevoked(keyuid [32]byte) (bool, error) {
	return _Store.Contract.IsKeyRevoked(&_Store.CallOpts, keyuid)
}

// IsKeyRevoked is a free data retrieval call binding the contract method 0x57254421.
//
// Solidity: function isKeyRevoked(keyuid bytes32) constant returns(bool)
func (_Store *StoreCallerSession) IsKeyRevoked(keyuid [32]byte) (bool, error) {
	return _Store.Contract.IsKeyRevoked(&_Store.CallOpts, keyuid)
}

// IsKeyValid is a free data retrieval call binding the contract method 0xee863317.
//
// Solidity: function isKeyValid(keyuid bytes32) constant returns(bool)
func (_Store *StoreCaller) IsKeyValid(opts *bind.CallOpts, keyuid [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "isKeyValid", keyuid)
	return *ret0, err
}

// IsKeyValid is a free data retrieval call binding the contract method 0xee863317.
//
// Solidity: function isKeyValid(keyuid bytes32) constant returns(bool)
func (_Store *StoreSession) IsKeyValid(keyuid [32]byte) (bool, error) {
	return _Store.Contract.IsKeyValid(&_Store.CallOpts, keyuid)
}

// IsKeyValid is a free data retrieval call binding the contract method 0xee863317.
//
// Solidity: function isKeyValid(keyuid bytes32) constant returns(bool)
func (_Store *StoreCallerSession) IsKeyValid(keyuid [32]byte) (bool, error) {
	return _Store.Contract.IsKeyValid(&_Store.CallOpts, keyuid)
}

// IsKeyValidAt is a free data retrieval call binding the contract method 0x5fb1d8f2.
//
// Solidity: function isKeyValidAt(keyuid bytes32, date int64) constant returns(bool)
func (_Store *StoreCaller) IsKeyValidAt(opts *bind.CallOpts, keyuid [32]byte, date int64) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "isKeyValidAt", keyuid, date)
	return *ret0, err
}

// IsKeyValidAt is a free data retrieval call binding the contract method 0x5fb1d8f2.
//
// Solidity: function isKeyValidAt(keyuid bytes32, date int64) constant returns(bool)
func (_Store *StoreSession) IsKeyValidAt(keyuid [32]byte, date int64) (bool, error) {
	return _Store.Contract.IsKeyValidAt(&_Store.CallOpts, keyuid, date)
}

// IsKeyValidAt is a free data retrieval call binding the contract method 0x5fb1d8f2.
//
// Solidity: function isKeyValidAt(keyuid bytes32, date int64) constant returns(bool)
func (_Store *StoreCallerSession) IsKeyValidAt(keyuid [32]byte, date int64) (bool, error) {
	return _Store.Contract.IsKeyValidAt(&_Store.CallOpts, keyuid, date)
}

// IsProofRevoked is a free data retrieval call binding the contract method 0xb93ac300.
//
// Solidity: function isProofRevoked(k bytes32) constant returns(revoked bool)
func (_Store *StoreCaller) IsProofRevoked(opts *bind.CallOpts, k [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "isProofRevoked", k)
	return *ret0, err
}

// IsProofRevoked is a free data retrieval call binding the contract method 0xb93ac300.
//
// Solidity: function isProofRevoked(k bytes32) constant returns(revoked bool)
func (_Store *StoreSession) IsProofRevoked(k [32]byte) (bool, error) {
	return _Store.Contract.IsProofRevoked(&_Store.CallOpts, k)
}

// IsProofRevoked is a free data retrieval call binding the contract method 0xb93ac300.
//
// Solidity: function isProofRevoked(k bytes32) constant returns(revoked bool)
func (_Store *StoreCallerSession) IsProofRevoked(k [32]byte) (bool, error) {
	return _Store.Contract.IsProofRevoked(&_Store.CallOpts, k)
}

// Keccak256StringToB32 is a free data retrieval call binding the contract method 0x4e70ef64.
//
// Solidity: function keccak256StringToB32(v string) constant returns(key bytes32)
func (_Store *StoreCaller) Keccak256StringToB32(opts *bind.CallOpts, v string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "keccak256StringToB32", v)
	return *ret0, err
}

// Keccak256StringToB32 is a free data retrieval call binding the contract method 0x4e70ef64.
//
// Solidity: function keccak256StringToB32(v string) constant returns(key bytes32)
func (_Store *StoreSession) Keccak256StringToB32(v string) ([32]byte, error) {
	return _Store.Contract.Keccak256StringToB32(&_Store.CallOpts, v)
}

// Keccak256StringToB32 is a free data retrieval call binding the contract method 0x4e70ef64.
//
// Solidity: function keccak256StringToB32(v string) constant returns(key bytes32)
func (_Store *StoreCallerSession) Keccak256StringToB32(v string) ([32]byte, error) {
	return _Store.Contract.Keccak256StringToB32(&_Store.CallOpts, v)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Store *StoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Store *StoreSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Store *StoreCallerSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// CloseKey is a paid mutator transaction binding the contract method 0x9a734c0a.
//
// Solidity: function closeKey(keyuid bytes32) returns()
func (_Store *StoreTransactor) CloseKey(opts *bind.TransactOpts, keyuid [32]byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "closeKey", keyuid)
}

// CloseKey is a paid mutator transaction binding the contract method 0x9a734c0a.
//
// Solidity: function closeKey(keyuid bytes32) returns()
func (_Store *StoreSession) CloseKey(keyuid [32]byte) (*types.Transaction, error) {
	return _Store.Contract.CloseKey(&_Store.TransactOpts, keyuid)
}

// CloseKey is a paid mutator transaction binding the contract method 0x9a734c0a.
//
// Solidity: function closeKey(keyuid bytes32) returns()
func (_Store *StoreTransactorSession) CloseKey(keyuid [32]byte) (*types.Transaction, error) {
	return _Store.Contract.CloseKey(&_Store.TransactOpts, keyuid)
}

// ResetUpgradeCard is a paid mutator transaction binding the contract method 0x6b456785.
//
// Solidity: function resetUpgradeCard(cardId bytes32) returns()
func (_Store *StoreTransactor) ResetUpgradeCard(opts *bind.TransactOpts, cardId [32]byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "resetUpgradeCard", cardId)
}

// ResetUpgradeCard is a paid mutator transaction binding the contract method 0x6b456785.
//
// Solidity: function resetUpgradeCard(cardId bytes32) returns()
func (_Store *StoreSession) ResetUpgradeCard(cardId [32]byte) (*types.Transaction, error) {
	return _Store.Contract.ResetUpgradeCard(&_Store.TransactOpts, cardId)
}

// ResetUpgradeCard is a paid mutator transaction binding the contract method 0x6b456785.
//
// Solidity: function resetUpgradeCard(cardId bytes32) returns()
func (_Store *StoreTransactorSession) ResetUpgradeCard(cardId [32]byte) (*types.Transaction, error) {
	return _Store.Contract.ResetUpgradeCard(&_Store.TransactOpts, cardId)
}

// RevokeKey is a paid mutator transaction binding the contract method 0x572f2210.
//
// Solidity: function revokeKey(keyuid bytes32) returns()
func (_Store *StoreTransactor) RevokeKey(opts *bind.TransactOpts, keyuid [32]byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "revokeKey", keyuid)
}

// RevokeKey is a paid mutator transaction binding the contract method 0x572f2210.
//
// Solidity: function revokeKey(keyuid bytes32) returns()
func (_Store *StoreSession) RevokeKey(keyuid [32]byte) (*types.Transaction, error) {
	return _Store.Contract.RevokeKey(&_Store.TransactOpts, keyuid)
}

// RevokeKey is a paid mutator transaction binding the contract method 0x572f2210.
//
// Solidity: function revokeKey(keyuid bytes32) returns()
func (_Store *StoreTransactorSession) RevokeKey(keyuid [32]byte) (*types.Transaction, error) {
	return _Store.Contract.RevokeKey(&_Store.TransactOpts, keyuid)
}

// RevokeProof is a paid mutator transaction binding the contract method 0x426a8745.
//
// Solidity: function revokeProof(k bytes32) returns()
func (_Store *StoreTransactor) RevokeProof(opts *bind.TransactOpts, k [32]byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "revokeProof", k)
}

// RevokeProof is a paid mutator transaction binding the contract method 0x426a8745.
//
// Solidity: function revokeProof(k bytes32) returns()
func (_Store *StoreSession) RevokeProof(k [32]byte) (*types.Transaction, error) {
	return _Store.Contract.RevokeProof(&_Store.TransactOpts, k)
}

// RevokeProof is a paid mutator transaction binding the contract method 0x426a8745.
//
// Solidity: function revokeProof(k bytes32) returns()
func (_Store *StoreTransactorSession) RevokeProof(k [32]byte) (*types.Transaction, error) {
	return _Store.Contract.RevokeProof(&_Store.TransactOpts, k)
}

// SetCard is a paid mutator transaction binding the contract method 0x614ed9ba.
//
// Solidity: function setCard(cardId bytes32, useruid bytes32, key bytes32, owner address) returns()
func (_Store *StoreTransactor) SetCard(opts *bind.TransactOpts, cardId [32]byte, useruid [32]byte, key [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setCard", cardId, useruid, key, owner)
}

// SetCard is a paid mutator transaction binding the contract method 0x614ed9ba.
//
// Solidity: function setCard(cardId bytes32, useruid bytes32, key bytes32, owner address) returns()
func (_Store *StoreSession) SetCard(cardId [32]byte, useruid [32]byte, key [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetCard(&_Store.TransactOpts, cardId, useruid, key, owner)
}

// SetCard is a paid mutator transaction binding the contract method 0x614ed9ba.
//
// Solidity: function setCard(cardId bytes32, useruid bytes32, key bytes32, owner address) returns()
func (_Store *StoreTransactorSession) SetCard(cardId [32]byte, useruid [32]byte, key [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetCard(&_Store.TransactOpts, cardId, useruid, key, owner)
}

// SetCardX is a paid mutator transaction binding the contract method 0x770b9429.
//
// Solidity: function setCardX(cardId bytes32, value bytes32, index int256) returns()
func (_Store *StoreTransactor) SetCardX(opts *bind.TransactOpts, cardId [32]byte, value [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setCardX", cardId, value, index)
}

// SetCardX is a paid mutator transaction binding the contract method 0x770b9429.
//
// Solidity: function setCardX(cardId bytes32, value bytes32, index int256) returns()
func (_Store *StoreSession) SetCardX(cardId [32]byte, value [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetCardX(&_Store.TransactOpts, cardId, value, index)
}

// SetCardX is a paid mutator transaction binding the contract method 0x770b9429.
//
// Solidity: function setCardX(cardId bytes32, value bytes32, index int256) returns()
func (_Store *StoreTransactorSession) SetCardX(cardId [32]byte, value [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetCardX(&_Store.TransactOpts, cardId, value, index)
}

// SetContract is a paid mutator transaction binding the contract method 0x75f890ab.
//
// Solidity: function setContract(contract_ address) returns()
func (_Store *StoreTransactor) SetContract(opts *bind.TransactOpts, contract_ common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setContract", contract_)
}

// SetContract is a paid mutator transaction binding the contract method 0x75f890ab.
//
// Solidity: function setContract(contract_ address) returns()
func (_Store *StoreSession) SetContract(contract_ common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetContract(&_Store.TransactOpts, contract_)
}

// SetContract is a paid mutator transaction binding the contract method 0x75f890ab.
//
// Solidity: function setContract(contract_ address) returns()
func (_Store *StoreTransactorSession) SetContract(contract_ common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetContract(&_Store.TransactOpts, contract_)
}

// SetKey is a paid mutator transaction binding the contract method 0x8231d541.
//
// Solidity: function setKey(keyuid bytes32, keySID bytes32, begin int64, end int64) returns()
func (_Store *StoreTransactor) SetKey(opts *bind.TransactOpts, keyuid [32]byte, keySID [32]byte, begin int64, end int64) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setKey", keyuid, keySID, begin, end)
}

// SetKey is a paid mutator transaction binding the contract method 0x8231d541.
//
// Solidity: function setKey(keyuid bytes32, keySID bytes32, begin int64, end int64) returns()
func (_Store *StoreSession) SetKey(keyuid [32]byte, keySID [32]byte, begin int64, end int64) (*types.Transaction, error) {
	return _Store.Contract.SetKey(&_Store.TransactOpts, keyuid, keySID, begin, end)
}

// SetKey is a paid mutator transaction binding the contract method 0x8231d541.
//
// Solidity: function setKey(keyuid bytes32, keySID bytes32, begin int64, end int64) returns()
func (_Store *StoreTransactorSession) SetKey(keyuid [32]byte, keySID [32]byte, begin int64, end int64) (*types.Transaction, error) {
	return _Store.Contract.SetKey(&_Store.TransactOpts, keyuid, keySID, begin, end)
}

// SetKeyX is a paid mutator transaction binding the contract method 0x3733c2d7.
//
// Solidity: function setKeyX(keyuid bytes32, value bytes32, index int256) returns()
func (_Store *StoreTransactor) SetKeyX(opts *bind.TransactOpts, keyuid [32]byte, value [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setKeyX", keyuid, value, index)
}

// SetKeyX is a paid mutator transaction binding the contract method 0x3733c2d7.
//
// Solidity: function setKeyX(keyuid bytes32, value bytes32, index int256) returns()
func (_Store *StoreSession) SetKeyX(keyuid [32]byte, value [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetKeyX(&_Store.TransactOpts, keyuid, value, index)
}

// SetKeyX is a paid mutator transaction binding the contract method 0x3733c2d7.
//
// Solidity: function setKeyX(keyuid bytes32, value bytes32, index int256) returns()
func (_Store *StoreTransactorSession) SetKeyX(keyuid [32]byte, value [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetKeyX(&_Store.TransactOpts, keyuid, value, index)
}

// SetProof is a paid mutator transaction binding the contract method 0x4966bfc7.
//
// Solidity: function setProof(hash bytes32, cardId bytes32, sig bytes32, metadata bytes32) returns(k bytes32)
func (_Store *StoreTransactor) SetProof(opts *bind.TransactOpts, hash [32]byte, cardId [32]byte, sig [32]byte, metadata [32]byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setProof", hash, cardId, sig, metadata)
}

// SetProof is a paid mutator transaction binding the contract method 0x4966bfc7.
//
// Solidity: function setProof(hash bytes32, cardId bytes32, sig bytes32, metadata bytes32) returns(k bytes32)
func (_Store *StoreSession) SetProof(hash [32]byte, cardId [32]byte, sig [32]byte, metadata [32]byte) (*types.Transaction, error) {
	return _Store.Contract.SetProof(&_Store.TransactOpts, hash, cardId, sig, metadata)
}

// SetProof is a paid mutator transaction binding the contract method 0x4966bfc7.
//
// Solidity: function setProof(hash bytes32, cardId bytes32, sig bytes32, metadata bytes32) returns(k bytes32)
func (_Store *StoreTransactorSession) SetProof(hash [32]byte, cardId [32]byte, sig [32]byte, metadata [32]byte) (*types.Transaction, error) {
	return _Store.Contract.SetProof(&_Store.TransactOpts, hash, cardId, sig, metadata)
}

// SetProofX is a paid mutator transaction binding the contract method 0x5ff0605c.
//
// Solidity: function setProofX(k bytes32, value bytes32, index int256) returns()
func (_Store *StoreTransactor) SetProofX(opts *bind.TransactOpts, k [32]byte, value [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setProofX", k, value, index)
}

// SetProofX is a paid mutator transaction binding the contract method 0x5ff0605c.
//
// Solidity: function setProofX(k bytes32, value bytes32, index int256) returns()
func (_Store *StoreSession) SetProofX(k [32]byte, value [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetProofX(&_Store.TransactOpts, k, value, index)
}

// SetProofX is a paid mutator transaction binding the contract method 0x5ff0605c.
//
// Solidity: function setProofX(k bytes32, value bytes32, index int256) returns()
func (_Store *StoreTransactorSession) SetProofX(k [32]byte, value [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetProofX(&_Store.TransactOpts, k, value, index)
}

// SetString1 is a paid mutator transaction binding the contract method 0xbb3da883.
//
// Solidity: function setString1(str string) returns(bytes32)
func (_Store *StoreTransactor) SetString1(opts *bind.TransactOpts, str string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setString1", str)
}

// SetString1 is a paid mutator transaction binding the contract method 0xbb3da883.
//
// Solidity: function setString1(str string) returns(bytes32)
func (_Store *StoreSession) SetString1(str string) (*types.Transaction, error) {
	return _Store.Contract.SetString1(&_Store.TransactOpts, str)
}

// SetString1 is a paid mutator transaction binding the contract method 0xbb3da883.
//
// Solidity: function setString1(str string) returns(bytes32)
func (_Store *StoreTransactorSession) SetString1(str string) (*types.Transaction, error) {
	return _Store.Contract.SetString1(&_Store.TransactOpts, str)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Store *StoreTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Store *StoreSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Store.Contract.TransferOwnership(&_Store.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Store *StoreTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Store.Contract.TransferOwnership(&_Store.TransactOpts, newOwner)
}

// UpgradeCard is a paid mutator transaction binding the contract method 0x5447f57a.
//
// Solidity: function upgradeCard(cardId bytes32, newCardId bytes32) returns()
func (_Store *StoreTransactor) UpgradeCard(opts *bind.TransactOpts, cardId [32]byte, newCardId [32]byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "upgradeCard", cardId, newCardId)
}

// UpgradeCard is a paid mutator transaction binding the contract method 0x5447f57a.
//
// Solidity: function upgradeCard(cardId bytes32, newCardId bytes32) returns()
func (_Store *StoreSession) UpgradeCard(cardId [32]byte, newCardId [32]byte) (*types.Transaction, error) {
	return _Store.Contract.UpgradeCard(&_Store.TransactOpts, cardId, newCardId)
}

// UpgradeCard is a paid mutator transaction binding the contract method 0x5447f57a.
//
// Solidity: function upgradeCard(cardId bytes32, newCardId bytes32) returns()
func (_Store *StoreTransactorSession) UpgradeCard(cardId [32]byte, newCardId [32]byte) (*types.Transaction, error) {
	return _Store.Contract.UpgradeCard(&_Store.TransactOpts, cardId, newCardId)
}
