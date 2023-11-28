// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package merchants

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MerchantRegistryMetaData contains all meta data concerning the MerchantRegistry contract.
var MerchantRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_merchantIDAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"merchantAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"merchantId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"arweaveID\",\"type\":\"string\"}],\"name\":\"NewMerchantRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"merchantAddress\",\"type\":\"address\"}],\"name\":\"getMerchantID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merchantIDContract\",\"outputs\":[{\"internalType\":\"contractMerchantID\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"merchantIDsRegistry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"merchantAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"arweaveID\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610da1380380610da1833981810160405281019061003291906101c4565b61004e61004361009560201b60201c565b61009d60201b60201c565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506101f1565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061019182610166565b9050919050565b6101a181610186565b81146101ac57600080fd5b50565b6000815190506101be81610198565b92915050565b6000602082840312156101da576101d9610161565b5b60006101e8848285016101af565b91505092915050565b610ba1806102006000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80636a34e3fb1161005b5780636a34e3fb146100fe578063715018a61461011c5780638da5cb5b14610126578063f2fde38b146101445761007d565b8063101b66231461008257806332434a2e146100b257806359825f9e146100ce575b600080fd5b61009c60048036038101906100979190610622565b610160565b6040516100a99190610668565b60405180910390f35b6100cc60048036038101906100c791906107c9565b6101a9565b005b6100e860048036038101906100e39190610622565b610368565b6040516100f59190610668565b60405180910390f35b610106610380565b6040516101139190610884565b60405180910390f35b6101246103a6565b005b61012e6103ba565b60405161013b91906108ae565b60405180910390f35b61015e60048036038101906101599190610622565b6103e3565b005b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541461022b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161022290610926565b60405180910390fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d0def52184846040518363ffffffff1660e01b815260040161028a9291906109b4565b6020604051808303816000875af11580156102a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102cd9190610a10565b905080600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff167f16770655f5e10e075410826e0cbd3067ba5e87639a49407124ec270af265766e828460405161035b929190610a3d565b60405180910390a2505050565b60026020528060005260406000206000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6103ae610466565b6103b860006104e4565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6103eb610466565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361045a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045190610adf565b60405180910390fd5b610463816104e4565b50565b61046e6105a8565b73ffffffffffffffffffffffffffffffffffffffff1661048c6103ba565b73ffffffffffffffffffffffffffffffffffffffff16146104e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d990610b4b565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105ef826105c4565b9050919050565b6105ff816105e4565b811461060a57600080fd5b50565b60008135905061061c816105f6565b92915050565b600060208284031215610638576106376105ba565b5b60006106468482850161060d565b91505092915050565b6000819050919050565b6106628161064f565b82525050565b600060208201905061067d6000830184610659565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6106d68261068d565b810181811067ffffffffffffffff821117156106f5576106f461069e565b5b80604052505050565b60006107086105b0565b905061071482826106cd565b919050565b600067ffffffffffffffff8211156107345761073361069e565b5b61073d8261068d565b9050602081019050919050565b82818337600083830152505050565b600061076c61076784610719565b6106fe565b90508281526020810184848401111561078857610787610688565b5b61079384828561074a565b509392505050565b600082601f8301126107b0576107af610683565b5b81356107c0848260208601610759565b91505092915050565b600080604083850312156107e0576107df6105ba565b5b60006107ee8582860161060d565b925050602083013567ffffffffffffffff81111561080f5761080e6105bf565b5b61081b8582860161079b565b9150509250929050565b6000819050919050565b600061084a610845610840846105c4565b610825565b6105c4565b9050919050565b600061085c8261082f565b9050919050565b600061086e82610851565b9050919050565b61087e81610863565b82525050565b60006020820190506108996000830184610875565b92915050565b6108a8816105e4565b82525050565b60006020820190506108c3600083018461089f565b92915050565b600082825260208201905092915050565b7f4d65726368616e7420616c726561647920726567697374657265640000000000600082015250565b6000610910601b836108c9565b915061091b826108da565b602082019050919050565b6000602082019050818103600083015261093f81610903565b9050919050565b600081519050919050565b60005b8381101561096f578082015181840152602081019050610954565b60008484015250505050565b600061098682610946565b61099081856108c9565b93506109a0818560208601610951565b6109a98161068d565b840191505092915050565b60006040820190506109c9600083018561089f565b81810360208301526109db818461097b565b90509392505050565b6109ed8161064f565b81146109f857600080fd5b50565b600081519050610a0a816109e4565b92915050565b600060208284031215610a2657610a256105ba565b5b6000610a34848285016109fb565b91505092915050565b6000604082019050610a526000830185610659565b8181036020830152610a64818461097b565b90509392505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000610ac96026836108c9565b9150610ad482610a6d565b604082019050919050565b60006020820190508181036000830152610af881610abc565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000610b356020836108c9565b9150610b4082610aff565b602082019050919050565b60006020820190508181036000830152610b6481610b28565b905091905056fea264697066735822122028edc6fef24da332c7f9b2af171ce57d696cf32d239eaba15be519fc00be181864736f6c63430008130033",
}

// MerchantRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use MerchantRegistryMetaData.ABI instead.
var MerchantRegistryABI = MerchantRegistryMetaData.ABI

// MerchantRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerchantRegistryMetaData.Bin instead.
var MerchantRegistryBin = MerchantRegistryMetaData.Bin

// DeployMerchantRegistry deploys a new Ethereum contract, binding an instance of MerchantRegistry to it.
func DeployMerchantRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _merchantIDAddress common.Address) (common.Address, *types.Transaction, *MerchantRegistry, error) {
	parsed, err := MerchantRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerchantRegistryBin), backend, _merchantIDAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerchantRegistry{MerchantRegistryCaller: MerchantRegistryCaller{contract: contract}, MerchantRegistryTransactor: MerchantRegistryTransactor{contract: contract}, MerchantRegistryFilterer: MerchantRegistryFilterer{contract: contract}}, nil
}

// MerchantRegistry is an auto generated Go binding around an Ethereum contract.
type MerchantRegistry struct {
	MerchantRegistryCaller     // Read-only binding to the contract
	MerchantRegistryTransactor // Write-only binding to the contract
	MerchantRegistryFilterer   // Log filterer for contract events
}

// MerchantRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerchantRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerchantRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerchantRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerchantRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerchantRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerchantRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerchantRegistrySession struct {
	Contract     *MerchantRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerchantRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerchantRegistryCallerSession struct {
	Contract *MerchantRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MerchantRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerchantRegistryTransactorSession struct {
	Contract     *MerchantRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MerchantRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerchantRegistryRaw struct {
	Contract *MerchantRegistry // Generic contract binding to access the raw methods on
}

// MerchantRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerchantRegistryCallerRaw struct {
	Contract *MerchantRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// MerchantRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerchantRegistryTransactorRaw struct {
	Contract *MerchantRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerchantRegistry creates a new instance of MerchantRegistry, bound to a specific deployed contract.
func NewMerchantRegistry(address common.Address, backend bind.ContractBackend) (*MerchantRegistry, error) {
	contract, err := bindMerchantRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerchantRegistry{MerchantRegistryCaller: MerchantRegistryCaller{contract: contract}, MerchantRegistryTransactor: MerchantRegistryTransactor{contract: contract}, MerchantRegistryFilterer: MerchantRegistryFilterer{contract: contract}}, nil
}

// NewMerchantRegistryCaller creates a new read-only instance of MerchantRegistry, bound to a specific deployed contract.
func NewMerchantRegistryCaller(address common.Address, caller bind.ContractCaller) (*MerchantRegistryCaller, error) {
	contract, err := bindMerchantRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerchantRegistryCaller{contract: contract}, nil
}

// NewMerchantRegistryTransactor creates a new write-only instance of MerchantRegistry, bound to a specific deployed contract.
func NewMerchantRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*MerchantRegistryTransactor, error) {
	contract, err := bindMerchantRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerchantRegistryTransactor{contract: contract}, nil
}

// NewMerchantRegistryFilterer creates a new log filterer instance of MerchantRegistry, bound to a specific deployed contract.
func NewMerchantRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*MerchantRegistryFilterer, error) {
	contract, err := bindMerchantRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerchantRegistryFilterer{contract: contract}, nil
}

