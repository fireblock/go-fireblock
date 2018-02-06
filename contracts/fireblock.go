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

// FireblockABI is the input ABI used to generate the binding from.
const FireblockABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"revokeValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"writer\",\"type\":\"address\"}],\"name\":\"revokeWriter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"revokeMaker\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"indexes\",\"type\":\"int16[]\"}],\"name\":\"validate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"keyuid\",\"type\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"string\"}],\"name\":\"closeKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"keyuid\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"sig\",\"type\":\"string\"}],\"name\":\"registerKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operationQty\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"irole\",\"type\":\"uint16\"}],\"name\":\"roles\",\"outputs\":[{\"name\":\"res\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"tick\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"revokeAdministrator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"cardId\",\"type\":\"bytes32\"}],\"name\":\"verify\",\"outputs\":[{\"name\":\"validity\",\"type\":\"bool\"},{\"name\":\"key\",\"type\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes32\"},{\"name\":\"upgrade\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"store_\",\"type\":\"address\"}],\"name\":\"setDataStore\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"irole\",\"type\":\"uint16\"}],\"name\":\"roleQty\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"useruid\",\"type\":\"bytes32\"},{\"name\":\"card\",\"type\":\"string\"},{\"name\":\"keyuid\",\"type\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"string\"},{\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"sign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cardId\",\"type\":\"bytes32\"}],\"name\":\"resetUpgradeCardAdm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"useruid\",\"type\":\"bytes32\"}],\"name\":\"verifyByUserUID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[64]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"oldCardId\",\"type\":\"bytes32\"},{\"name\":\"newCardId\",\"type\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"string\"}],\"name\":\"upgradeCard\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operationIndexes\",\"outputs\":[{\"name\":\"res\",\"type\":\"int16[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"oldCardId\",\"type\":\"bytes32\"},{\"name\":\"newCardId\",\"type\":\"bytes32\"}],\"name\":\"upgradeCardEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"cardId\",\"type\":\"bytes32\"}],\"name\":\"revokeProofAdm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"int16\"}],\"name\":\"operationAt\",\"outputs\":[{\"name\":\"otype\",\"type\":\"int16\"},{\"name\":\"blockWrite\",\"type\":\"int64\"},{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"useruid\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"cardId\",\"type\":\"bytes32\"},{\"name\":\"keyuid\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes32\"},{\"name\":\"metadata\",\"type\":\"bytes32\"},{\"name\":\"used\",\"type\":\"bool\"},{\"name\":\"seals\",\"type\":\"address[16]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"registerValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"registerAdministrator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"signEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"registerMaker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"role\",\"type\":\"address\"},{\"name\":\"irole\",\"type\":\"uint16\"}],\"name\":\"isRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"revokeKeyEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"revokeKeyAdm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"closeKeyEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"useruid\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"card\",\"type\":\"string\"}],\"name\":\"registerCardEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"revokeProofEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"cardId\",\"type\":\"bytes32\"},{\"name\":\"keyuid\",\"type\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"string\"}],\"name\":\"revokeProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"keyuid\",\"type\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"string\"}],\"name\":\"revokeKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"writer\",\"type\":\"address\"}],\"name\":\"registerWriter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"irole\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"RoleRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"op\",\"type\":\"uint256\"}],\"name\":\"OperationUnknown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"op\",\"type\":\"uint256\"}],\"name\":\"OperationAlreadyDone\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"op\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"arg\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"arg2\",\"type\":\"bytes32\"}],\"name\":\"OperationInvalid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Fireblock is an auto generated Go binding around an Ethereum contract.
type Fireblock struct {
	FireblockCaller     // Read-only binding to the contract
	FireblockTransactor // Write-only binding to the contract
}

