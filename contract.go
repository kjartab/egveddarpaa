// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// WinnerTakesAllABI is the input ABI used to generate the binding from.
const WinnerTakesAllABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"deadlineCampaign\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"winningAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getProjectInfo\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"url\",\"type\":\"string\"},{\"name\":\"funds\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"winningFunds\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"url\",\"type\":\"string\"}],\"name\":\"submitProject\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"projectAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"supportProject\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finish\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberOfProjects\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deadlineProjects\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_minimumEntryFee\",\"type\":\"uint256\"},{\"name\":\"_durationProjects\",\"type\":\"uint256\"},{\"name\":\"_durationCampaign\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"url\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"initialized\",\"type\":\"bool\"}],\"name\":\"ProjectSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ProjectSupported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winningFunds\",\"type\":\"uint256\"}],\"name\":\"PayedOutTo\",\"type\":\"event\"}]"

// WinnerTakesAllBin is the compiled bytecode used for deploying new contracts.
const WinnerTakesAllBin = `0x6060604052341561000f57600080fd5b604051606080610b0a8339810160405280805191906020018051919060200180519150505b81811161004057600080fd5b600083815542838101600155820160025560048054600160a060020a03191633600160a060020a03161790556003555b5050505b610a87806100836000396000f300606060405236156100a15763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633356331d81146100a65780633ea8b3bc146100cb57806347ee45b9146100fa5780637ecdf746146102035780637ee3e74f14610228578063cc13fd3b146102c4578063cf5e9747146102f6578063d56b28891461031e578063db5e4c7f14610333578063eb1f9d3914610358575b600080fd5b34156100b157600080fd5b6100b961037d565b60405190815260200160405180910390f35b34156100d657600080fd5b6100de610383565b604051600160a060020a03909116815260200160405180910390f35b341561010557600080fd5b610119600160a060020a0360043516610392565b604051808060200180602001848152602001838103835286818151815260200191508051906020019080838360005b838110156101615780820151818401525b602001610148565b50505050905090810190601f16801561018e5780820380516001836020036101000a031916815260200191505b50838103825285818151815260200191508051906020019080838360005b838110156101c55780820151818401525b6020016101ac565b50505050905090810190601f1680156101f25780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b341561020e57600080fd5b6100b9610521565b60405190815260200160405180910390f35b6102b060046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061052795505050505050565b604051901515815260200160405180910390f35b34156102cf57600080fd5b6100de6004356107c2565b604051600160a060020a03909116815260200160405180910390f35b6102b0600160a060020a03600435166107f4565b604051901515815260200160405180910390f35b341561032957600080fd5b610331610906565b005b341561033e57600080fd5b6100b9610973565b60405190815260200160405180910390f35b341561036357600080fd5b6100b9610979565b60405190815260200160405180910390f35b60025481565b600454600160a060020a031681565b61039a61097f565b6103a261097f565b600160a060020a0383166000908152600560205260408120600481015460ff1615156103cd57600080fd5b80600101816002018260030154828054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561046f5780601f106104445761010080835404028352916020019161046f565b820191906000526020600020905b81548152906001019060200180831161045257829003601f168201915b50505050509250818054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561050b5780601f106104e05761010080835404028352916020019161050b565b820191906000526020600020905b8154815290600101906020018083116104ee57829003601f168201915b505050505091509350935093505b509193909250565b60035481565b6000805434101561053757600080fd5b60015442111561054657600080fd5b600160a060020a03331660009081526005602052604090206004015460ff1615156107b85760a06040519081016040908152600160a060020a03331680835260208084018790528284018690526000606085018190526001608086015291825260059052208151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03919091161781556020820151816001019080516105ef929160200190610991565b5060408201518160020190805161060a929160200190610991565b50606082015181600301556080820151600491909101805460ff19169115159190911790555060068054600181016106428382610a10565b916000526020600020900160005b8154600160a060020a03338181166101009490940a8481029202199092161790925560065460075560009081526005602052604090819020600401547fa31614ce68f0443c0017f17939642084995e6a8b91736cac18520b829b157c3393508691869160ff169051600160a060020a03851681528115156060820152608060208201818152906040830190830186818151815260200191508051906020019080838360005b8381101561070e5780820151818401525b6020016106f5565b50505050905090810190601f16801561073b5780820380516001836020036101000a031916815260200191505b50838103825285818151815260200191508051906020019080838360005b838110156107725780820151818401525b602001610759565b50505050905090810190601f16801561079f5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a15060016107bc565b5060005b92915050565b60068054829081106107d057fe5b906000526020600020900160005b915054906101000a9004600160a060020a031681565b60003481901161080357600080fd5b60025442118061081557506001544211155b1561081f57600080fd5b600160a060020a03821660009081526005602052604090206004015460ff16151561084957600080fd5b600160a060020a03821660009081526005602052604090206003908101805434019081905590549011156108b9576004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038416908117909155600090815260056020526040902060039081015490555b7fdaf9d988cff8579fb76d980affd31da41317fddf02e5134b4c00220b2cb2f3748234604051600160a060020a03909216825260208201526040908101905180910390a15060015b919050565b6002544210610970576004546003547f8fba92fdc99409a91fdfb6b6b1da7ddf633a9c3cc3fc11f237866e54321d39a991600160a060020a031690604051600160a060020a03909216825260208201526040908101905180910390a1600454600160a060020a0316ff5b5b565b60075481565b60015481565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106109d257805160ff19168380011785556109ff565b828001600101855582156109ff579182015b828111156109ff5782518255916020019190600101906109e4565b5b50610a0c929150610a3a565b5090565b815481835581811511610a3457600083815260209020610a34918101908301610a3a565b5b505050565b610a5891905b80821115610a0c5760008155600101610a40565b5090565b905600a165627a7a7230582093f81d5acd2286b45ce303f13ba321b024dec9f38e25dad4ee14defcfa0398970029`

