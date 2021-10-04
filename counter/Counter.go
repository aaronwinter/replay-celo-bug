// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package counter

import (
	"math/big"
	"strings"

	ethereum "github.com/aaronwinter/celo-blockchain"
	"github.com/aaronwinter/celo-blockchain/accounts/abi"
	bind "github.com/aaronwinter/celo-blockchain/accounts/abi/bind_v2"
	"github.com/aaronwinter/celo-blockchain/common"
	"github.com/aaronwinter/celo-blockchain/core/types"
	"github.com/aaronwinter/celo-blockchain/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	// _ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CounterABI is the input ABI used to generate the binding from.
const CounterABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"init\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Counter is an auto generated Go binding around an Ethereum contract.
type Counter struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterSession struct {
	Contract     *Counter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterRaw struct {
	Contract *Counter // Generic contract binding to access the raw methods on
}

// NewCounter creates a new instance of Counter, bound to a specific deployed contract.
func NewCounter(address common.Address, backend bind.ContractBackend) (*Counter, error) {
	contract, err := bindCounter(address, backend)
	if err != nil {
		return nil, err
	}
	return &Counter{contract: contract}, nil
}

// bindCounter binds a generic wrapper to an already deployed contract.
func bindCounter(address common.Address, contract bind.ContractBackend) (*bind.BoundContract, error) {
	parsed, err := ParseCounterABI()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, contract), nil
}

// ParseCounterABI parses the ABI
func ParseCounterABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(CounterABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transact(opts, method, params...)
}

// TxObj returns an obj that can be used to get the transaction, send it, estimate it.
func (_Counter *CounterRaw) TxObj(opts *bind.TransactOpts, method string, params ...interface{}) *bind.TxObject {
	return _Counter.Contract.contract.TxObj(opts, method, params...)
}

// EstimateGas obtains an estimate for calling contract method with params as input values.
func (_Counter *CounterRaw) EstimateGas(opts *bind.TransactOpts, method string, params ...interface{}) (uint64, error) {
	return _Counter.Contract.contract.EstimateGas(opts, method, params...)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_Counter *Counter) Value(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Counter.contract.Call(opts, out, "value")
	return *ret0, err
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_Counter *CounterSession) Value() (*big.Int, error) {
	return _Counter.Contract.Value(&_Counter.CallOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns(bool)
func (_Counter *Counter) Increment(opts *bind.TransactOpts) *bind.TxObject {
	return _Counter.contract.TxObj(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns(bool)
func (_Counter *CounterSession) Increment() *bind.TxObject {
	return _Counter.Contract.Increment(&_Counter.TransactOpts)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_Counter *Counter) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _Counter.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}