// FireblockCaller is an auto generated read-only Go binding around an Ethereum contract.
type FireblockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FireblockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FireblockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FireblockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FireblockSession struct {
	Contract     *Fireblock        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FireblockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FireblockCallerSession struct {
	Contract *FireblockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FireblockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FireblockTransactorSession struct {
	Contract     *FireblockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FireblockRaw is an auto generated low-level Go binding around an Ethereum contract.
type FireblockRaw struct {
	Contract *Fireblock // Generic contract binding to access the raw methods on
}

// FireblockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FireblockCallerRaw struct {
	Contract *FireblockCaller // Generic read-only contract binding to access the raw methods on
}

// FireblockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FireblockTransactorRaw struct {
	Contract *FireblockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFireblock creates a new instance of Fireblock, bound to a specific deployed contract.
func NewFireblock(address common.Address, backend bind.ContractBackend) (*Fireblock, error) {
	contract, err := bindFireblock(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fireblock{FireblockCaller: FireblockCaller{contract: contract}, FireblockTransactor: FireblockTransactor{contract: contract}}, nil
}

// NewFireblockCaller creates a new read-only instance of Fireblock, bound to a specific deployed contract.
func NewFireblockCaller(address common.Address, caller bind.ContractCaller) (*FireblockCaller, error) {
	contract, err := bindFireblock(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &FireblockCaller{contract: contract}, nil
}

// NewFireblockTransactor creates a new write-only instance of Fireblock, bound to a specific deployed contract.
func NewFireblockTransactor(address common.Address, transactor bind.ContractTransactor) (*FireblockTransactor, error) {
	contract, err := bindFireblock(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &FireblockTransactor{contract: contract}, nil
}

// bindFireblock binds a generic wrapper to an already deployed contract.
func bindFireblock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FireblockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fireblock *FireblockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Fireblock.Contract.FireblockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fireblock *FireblockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fireblock.Contract.FireblockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fireblock *FireblockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fireblock.Contract.FireblockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fireblock *FireblockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Fireblock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fireblock *FireblockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fireblock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fireblock *FireblockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fireblock.Contract.contract.Transact(opts, method, params...)
}

// IsRole is a free data retrieval call binding the contract method 0xcdf22019.
//
// Solidity: function isRole(role address, irole uint16) constant returns(bool)
func (_Fireblock *FireblockCaller) IsRole(opts *bind.CallOpts, role common.Address, irole uint16) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Fireblock.contract.Call(opts, out, "isRole", role, irole)
	return *ret0, err
}

// IsRole is a free data retrieval call binding the contract method 0xcdf22019.
//
// Solidity: function isRole(role address, irole uint16) constant returns(bool)
func (_Fireblock *FireblockSession) IsRole(role common.Address, irole uint16) (bool, error) {
	return _Fireblock.Contract.IsRole(&_Fireblock.CallOpts, role, irole)
}

// IsRole is a free data retrieval call binding the contract method 0xcdf22019.
//
// Solidity: function isRole(role address, irole uint16) constant returns(bool)
func (_Fireblock *FireblockCallerSession) IsRole(role common.Address, irole uint16) (bool, error) {
	return _Fireblock.Contract.IsRole(&_Fireblock.CallOpts, role, irole)
}

// OperationAt is a free data retrieval call binding the contract method 0x8ed2c1b1.
//
// Solidity: function operationAt(index int16) constant returns(otype int16, blockWrite int64, sender address, useruid bytes32, hash bytes32, cardId bytes32, keyuid bytes32, key bytes32, signature bytes32, metadata bytes32, used bool, seals address[16])
func (_Fireblock *FireblockCaller) OperationAt(opts *bind.CallOpts, index int16) (struct {
	Otype      int16
	BlockWrite int64
	Sender     common.Address
	Useruid    [32]byte
	Hash       [32]byte
	CardId     [32]byte
	Keyuid     [32]byte
	Key        [32]byte
	Signature  [32]byte
	Metadata   [32]byte
	Used       bool
	Seals      [16]common.Address
}, error) {
	ret := new(struct {
		Otype      int16
		BlockWrite int64
		Sender     common.Address
		Useruid    [32]byte
		Hash       [32]byte
		CardId     [32]byte
		Keyuid     [32]byte
		Key        [32]byte
		Signature  [32]byte
		Metadata   [32]byte
		Used       bool
		Seals      [16]common.Address
	})
	out := ret
	err := _Fireblock.contract.Call(opts, out, "operationAt", index)
	return *ret, err
}

// OperationAt is a free data retrieval call binding the contract method 0x8ed2c1b1.
//
// Solidity: function operationAt(index int16) constant returns(otype int16, blockWrite int64, sender address, useruid bytes32, hash bytes32, cardId bytes32, keyuid bytes32, key bytes32, signature bytes32, metadata bytes32, used bool, seals address[16])
func (_Fireblock *FireblockSession) OperationAt(index int16) (struct {
	Otype      int16
	BlockWrite int64
	Sender     common.Address
	Useruid    [32]byte
	Hash       [32]byte
	CardId     [32]byte
	Keyuid     [32]byte
	Key        [32]byte
	Signature  [32]byte
	Metadata   [32]byte
	Used       bool
	Seals      [16]common.Address
}, error) {
	return _Fireblock.Contract.OperationAt(&_Fireblock.CallOpts, index)
}

// OperationAt is a free data retrieval call binding the contract method 0x8ed2c1b1.
//
// Solidity: function operationAt(index int16) constant returns(otype int16, blockWrite int64, sender address, useruid bytes32, hash bytes32, cardId bytes32, keyuid bytes32, key bytes32, signature bytes32, metadata bytes32, used bool, seals address[16])
func (_Fireblock *FireblockCallerSession) OperationAt(index int16) (struct {
	Otype      int16
	BlockWrite int64
	Sender     common.Address
	Useruid    [32]byte
	Hash       [32]byte
	CardId     [32]byte
	Keyuid     [32]byte
	Key        [32]byte
	Signature  [32]byte
	Metadata   [32]byte
	Used       bool
	Seals      [16]common.Address
}, error) {
	return _Fireblock.Contract.OperationAt(&_Fireblock.CallOpts, index)
}

// OperationIndexes is a free data retrieval call binding the contract method 0x87d13a1c.
//
// Solidity: function operationIndexes() constant returns(res int16[])
func (_Fireblock *FireblockCaller) OperationIndexes(opts *bind.CallOpts) ([]int16, error) {
	var (
		ret0 = new([]int16)
	)
	out := ret0
	err := _Fireblock.contract.Call(opts, out, "operationIndexes")
	return *ret0, err
}

// OperationIndexes is a free data retrieval call binding the contract method 0x87d13a1c.
//
// Solidity: function operationIndexes() constant returns(res int16[])
func (_Fireblock *FireblockSession) OperationIndexes() ([]int16, error) {
	return _Fireblock.Contract.OperationIndexes(&_Fireblock.CallOpts)
}

// OperationIndexes is a free data retrieval call binding the contract method 0x87d13a1c.
//
// Solidity: function operationIndexes() constant returns(res int16[])
func (_Fireblock *FireblockCallerSession) OperationIndexes() ([]int16, error) {
	return _Fireblock.Contract.OperationIndexes(&_Fireblock.CallOpts)
}

// OperationQty is a free data retrieval call binding the contract method 0x2fc8996a.
//
// Solidity: function operationQty() constant returns(uint256)
func (_Fireblock *FireblockCaller) OperationQty(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fireblock.contract.Call(opts, out, "operationQty")
	return *ret0, err
}

// OperationQty is a free data retrieval call binding the contract method 0x2fc8996a.
//
// Solidity: function operationQty() constant returns(uint256)
func (_Fireblock *FireblockSession) OperationQty() (*big.Int, error) {
	return _Fireblock.Contract.OperationQty(&_Fireblock.CallOpts)
}

// OperationQty is a free data retrieval call binding the contract method 0x2fc8996a.
//
// Solidity: function operationQty() constant returns(uint256)
func (_Fireblock *FireblockCallerSession) OperationQty() (*big.Int, error) {
	return _Fireblock.Contract.OperationQty(&_Fireblock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Fireblock *FireblockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Fireblock.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Fireblock *FireblockSession) Owner() (common.Address, error) {
	return _Fireblock.Contract.Owner(&_Fireblock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Fireblock *FireblockCallerSession) Owner() (common.Address, error) {
	return _Fireblock.Contract.Owner(&_Fireblock.CallOpts)
}

// RoleQty is a free data retrieval call binding the contract method 0x54f97b45.
//
// Solidity: function roleQty(irole uint16) constant returns(uint256)
func (_Fireblock *FireblockCaller) RoleQty(opts *bind.CallOpts, irole uint16) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fireblock.contract.Call(opts, out, "roleQty", irole)
	return *ret0, err
}

// RoleQty is a free data retrieval call binding the contract method 0x54f97b45.
//
// Solidity: function roleQty(irole uint16) constant returns(uint256)
func (_Fireblock *FireblockSession) RoleQty(irole uint16) (*big.Int, error) {
	return _Fireblock.Contract.RoleQty(&_Fireblock.CallOpts, irole)
}

// RoleQty is a free data retrieval call binding the contract method 0x54f97b45.
//
// Solidity: function roleQty(irole uint16) constant returns(uint256)
func (_Fireblock *FireblockCallerSession) RoleQty(irole uint16) (*big.Int, error) {
	return _Fireblock.Contract.RoleQty(&_Fireblock.CallOpts, irole)
}

// Roles is a free data retrieval call binding the contract method 0x31479bad.
//
// Solidity: function roles(irole uint16) constant returns(res address[])
func (_Fireblock *FireblockCaller) Roles(opts *bind.CallOpts, irole uint16) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Fireblock.contract.Call(opts, out, "roles", irole)
	return *ret0, err
}

// Roles is a free data retrieval call binding the contract method 0x31479bad.
//
// Solidity: function roles(irole uint16) constant returns(res address[])
func (_Fireblock *FireblockSession) Roles(irole uint16) ([]common.Address, error) {
	return _Fireblock.Contract.Roles(&_Fireblock.CallOpts, irole)
}

// Roles is a free data retrieval call binding the contract method 0x31479bad.
//
// Solidity: function roles(irole uint16) constant returns(res address[])
func (_Fireblock *FireblockCallerSession) Roles(irole uint16) ([]common.Address, error) {
	return _Fireblock.Contract.Roles(&_Fireblock.CallOpts, irole)
}

// Verify is a free data retrieval call binding the contract method 0x4e8fee00.
//
// Solidity: function verify(hash bytes32, cardId bytes32) constant returns(validity bool, key bytes32, sig bytes32, upgrade bytes32)
func (_Fireblock *FireblockCaller) Verify(opts *bind.CallOpts, hash [32]byte, cardId [32]byte) (struct {
	Validity bool
	Key      [32]byte
	Sig      [32]byte
	Upgrade  [32]byte
}, error) {
	ret := new(struct {
		Validity bool
		Key      [32]byte
		Sig      [32]byte
		Upgrade  [32]byte
	})
	out := ret
	err := _Fireblock.contract.Call(opts, out, "verify", hash, cardId)
	return *ret, err
}

// Verify is a free data retrieval call binding the contract method 0x4e8fee00.
//
// Solidity: function verify(hash bytes32, cardId bytes32) constant returns(validity bool, key bytes32, sig bytes32, upgrade bytes32)
func (_Fireblock *FireblockSession) Verify(hash [32]byte, cardId [32]byte) (struct {
	Validity bool
	Key      [32]byte
	Sig      [32]byte
	Upgrade  [32]byte
}, error) {
	return _Fireblock.Contract.Verify(&_Fireblock.CallOpts, hash, cardId)
}

// Verify is a free data retrieval call binding the contract method 0x4e8fee00.
//
// Solidity: function verify(hash bytes32, cardId bytes32) constant returns(validity bool, key bytes32, sig bytes32, upgrade bytes32)
func (_Fireblock *FireblockCallerSession) Verify(hash [32]byte, cardId [32]byte) (struct {
	Validity bool
	Key      [32]byte
	Sig      [32]byte
	Upgrade  [32]byte
}, error) {
	return _Fireblock.Contract.Verify(&_Fireblock.CallOpts, hash, cardId)
}

// VerifyByUserUID is a free data retrieval call binding the contract method 0x81835871.
//
// Solidity: function verifyByUserUID(hash bytes32, useruid bytes32) constant returns(bytes32[64])
func (_Fireblock *FireblockCaller) VerifyByUserUID(opts *bind.CallOpts, hash [32]byte, useruid [32]byte) ([64][32]byte, error) {
	var (
		ret0 = new([64][32]byte)
	)
	out := ret0
	err := _Fireblock.contract.Call(opts, out, "verifyByUserUID", hash, useruid)
	return *ret0, err
}

// VerifyByUserUID is a free data retrieval call binding the contract method 0x81835871.
//
// Solidity: function verifyByUserUID(hash bytes32, useruid bytes32) constant returns(bytes32[64])
func (_Fireblock *FireblockSession) VerifyByUserUID(hash [32]byte, useruid [32]byte) ([64][32]byte, error) {
	return _Fireblock.Contract.VerifyByUserUID(&_Fireblock.CallOpts, hash, useruid)
}

// VerifyByUserUID is a free data retrieval call binding the contract method 0x81835871.
//
// Solidity: function verifyByUserUID(hash bytes32, useruid bytes32) constant returns(bytes32[64])
func (_Fireblock *FireblockCallerSession) VerifyByUserUID(hash [32]byte, useruid [32]byte) ([64][32]byte, error) {
	return _Fireblock.Contract.VerifyByUserUID(&_Fireblock.CallOpts, hash, useruid)
}

// CloseKey is a paid mutator transaction binding the contract method 0x1de56798.
//
// Solidity: function closeKey(keyuid bytes32, sig string) returns()
func (_Fireblock *FireblockTransactor) CloseKey(opts *bind.TransactOpts, keyuid [32]byte, sig string) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "closeKey", keyuid, sig)
}

// CloseKey is a paid mutator transaction binding the contract method 0x1de56798.
//
// Solidity: function closeKey(keyuid bytes32, sig string) returns()
func (_Fireblock *FireblockSession) CloseKey(keyuid [32]byte, sig string) (*types.Transaction, error) {
	return _Fireblock.Contract.CloseKey(&_Fireblock.TransactOpts, keyuid, sig)
}

// CloseKey is a paid mutator transaction binding the contract method 0x1de56798.
//
// Solidity: function closeKey(keyuid bytes32, sig string) returns()
func (_Fireblock *FireblockTransactorSession) CloseKey(keyuid [32]byte, sig string) (*types.Transaction, error) {
	return _Fireblock.Contract.CloseKey(&_Fireblock.TransactOpts, keyuid, sig)
}

// CloseKeyEth is a paid mutator transaction binding the contract method 0xdc7ef0d9.
//
// Solidity: function closeKeyEth() returns()
func (_Fireblock *FireblockTransactor) CloseKeyEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "closeKeyEth")
}

