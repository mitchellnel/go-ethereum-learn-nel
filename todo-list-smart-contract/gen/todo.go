// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package todo

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
)

// TodoTask is an auto generated low-level Go binding around an user-defined struct.
type TodoTask struct {
	Content   string
	Completed bool
}

// TodoMetaData contains all meta data concerning the Todo contract.
var TodoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"addTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"deleteTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getTask\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listTasks\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"toggleTaskCompletion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"updateTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611042806100606000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80636a2b58841161005b5780636a2b5884146100ea57806386533f6c146101085780638da5cb5b14610124578063a9bda587146101425761007d565b80631d65e77e14610082578063560f3192146100b257806367238562146100ce575b600080fd5b61009c600480360381019061009791906109b6565b61015e565b6040516100a99190610ad4565b60405180910390f35b6100cc60048036038101906100c791906109b6565b61029f565b005b6100e860048036038101906100e39190610c2b565b61041a565b005b6100f26104f7565b6040516100ff9190610d73565b60405180910390f35b610122600480360381019061011d9190610d95565b61065b565b005b61012c6106f1565b6040516101399190610e32565b60405180910390f35b61015c600480360381019061015791906109b6565b610715565b005b6101666107e0565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101be57600080fd5b600182815481106101d2576101d1610e4d565b5b90600052602060002090600202016040518060400160405290816000820180546101fb90610eab565b80601f016020809104026020016040519081016040528092919081815260200182805461022790610eab565b80156102745780601f1061024957610100808354040283529160200191610274565b820191906000526020600020905b81548152906001019060200180831161025757829003601f168201915b505050505081526020016001820160009054906101000a900460ff1615151515815250509050919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102f757600080fd5b60008190505b6001808054905061030e9190610f0b565b8110156103c757600180836103239190610f3f565b8154811061033457610333610e4d565b5b90600052602060002090600202016001838154811061035657610355610e4d565b5b9060005260206000209060020201600082018160000190805461037890610eab565b6103839291906107fc565b506001820160009054906101000a900460ff168160010160006101000a81548160ff02191690831515021790555090505080806103bf90610f95565b9150506102fd565b5060018054806103da576103d9610fdd565b5b6001900381819060005260206000209060020201600080820160006103ff9190610889565b6001820160006101000a81549060ff02191690555050905550565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461047257600080fd5b6001604051806040016040528083815260200160001515815250908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000190805190602001906104d19291906108c9565b5060208201518160010160006101000a81548160ff021916908315150217905550505050565b606060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461055157600080fd5b6001805480602002602001604051908101604052809291908181526020016000905b8282101561065257838290600052602060002090600202016040518060400160405290816000820180546105a690610eab565b80601f01602080910402602001604051908101604052809291908181526020018280546105d290610eab565b801561061f5780601f106105f45761010080835404028352916020019161061f565b820191906000526020600020905b81548152906001019060200180831161060257829003601f168201915b505050505081526020016001820160009054906101000a900460ff16151515158152505081526020019060010190610573565b50505050905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106b357600080fd5b80600183815481106106c8576106c7610e4d565b5b906000526020600020906002020160000190805190602001906106ec9291906108c9565b505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461076d57600080fd5b6001818154811061078157610780610e4d565b5b906000526020600020906002020160010160009054906101000a900460ff1615600182815481106107b5576107b4610e4d565b5b906000526020600020906002020160010160006101000a81548160ff02191690831515021790555050565b6040518060400160405280606081526020016000151581525090565b82805461080890610eab565b90600052602060002090601f01602090048101928261082a5760008555610878565b82601f1061083b5780548555610878565b8280016001018555821561087857600052602060002091601f016020900482015b8281111561087757825482559160010191906001019061085c565b5b509050610885919061094f565b5090565b50805461089590610eab565b6000825580601f106108a757506108c6565b601f0160209004906000526020600020908101906108c5919061094f565b5b50565b8280546108d590610eab565b90600052602060002090601f0160209004810192826108f7576000855561093e565b82601f1061091057805160ff191683800117855561093e565b8280016001018555821561093e579182015b8281111561093d578251825591602001919060010190610922565b5b50905061094b919061094f565b5090565b5b80821115610968576000816000905550600101610950565b5090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b61099381610980565b811461099e57600080fd5b50565b6000813590506109b08161098a565b92915050565b6000602082840312156109cc576109cb610976565b5b60006109da848285016109a1565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610a1d578082015181840152602081019050610a02565b83811115610a2c576000848401525b50505050565b6000601f19601f8301169050919050565b6000610a4e826109e3565b610a5881856109ee565b9350610a688185602086016109ff565b610a7181610a32565b840191505092915050565b60008115159050919050565b610a9181610a7c565b82525050565b60006040830160008301518482036000860152610ab48282610a43565b9150506020830151610ac96020860182610a88565b508091505092915050565b60006020820190508181036000830152610aee8184610a97565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610b3882610a32565b810181811067ffffffffffffffff82111715610b5757610b56610b00565b5b80604052505050565b6000610b6a61096c565b9050610b768282610b2f565b919050565b600067ffffffffffffffff821115610b9657610b95610b00565b5b610b9f82610a32565b9050602081019050919050565b82818337600083830152505050565b6000610bce610bc984610b7b565b610b60565b905082815260208101848484011115610bea57610be9610afb565b5b610bf5848285610bac565b509392505050565b600082601f830112610c1257610c11610af6565b5b8135610c22848260208601610bbb565b91505092915050565b600060208284031215610c4157610c40610976565b5b600082013567ffffffffffffffff811115610c5f57610c5e61097b565b5b610c6b84828501610bfd565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60006040830160008301518482036000860152610cbd8282610a43565b9150506020830151610cd26020860182610a88565b508091505092915050565b6000610ce98383610ca0565b905092915050565b6000602082019050919050565b6000610d0982610c74565b610d138185610c7f565b935083602082028501610d2585610c90565b8060005b85811015610d615784840389528151610d428582610cdd565b9450610d4d83610cf1565b925060208a01995050600181019050610d29565b50829750879550505050505092915050565b60006020820190508181036000830152610d8d8184610cfe565b905092915050565b60008060408385031215610dac57610dab610976565b5b6000610dba858286016109a1565b925050602083013567ffffffffffffffff811115610ddb57610dda61097b565b5b610de785828601610bfd565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610e1c82610df1565b9050919050565b610e2c81610e11565b82525050565b6000602082019050610e476000830184610e23565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610ec357607f821691505b602082108103610ed657610ed5610e7c565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610f1682610980565b9150610f2183610980565b925082821015610f3457610f33610edc565b5b828203905092915050565b6000610f4a82610980565b9150610f5583610980565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610f8a57610f89610edc565b5b828201905092915050565b6000610fa082610980565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610fd257610fd1610edc565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea2646970667358221220218832d8194af190ba1186afbdf1d17f387c6a33a288cb5aa015bd6ad42223e564736f6c634300080e0033",
}