// bindMerchantRegistry binds a generic wrapper to an already deployed contract.
func bindMerchantRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MerchantRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerchantRegistry *MerchantRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerchantRegistry.Contract.MerchantRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerchantRegistry *MerchantRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.MerchantRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerchantRegistry *MerchantRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.MerchantRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerchantRegistry *MerchantRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerchantRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerchantRegistry *MerchantRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerchantRegistry *MerchantRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetMerchantID is a free data retrieval call binding the contract method 0x101b6623.
//
// Solidity: function getMerchantID(address merchantAddress) view returns(uint256)
func (_MerchantRegistry *MerchantRegistryCaller) GetMerchantID(opts *bind.CallOpts, merchantAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MerchantRegistry.contract.Call(opts, &out, "getMerchantID", merchantAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMerchantID is a free data retrieval call binding the contract method 0x101b6623.
//
// Solidity: function getMerchantID(address merchantAddress) view returns(uint256)
func (_MerchantRegistry *MerchantRegistrySession) GetMerchantID(merchantAddress common.Address) (*big.Int, error) {
	return _MerchantRegistry.Contract.GetMerchantID(&_MerchantRegistry.CallOpts, merchantAddress)
}

// GetMerchantID is a free data retrieval call binding the contract method 0x101b6623.
//
// Solidity: function getMerchantID(address merchantAddress) view returns(uint256)
func (_MerchantRegistry *MerchantRegistryCallerSession) GetMerchantID(merchantAddress common.Address) (*big.Int, error) {
	return _MerchantRegistry.Contract.GetMerchantID(&_MerchantRegistry.CallOpts, merchantAddress)
}

// MerchantIDContract is a free data retrieval call binding the contract method 0x6a34e3fb.
//
// Solidity: function merchantIDContract() view returns(address)
func (_MerchantRegistry *MerchantRegistryCaller) MerchantIDContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MerchantRegistry.contract.Call(opts, &out, "merchantIDContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MerchantIDContract is a free data retrieval call binding the contract method 0x6a34e3fb.
//
// Solidity: function merchantIDContract() view returns(address)
func (_MerchantRegistry *MerchantRegistrySession) MerchantIDContract() (common.Address, error) {
	return _MerchantRegistry.Contract.MerchantIDContract(&_MerchantRegistry.CallOpts)
}

// MerchantIDContract is a free data retrieval call binding the contract method 0x6a34e3fb.
//
// Solidity: function merchantIDContract() view returns(address)
func (_MerchantRegistry *MerchantRegistryCallerSession) MerchantIDContract() (common.Address, error) {
	return _MerchantRegistry.Contract.MerchantIDContract(&_MerchantRegistry.CallOpts)
}

// MerchantIDsRegistry is a free data retrieval call binding the contract method 0x59825f9e.
//
// Solidity: function merchantIDsRegistry(address ) view returns(uint256)
func (_MerchantRegistry *MerchantRegistryCaller) MerchantIDsRegistry(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MerchantRegistry.contract.Call(opts, &out, "merchantIDsRegistry", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MerchantIDsRegistry is a free data retrieval call binding the contract method 0x59825f9e.
//
// Solidity: function merchantIDsRegistry(address ) view returns(uint256)
func (_MerchantRegistry *MerchantRegistrySession) MerchantIDsRegistry(arg0 common.Address) (*big.Int, error) {
	return _MerchantRegistry.Contract.MerchantIDsRegistry(&_MerchantRegistry.CallOpts, arg0)
}

// MerchantIDsRegistry is a free data retrieval call binding the contract method 0x59825f9e.
//
// Solidity: function merchantIDsRegistry(address ) view returns(uint256)
func (_MerchantRegistry *MerchantRegistryCallerSession) MerchantIDsRegistry(arg0 common.Address) (*big.Int, error) {
	return _MerchantRegistry.Contract.MerchantIDsRegistry(&_MerchantRegistry.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MerchantRegistry *MerchantRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MerchantRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MerchantRegistry *MerchantRegistrySession) Owner() (common.Address, error) {
	return _MerchantRegistry.Contract.Owner(&_MerchantRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MerchantRegistry *MerchantRegistryCallerSession) Owner() (common.Address, error) {
	return _MerchantRegistry.Contract.Owner(&_MerchantRegistry.CallOpts)
}

// Register is a paid mutator transaction binding the contract method 0x32434a2e.
//
// Solidity: function register(address merchantAddress, string arweaveID) returns()
func (_MerchantRegistry *MerchantRegistryTransactor) Register(opts *bind.TransactOpts, merchantAddress common.Address, arweaveID string) (*types.Transaction, error) {
	return _MerchantRegistry.contract.Transact(opts, "register", merchantAddress, arweaveID)
}

// Register is a paid mutator transaction binding the contract method 0x32434a2e.
//
// Solidity: function register(address merchantAddress, string arweaveID) returns()
func (_MerchantRegistry *MerchantRegistrySession) Register(merchantAddress common.Address, arweaveID string) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.Register(&_MerchantRegistry.TransactOpts, merchantAddress, arweaveID)
}

// Register is a paid mutator transaction binding the contract method 0x32434a2e.
//
// Solidity: function register(address merchantAddress, string arweaveID) returns()
func (_MerchantRegistry *MerchantRegistryTransactorSession) Register(merchantAddress common.Address, arweaveID string) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.Register(&_MerchantRegistry.TransactOpts, merchantAddress, arweaveID)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MerchantRegistry *MerchantRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MerchantRegistry *MerchantRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _MerchantRegistry.Contract.RenounceOwnership(&_MerchantRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MerchantRegistry *MerchantRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MerchantRegistry.Contract.RenounceOwnership(&_MerchantRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MerchantRegistry *MerchantRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MerchantRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MerchantRegistry *MerchantRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.TransferOwnership(&_MerchantRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MerchantRegistry *MerchantRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.TransferOwnership(&_MerchantRegistry.TransactOpts, newOwner)
}

// MerchantRegistryNewMerchantRegisteredIterator is returned from FilterNewMerchantRegistered and is used to iterate over the raw logs and unpacked data for NewMerchantRegistered events raised by the MerchantRegistry contract.
type MerchantRegistryNewMerchantRegisteredIterator struct {
	Event *MerchantRegistryNewMerchantRegistered // Event containing the contract specifics and raw log

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
func (it *MerchantRegistryNewMerchantRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantRegistryNewMerchantRegistered)
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
		it.Event = new(MerchantRegistryNewMerchantRegistered)
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
func (it *MerchantRegistryNewMerchantRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantRegistryNewMerchantRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantRegistryNewMerchantRegistered represents a NewMerchantRegistered event raised by the MerchantRegistry contract.
type MerchantRegistryNewMerchantRegistered struct {
	MerchantAddress common.Address
	MerchantId      *big.Int
	ArweaveID       string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewMerchantRegistered is a free log retrieval operation binding the contract event 0x16770655f5e10e075410826e0cbd3067ba5e87639a49407124ec270af265766e.
//
// Solidity: event NewMerchantRegistered(address indexed merchantAddress, uint256 merchantId, string arweaveID)
func (_MerchantRegistry *MerchantRegistryFilterer) FilterNewMerchantRegistered(opts *bind.FilterOpts, merchantAddress []common.Address) (*MerchantRegistryNewMerchantRegisteredIterator, error) {

	var merchantAddressRule []interface{}
	for _, merchantAddressItem := range merchantAddress {
		merchantAddressRule = append(merchantAddressRule, merchantAddressItem)
	}

	logs, sub, err := _MerchantRegistry.contract.FilterLogs(opts, "NewMerchantRegistered", merchantAddressRule)
	if err != nil {
		return nil, err
	}
	return &MerchantRegistryNewMerchantRegisteredIterator{contract: _MerchantRegistry.contract, event: "NewMerchantRegistered", logs: logs, sub: sub}, nil
}

// WatchNewMerchantRegistered is a free log subscription operation binding the contract event 0x16770655f5e10e075410826e0cbd3067ba5e87639a49407124ec270af265766e.
//
// Solidity: event NewMerchantRegistered(address indexed merchantAddress, uint256 merchantId, string arweaveID)
func (_MerchantRegistry *MerchantRegistryFilterer) WatchNewMerchantRegistered(opts *bind.WatchOpts, sink chan<- *MerchantRegistryNewMerchantRegistered, merchantAddress []common.Address) (event.Subscription, error) {

	var merchantAddressRule []interface{}
	for _, merchantAddressItem := range merchantAddress {
		merchantAddressRule = append(merchantAddressRule, merchantAddressItem)
	}

	logs, sub, err := _MerchantRegistry.contract.WatchLogs(opts, "NewMerchantRegistered", merchantAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantRegistryNewMerchantRegistered)
				if err := _MerchantRegistry.contract.UnpackLog(event, "NewMerchantRegistered", log); err != nil {
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

// ParseNewMerchantRegistered is a log parse operation binding the contract event 0x16770655f5e10e075410826e0cbd3067ba5e87639a49407124ec270af265766e.
//
// Solidity: event NewMerchantRegistered(address indexed merchantAddress, uint256 merchantId, string arweaveID)
func (_MerchantRegistry *MerchantRegistryFilterer) ParseNewMerchantRegistered(log types.Log) (*MerchantRegistryNewMerchantRegistered, error) {
	event := new(MerchantRegistryNewMerchantRegistered)
	if err := _MerchantRegistry.contract.UnpackLog(event, "NewMerchantRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerchantRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MerchantRegistry contract.
type MerchantRegistryOwnershipTransferredIterator struct {
	Event *MerchantRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MerchantRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantRegistryOwnershipTransferred)
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
		it.Event = new(MerchantRegistryOwnershipTransferred)
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
func (it *MerchantRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the MerchantRegistry contract.
type MerchantRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MerchantRegistry *MerchantRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MerchantRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MerchantRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MerchantRegistryOwnershipTransferredIterator{contract: _MerchantRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MerchantRegistry *MerchantRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MerchantRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MerchantRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantRegistryOwnershipTransferred)
				if err := _MerchantRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MerchantRegistry *MerchantRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*MerchantRegistryOwnershipTransferred, error) {
	event := new(MerchantRegistryOwnershipTransferred)
	if err := _MerchantRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