// DeployWinnerTakesAll deploys a new Ethereum contract, binding an instance of WinnerTakesAll to it.
func DeployWinnerTakesAll(auth *bind.TransactOpts, backend bind.ContractBackend, _minimumEntryFee *big.Int, _durationProjects *big.Int, _durationCampaign *big.Int) (common.Address, *types.Transaction, *WinnerTakesAll, error) {
	parsed, err := abi.JSON(strings.NewReader(WinnerTakesAllABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WinnerTakesAllBin), backend, _minimumEntryFee, _durationProjects, _durationCampaign)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WinnerTakesAll{WinnerTakesAllCaller: WinnerTakesAllCaller{contract: contract}, WinnerTakesAllTransactor: WinnerTakesAllTransactor{contract: contract}}, nil
}

// WinnerTakesAll is an auto generated Go binding around an Ethereum contract.
type WinnerTakesAll struct {
	WinnerTakesAllCaller     // Read-only binding to the contract
	WinnerTakesAllTransactor // Write-only binding to the contract
}

// WinnerTakesAllCaller is an auto generated read-only Go binding around an Ethereum contract.
type WinnerTakesAllCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WinnerTakesAllTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WinnerTakesAllTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WinnerTakesAllSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WinnerTakesAllSession struct {
	Contract     *WinnerTakesAll   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WinnerTakesAllCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WinnerTakesAllCallerSession struct {
	Contract *WinnerTakesAllCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// WinnerTakesAllTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WinnerTakesAllTransactorSession struct {
	Contract     *WinnerTakesAllTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// WinnerTakesAllRaw is an auto generated low-level Go binding around an Ethereum contract.
type WinnerTakesAllRaw struct {
	Contract *WinnerTakesAll // Generic contract binding to access the raw methods on
}

// WinnerTakesAllCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WinnerTakesAllCallerRaw struct {
	Contract *WinnerTakesAllCaller // Generic read-only contract binding to access the raw methods on
}

// WinnerTakesAllTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WinnerTakesAllTransactorRaw struct {
	Contract *WinnerTakesAllTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWinnerTakesAll creates a new instance of WinnerTakesAll, bound to a specific deployed contract.
func NewWinnerTakesAll(address common.Address, backend bind.ContractBackend) (*WinnerTakesAll, error) {
	contract, err := bindWinnerTakesAll(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WinnerTakesAll{WinnerTakesAllCaller: WinnerTakesAllCaller{contract: contract}, WinnerTakesAllTransactor: WinnerTakesAllTransactor{contract: contract}}, nil
}

// NewWinnerTakesAllCaller creates a new read-only instance of WinnerTakesAll, bound to a specific deployed contract.
func NewWinnerTakesAllCaller(address common.Address, caller bind.ContractCaller) (*WinnerTakesAllCaller, error) {
	contract, err := bindWinnerTakesAll(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &WinnerTakesAllCaller{contract: contract}, nil
}

// NewWinnerTakesAllTransactor creates a new write-only instance of WinnerTakesAll, bound to a specific deployed contract.
func NewWinnerTakesAllTransactor(address common.Address, transactor bind.ContractTransactor) (*WinnerTakesAllTransactor, error) {
	contract, err := bindWinnerTakesAll(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &WinnerTakesAllTransactor{contract: contract}, nil
}

// bindWinnerTakesAll binds a generic wrapper to an already deployed contract.
func bindWinnerTakesAll(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WinnerTakesAllABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WinnerTakesAll *WinnerTakesAllRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WinnerTakesAll.Contract.WinnerTakesAllCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WinnerTakesAll *WinnerTakesAllRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WinnerTakesAll.Contract.WinnerTakesAllTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WinnerTakesAll *WinnerTakesAllRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WinnerTakesAll.Contract.WinnerTakesAllTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WinnerTakesAll *WinnerTakesAllCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WinnerTakesAll.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WinnerTakesAll *WinnerTakesAllTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WinnerTakesAll.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WinnerTakesAll *WinnerTakesAllTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WinnerTakesAll.Contract.contract.Transact(opts, method, params...)
}

// DeadlineCampaign is a free data retrieval call binding the contract method 0x3356331d.
//
// Solidity: function deadlineCampaign() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllCaller) DeadlineCampaign(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WinnerTakesAll.contract.Call(opts, out, "deadlineCampaign")
	return *ret0, err
}

// DeadlineCampaign is a free data retrieval call binding the contract method 0x3356331d.
//
// Solidity: function deadlineCampaign() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllSession) DeadlineCampaign() (*big.Int, error) {
	return _WinnerTakesAll.Contract.DeadlineCampaign(&_WinnerTakesAll.CallOpts)
}

// DeadlineCampaign is a free data retrieval call binding the contract method 0x3356331d.
//
// Solidity: function deadlineCampaign() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllCallerSession) DeadlineCampaign() (*big.Int, error) {
	return _WinnerTakesAll.Contract.DeadlineCampaign(&_WinnerTakesAll.CallOpts)
}

// DeadlineProjects is a free data retrieval call binding the contract method 0xeb1f9d39.
//
// Solidity: function deadlineProjects() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllCaller) DeadlineProjects(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WinnerTakesAll.contract.Call(opts, out, "deadlineProjects")
	return *ret0, err
}