// TodoABI is the input ABI used to generate the binding from.
// Deprecated: Use TodoMetaData.ABI instead.
var TodoABI = TodoMetaData.ABI

// TodoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TodoMetaData.Bin instead.
var TodoBin = TodoMetaData.Bin

// DeployTodo deploys a new Ethereum contract, binding an instance of Todo to it.
func DeployTodo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Todo, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TodoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// Todo is an auto generated Go binding around an Ethereum contract.
type Todo struct {
	TodoCaller     // Read-only binding to the contract
	TodoTransactor // Write-only binding to the contract
	TodoFilterer   // Log filterer for contract events
}

// TodoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TodoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TodoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TodoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TodoSession struct {
	Contract     *Todo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TodoCallerSession struct {
	Contract *TodoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TodoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TodoTransactorSession struct {
	Contract     *TodoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TodoRaw struct {
	Contract *Todo // Generic contract binding to access the raw methods on
}

// TodoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TodoCallerRaw struct {
	Contract *TodoCaller // Generic read-only contract binding to access the raw methods on
}

// TodoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TodoTransactorRaw struct {
	Contract *TodoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTodo creates a new instance of Todo, bound to a specific deployed contract.
func NewTodo(address common.Address, backend bind.ContractBackend) (*Todo, error) {
	contract, err := bindTodo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// NewTodoCaller creates a new read-only instance of Todo, bound to a specific deployed contract.
func NewTodoCaller(address common.Address, caller bind.ContractCaller) (*TodoCaller, error) {
	contract, err := bindTodo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TodoCaller{contract: contract}, nil
}

// NewTodoTransactor creates a new write-only instance of Todo, bound to a specific deployed contract.
func NewTodoTransactor(address common.Address, transactor bind.ContractTransactor) (*TodoTransactor, error) {
	contract, err := bindTodo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TodoTransactor{contract: contract}, nil
}

// NewTodoFilterer creates a new log filterer instance of Todo, bound to a specific deployed contract.
func NewTodoFilterer(address common.Address, filterer bind.ContractFilterer) (*TodoFilterer, error) {
	contract, err := bindTodo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TodoFilterer{contract: contract}, nil
}

// bindTodo binds a generic wrapper to an already deployed contract.
func bindTodo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TodoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.TodoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transact(opts, method, params...)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 _id) view returns((string,bool))
func (_Todo *TodoCaller) GetTask(opts *bind.CallOpts, _id *big.Int) (TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "getTask", _id)

	if err != nil {
		return *new(TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new(TodoTask)).(*TodoTask)

	return out0, err

}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 _id) view returns((string,bool))
func (_Todo *TodoSession) GetTask(_id *big.Int) (TodoTask, error) {
	return _Todo.Contract.GetTask(&_Todo.CallOpts, _id)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 _id) view returns((string,bool))
