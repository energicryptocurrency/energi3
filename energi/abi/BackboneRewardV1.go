// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// BackboneRewardV1ABI is the input ABI used to generate the binding from.
const BackboneRewardV1ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getReward\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reward\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_oldImpl\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// BackboneRewardV1Bin is the compiled bytecode used for deploying new contracts.
const BackboneRewardV1Bin = `608060405234801561001057600080fd5b5060405161043e38038061043e8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556103d9806100656000396000f3fe6080604052600436106100595760003560e01c8063228cb73311610043578063228cb7331461013e578063ce5494bb14610146578063ec5568891461018657610059565b8062f55d9d146100c05780631c4b774b14610102575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b3480156100cc57600080fd5b50610100600480360360208110156100e357600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101c4565b005b34801561010e57600080fd5b5061012c6004803603602081101561012557600080fd5b503561026c565b60408051918252519081900360200190f35b610100610284565b34801561015257600080fd5b506101006004803603602081101561016957600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166102fa565b34801561019257600080fd5b5061019b610388565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff16331461024a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b61025381610385565b8073ffffffffffffffffffffffffffffffffffffffff16ff5b6000811561027f5750671fa42feb87e400005b919050565b600154156102f357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6000600155565b60005473ffffffffffffffffffffffffffffffffffffffff16331461038057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b610385815b50565b60005473ffffffffffffffffffffffffffffffffffffffff168156fea265627a7a72305820bb09a6392ea8ac8f85bb713bc0de5dc5b6b1c0bd6c631334e1c9583dffd43a4664736f6c63430005090032`