// CloseKeyEth is a paid mutator transaction binding the contract method 0xdc7ef0d9.
//
// Solidity: function closeKeyEth() returns()
func (_Fireblock *FireblockSession) CloseKeyEth() (*types.Transaction, error) {
	return _Fireblock.Contract.CloseKeyEth(&_Fireblock.TransactOpts)
}

// CloseKeyEth is a paid mutator transaction binding the contract method 0xdc7ef0d9.
//
// Solidity: function closeKeyEth() returns()
func (_Fireblock *FireblockTransactorSession) CloseKeyEth() (*types.Transaction, error) {
	return _Fireblock.Contract.CloseKeyEth(&_Fireblock.TransactOpts)
}

// RegisterAdministrator is a paid mutator transaction binding the contract method 0xa32c706a.
//
// Solidity: function registerAdministrator(admin address) returns()
func (_Fireblock *FireblockTransactor) RegisterAdministrator(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "registerAdministrator", admin)
}

// RegisterAdministrator is a paid mutator transaction binding the contract method 0xa32c706a.
//
// Solidity: function registerAdministrator(admin address) returns()
func (_Fireblock *FireblockSession) RegisterAdministrator(admin common.Address) (*types.Transaction, error) {
	return _Fireblock.Contract.RegisterAdministrator(&_Fireblock.TransactOpts, admin)
}