// DeadlineProjects is a free data retrieval call binding the contract method 0xeb1f9d39.
//
// Solidity: function deadlineProjects() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllSession) DeadlineProjects() (*big.Int, error) {
	return _WinnerTakesAll.Contract.DeadlineProjects(&_WinnerTakesAll.CallOpts)
}

// DeadlineProjects is a free data retrieval call binding the contract method 0xeb1f9d39.
//
// Solidity: function deadlineProjects() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllCallerSession) DeadlineProjects() (*big.Int, error) {
	return _WinnerTakesAll.Contract.DeadlineProjects(&_WinnerTakesAll.CallOpts)
}

// GetProjectInfo is a free data retrieval call binding the contract method 0x47ee45b9.
//
// Solidity: function getProjectInfo(addr address) constant returns(name string, url string, funds uint256)
func (_WinnerTakesAll *WinnerTakesAllCaller) GetProjectInfo(opts *bind.CallOpts, addr common.Address) (struct {
	Name  string
	Url   string
	Funds *big.Int
}, error) {
	ret := new(struct {
		Name  string
		Url   string
		Funds *big.Int
	})
	out := ret
	err := _WinnerTakesAll.contract.Call(opts, out, "getProjectInfo", addr)
	return *ret, err
}

// GetProjectInfo is a free data retrieval call binding the contract method 0x47ee45b9.
//
// Solidity: function getProjectInfo(addr address) constant returns(name string, url string, funds uint256)
func (_WinnerTakesAll *WinnerTakesAllSession) GetProjectInfo(addr common.Address) (struct {
	Name  string
	Url   string
	Funds *big.Int
}, error) {
	return _WinnerTakesAll.Contract.GetProjectInfo(&_WinnerTakesAll.CallOpts, addr)
}

