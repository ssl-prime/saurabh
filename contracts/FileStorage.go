// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// FileStorageABI is the input ABI used to generate the binding from.
const FileStorageABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"addFile\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getHash\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"added\",\"type\":\"event\"}]"

// FileStorageBin is the compiled bytecode used for deploying new contracts.
const FileStorageBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a0319163317905561040d806100326000396000f3fe608060405234801561001057600080fd5b506004361061005d577c01000000000000000000000000000000000000000000000000000000006000350463248bfc3b81146100625780638da5cb5b14610191578063d13319c4146101c2575b600080fd5b61018f6004803603604081101561007857600080fd5b81019060208101813564010000000081111561009357600080fd5b8201836020820111156100a557600080fd5b803590602001918460018302840111640100000000831117156100c757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561011a57600080fd5b82018360208201111561012c57600080fd5b8035906020019184600183028401116401000000008311171561014e57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061023f945050505050565b005b61019961028b565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6101ca6102a7565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102045781810151838201526020016101ec565b50505050905090810190601f1680156102315780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b336000908152600160209081526040909120825161025f92840190610349565b50336000908152600160208181526040909220845161028693919092019190850190610349565b505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b33600090815260016020818152604092839020820180548451600294821615610100026000190190911693909304601f8101839004830284018301909452838352606093909183018282801561033e5780601f106103135761010080835404028352916020019161033e565b820191906000526020600020905b81548152906001019060200180831161032157829003601f168201915b505050505090505b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061038a57805160ff19168380011785556103b7565b828001600101855582156103b7579182015b828111156103b757825182559160200191906001019061039c565b506103c39291506103c7565b5090565b61034691905b808211156103c357600081556001016103cd56fea165627a7a7230582036135af7f030ab498ad266ec75fe496fa3aa07cd471504c3c817fe95478022f70029`

// DeployFileStorage deploys a new Ethereum contract, binding an instance of FileStorage to it.
func DeployFileStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FileStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(FileStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FileStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FileStorage{FileStorageCaller: FileStorageCaller{contract: contract}, FileStorageTransactor: FileStorageTransactor{contract: contract}, FileStorageFilterer: FileStorageFilterer{contract: contract}}, nil
}

// FileStorage is an auto generated Go binding around an Ethereum contract.
type FileStorage struct {
	FileStorageCaller     // Read-only binding to the contract
	FileStorageTransactor // Write-only binding to the contract
	FileStorageFilterer   // Log filterer for contract events
}

// FileStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type FileStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FileStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FileStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FileStorageSession struct {
	Contract     *FileStorage      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FileStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FileStorageCallerSession struct {
	Contract *FileStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FileStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FileStorageTransactorSession struct {
	Contract     *FileStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FileStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type FileStorageRaw struct {
	Contract *FileStorage // Generic contract binding to access the raw methods on
}

// FileStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FileStorageCallerRaw struct {
	Contract *FileStorageCaller // Generic read-only contract binding to access the raw methods on
}

// FileStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FileStorageTransactorRaw struct {
	Contract *FileStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFileStorage creates a new instance of FileStorage, bound to a specific deployed contract.
func NewFileStorage(address common.Address, backend bind.ContractBackend) (*FileStorage, error) {
	contract, err := bindFileStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FileStorage{FileStorageCaller: FileStorageCaller{contract: contract}, FileStorageTransactor: FileStorageTransactor{contract: contract}, FileStorageFilterer: FileStorageFilterer{contract: contract}}, nil
}

// NewFileStorageCaller creates a new read-only instance of FileStorage, bound to a specific deployed contract.
func NewFileStorageCaller(address common.Address, caller bind.ContractCaller) (*FileStorageCaller, error) {
	contract, err := bindFileStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FileStorageCaller{contract: contract}, nil
}

// NewFileStorageTransactor creates a new write-only instance of FileStorage, bound to a specific deployed contract.
func NewFileStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*FileStorageTransactor, error) {
	contract, err := bindFileStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FileStorageTransactor{contract: contract}, nil
}

// NewFileStorageFilterer creates a new log filterer instance of FileStorage, bound to a specific deployed contract.
func NewFileStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*FileStorageFilterer, error) {
	contract, err := bindFileStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FileStorageFilterer{contract: contract}, nil
}

// bindFileStorage binds a generic wrapper to an already deployed contract.
func bindFileStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FileStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileStorage *FileStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FileStorage.Contract.FileStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileStorage *FileStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileStorage.Contract.FileStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileStorage *FileStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FileStorage.Contract.FileStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileStorage *FileStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FileStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileStorage *FileStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileStorage *FileStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FileStorage.Contract.contract.Transact(opts, method, params...)
}

// GetHash is a free data retrieval call binding the contract method 0xd13319c4.
//
// Solidity: function getHash() constant returns(string)
func (_FileStorage *FileStorageCaller) GetHash(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FileStorage.contract.Call(opts, out, "getHash")
	return *ret0, err
}

// GetHash is a free data retrieval call binding the contract method 0xd13319c4.
//
// Solidity: function getHash() constant returns(string)
func (_FileStorage *FileStorageSession) GetHash() (string, error) {
	return _FileStorage.Contract.GetHash(&_FileStorage.CallOpts)
}

// GetHash is a free data retrieval call binding the contract method 0xd13319c4.
//
// Solidity: function getHash() constant returns(string)
func (_FileStorage *FileStorageCallerSession) GetHash() (string, error) {
	return _FileStorage.Contract.GetHash(&_FileStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FileStorage *FileStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FileStorage.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FileStorage *FileStorageSession) Owner() (common.Address, error) {
	return _FileStorage.Contract.Owner(&_FileStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FileStorage *FileStorageCallerSession) Owner() (common.Address, error) {
	return _FileStorage.Contract.Owner(&_FileStorage.CallOpts)
}

// AddFile is a paid mutator transaction binding the contract method 0x248bfc3b.
//
// Solidity: function addFile(string hash, string name) returns()
func (_FileStorage *FileStorageTransactor) AddFile(opts *bind.TransactOpts, hash string, name string) (*types.Transaction, error) {
	return _FileStorage.contract.Transact(opts, "addFile", hash, name)
}

// AddFile is a paid mutator transaction binding the contract method 0x248bfc3b.
//
// Solidity: function addFile(string hash, string name) returns()
func (_FileStorage *FileStorageSession) AddFile(hash string, name string) (*types.Transaction, error) {
	return _FileStorage.Contract.AddFile(&_FileStorage.TransactOpts, hash, name)
}

// AddFile is a paid mutator transaction binding the contract method 0x248bfc3b.
//
// Solidity: function addFile(string hash, string name) returns()
func (_FileStorage *FileStorageTransactorSession) AddFile(hash string, name string) (*types.Transaction, error) {
	return _FileStorage.Contract.AddFile(&_FileStorage.TransactOpts, hash, name)
}

// FileStorageAddedIterator is returned from FilterAdded and is used to iterate over the raw logs and unpacked data for Added events raised by the FileStorage contract.
type FileStorageAddedIterator struct {
	Event *FileStorageAdded // Event containing the contract specifics and raw log

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
func (it *FileStorageAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileStorageAdded)
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
		it.Event = new(FileStorageAdded)
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
func (it *FileStorageAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileStorageAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileStorageAdded represents a Added event raised by the FileStorage contract.
type FileStorageAdded struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAdded is a free log retrieval operation binding the contract event 0xc87542064bc1930c362cb7f85a979ab1051627291e7db73dfda0f48bca405481.
//
// Solidity: event added()
func (_FileStorage *FileStorageFilterer) FilterAdded(opts *bind.FilterOpts) (*FileStorageAddedIterator, error) {

	logs, sub, err := _FileStorage.contract.FilterLogs(opts, "added")
	if err != nil {
		return nil, err
	}
	return &FileStorageAddedIterator{contract: _FileStorage.contract, event: "added", logs: logs, sub: sub}, nil
}

// WatchAdded is a free log subscription operation binding the contract event 0xc87542064bc1930c362cb7f85a979ab1051627291e7db73dfda0f48bca405481.
//
// Solidity: event added()
func (_FileStorage *FileStorageFilterer) WatchAdded(opts *bind.WatchOpts, sink chan<- *FileStorageAdded) (event.Subscription, error) {

	logs, sub, err := _FileStorage.contract.WatchLogs(opts, "added")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileStorageAdded)
				if err := _FileStorage.contract.UnpackLog(event, "added", log); err != nil {
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