// RegisterAdministrator is a paid mutator transaction binding the contract method 0xa32c706a.
//
// Solidity: function registerAdministrator(admin address) returns()
func (_Fireblock *FireblockTransactorSession) RegisterAdministrator(admin common.Address) (*types.Transaction, error) {
	return _Fireblock.Contract.RegisterAdministrator(&_Fireblock.TransactOpts, admin)
}

// RegisterCardEth is a paid mutator transaction binding the contract method 0xde57d0e0.
//
// Solidity: function registerCardEth(useruid bytes32, owner address, card string) returns()
func (_Fireblock *FireblockTransactor) RegisterCardEth(opts *bind.TransactOpts, useruid [32]byte, owner common.Address, card string) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "registerCardEth", useruid, owner, card)
}

// RegisterCardEth is a paid mutator transaction binding the contract method 0xde57d0e0.
//
// Solidity: function registerCardEth(useruid bytes32, owner address, card string) returns()
func (_Fireblock *FireblockSession) RegisterCardEth(useruid [32]byte, owner common.Address, card string) (*types.Transaction, error) {
	return _Fireblock.Contract.RegisterCardEth(&_Fireblock.TransactOpts, useruid, owner, card)
}

// RegisterCardEth is a paid mutator transaction binding the contract method 0xde57d0e0.
//
// Solidity: function registerCardEth(useruid bytes32, owner address, card string) returns()
func (_Fireblock *FireblockTransactorSession) RegisterCardEth(useruid [32]byte, owner common.Address, card string) (*types.Transaction, error) {
	return _Fireblock.Contract.RegisterCardEth(&_Fireblock.TransactOpts, useruid, owner, card)
}

// RegisterKey is a paid mutator transaction binding the contract method 0x2ced3b4b.
//
// Solidity: function registerKey(keyuid bytes32, key string, sig string) returns()
func (_Fireblock *FireblockTransactor) RegisterKey(opts *bind.TransactOpts, keyuid [32]byte, key string, sig string) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "registerKey", keyuid, key, sig)
}

// RegisterKey is a paid mutator transaction binding the contract method 0x2ced3b4b.
//
// Solidity: function registerKey(keyuid bytes32, key string, sig string) returns()
func (_Fireblock *FireblockSession) RegisterKey(keyuid [32]byte, key string, sig string) (*types.Transaction, error) {
	return _Fireblock.Contract.RegisterKey(&_Fireblock.TransactOpts, keyuid, key, sig)
}

// RegisterKey is a paid mutator transaction binding the contract method 0x2ced3b4b.
//
// Solidity: function registerKey(keyuid bytes32, key string, sig string) returns()
func (_Fireblock *FireblockTransactorSession) RegisterKey(keyuid [32]byte, key string, sig string) (*types.Transaction, error) {
	return _Fireblock.Contract.RegisterKey(&_Fireblock.TransactOpts, keyuid, key, sig)
}

// RegisterMaker is a paid mutator transaction binding the contract method 0xbaf167fa.
//
// Solidity: function registerMaker(maker address) returns()
func (_Fireblock *FireblockTransactor) RegisterMaker(opts *bind.TransactOpts, maker common.Address) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "registerMaker", maker)
}