// GetProjectInfo is a free data retrieval call binding the contract method 0x47ee45b9.
//
// Solidity: function getProjectInfo(addr address) constant returns(name string, url string, funds uint256)
func (_WinnerTakesAll *WinnerTakesAllCallerSession) GetProjectInfo(addr common.Address) (struct {
	Name  string
	Url   string
	Funds *big.Int
}, error) {
	return _WinnerTakesAll.Contract.GetProjectInfo(&_WinnerTakesAll.CallOpts, addr)
}

// NumberOfProjects is a free data retrieval call binding the contract method 0xdb5e4c7f.
//
// Solidity: function numberOfProjects() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllCaller) NumberOfProjects(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WinnerTakesAll.contract.Call(opts, out, "numberOfProjects")
	return *ret0, err
}

// NumberOfProjects is a free data retrieval call binding the contract method 0xdb5e4c7f.
//
// Solidity: function numberOfProjects() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllSession) NumberOfProjects() (*big.Int, error) {
	return _WinnerTakesAll.Contract.NumberOfProjects(&_WinnerTakesAll.CallOpts)
}

// NumberOfProjects is a free data retrieval call binding the contract method 0xdb5e4c7f.
//
// Solidity: function numberOfProjects() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllCallerSession) NumberOfProjects() (*big.Int, error) {
	return _WinnerTakesAll.Contract.NumberOfProjects(&_WinnerTakesAll.CallOpts)
}

// ProjectAddresses is a free data retrieval call binding the contract method 0xcc13fd3b.
//
// Solidity: function projectAddresses( uint256) constant returns(address)
func (_WinnerTakesAll *WinnerTakesAllCaller) ProjectAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WinnerTakesAll.contract.Call(opts, out, "projectAddresses", arg0)
	return *ret0, err
}

// ProjectAddresses is a free data retrieval call binding the contract method 0xcc13fd3b.
//
// Solidity: function projectAddresses( uint256) constant returns(address)
func (_WinnerTakesAll *WinnerTakesAllSession) ProjectAddresses(arg0 *big.Int) (common.Address, error) {
	return _WinnerTakesAll.Contract.ProjectAddresses(&_WinnerTakesAll.CallOpts, arg0)
}

// ProjectAddresses is a free data retrieval call binding the contract method 0xcc13fd3b.
//
// Solidity: function projectAddresses( uint256) constant returns(address)
func (_WinnerTakesAll *WinnerTakesAllCallerSession) ProjectAddresses(arg0 *big.Int) (common.Address, error) {
	return _WinnerTakesAll.Contract.ProjectAddresses(&_WinnerTakesAll.CallOpts, arg0)
}

// WinningAddress is a free data retrieval call binding the contract method 0x3ea8b3bc.
//
// Solidity: function winningAddress() constant returns(address)
func (_WinnerTakesAll *WinnerTakesAllCaller) WinningAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WinnerTakesAll.contract.Call(opts, out, "winningAddress")
	return *ret0, err
}

// WinningAddress is a free data retrieval call binding the contract method 0x3ea8b3bc.
//
// Solidity: function winningAddress() constant returns(address)
func (_WinnerTakesAll *WinnerTakesAllSession) WinningAddress() (common.Address, error) {
	return _WinnerTakesAll.Contract.WinningAddress(&_WinnerTakesAll.CallOpts)
}