// DeployBackboneRewardV1 deploys a new Ethereum contract, binding an instance of BackboneRewardV1 to it.
func DeployBackboneRewardV1(auth *bind.TransactOpts, backend bind.ContractBackend, _proxy common.Address) (common.Address, *types.Transaction, *BackboneRewardV1, error) {
	parsed, err := abi.JSON(strings.NewReader(BackboneRewardV1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BackboneRewardV1Bin), backend, _proxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BackboneRewardV1{BackboneRewardV1Caller: BackboneRewardV1Caller{contract: contract}, BackboneRewardV1Transactor: BackboneRewardV1Transactor{contract: contract}, BackboneRewardV1Filterer: BackboneRewardV1Filterer{contract: contract}}, nil
}

// BackboneRewardV1 is an auto generated Go binding around an Ethereum contract.
type BackboneRewardV1 struct {
	BackboneRewardV1Caller     // Read-only binding to the contract
	BackboneRewardV1Transactor // Write-only binding to the contract
	BackboneRewardV1Filterer   // Log filterer for contract events
}

// BackboneRewardV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type BackboneRewardV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BackboneRewardV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BackboneRewardV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BackboneRewardV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BackboneRewardV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BackboneRewardV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BackboneRewardV1Session struct {
	Contract     *BackboneRewardV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BackboneRewardV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BackboneRewardV1CallerSession struct {
	Contract *BackboneRewardV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BackboneRewardV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BackboneRewardV1TransactorSession struct {
	Contract     *BackboneRewardV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BackboneRewardV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type BackboneRewardV1Raw struct {
	Contract *BackboneRewardV1 // Generic contract binding to access the raw methods on
}

// BackboneRewardV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BackboneRewardV1CallerRaw struct {
	Contract *BackboneRewardV1Caller // Generic read-only contract binding to access the raw methods on
}

// BackboneRewardV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BackboneRewardV1TransactorRaw struct {
	Contract *BackboneRewardV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBackboneRewardV1 creates a new instance of BackboneRewardV1, bound to a specific deployed contract.
func NewBackboneRewardV1(address common.Address, backend bind.ContractBackend) (*BackboneRewardV1, error) {
	contract, err := bindBackboneRewardV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BackboneRewardV1{BackboneRewardV1Caller: BackboneRewardV1Caller{contract: contract}, BackboneRewardV1Transactor: BackboneRewardV1Transactor{contract: contract}, BackboneRewardV1Filterer: BackboneRewardV1Filterer{contract: contract}}, nil
}

// NewBackboneRewardV1Caller creates a new read-only instance of BackboneRewardV1, bound to a specific deployed contract.
func NewBackboneRewardV1Caller(address common.Address, caller bind.ContractCaller) (*BackboneRewardV1Caller, error) {
	contract, err := bindBackboneRewardV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BackboneRewardV1Caller{contract: contract}, nil
}

// NewBackboneRewardV1Transactor creates a new write-only instance of BackboneRewardV1, bound to a specific deployed contract.
func NewBackboneRewardV1Transactor(address common.Address, transactor bind.ContractTransactor) (*BackboneRewardV1Transactor, error) {
	contract, err := bindBackboneRewardV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BackboneRewardV1Transactor{contract: contract}, nil
}

// NewBackboneRewardV1Filterer creates a new log filterer instance of BackboneRewardV1, bound to a specific deployed contract.
func NewBackboneRewardV1Filterer(address common.Address, filterer bind.ContractFilterer) (*BackboneRewardV1Filterer, error) {
	contract, err := bindBackboneRewardV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BackboneRewardV1Filterer{contract: contract}, nil
}

// bindBackboneRewardV1 binds a generic wrapper to an already deployed contract.
func bindBackboneRewardV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BackboneRewardV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BackboneRewardV1 *BackboneRewardV1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BackboneRewardV1.Contract.BackboneRewardV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BackboneRewardV1 *BackboneRewardV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BackboneRewardV1.Contract.BackboneRewardV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BackboneRewardV1 *BackboneRewardV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BackboneRewardV1.Contract.BackboneRewardV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BackboneRewardV1 *BackboneRewardV1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BackboneRewardV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BackboneRewardV1 *BackboneRewardV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BackboneRewardV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BackboneRewardV1 *BackboneRewardV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BackboneRewardV1.Contract.contract.Transact(opts, method, params...)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_BackboneRewardV1 *BackboneRewardV1Caller) GetReward(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BackboneRewardV1.contract.Call(opts, out, "getReward", _blockNumber)
	return *ret0, err
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_BackboneRewardV1 *BackboneRewardV1Session) GetReward(_blockNumber *big.Int) (*big.Int, error) {
	return _BackboneRewardV1.Contract.GetReward(&_BackboneRewardV1.CallOpts, _blockNumber)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_BackboneRewardV1 *BackboneRewardV1CallerSession) GetReward(_blockNumber *big.Int) (*big.Int, error) {
	return _BackboneRewardV1.Contract.GetReward(&_BackboneRewardV1.CallOpts, _blockNumber)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_BackboneRewardV1 *BackboneRewardV1Caller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BackboneRewardV1.contract.Call(opts, out, "proxy")
	return *ret0, err
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_BackboneRewardV1 *BackboneRewardV1Session) Proxy() (common.Address, error) {
	return _BackboneRewardV1.Contract.Proxy(&_BackboneRewardV1.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_BackboneRewardV1 *BackboneRewardV1CallerSession) Proxy() (common.Address, error) {
	return _BackboneRewardV1.Contract.Proxy(&_BackboneRewardV1.CallOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_BackboneRewardV1 *BackboneRewardV1Transactor) Destroy(opts *bind.TransactOpts, _newImpl common.Address) (*types.Transaction, error) {
	return _BackboneRewardV1.contract.Transact(opts, "destroy", _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_BackboneRewardV1 *BackboneRewardV1Session) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _BackboneRewardV1.Contract.Destroy(&_BackboneRewardV1.TransactOpts, _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_BackboneRewardV1 *BackboneRewardV1TransactorSession) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _BackboneRewardV1.Contract.Destroy(&_BackboneRewardV1.TransactOpts, _newImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_BackboneRewardV1 *BackboneRewardV1Transactor) Migrate(opts *bind.TransactOpts, _oldImpl common.Address) (*types.Transaction, error) {
	return _BackboneRewardV1.contract.Transact(opts, "migrate", _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_BackboneRewardV1 *BackboneRewardV1Session) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _BackboneRewardV1.Contract.Migrate(&_BackboneRewardV1.TransactOpts, _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_BackboneRewardV1 *BackboneRewardV1TransactorSession) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _BackboneRewardV1.Contract.Migrate(&_BackboneRewardV1.TransactOpts, _oldImpl)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_BackboneRewardV1 *BackboneRewardV1Transactor) Reward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BackboneRewardV1.contract.Transact(opts, "reward")
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_BackboneRewardV1 *BackboneRewardV1Session) Reward() (*types.Transaction, error) {
	return _BackboneRewardV1.Contract.Reward(&_BackboneRewardV1.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_BackboneRewardV1 *BackboneRewardV1TransactorSession) Reward() (*types.Transaction, error) {
	return _BackboneRewardV1.Contract.Reward(&_BackboneRewardV1.TransactOpts)
}