func (_Todo *TodoCallerSession) GetTask(_id *big.Int) (TodoTask, error) {
	return _Todo.Contract.GetTask(&_Todo.CallOpts, _id)
}

// ListTasks is a free data retrieval call binding the contract method 0x6a2b5884.
//
// Solidity: function listTasks() view returns((string,bool)[])
func (_Todo *TodoCaller) ListTasks(opts *bind.CallOpts) ([]TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "listTasks")

	if err != nil {
		return *new([]TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]TodoTask)).(*[]TodoTask)

	return out0, err

}

// ListTasks is a free data retrieval call binding the contract method 0x6a2b5884.
//
// Solidity: function listTasks() view returns((string,bool)[])
func (_Todo *TodoSession) ListTasks() ([]TodoTask, error) {
	return _Todo.Contract.ListTasks(&_Todo.CallOpts)
}

// ListTasks is a free data retrieval call binding the contract method 0x6a2b5884.
//
// Solidity: function listTasks() view returns((string,bool)[])
func (_Todo *TodoCallerSession) ListTasks() ([]TodoTask, error) {
	return _Todo.Contract.ListTasks(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCallerSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// AddTask is a paid mutator transaction binding the contract method 0x67238562.
//
// Solidity: function addTask(string _content) returns()
func (_Todo *TodoTransactor) AddTask(opts *bind.TransactOpts, _content string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "addTask", _content)
}

// AddTask is a paid mutator transaction binding the contract method 0x67238562.
//
// Solidity: function addTask(string _content) returns()
func (_Todo *TodoSession) AddTask(_content string) (*types.Transaction, error) {
	return _Todo.Contract.AddTask(&_Todo.TransactOpts, _content)
}

// AddTask is a paid mutator transaction binding the contract method 0x67238562.
//
// Solidity: function addTask(string _content) returns()
func (_Todo *TodoTransactorSession) AddTask(_content string) (*types.Transaction, error) {
	return _Todo.Contract.AddTask(&_Todo.TransactOpts, _content)
}

// DeleteTask is a paid mutator transaction binding the contract method 0x560f3192.
//
// Solidity: function deleteTask(uint256 _id) returns()
func (_Todo *TodoTransactor) DeleteTask(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "deleteTask", _id)
}

// DeleteTask is a paid mutator transaction binding the contract method 0x560f3192.
//
// Solidity: function deleteTask(uint256 _id) returns()
func (_Todo *TodoSession) DeleteTask(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.DeleteTask(&_Todo.TransactOpts, _id)
}

// DeleteTask is a paid mutator transaction binding the contract method 0x560f3192.
//
// Solidity: function deleteTask(uint256 _id) returns()
func (_Todo *TodoTransactorSession) DeleteTask(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.DeleteTask(&_Todo.TransactOpts, _id)
}

// ToggleTaskCompletion is a paid mutator transaction binding the contract method 0xa9bda587.
//
// Solidity: function toggleTaskCompletion(uint256 _id) returns()
func (_Todo *TodoTransactor) ToggleTaskCompletion(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "toggleTaskCompletion", _id)
}

// ToggleTaskCompletion is a paid mutator transaction binding the contract method 0xa9bda587.
//
// Solidity: function toggleTaskCompletion(uint256 _id) returns()
func (_Todo *TodoSession) ToggleTaskCompletion(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.ToggleTaskCompletion(&_Todo.TransactOpts, _id)
}

// ToggleTaskCompletion is a paid mutator transaction binding the contract method 0xa9bda587.
//
// Solidity: function toggleTaskCompletion(uint256 _id) returns()
func (_Todo *TodoTransactorSession) ToggleTaskCompletion(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.ToggleTaskCompletion(&_Todo.TransactOpts, _id)
}

// UpdateTask is a paid mutator transaction binding the contract method 0x86533f6c.
//
// Solidity: function updateTask(uint256 _id, string _content) returns()
func (_Todo *TodoTransactor) UpdateTask(opts *bind.TransactOpts, _id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "updateTask", _id, _content)
}

// UpdateTask is a paid mutator transaction binding the contract method 0x86533f6c.
//
// Solidity: function updateTask(uint256 _id, string _content) returns()
func (_Todo *TodoSession) UpdateTask(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.Contract.UpdateTask(&_Todo.TransactOpts, _id, _content)
}

// UpdateTask is a paid mutator transaction binding the contract method 0x86533f6c.
//
// Solidity: function updateTask(uint256 _id, string _content) returns()
func (_Todo *TodoTransactorSession) UpdateTask(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.Contract.UpdateTask(&_Todo.TransactOpts, _id, _content)
}