// RegisterMaker is a paid mutator transaction binding the contract method 0xbaf167fa.
//
// Solidity: function registerMaker(maker address) returns()
func (_Fireblock *FireblockSession) RegisterMaker(maker common.Address) (*types.Transaction, error) {
	return _Fireblock.Contract.RegisterMaker(&_Fireblock.TransactOpts, maker)
}

// RegisterMaker is a paid mutator transaction binding the contract method 0xbaf167fa.
//
// Solidity: function registerMaker(maker address) returns()
func (_Fireblock *FireblockTransactorSession) RegisterMaker(maker common.Address) (*types.Transaction, error) {
	return _Fireblock.Contract.RegisterMaker(&_Fireblock.TransactOpts, maker)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x9fca5169.
//
// Solidity: function registerValidator(validator address) returns()
func (_Fireblock *FireblockTransactor) RegisterValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "registerValidator", validator)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x9fca5169.
//
// Solidity: function registerValidator(validator address) returns()
func (_Fireblock *FireblockSession) RegisterValidator(validator common.Address) (*types.Transaction, error) {
	return _Fireblock.Contract.RegisterValidator(&_Fireblock.TransactOpts, validator)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x9fca5169.
//
// Solidity: function registerValidator(validator address) returns()
func (_Fireblock *FireblockTransactorSession) RegisterValidator(validator common.Address) (*types.Transaction, error) {
	return _Fireblock.Contract.RegisterValidator(&_Fireblock.TransactOpts, validator)
}

// RegisterWriter is a paid mutator transaction binding the contract method 0xff48b017.
//
// Solidity: function registerWriter(writer address) returns()
func (_Fireblock *FireblockTransactor) RegisterWriter(opts *bind.TransactOpts, writer common.Address) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "registerWriter", writer)
}

// RegisterWriter is a paid mutator transaction binding the contract method 0xff48b017.
//
// Solidity: function registerWriter(writer address) returns()
func (_Fireblock *FireblockSession) RegisterWriter(writer common.Address) (*types.Transaction, error) {
	return _Fireblock.Contract.RegisterWriter(&_Fireblock.TransactOpts, writer)
}

// RegisterWriter is a paid mutator transaction binding the contract method 0xff48b017.
//
// Solidity: function registerWriter(writer address) returns()
func (_Fireblock *FireblockTransactorSession) RegisterWriter(writer common.Address) (*types.Transaction, error) {
	return _Fireblock.Contract.RegisterWriter(&_Fireblock.TransactOpts, writer)
}

// ResetUpgradeCardAdm is a paid mutator transaction binding the contract method 0x800e67f3.
//
// Solidity: function resetUpgradeCardAdm(cardId bytes32) returns()
func (_Fireblock *FireblockTransactor) ResetUpgradeCardAdm(opts *bind.TransactOpts, cardId [32]byte) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "resetUpgradeCardAdm", cardId)
}

// ResetUpgradeCardAdm is a paid mutator transaction binding the contract method 0x800e67f3.
//
// Solidity: function resetUpgradeCardAdm(cardId bytes32) returns()
func (_Fireblock *FireblockSession) ResetUpgradeCardAdm(cardId [32]byte) (*types.Transaction, error) {
	return _Fireblock.Contract.ResetUpgradeCardAdm(&_Fireblock.TransactOpts, cardId)
}

// ResetUpgradeCardAdm is a paid mutator transaction binding the contract method 0x800e67f3.
//
// Solidity: function resetUpgradeCardAdm(cardId bytes32) returns()
func (_Fireblock *FireblockTransactorSession) ResetUpgradeCardAdm(cardId [32]byte) (*types.Transaction, error) {
	return _Fireblock.Contract.ResetUpgradeCardAdm(&_Fireblock.TransactOpts, cardId)
}

// RevokeAdministrator is a paid mutator transaction binding the contract method 0x4402bba3.
//
// Solidity: function revokeAdministrator(admin address) returns(bool)
func (_Fireblock *FireblockTransactor) RevokeAdministrator(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "revokeAdministrator", admin)
}

// RevokeAdministrator is a paid mutator transaction binding the contract method 0x4402bba3.
//
// Solidity: function revokeAdministrator(admin address) returns(bool)
func (_Fireblock *FireblockSession) RevokeAdministrator(admin common.Address) (*types.Transaction, error) {
	return _Fireblock.Contract.RevokeAdministrator(&_Fireblock.TransactOpts, admin)
}

// RevokeAdministrator is a paid mutator transaction binding the contract method 0x4402bba3.
//
// Solidity: function revokeAdministrator(admin address) returns(bool)
func (_Fireblock *FireblockTransactorSession) RevokeAdministrator(admin common.Address) (*types.Transaction, error) {
	return _Fireblock.Contract.RevokeAdministrator(&_Fireblock.TransactOpts, admin)
}

// RevokeKey is a paid mutator transaction binding the contract method 0xfb55043e.
//
// Solidity: function revokeKey(keyuid bytes32, sig string) returns()
func (_Fireblock *FireblockTransactor) RevokeKey(opts *bind.TransactOpts, keyuid [32]byte, sig string) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "revokeKey", keyuid, sig)
}

// RevokeKey is a paid mutator transaction binding the contract method 0xfb55043e.
//
// Solidity: function revokeKey(keyuid bytes32, sig string) returns()
func (_Fireblock *FireblockSession) RevokeKey(keyuid [32]byte, sig string) (*types.Transaction, error) {
	return _Fireblock.Contract.RevokeKey(&_Fireblock.TransactOpts, keyuid, sig)
}