// WinningAddress is a free data retrieval call binding the contract method 0x3ea8b3bc.
//
// Solidity: function winningAddress() constant returns(address)
func (_WinnerTakesAll *WinnerTakesAllCallerSession) WinningAddress() (common.Address, error) {
	return _WinnerTakesAll.Contract.WinningAddress(&_WinnerTakesAll.CallOpts)
}

// WinningFunds is a free data retrieval call binding the contract method 0x7ecdf746.
//
// Solidity: function winningFunds() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllCaller) WinningFunds(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WinnerTakesAll.contract.Call(opts, out, "winningFunds")
	return *ret0, err
}

// WinningFunds is a free data retrieval call binding the contract method 0x7ecdf746.
//
// Solidity: function winningFunds() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllSession) WinningFunds() (*big.Int, error) {
	return _WinnerTakesAll.Contract.WinningFunds(&_WinnerTakesAll.CallOpts)
}

// WinningFunds is a free data retrieval call binding the contract method 0x7ecdf746.
//
// Solidity: function winningFunds() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllCallerSession) WinningFunds() (*big.Int, error) {
	return _WinnerTakesAll.Contract.WinningFunds(&_WinnerTakesAll.CallOpts)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_WinnerTakesAll *WinnerTakesAllTransactor) Finish(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WinnerTakesAll.contract.Transact(opts, "finish")
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_WinnerTakesAll *WinnerTakesAllSession) Finish() (*types.Transaction, error) {
	return _WinnerTakesAll.Contract.Finish(&_WinnerTakesAll.TransactOpts)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_WinnerTakesAll *WinnerTakesAllTransactorSession) Finish() (*types.Transaction, error) {
	return _WinnerTakesAll.Contract.Finish(&_WinnerTakesAll.TransactOpts)
}

// SubmitProject is a paid mutator transaction binding the contract method 0x7ee3e74f.
//
// Solidity: function submitProject(name string, url string) returns(success bool)
func (_WinnerTakesAll *WinnerTakesAllTransactor) SubmitProject(opts *bind.TransactOpts, name string, url string) (*types.Transaction, error) {
	return _WinnerTakesAll.contract.Transact(opts, "submitProject", name, url)
}

// SubmitProject is a paid mutator transaction binding the contract method 0x7ee3e74f.
//
// Solidity: function submitProject(name string, url string) returns(success bool)
func (_WinnerTakesAll *WinnerTakesAllSession) SubmitProject(name string, url string) (*types.Transaction, error) {
	return _WinnerTakesAll.Contract.SubmitProject(&_WinnerTakesAll.TransactOpts, name, url)
}

// SubmitProject is a paid mutator transaction binding the contract method 0x7ee3e74f.
//
// Solidity: function submitProject(name string, url string) returns(success bool)
func (_WinnerTakesAll *WinnerTakesAllTransactorSession) SubmitProject(name string, url string) (*types.Transaction, error) {
	return _WinnerTakesAll.Contract.SubmitProject(&_WinnerTakesAll.TransactOpts, name, url)
}

// SupportProject is a paid mutator transaction binding the contract method 0xcf5e9747.
//
// Solidity: function supportProject(addr address) returns(success bool)
func (_WinnerTakesAll *WinnerTakesAllTransactor) SupportProject(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _WinnerTakesAll.contract.Transact(opts, "supportProject", addr)
}

// SupportProject is a paid mutator transaction binding the contract method 0xcf5e9747.
//
// Solidity: function supportProject(addr address) returns(success bool)
func (_WinnerTakesAll *WinnerTakesAllSession) SupportProject(addr common.Address) (*types.Transaction, error) {
	return _WinnerTakesAll.Contract.SupportProject(&_WinnerTakesAll.TransactOpts, addr)
}

// SupportProject is a paid mutator transaction binding the contract method 0xcf5e9747.
//
// Solidity: function supportProject(addr address) returns(success bool)
func (_WinnerTakesAll *WinnerTakesAllTransactorSession) SupportProject(addr common.Address) (*types.Transaction, error) {
	return _WinnerTakesAll.Contract.SupportProject(&_WinnerTakesAll.TransactOpts, addr)
}