// RevokeKey is a paid mutator transaction binding the contract method 0xfb55043e.
//
// Solidity: function revokeKey(keyuid bytes32, sig string) returns()
func (_Fireblock *FireblockTransactorSession) RevokeKey(keyuid [32]byte, sig string) (*types.Transaction, error) {
	return _Fireblock.Contract.RevokeKey(&_Fireblock.TransactOpts, keyuid, sig)
}

// RevokeKeyAdm is a paid mutator transaction binding the contract method 0xd91f4d35.
//
// Solidity: function revokeKeyAdm(key bytes32) returns()
func (_Fireblock *FireblockTransactor) RevokeKeyAdm(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "revokeKeyAdm", key)
}

// RevokeKeyAdm is a paid mutator transaction binding the contract method 0xd91f4d35.
//
// Solidity: function revokeKeyAdm(key bytes32) returns()
func (_Fireblock *FireblockSession) RevokeKeyAdm(key [32]byte) (*types.Transaction, error) {
	return _Fireblock.Contract.RevokeKeyAdm(&_Fireblock.TransactOpts, key)
}

// RevokeKeyAdm is a paid mutator transaction binding the contract method 0xd91f4d35.
//
// Solidity: function revokeKeyAdm(key bytes32) returns()
func (_Fireblock *FireblockTransactorSession) RevokeKeyAdm(key [32]byte) (*types.Transaction, error) {
	return _Fireblock.Contract.RevokeKeyAdm(&_Fireblock.TransactOpts, key)
}

// RevokeKeyEth is a paid mutator transaction binding the contract method 0xd421dafd.
//
// Solidity: function revokeKeyEth() returns()
func (_Fireblock *FireblockTransactor) RevokeKeyEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "revokeKeyEth")
}

// RevokeKeyEth is a paid mutator transaction binding the contract method 0xd421dafd.
//
// Solidity: function revokeKeyEth() returns()
func (_Fireblock *FireblockSession) RevokeKeyEth() (*types.Transaction, error) {
	return _Fireblock.Contract.RevokeKeyEth(&_Fireblock.TransactOpts)
}

// RevokeKeyEth is a paid mutator transaction binding the contract method 0xd421dafd.
//
// Solidity: function revokeKeyEth() returns()
func (_Fireblock *FireblockTransactorSession) RevokeKeyEth() (*types.Transaction, error) {
	return _Fireblock.Contract.RevokeKeyEth(&_Fireblock.TransactOpts)
}

// RevokeMaker is a paid mutator transaction binding the contract method 0x17f094f9.
//
// Solidity: function revokeMaker(maker address) returns(bool)
func (_Fireblock *FireblockTransactor) RevokeMaker(opts *bind.TransactOpts, maker common.Address) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "revokeMaker", maker)
}

// RevokeMaker is a paid mutator transaction binding the contract method 0x17f094f9.
//
// Solidity: function revokeMaker(maker address) returns(bool)
func (_Fireblock *FireblockSession) RevokeMaker(maker common.Address) (*types.Transaction, error) {
	return _Fireblock.Contract.RevokeMaker(&_Fireblock.TransactOpts, maker)
}

// RevokeMaker is a paid mutator transaction binding the contract method 0x17f094f9.
//
// Solidity: function revokeMaker(maker address) returns(bool)
func (_Fireblock *FireblockTransactorSession) RevokeMaker(maker common.Address) (*types.Transaction, error) {
	return _Fireblock.Contract.RevokeMaker(&_Fireblock.TransactOpts, maker)
}

// RevokeProof is a paid mutator transaction binding the contract method 0xf3111503.
//
// Solidity: function revokeProof(hash bytes32, cardId bytes32, keyuid bytes32, sig string) returns()
func (_Fireblock *FireblockTransactor) RevokeProof(opts *bind.TransactOpts, hash [32]byte, cardId [32]byte, keyuid [32]byte, sig string) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "revokeProof", hash, cardId, keyuid, sig)
}

// RevokeProof is a paid mutator transaction binding the contract method 0xf3111503.
//
// Solidity: function revokeProof(hash bytes32, cardId bytes32, keyuid bytes32, sig string) returns()
func (_Fireblock *FireblockSession) RevokeProof(hash [32]byte, cardId [32]byte, keyuid [32]byte, sig string) (*types.Transaction, error) {
	return _Fireblock.Contract.RevokeProof(&_Fireblock.TransactOpts, hash, cardId, keyuid, sig)
}

// RevokeProof is a paid mutator transaction binding the contract method 0xf3111503.
//
// Solidity: function revokeProof(hash bytes32, cardId bytes32, keyuid bytes32, sig string) returns()
func (_Fireblock *FireblockTransactorSession) RevokeProof(hash [32]byte, cardId [32]byte, keyuid [32]byte, sig string) (*types.Transaction, error) {
	return _Fireblock.Contract.RevokeProof(&_Fireblock.TransactOpts, hash, cardId, keyuid, sig)
}

// RevokeProofAdm is a paid mutator transaction binding the contract method 0x8c737ceb.
//
// Solidity: function revokeProofAdm(hash bytes32, cardId bytes32) returns()
func (_Fireblock *FireblockTransactor) RevokeProofAdm(opts *bind.TransactOpts, hash [32]byte, cardId [32]byte) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "revokeProofAdm", hash, cardId)
}

// RevokeProofAdm is a paid mutator transaction binding the contract method 0x8c737ceb.
//
// Solidity: function revokeProofAdm(hash bytes32, cardId bytes32) returns()
func (_Fireblock *FireblockSession) RevokeProofAdm(hash [32]byte, cardId [32]byte) (*types.Transaction, error) {
	return _Fireblock.Contract.RevokeProofAdm(&_Fireblock.TransactOpts, hash, cardId)
}

// RevokeProofAdm is a paid mutator transaction binding the contract method 0x8c737ceb.
//
// Solidity: function revokeProofAdm(hash bytes32, cardId bytes32) returns()
func (_Fireblock *FireblockTransactorSession) RevokeProofAdm(hash [32]byte, cardId [32]byte) (*types.Transaction, error) {
	return _Fireblock.Contract.RevokeProofAdm(&_Fireblock.TransactOpts, hash, cardId)
}

// RevokeProofEth is a paid mutator transaction binding the contract method 0xf2dd7f32.
//
// Solidity: function revokeProofEth(hash bytes32) returns()
func (_Fireblock *FireblockTransactor) RevokeProofEth(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "revokeProofEth", hash)
}

// RevokeProofEth is a paid mutator transaction binding the contract method 0xf2dd7f32.
//
// Solidity: function revokeProofEth(hash bytes32) returns()
func (_Fireblock *FireblockSession) RevokeProofEth(hash [32]byte) (*types.Transaction, error) {
	return _Fireblock.Contract.RevokeProofEth(&_Fireblock.TransactOpts, hash)
}

// RevokeProofEth is a paid mutator transaction binding the contract method 0xf2dd7f32.
//
// Solidity: function revokeProofEth(hash bytes32) returns()
func (_Fireblock *FireblockTransactorSession) RevokeProofEth(hash [32]byte) (*types.Transaction, error) {
	return _Fireblock.Contract.RevokeProofEth(&_Fireblock.TransactOpts, hash)
}

// RevokeValidator is a paid mutator transaction binding the contract method 0x047564b7.
//
// Solidity: function revokeValidator(validator address) returns(bool)
func (_Fireblock *FireblockTransactor) RevokeValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "revokeValidator", validator)
}

// RevokeValidator is a paid mutator transaction binding the contract method 0x047564b7.
//
// Solidity: function revokeValidator(validator address) returns(bool)
func (_Fireblock *FireblockSession) RevokeValidator(validator common.Address) (*types.Transaction, error) {
	return _Fireblock.Contract.RevokeValidator(&_Fireblock.TransactOpts, validator)
}

// RevokeValidator is a paid mutator transaction binding the contract method 0x047564b7.
//
// Solidity: function revokeValidator(validator address) returns(bool)
func (_Fireblock *FireblockTransactorSession) RevokeValidator(validator common.Address) (*types.Transaction, error) {
	return _Fireblock.Contract.RevokeValidator(&_Fireblock.TransactOpts, validator)
}

// RevokeWriter is a paid mutator transaction binding the contract method 0x06a10eb3.
//
// Solidity: function revokeWriter(writer address) returns(bool)
func (_Fireblock *FireblockTransactor) RevokeWriter(opts *bind.TransactOpts, writer common.Address) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "revokeWriter", writer)
}

// RevokeWriter is a paid mutator transaction binding the contract method 0x06a10eb3.
//
// Solidity: function revokeWriter(writer address) returns(bool)
func (_Fireblock *FireblockSession) RevokeWriter(writer common.Address) (*types.Transaction, error) {
	return _Fireblock.Contract.RevokeWriter(&_Fireblock.TransactOpts, writer)
}

// RevokeWriter is a paid mutator transaction binding the contract method 0x06a10eb3.
//
// Solidity: function revokeWriter(writer address) returns(bool)
func (_Fireblock *FireblockTransactorSession) RevokeWriter(writer common.Address) (*types.Transaction, error) {
	return _Fireblock.Contract.RevokeWriter(&_Fireblock.TransactOpts, writer)
}

// SetDataStore is a paid mutator transaction binding the contract method 0x500065d2.
//
// Solidity: function setDataStore(store_ address) returns()
func (_Fireblock *FireblockTransactor) SetDataStore(opts *bind.TransactOpts, store_ common.Address) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "setDataStore", store_)
}

// SetDataStore is a paid mutator transaction binding the contract method 0x500065d2.
//
// Solidity: function setDataStore(store_ address) returns()
func (_Fireblock *FireblockSession) SetDataStore(store_ common.Address) (*types.Transaction, error) {
	return _Fireblock.Contract.SetDataStore(&_Fireblock.TransactOpts, store_)
}

// SetDataStore is a paid mutator transaction binding the contract method 0x500065d2.
//
// Solidity: function setDataStore(store_ address) returns()
func (_Fireblock *FireblockTransactorSession) SetDataStore(store_ common.Address) (*types.Transaction, error) {
	return _Fireblock.Contract.SetDataStore(&_Fireblock.TransactOpts, store_)
}

// Sign is a paid mutator transaction binding the contract method 0x647fd1a5.
//
// Solidity: function sign(hash bytes32, useruid bytes32, card string, keyuid bytes32, sig string, metadata string) returns()
func (_Fireblock *FireblockTransactor) Sign(opts *bind.TransactOpts, hash [32]byte, useruid [32]byte, card string, keyuid [32]byte, sig string, metadata string) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "sign", hash, useruid, card, keyuid, sig, metadata)
}

// Sign is a paid mutator transaction binding the contract method 0x647fd1a5.
//
// Solidity: function sign(hash bytes32, useruid bytes32, card string, keyuid bytes32, sig string, metadata string) returns()
func (_Fireblock *FireblockSession) Sign(hash [32]byte, useruid [32]byte, card string, keyuid [32]byte, sig string, metadata string) (*types.Transaction, error) {
	return _Fireblock.Contract.Sign(&_Fireblock.TransactOpts, hash, useruid, card, keyuid, sig, metadata)
}

// Sign is a paid mutator transaction binding the contract method 0x647fd1a5.
//
// Solidity: function sign(hash bytes32, useruid bytes32, card string, keyuid bytes32, sig string, metadata string) returns()
func (_Fireblock *FireblockTransactorSession) Sign(hash [32]byte, useruid [32]byte, card string, keyuid [32]byte, sig string, metadata string) (*types.Transaction, error) {
	return _Fireblock.Contract.Sign(&_Fireblock.TransactOpts, hash, useruid, card, keyuid, sig, metadata)
}

// SignEth is a paid mutator transaction binding the contract method 0xb0d1d6c7.
//
// Solidity: function signEth(hash bytes32, metadata string) returns()
func (_Fireblock *FireblockTransactor) SignEth(opts *bind.TransactOpts, hash [32]byte, metadata string) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "signEth", hash, metadata)
}

// SignEth is a paid mutator transaction binding the contract method 0xb0d1d6c7.
//
// Solidity: function signEth(hash bytes32, metadata string) returns()
func (_Fireblock *FireblockSession) SignEth(hash [32]byte, metadata string) (*types.Transaction, error) {
	return _Fireblock.Contract.SignEth(&_Fireblock.TransactOpts, hash, metadata)
}

// SignEth is a paid mutator transaction binding the contract method 0xb0d1d6c7.
//
// Solidity: function signEth(hash bytes32, metadata string) returns()
func (_Fireblock *FireblockTransactorSession) SignEth(hash [32]byte, metadata string) (*types.Transaction, error) {
	return _Fireblock.Contract.SignEth(&_Fireblock.TransactOpts, hash, metadata)
}

// Tick is a paid mutator transaction binding the contract method 0x3eaf5d9f.
//
// Solidity: function tick() returns()
func (_Fireblock *FireblockTransactor) Tick(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "tick")
}

// Tick is a paid mutator transaction binding the contract method 0x3eaf5d9f.
//
// Solidity: function tick() returns()
func (_Fireblock *FireblockSession) Tick() (*types.Transaction, error) {
	return _Fireblock.Contract.Tick(&_Fireblock.TransactOpts)
}

// Tick is a paid mutator transaction binding the contract method 0x3eaf5d9f.
//
// Solidity: function tick() returns()
func (_Fireblock *FireblockTransactorSession) Tick() (*types.Transaction, error) {
	return _Fireblock.Contract.Tick(&_Fireblock.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Fireblock *FireblockTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Fireblock *FireblockSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Fireblock.Contract.TransferOwnership(&_Fireblock.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Fireblock *FireblockTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Fireblock.Contract.TransferOwnership(&_Fireblock.TransactOpts, newOwner)
}

// UpgradeCard is a paid mutator transaction binding the contract method 0x841d5385.
//
// Solidity: function upgradeCard(oldCardId bytes32, newCardId bytes32, sig string) returns()
func (_Fireblock *FireblockTransactor) UpgradeCard(opts *bind.TransactOpts, oldCardId [32]byte, newCardId [32]byte, sig string) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "upgradeCard", oldCardId, newCardId, sig)
}

// UpgradeCard is a paid mutator transaction binding the contract method 0x841d5385.
//
// Solidity: function upgradeCard(oldCardId bytes32, newCardId bytes32, sig string) returns()
func (_Fireblock *FireblockSession) UpgradeCard(oldCardId [32]byte, newCardId [32]byte, sig string) (*types.Transaction, error) {
	return _Fireblock.Contract.UpgradeCard(&_Fireblock.TransactOpts, oldCardId, newCardId, sig)
}

// UpgradeCard is a paid mutator transaction binding the contract method 0x841d5385.
//
// Solidity: function upgradeCard(oldCardId bytes32, newCardId bytes32, sig string) returns()
func (_Fireblock *FireblockTransactorSession) UpgradeCard(oldCardId [32]byte, newCardId [32]byte, sig string) (*types.Transaction, error) {
	return _Fireblock.Contract.UpgradeCard(&_Fireblock.TransactOpts, oldCardId, newCardId, sig)
}

// UpgradeCardEth is a paid mutator transaction binding the contract method 0x8aea63ee.
//
// Solidity: function upgradeCardEth(oldCardId bytes32, newCardId bytes32) returns()
func (_Fireblock *FireblockTransactor) UpgradeCardEth(opts *bind.TransactOpts, oldCardId [32]byte, newCardId [32]byte) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "upgradeCardEth", oldCardId, newCardId)
}

// UpgradeCardEth is a paid mutator transaction binding the contract method 0x8aea63ee.
//
// Solidity: function upgradeCardEth(oldCardId bytes32, newCardId bytes32) returns()
func (_Fireblock *FireblockSession) UpgradeCardEth(oldCardId [32]byte, newCardId [32]byte) (*types.Transaction, error) {
	return _Fireblock.Contract.UpgradeCardEth(&_Fireblock.TransactOpts, oldCardId, newCardId)
}

// UpgradeCardEth is a paid mutator transaction binding the contract method 0x8aea63ee.
//
// Solidity: function upgradeCardEth(oldCardId bytes32, newCardId bytes32) returns()
func (_Fireblock *FireblockTransactorSession) UpgradeCardEth(oldCardId [32]byte, newCardId [32]byte) (*types.Transaction, error) {
	return _Fireblock.Contract.UpgradeCardEth(&_Fireblock.TransactOpts, oldCardId, newCardId)
}

// Validate is a paid mutator transaction binding the contract method 0x1a25d7aa.
//
// Solidity: function validate(indexes int16[]) returns()
func (_Fireblock *FireblockTransactor) Validate(opts *bind.TransactOpts, indexes []int16) (*types.Transaction, error) {
	return _Fireblock.contract.Transact(opts, "validate", indexes)
}

// Validate is a paid mutator transaction binding the contract method 0x1a25d7aa.
//
// Solidity: function validate(indexes int16[]) returns()
func (_Fireblock *FireblockSession) Validate(indexes []int16) (*types.Transaction, error) {
	return _Fireblock.Contract.Validate(&_Fireblock.TransactOpts, indexes)
}

// Validate is a paid mutator transaction binding the contract method 0x1a25d7aa.
//
// Solidity: function validate(indexes int16[]) returns()
func (_Fireblock *FireblockTransactorSession) Validate(indexes []int16) (*types.Transaction, error) {
	return _Fireblock.Contract.Validate(&_Fireblock.TransactOpts, indexes)
}
